// https://github.com/go-kratos/kratos/pull/2811

package libcodec

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/go-kratos/kratos/v2/encoding"
)

func init() { //nolint: gochecknoinits // required by the encoding package
	encoding.RegisterCodec(tomlCodec{})
}

// tomlCodec is a Codec implementation with toml.
type tomlCodec struct{}

func (c tomlCodec) Marshal(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	encoder := toml.NewEncoder(buf)

	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	data := buf.Bytes()
	return data, nil
}

func (c tomlCodec) Unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}

func (c tomlCodec) Name() string {
	return string(TOML)
}
