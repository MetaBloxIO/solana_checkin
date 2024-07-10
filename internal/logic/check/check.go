package check

import (
	"context"
	"encoding/base64"
	"time"

	"check/internal/service"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/pkg/bincode"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

type sCheck struct {
}

func init() {
	network := g.Cfg().MustGet(context.Background(), "solana.network").String()
	c = client.NewClient(network)

	adminAccountStr := g.Cfg().MustGet(context.Background(), "solana.adminAccount").String()
	adminAccount, _ = types.AccountFromBase58(adminAccountStr)

	programIdStr := g.Cfg().MustGet(context.Background(), "solana.programId").String()
	programId = common.PublicKeyFromString(programIdStr)

	configAccount, _, _ = common.FindProgramAddress([][]byte{[]byte("config")}, programId)

	service.RegisterCheck(&sCheck{})
}

var (
	c             *client.Client
	adminAccount  types.Account
	programId     common.PublicKey
	configAccount common.PublicKey
)

type BuildCheckInInstCheckinParam struct {
	User    common.PublicKey
	Admin   common.PublicKey
	BizHash string
}

func buildCheckInInst(param BuildCheckInInstCheckinParam) types.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction [8]byte
		BizHash     string
	}{
		Instruction: [8]byte{209, 253, 4, 217, 250, 241, 207, 50},
		BizHash:     param.BizHash,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		Accounts: []types.AccountMeta{
			{PubKey: param.Admin, IsSigner: true, IsWritable: false},
			{PubKey: param.User, IsSigner: true, IsWritable: true},
			{PubKey: configAccount, IsSigner: false, IsWritable: false},
		},
		ProgramID: programId,
		Data:      data,
	}
}

func buildCheckInTx(ctx context.Context, user string, hashData string) (tx types.Transaction, err error) {
	var (
		cacheTime = time.Second * 10
		cacheKey  = "GetLatestBlockhash"
		cacheFunc = func(ctx context.Context) (value interface{}, err error) {
			resp, err := c.GetLatestBlockhash(ctx)
			if err != nil {
				return types.Transaction{}, err
			}
			return resp.Blockhash, nil
		}
	)

	v, err := gcache.GetOrSetFuncLock(ctx, cacheKey, cacheFunc, cacheTime)
	if err != nil {
		return types.Transaction{}, err
	}
	blockHash := v.String()

	userPublicKey := common.PublicKeyFromString(user)

	tx, err = types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        adminAccount.PublicKey,
			RecentBlockhash: blockHash,
			Instructions: []types.Instruction{
				buildCheckInInst(BuildCheckInInstCheckinParam{
					Admin:   adminAccount.PublicKey,
					User:    userPublicKey,
					BizHash: hashData,
				}),
			},
		}),
		Signers: []types.Account{adminAccount},
	})

	return tx, err
}

// NewTransaction creates a new transaction for checking in.
//
// This method takes a context, wallet information, and data as inputs, constructs a check-in transaction,
// and returns the serialized transaction data in base64 encoding. If the transaction cannot be created or serialized,
// an error is returned.
//
// Parameters:
// - ctx: Context for controlling the lifecycle of the request, such as cancellation.
// - wallet: The wallet information used for the transaction, typically representing the sender or payer.
// - data: Additional data associated with the transaction, specific to the application's needs.
//
// Returns:
// - txBytes: The base64 encoded string of the serialized transaction.
// - err: An error object indicating failure, nil if the operation is successful.
func (s *sCheck) NewTransaction(ctx context.Context, wallet, data string) (txBytes string, err error) {
	// Build the check-in transaction.
	tx, err := buildCheckInTx(ctx, wallet, data)
	if err != nil {
		// Return immediately if an error occurs during transaction building.
		return "", err
	}

	// Serialize the transaction.
	serial, err := tx.Serialize()
	if err != nil {
		// Return immediately if an error occurs during serialization.
		return "", err
	}

	// Encode the serialized transaction data into base64 and return.
	return base64.StdEncoding.EncodeToString(serial), nil
}
