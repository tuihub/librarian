package converter

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --ignoreUnexportedFields --packagePath github.com/tuihub/librarian/app/sephirah/internal/service/converter --packageName converter --output ./generated.go ./

type Converter struct {
	toPBConverter
	toBizConverter
}

func NewConverter() Converter {
	return Converter{
		&toPBConverterImpl{},
		&toBizConverterImpl{},
	}
}

func PtrToString(u *string) string {
	if u == nil {
		return ""
	}
	return *u
}

func TimeStampToTime(t *timestamppb.Timestamp) time.Time {
	if t == nil {
		return time.UnixMicro(0)
	}
	return t.AsTime()
}
