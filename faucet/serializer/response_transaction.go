package serializer

import (
	"github.com/hoangnguyen-1312/faucet/domain"
	"time"
)

type Transactions struct {
	TransactionLogDetails 	[]TransactionLogDetail 	`json:"transactionLogDetails" `
}

type TransactionLogDetail struct {
	TransactionID      		uint64         `json:"transactionID"`
	PaymentAddress       	string         `json:"paymentAddress"`
	Hash     				string         `json:"hash"`
	CreatedDate 			time.Time      `json:"createdDate"`
	Status					int		   	   `json:"status"`
}

func NewResponseTransactions(transactions domain.TxLogs) Transactions {
	responseTransactions := Transactions{}
	responseTransactions.TransactionLogDetails = []TransactionLogDetail{}
	for _, tx := range transactions {
		responseTransactions.TransactionLogDetails = append(responseTransactions.TransactionLogDetails, NewResponseTransactionDetail(&tx))
	}
	return responseTransactions
}


func NewResponseTransactionDetail(tx *domain.TransactionLog) TransactionLogDetail {
	responseTransactionInfo := TransactionLogDetail{}
	responseTransactionInfo.PaymentAddress = tx.PaymentAddress
	responseTransactionInfo.CreatedDate = tx.CreateAt
	responseTransactionInfo.TransactionID = tx.ID
	responseTransactionInfo.Hash = tx.TxHash
	responseTransactionInfo.Status = tx.Status
	return responseTransactionInfo
}