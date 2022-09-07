package data

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
)

type greeterRepo struct {
	data *Data
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
	}
}
