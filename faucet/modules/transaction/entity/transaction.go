package entity

import (
	"context"
	_ "fmt"
	"log"
	"os"

	"github.com/hoangnguyen-1312/faucet/domain"
	"github.com/hoangnguyen-1312/faucet/internal/incognito"
	"github.com/joho/godotenv"
)

const SUCCESS = "success"
const FAILURE = "failure"

type LogTransactionUsecase struct {
	r domain.IRepositoryTransactionLog
	client incognito.Client
	
}

func NewLogTransactionUsecase(repo domain.IRepositoryTransactionLog, client incognito.Client) domain.IUsecaseTransactionLog {
	return &LogTransactionUsecase{
			r: repo,
			client: client,
		}
}

func (l *LogTransactionUsecase) CreateTransactionLog(ctx context.Context, paymentAddress string) (res domain.TransactionLog, err *domain.Error) {
	tx := domain.TransactionLog{
		PaymentAddress: paymentAddress,
		TxHash: "",
		Amount: 2,
		Status: 0,
	}
	totalTransaction, err4 := l.r.CountAllTransactions(ctx)
	if err4 != nil {
		return domain.TransactionLog{}, domain.ErrInvalidTransaction.NewError(err4)
	}
	if totalTransaction > 10 {
		return domain.TransactionLog{}, nil
	}

	countNumberTransactionByAddress, err3 := l.r.CountTotalTransactionLogsByAddress(ctx, paymentAddress)
	if err3 != nil {
		return domain.TransactionLog{}, domain.ErrInvalidTransaction.NewError(err3)
	}
	if countNumberTransactionByAddress > 1 {
		return domain.TransactionLog{}, nil
	}

	res, err2 := l.r.CreateTransactionLog(ctx, tx)
	if err2 != nil {
		return domain.TransactionLog{}, domain.ErrTransactionLogRepository.NewError(err2)
	}
	return tx, nil
}

func (l *LogTransactionUsecase) OpeningTransactionLog(ctx context.Context, tx domain.TransactionLog) (err *domain.Error) {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
	transactionLog, err1 := l.r.GetTransactionLogByID(ctx, tx.ID)
	if err1 != nil {
		return domain.ErrTransactionLogRepository.NewError(err1)
	}
	receiver := make(map[string]uint64)
	receiver[transactionLog.PaymentAddress] = 2
	params := []interface{}{
		os.Getenv("PRIVATE_KEY"),
		receiver,
		15,
		0,
	}
	txLogResult, err2 := l.client.CreateAndSendTransaction(params)
	if err2 != nil {
		return domain.ErrIncognitoService.NewError(err2)
	}
	err3 := l.r.UpdateTransactionLog(ctx, tx.ID, 1, txLogResult.TxID)
	if err3 != nil {
		return domain.ErrTransactionLogRepository.NewError(err3)
	}
	return nil
}

func (l *LogTransactionUsecase) FinalizingTransactionLog(ctx context.Context, tx domain.TransactionLog) (err *domain.Error, notification string) {
	transactionInfo, err2 := l.client.GetTransactionByHash(tx.TxHash)
	if err2 != nil {
		return domain.ErrIncognitoService.NewError(err2), FAILURE
	}
	if transactionInfo.IsInBlock == true {
		err3 := l.r.UpdateTransactionLog(ctx, tx.ID, 2, tx.TxHash)
		if err3 != nil {
			return domain.ErrTransactionLogRepository.NewError(err3), FAILURE
		}
		return nil, SUCCESS
	}
    return nil, FAILURE
}

func (l *LogTransactionUsecase) DeletingSuccessfulTransactionLog(ctx context.Context, tx domain.TransactionLog) (err *domain.Error, notification string) {

	err1 := l.r.DeleteTransactionLogByTxHash(ctx, tx.TxHash)
	if err1 != nil {
		return domain.ErrTransactionLogRepository.NewError(err1), FAILURE
	}
	return nil, SUCCESS
}


func (l *LogTransactionUsecase) GetAllTransaction(ctx context.Context) (res domain.TxLogs, err *domain.Error) {
	res, err1 := l.r.GetAllTransaction(ctx)
	if err1 != nil {
		return nil, domain.ErrTransactionLogRepository.NewError(err1)
	}
	return res, nil
}

func (l *LogTransactionUsecase) GetPRVBalance(xtx context.Context) (res uint64, err *domain.Error) {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
	res, err1 := l.client.GetBalanceByPrivateKey(os.Getenv("PRIVATE_KEY"))
	if err1 != nil {
		return 0, domain.ErrIncognitoService.NewError(err1)
	}
	return res, nil
}

func (l *LogTransactionUsecase) GetTransactionLogsByStatus(ctx context.Context, status int)(domain.TxLogs, *domain.Error) {
	var listTransactionLogs domain.TxLogs
	transactions, err := l.r.GetTransactionLogsByStatus(ctx, status)
	if err != nil {
		return nil, domain.ErrTransactionLog.NewError(err)
	}
	for _, transaction := range (transactions) {
		listTransactionLogs = append(listTransactionLogs, transaction)
	}
	return listTransactionLogs, nil

}
func (l *LogTransactionUsecase)	DeleteTransactionLogByTxHash(ctx context.Context, txHash string) *domain.Error {
	err := l.r.DeleteTransactionLogByTxHash(ctx, txHash)
	if err != nil {
		return domain.ErrTransactionLogRepository.NewError(err)
	}
	return nil
}

func (l *LogTransactionUsecase)	DeleteTransactionLogByPaymentAddress(ctx context.Context, paymentAddress string) *domain.Error {
	err := l.r.DeleteTransactionLogByPaymentAddress(ctx, paymentAddress)
	if err != nil {
		return domain.ErrTransactionLogRepository.NewError(err)
	}
	return nil
}

func (l *LogTransactionUsecase)	GetTotalTransactionLogsByAddress(ctx context.Context, paymentAddress string) (int64, *domain.Error) {
	totalTxs, err := l.r.CountTotalTransactionLogsByAddress(ctx, paymentAddress)
	if err != nil {
		return 0, domain.ErrTransactionLogRepository.NewError(err)
	}
	return totalTxs, nil
}

func (l *LogTransactionUsecase) CountAllTransactions(ctx context.Context) (int64, *domain.Error) {
	totalTxs, err := l.r.CountAllTransactions(ctx)
	if err != nil {
		return 0, domain.ErrTransactionLogRepository.NewError(err)
	}
	return totalTxs, nil
}


