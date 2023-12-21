package bizutils

import (
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func NoPermissionError() *errors.Error {
	return pb.ErrorErrorReasonForbidden("no permission")
}
