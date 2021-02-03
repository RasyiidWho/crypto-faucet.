package repository

import (
	"context"
	"errors"

	"github.com/hoangnguyen-1312/faucet/domain"
	"github.com/jinzhu/gorm"
)

type LogTransactionRepository struct {
	db *gorm.DB
}

func NewLogTransactionRepository(db *gorm.DB) domain.IRepositoryTransactionLog {
	return &LogTransactionRepository{db}
}

func (l *LogTransactionRepository) CreateTransactionLog( ctx context.Context, tl domain.TransactionLog) (domain.TransactionLog, error) {
	if err := l.db.Debug().Create(&tl).Error; err != nil {
		return domain.TransactionLog{}, err
	}
	return tl, nil
}

func (l *LogTransactionRepository) GetAllTransaction(ctx context.Context) (domain.TxLogs, error) {
	var txLogs domain.TxLogs
	err := l.db.Debug().Find(&txLogs).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("No Transactions")
	}
	return txLogs, nil
}

func (l *LogTransactionRepository) GetTransactionLogsByStatus(ctx context.Context, status int) (domain.TxLogs, error) {
	var txLogs domain.TxLogs
	err := l.db.Where("status = ?", status).Find(&txLogs).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("No Transactions")
	}
	return txLogs, nil
}
func (l *LogTransactionRepository) GetTransactionLogsByAddress(ctx context.Context, paymentAddress string) (domain.TxLogs, error) {
	var txLogs domain.TxLogs
	err := l.db.Where("address = ?", paymentAddress).Find(&txLogs).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("No Transactions")
	}
	return txLogs, nil
}

func (l *LogTransactionRepository) GetTransactionLogByID(ctx context.Context, ID uint64) (domain.TransactionLog, error) {
	var txLog domain.TransactionLog
	err := l.db.Where("id = ?", ID).First(&txLog).Error
	if err != nil {
		return domain.TransactionLog{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return domain.TransactionLog{}, errors.New("No Transactions")
	}
	return txLog, nil
}


func (l *LogTransactionRepository) UpdateTransactionLog(ctx context.Context, ID uint64, status int, txHash string) error {
	var transaction domain.TransactionLog
	err1 := l.db.Model(transaction).Where("id = ?", ID).Update("status", status).Error
	err2 := l.db.Model(transaction).Where("id = ?", ID).Update("hash", txHash).Error
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func (l *LogTransactionRepository) DeleteTransactionLogByTxHash(ctx context.Context, txHash string) error {
	var transaction domain.TransactionLog
	err := l.db.Where("hash = ?", txHash).Delete(transaction).Error
	if err != nil {
		return err
	}
	return nil

}
func (l *LogTransactionRepository) CountTotalTransactionLogsByAddress(ctx context.Context, paymentAddress string) (int64, error) {
	var result int64
	err := l.db.Table("transaction_logs").Where("address = ?", paymentAddress).Count(&result).Error
	if err != nil {
		return 0, err
	}
	return result, nil

}

func (l *LogTransactionRepository) CountAllTransactions(ctx context.Context) (int64, error) {
	var res int64
	err := l.db.Table("transaction_logs").Where("1=1").Count(&res).Error
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (l *LogTransactionRepository) 	DeleteTransactionLogByPaymentAddress(ctx context.Context, paymentAddress string) error {
	var transaction domain.TransactionLog
	err := l.db.Where("address = ?", paymentAddress).Delete(transaction).Error
	if err != nil {
		return err
	}
	return nil
}


