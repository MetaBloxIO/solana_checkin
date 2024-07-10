package check

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"check/api/check/v1"
	"check/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Check(ctx context.Context, req *v1.CheckReq) (res *v1.CheckRes, err error) {
	m := map[string]interface{}{
		"did":       req.Did,
		"location":  req.Location,
		"timestamp": req.Timestamp.Format(time.RFC3339),
	}

	j, _ := json.Marshal(m)

	g.Log().Infof(ctx, "3W data: %s", string(j))

	hash := sha256.New()
	hash.Write(j)
	hashString := hex.EncodeToString(hash.Sum(nil))

	tx, err := service.Check().NewTransaction(ctx, req.UserAddress, hashString)
	if err != nil {
		return nil, err
	}

	return &v1.CheckRes{
		Tx: tx,
	}, nil
}
