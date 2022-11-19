package service

import (
	"github.com/tuihub/librarian/app/mapper/internal/biz"
	pb "github.com/tuihub/protos/pkg/librarian/mapper/v1"
)

func toBizVertexType(t pb.VertexType) biz.VertexType {
	switch t {
	case pb.VertexType_VERTEX_TYPE_ABSTRACT:
		return biz.VertexTypeAbstract
	case pb.VertexType_VERTEX_TYPE_ENTITY:
		return biz.VertexTypeEntity
	case pb.VertexType_VERTEX_TYPE_MESSAGE:
		return biz.VertexTypeMessage
	case pb.VertexType_VERTEX_TYPE_OBJECT:
		return biz.VertexTypeObject
	case pb.VertexType_VERTEX_TYPE_METADATA:
		return biz.VertexTypeMetadata
	default:
		return biz.VertexTypeUnspecified
	}
}

func toPBVertexType(t biz.VertexType) pb.VertexType {
	switch t {
	case biz.VertexTypeAbstract:
		return pb.VertexType_VERTEX_TYPE_ABSTRACT
	case biz.VertexTypeEntity:
		return pb.VertexType_VERTEX_TYPE_ENTITY
	case biz.VertexTypeMessage:
		return pb.VertexType_VERTEX_TYPE_MESSAGE
	case biz.VertexTypeObject:
		return pb.VertexType_VERTEX_TYPE_OBJECT
	case biz.VertexTypeMetadata:
		return pb.VertexType_VERTEX_TYPE_METADATA
	default:
		return pb.VertexType_VERTEX_TYPE_UNSPECIFIED
	}
}

func toBizEdgeType(t pb.EdgeType) biz.EdgeType {
	switch t {
	case pb.EdgeType_EDGE_TYPE_GENERAL:
		return biz.EdgeTypeGeneral
	case pb.EdgeType_EDGE_TYPE_EQUAL:
		return biz.EdgeTypeEqual
	case pb.EdgeType_EDGE_TYPE_CREATE:
		return biz.EdgeTypeCreate
	case pb.EdgeType_EDGE_TYPE_ENJOY:
		return biz.EdgeTypeEnjoy
	case pb.EdgeType_EDGE_TYPE_MENTION:
		return biz.EdgeTypeMention
	case pb.EdgeType_EDGE_TYPE_DERIVE:
		return biz.EdgeTypeDerive
	case pb.EdgeType_EDGE_TYPE_CONTROL:
		return biz.EdgeTypeControl
	case pb.EdgeType_EDGE_TYPE_FOLLOW:
		return biz.EdgeTypeFollow
	case pb.EdgeType_EDGE_TYPE_DESCRIBE:
		return biz.EdgeTypeDescribe
	default:
		return biz.EdgeTypeUnspecified
	}
}

func toBizVertex(v *pb.Vertex) biz.Vertex {
	return biz.Vertex{
		ID:   v.GetVid(),
		Type: toBizVertexType(v.GetType()),
	}
}

func toBizVertexList(vl []*pb.Vertex) []*biz.Vertex {
	res := make([]*biz.Vertex, len(vl))
	for i, v := range vl {
		r := toBizVertex(v)
		res[i] = &r
	}
	return res
}

func toPBVertex(v biz.Vertex) pb.Vertex {
	return pb.Vertex{
		Vid:  v.ID,
		Type: toPBVertexType(v.Type),
		Prop: nil,
	}
}

func toPBVertexList(vl []*biz.Vertex) []*pb.Vertex {
	res := make([]*pb.Vertex, len(vl))
	for i, v := range vl {
		r := toPBVertex(*v)
		res[i] = &r
	}
	return res
}

func toBizEdge(e *pb.Edge) biz.Edge {
	return biz.Edge{
		SourceID:      e.GetSrcVid(),
		DestinationID: e.GetDstVid(),
		Type:          toBizEdgeType(e.GetType()),
	}
}

func toBizEdgeList(el []*pb.Edge) []*biz.Edge {
	res := make([]*biz.Edge, len(el))
	for i, e := range el {
		r := toBizEdge(e)
		res[i] = &r
	}
	return res
}
