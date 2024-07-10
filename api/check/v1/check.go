package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type CheckReq struct {
	g.Meta      `path:"/solana/3w/tx" method:"post" tag:"Solana" sm:"3W Data On-Chain Transaction"`
	UserAddress string    `json:"userAddress" v:"required" dc:"solana wallet address" example:"3PcmXDanBD2wohL4zVoafAKwNm5whM5sSgL9QPnuw6oc"`
	Did         string    `json:"did" v:"required" dc:"user did" example:"did:metablox:solana:3PcmXDanBD2wohL4zVoafAKwNm5whM5sSgL9QPnuw6oc"`
	Location    string    `json:"location" v:"required" dc:"latitude,longitude" example:"40.748817,-73.985428"`
	Timestamp   time.Time `json:"timestamp" v:"required" dc:"check time" example:"2006-01-02T15:04:05+07:00"`
}

type CheckRes struct {
	Tx string `json:"tx"`
}
