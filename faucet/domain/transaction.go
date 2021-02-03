package domain

import (
	"context"
	"time"
)
type TransactionLog struct {
	ID        		uint64     	`gorm:"primary_key;auto_increment" json:"id"`
	CreateAt    	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	PaymentAddress  string    	`gorm:"column:address;not null"`
	TxHash			string    	`gorm:"column:hash;not null"`
	Amount    		uint64    	`gorm:"column:amount;not null"`
	Status    		int       	`gorm:"column:status;not null"`
}
func (tx *TransactionLog) NewTransactionLog() interface{} {
	return &TransactionLog{
		ID:			tx.ID,
		CreateAt:	tx.CreateAt,
		PaymentAddress:	tx.PaymentAddress,
		TxHash:		tx.TxHash,
		Amount:		tx.Amount,
		Status: 	tx.Status,
	}
}
type TxLogs []TransactionLog

func (txLogs TxLogs) AllTransactionLogs() []interface{} {
	setTransactionLogs := make([]interface{}, len(txLogs))
	for indexOfTransaction, txLog := range txLogs {
		setTransactionLogs[indexOfTransaction] = txLog.NewTransactionLog()
	}
	return setTransactionLogs
}

type IUsecaseTransactionLog interface {
	CreateTransactionLog(ctx context.Context, paymentAddress string) (TransactionLog, *Error)
	FinalizingTransactionLog(ctx context.Context, tx TransactionLog) (*Error, string)
	OpeningTransactionLog(ctx context.Context, tx TransactionLog) (*Error)
	DeletingSuccessfulTransactionLog(ctx context.Context, tx TransactionLog) (*Error, string)
	GetAllTransaction(ctx context.Context)(TxLogs, *Error)
	GetPRVBalance(ctx context.Context)(uint64, *Error)
	GetTransactionLogsByStatus(ctx context.Context, status int)(TxLogs, *Error)
	DeleteTransactionLogByTxHash(ctx context.Context, txHash string) *Error
	DeleteTransactionLogByPaymentAddress(ctx context.Context, paymentAddress string) *Error
	GetTotalTransactionLogsByAddress(ctx context.Context, paymentAddress string) (int64, *Error)
	CountAllTransactions(ctx context.Context) (int64, *Error)
}

type IRepositoryTransactionLog interface {
	CreateTransactionLog(ctx context.Context, tx TransactionLog) (TransactionLog, error)
	GetAllTransaction(ctx context.Context)(TxLogs, error)
	GetTransactionLogsByStatus(ctx context.Context, status int) (TxLogs, error)
	GetTransactionLogsByAddress(ctx context.Context, paymentAddress string) (TxLogs, error)
	UpdateTransactionLog(ctx context.Context, ID uint64, status int, txHash string) error
	DeleteTransactionLogByTxHash(ctx context.Context, txHash string) error
	DeleteTransactionLogByPaymentAddress(ctx context.Context, paymentAddress string) error
	CountTotalTransactionLogsByAddress(ctx context.Context, paymentAddress string) (int64, error)
	GetTransactionLogByID(ctx context.Context, ID uint64) (TransactionLog, error)
	CountAllTransactions(ctx context.Context) (int64, error)
}