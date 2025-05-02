package tuihub

import (
	"github.com/invopop/jsonschema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ReflectJSONSchema(v interface{}) (string, error) {
	r := new(jsonschema.Reflector)
	r.ExpandedStruct = true
	r.DoNotReference = true
	j, err := r.Reflect(v).MarshalJSON()
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func MustReflectJSONSchema(v interface{}) string {
	j, err := ReflectJSONSchema(v)
	if err != nil {
		panic(err)
	}
	return j
}

// isUnimplementedError checks if the error is a gRPC unimplemented error.
func isUnimplementedError(err error) bool {
	if err == nil {
		return false
	}
	st, ok := status.FromError(err)
	if !ok {
		return false
	}
	return st.Code() == codes.Unimplemented
}
