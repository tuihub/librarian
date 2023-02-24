package converter

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --packagePath github.com/tuihub/librarian/app/sephirah/internal/data/converter --packageName converter --output ./generated.go ./

func NewConverter() Converter {
	return &ConverterImpl{}
}
