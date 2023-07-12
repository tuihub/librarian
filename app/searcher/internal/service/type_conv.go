package service

import (
	"github.com/tuihub/librarian/app/searcher/internal/biz"
	pb "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

func toBizIndex(index pb.Index) biz.Index {
	switch index {
	case pb.Index_INDEX_UNSPECIFIED:
		return biz.IndexUnspecified
	case pb.Index_INDEX_GENERAL:
		return biz.IndexGeneral
	case pb.Index_INDEX_GEBURA_APP:
		return biz.IndexGeburaApp
	case pb.Index_INDEX_CHESED_IMAGE:
		return biz.IndexChesedImage
	default:
		return biz.IndexUnspecified
	}
}
