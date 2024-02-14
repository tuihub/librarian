package libcodec

import (
	"errors"

	"github.com/go-kratos/kratos/v2/encoding"

	_ "github.com/go-kratos/kratos/v2/encoding/json"
)

type ContentType string

const (
	JSON ContentType = "json"
	TOML ContentType = "toml"
)

func checkContentType(contentType ContentType) (string, error) {
	switch contentType {
	case JSON:
		return string(JSON), nil
	case TOML:
		return string(TOML), nil
	default:
		return "", errors.New("unsupported content type")
	}
}

func Marshal(contentType ContentType, v interface{}) ([]byte, error) {
	c, err := checkContentType(contentType)
	if err != nil {
		return nil, err
	}
	codec := encoding.GetCodec(c)
	return codec.Marshal(v)
}

func Unmarshal(contentType ContentType, data []byte, v interface{}) error {
	c, err := checkContentType(contentType)
	if err != nil {
		return err
	}
	codec := encoding.GetCodec(c)
	return codec.Unmarshal(data, v)
}
