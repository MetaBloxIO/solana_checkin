// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package check

import (
	"context"

	"check/api/check/v1"
)

type ICheckV1 interface {
	Check(ctx context.Context, req *v1.CheckReq) (res *v1.CheckRes, err error)
}
