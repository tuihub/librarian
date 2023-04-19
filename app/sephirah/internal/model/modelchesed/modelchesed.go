package modelchesed

import "github.com/tuihub/librarian/internal/model"

type Image struct {
	ID          model.InternalID
	Name        string
	Description string
}
