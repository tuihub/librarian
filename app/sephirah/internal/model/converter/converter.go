package converter

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --ignoreUnexportedFields --packagePath github.com/tuihub/librarian/app/sephirah/internal/model/converter --packageName converter --output ./generated.go ./

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

func DurationPBToDuration(t *durationpb.Duration) time.Duration {
	if t == nil {
		return time.Duration(0)
	}
	return t.AsDuration()
}
