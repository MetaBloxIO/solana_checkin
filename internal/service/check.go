// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICheck interface {
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
		NewTransaction(ctx context.Context, wallet string, data string) (txBytes string, err error)
	}
)

var (
	localCheck ICheck
)

func Check() ICheck {
	if localCheck == nil {
		panic("implement not found for interface ICheck, forgot register?")
	}
	return localCheck
}

func RegisterCheck(i ICheck) {
	localCheck = i
}
