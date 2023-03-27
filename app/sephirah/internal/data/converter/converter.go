package converter

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --ignoreUnexportedFields --packagePath github.com/tuihub/librarian/app/sephirah/internal/data/converter --packageName converter --output ./generated.go ./

type Converter struct {
	toEntConverter
	toBizConverter
}

func NewConverter() Converter {
	return Converter{
		&toEntConverterImpl{},
		&toBizConverterImpl{},
	}
}
