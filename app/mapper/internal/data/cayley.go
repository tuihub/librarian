package data

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tuihub/librarian/app/mapper/internal/biz"
	"github.com/tuihub/librarian/internal/conf"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/schema"
	"github.com/cayleygraph/quad"
	"github.com/cayleygraph/quad/voc"
	"github.com/go-kratos/kratos/v2/log"

	// Import RDF vocabulary definitions to be able to expand IRIs like rdf:label.
	_ "github.com/cayleygraph/quad/voc/core"
)

type cayleyMapperRepo struct {
	db *cayley.Handle
}

type Vertex struct {
	// dummy field to enforce all object to have a <id> <rdf:type> <ex:Person> relation
	rdfType struct{}       `quad:"@type > ex:Vertex"` //nolint:unused // means nothing for Go itself
	ID      quad.IRI       `json:"@id"`
	Type    biz.VertexType `json:"ex:type"`
}

func (r *cayleyMapperRepo) InsertVertex(ctx context.Context, v biz.Vertex) error {
	qw := graph.NewWriter(r.db)
	_, err := schema.NewConfig().WriteAsQuads(qw, Vertex{
		ID:   toIRI(v.ID),
		Type: v.Type,
	})
	if err != nil {
		return err
	}
	_ = qw.Close()
	return nil
}

func (r *cayleyMapperRepo) InsertEdge(ctx context.Context, e biz.Edge) error {
	return r.db.AddQuad(quad.MakeIRI(toString(e.SourceID), toString(e.Type), toString(e.DestinationID), toString(nil)))
}

func (r *cayleyMapperRepo) FetchEqualVertex(ctx context.Context, v biz.Vertex) ([]*biz.Vertex, error) {
	p := cayley.StartPath(r.db, toIRI(v.ID)).Out(toIRI(biz.EdgeTypeEqual))
	var vertices []Vertex
	err := schema.NewConfig().LoadPathTo(ctx, r.db, &vertices, p)
	if err != nil {
		return nil, err
	}
	res := make([]*biz.Vertex, len(vertices))
	for i, vertex := range vertices {
		var id int64
		id, err = toBizID(vertex.ID)
		if err != nil {
			return nil, err
		}
		res[i] = &biz.Vertex{
			ID:   id,
			Type: vertex.Type,
		}
	}
	return res, nil
}

// NewCayley .
func NewCayley(c *conf.Mapper_Data) (*cayley.Handle, func()) {
	if c == nil || c.GetCayley() == nil {
		return nil, func() {}
	}
	if c.GetCayley().GetStore() != "memory" {
		log.Errorf("Unsupported cayley store: %s, skip initialize", c.GetCayley().GetStore())
		return nil, func() {}
	}

	db, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Errorf("Failed to initialize Cayley DB, %s", err.Error())
	}
	voc.RegisterPrefix("ex:", "github.com/tuihub/librarian")

	return db, func() {
		log.Info("closing the data resources")
		db.Close()
	}
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
