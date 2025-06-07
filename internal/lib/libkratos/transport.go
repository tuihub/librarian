package libkratos

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/metadata"
)

type Transport struct {
	K   transport.Kind
	E   string
	O   string
	Req transport.Header
	Res transport.Header
}

func (tr *Transport) Kind() transport.Kind {
	return tr.K
}

func (tr *Transport) Endpoint() string {
	return tr.E
}

func (tr *Transport) Operation() string {
	return tr.O
}

func (tr *Transport) RequestHeader() transport.Header {
	return tr.Req
}

func (tr *Transport) ReplyHeader() transport.Header {
	return tr.Res
}

type HTTPHeaderCarrier http.Header

func (hc HTTPHeaderCarrier) Get(key string) string { return http.Header(hc).Get(key) }

func (hc HTTPHeaderCarrier) Set(key string, value string) { http.Header(hc).Set(key, value) }

func (hc HTTPHeaderCarrier) Add(key string, value string) { http.Header(hc).Add(key, value) }

// Keys lists the keys stored in this carrier.
func (hc HTTPHeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range http.Header(hc) {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice value associated with the passed key.
func (hc HTTPHeaderCarrier) Values(key string) []string {
	return http.Header(hc).Values(key)
}

type GRPCHeaderCarrier metadata.MD

// Get returns the value associated with the passed key.
func (mc GRPCHeaderCarrier) Get(key string) string {
	vals := metadata.MD(mc).Get(key)
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}

// Set stores the key-value pair.
func (mc GRPCHeaderCarrier) Set(key string, value string) {
	metadata.MD(mc).Set(key, value)
}

// Add append value to key-values pair.
func (mc GRPCHeaderCarrier) Add(key string, value string) {
	metadata.MD(mc).Append(key, value)
}

// Keys lists the keys stored in this carrier.
func (mc GRPCHeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(mc))
	for k := range metadata.MD(mc) {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of values associated with the passed key.
func (mc GRPCHeaderCarrier) Values(key string) []string {
	return metadata.MD(mc).Get(key)
}
