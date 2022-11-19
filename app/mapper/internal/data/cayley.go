package data

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/tuihub/librarian/app/mapper/internal/biz"
	"github.com/tuihub/librarian/internal/conf"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/schema"
	"github.com/cayleygraph/quad"
	"github.com/cayleygraph/quad/voc"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	// Import sqlite to save graph to sqlite file.
	_ "github.com/cayleygraph/cayley/graph/sql/sqlite"
	// Import RDF vocabulary definitions to be able to expand IRIs like rdf:label.
	_ "github.com/cayleygraph/quad/voc/core"
)

type cayleyMapperRepo struct {
	db *cayley.Handle
}

type Vertex struct {
	// dummy field to enforce all object to have a <id> <rdf:type> <ex:Person> relation
	rdfType      struct{}       `quad:"@type > ex:Vertex"` //nolint:unused // means nothing for Go itself
	ID           quad.IRI       `json:"@id"`
	Type         biz.VertexType `json:"ex:type"`
	OriginalType biz.VertexType `json:"ex:oType"`
}

// NewCayley .
func NewCayley(c *conf.Mapper_Data) (*cayley.Handle, func(), error) {
	if c == nil || c.GetCayley() == nil {
		return nil, func() {}, nil
	}
	var db *cayley.Handle
	var err error
	switch c.GetCayley().GetStore() {
	case "memory":
		db, err = cayley.NewMemoryGraph()
		if err != nil {
			return nil, func() {}, err
		}
	case "sqlite":
		dbpath := "cayley.db"
		_, err = os.Stat(dbpath)
		if err != nil {
			err = graph.InitQuadStore("sqlite", dbpath, nil)
			if err != nil {
				return nil, func() {}, err
			}
		}
		db, err = cayley.NewGraph("sqlite", dbpath, nil)
		if err != nil {
			return nil, func() {}, err
		}
	default:
		return nil, func() {}, err
	}

	voc.RegisterPrefix("ex:", "github.com/tuihub/librarian")

	return db, func() {
		log.Info("closing the data resources")
		_ = db.Close()
	}, nil
}

func (r *cayleyMapperRepo) InsertVertex(ctx context.Context, vl []*biz.Vertex) error {
	for _, v := range vl {
		err := r.writeVertex(Vertex{
			ID:           toIRI(v.ID),
			Type:         v.Type,
			OriginalType: v.Type,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *cayleyMapperRepo) InsertEdge(ctx context.Context, el []*biz.Edge) error {
	var q []quad.Quad
	for _, e := range el {
		src := Vertex{ID: toIRI(e.SourceID)}
		err := r.readVertex(ctx, &src)
		if err != nil {
			return err
		}
		dst := Vertex{ID: toIRI(e.DestinationID)}
		err = r.readVertex(ctx, &dst)
		if err != nil {
			return err
		}
		q = append(q, quad.MakeIRI(toString(e.SourceID), toString(e.Type), toString(e.DestinationID), toString(nil)))
		q, err = r.applyEdgeRules(q, e.Type, src, dst)
		if err != nil {
			return err
		}
	}
	return r.db.AddQuadSet(q)
}

func (r *cayleyMapperRepo) FetchEqualVertex(ctx context.Context, v biz.Vertex) ([]*biz.Vertex, error) {
	var res []*biz.Vertex
	viewed := map[quad.IRI]bool{toIRI(v.ID): true}
	waitingList, err := r.loadOutVertex(ctx, toIRI(v.ID), biz.EdgeTypeEqual)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(waitingList); i++ {
		if viewed[waitingList[i].ID] {
			continue
		}
		var id int64
		id, err = toBizID(waitingList[i].ID)
		if err != nil {
			return nil, err
		}
		res = append(res, &biz.Vertex{
			ID:   id,
			Type: waitingList[i].OriginalType,
		})
		viewed[waitingList[i].ID] = true
		var newList []Vertex
		newList, err = r.loadOutVertex(ctx, waitingList[i].ID, biz.EdgeTypeEqual)
		if err != nil {
			return nil, err
		}
		waitingList = append(waitingList, newList...)
	}
	return res, nil
}

func (r *cayleyMapperRepo) loadOutVertex(ctx context.Context, id quad.IRI, ty biz.EdgeType) ([]Vertex, error) {
	p := cayley.StartPath(r.db, id).Out(toIRI(ty))
	var v []Vertex
	err := schema.NewConfig().LoadPathTo(ctx, r.db, &v, p)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (r *cayleyMapperRepo) applyEdgeRules( //nolint:gocognit // edge rule is complex
	q []quad.Quad,
	edgeType biz.EdgeType,
	src, dst Vertex,
) ([]quad.Quad, error) {
	if src.Type == biz.VertexTypeAbstract || dst.Type == biz.VertexTypeAbstract {
		if edgeType == biz.EdgeTypeEqual {
			err := r.handleAbstractEqual(&src, &dst)
			if err != nil {
				return nil, err
			}
		}
	}
	switch edgeType {
	case biz.EdgeTypeEqual:
		if src.Type != dst.Type {
			return nil, errors.BadRequest("illegal edge", "EQUAL must link vertex with same type")
		}
		q = append(q, quad.MakeIRI(string(dst.ID), toString(edgeType), string(src.ID), toString(nil)))
	case biz.EdgeTypeCreate:
		if src.Type != biz.VertexTypeEntity {
			return nil, errors.BadRequest("illegal edge", "CREATE must start at ENTITY")
		}
	case biz.EdgeTypeEnjoy:
		if src.Type != biz.VertexTypeEntity {
			return nil, errors.BadRequest("illegal edge", "ENJOY must start at ENTITY")
		}
	case biz.EdgeTypeMention:
		if src.Type != biz.VertexTypeMessage && src.Type != biz.VertexTypeObject {
			return nil, errors.BadRequest("illegal edge", "MENTION must start at MESSAGE or OBJECT")
		}
	case biz.EdgeTypeDerive:
		if src.Type != biz.VertexTypeObject {
			return nil, errors.BadRequest("illegal edge", "DERIVE must start at OBJECT")
		}
	case biz.EdgeTypeControl:
		if src.Type != biz.VertexTypeEntity {
			return nil, errors.BadRequest("illegal edge", "CONTROL must start at ENTITY")
		}
		if dst.Type != biz.VertexTypeEntity {
			return nil, errors.BadRequest("illegal edge", "CONTROL must end at ENTITY")
		}
	case biz.EdgeTypeFollow:
		if src.Type != biz.VertexTypeEntity {
			return nil, errors.BadRequest("illegal edge", "FOLLOW must start at ENTITY")
		}
		if dst.Type != biz.VertexTypeEntity {
			return nil, errors.BadRequest("illegal edge", "CONTROL must end at ENTITY")
		}
	case biz.EdgeTypeDescribe:
		if src.Type != biz.VertexTypeMetadata {
			return nil, errors.BadRequest("illegal edge", "DESCRIBE must start at METADATA")
		}
	default:
		return nil, errors.BadRequest("Unsupported edge type", "")
	}
	return q, nil
}

func (r *cayleyMapperRepo) writeVertex(v Vertex) error {
	qw := graph.NewWriter(r.db)
	_, err := schema.NewConfig().WriteAsQuads(qw, v)
	_ = qw.Close()
	return err
}

func (r *cayleyMapperRepo) readVertex(ctx context.Context, v *Vertex) error {
	return schema.NewConfig().LoadTo(ctx, r.db, v, v.ID)
}

func (r *cayleyMapperRepo) handleAbstractEqual(src, dst *Vertex) error {
	if src.Type == dst.Type {
		return nil
	}
	if src.Type == biz.VertexTypeAbstract {
		// TODO check exist edges
		src.Type = dst.Type
		return r.writeVertex(*src)
	}
	if dst.Type == biz.VertexTypeAbstract {
		// TODO check exist edges
		dst.Type = src.Type
		return r.writeVertex(*dst)
	}
	return nil
}

func toString(i interface{}) string {
	return fmt.Sprint(i)
}

func toIRI(i interface{}) quad.IRI {
	return quad.IRI(fmt.Sprint(i))
}

func toBizID(id quad.IRI) (int64, error) {
	return strconv.ParseInt(string(id), 10, 64)
}
