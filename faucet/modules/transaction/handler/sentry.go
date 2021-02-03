package handler

import (
	"context"
	_"fmt"
	"sync"
	"time"

	"github.com/hoangnguyen-1312/faucet/domain"
	"github.com/hoangnguyen-1312/faucet/logger"
)

type Sentry struct {
	tx 	domain.IUsecaseTransactionLog
	mu 	sync.Mutex
}

func NewDefaultSentry(tx domain.IUsecaseTransactionLog){
	sentry := &Sentry{
		tx: tx,
	}
	go sentry.processTransaction()
}
func (sentry *Sentry) processTransaction() {
	go func() {
		ticker := time.Tick(40 * time.Second)
		for range ticker {
			ctx := context.TODO()
			sentry.tryCreateTransaction(ctx)
			sentry.tryFinalizingTransaction(ctx)
		}
	}()
	go func() {
		ticker := time.Tick(1 * time.Hour)
		for range ticker {
			ctx := context.TODO()
			sentry.tryDeleteAllCurrentTransaction(ctx)
		}
	}()
	
}

func (sentry *Sentry) tryCreateTransaction(ctx context.Context) {
	sentry.mu.Lock()
	defer sentry.mu.Unlock()
	newTransactions, err := sentry.tx.GetTransactionLogsByStatus(ctx, 0)
	if (len(newTransactions) == 0){
		return
	}
	if err != nil {
		return
	}
	for _, newTransaction := range newTransactions {
		logger.Log.Info().Msg("TRY OPENING TRANSACTION")
		ctx := context.TODO()
		err := sentry.tx.OpeningTransactionLog(ctx, newTransaction)
		if err != nil {
			continue
		}
		logger.Log.Info().Msg("TRY OPENING SUCCESS")
	}
}

func (sentry *Sentry) tryFinalizingTransaction(ctx context.Context) {
	sentry.mu.Lock()
	defer sentry.mu.Unlock()
	pendingTransactions, err := sentry.tx.GetTransactionLogsByStatus(ctx, 1)
	if (len(pendingTransactions) == 0){
		return
	}
	if err != nil {
		return
	}
	for _, pendingTransaction := range pendingTransactions {
		logger.Log.Info().Msg("TRY FINALIZING TRANSACTION")
		ctx := context.TODO()
		err, notification := sentry.tx.FinalizingTransactionLog(ctx, pendingTransaction)
		if err != nil || notification == "Failure" {
			continue
		}
		logger.Log.Info().Msg("TRY FINALIZING SUCCESS")
	}
}

func (sentry *Sentry) tryDeleteAllCurrentTransaction(ctx context.Context) {
	sentry.mu.Lock()
	defer sentry.mu.Unlock()
	allSuccessfulTransaction, err := sentry.tx.GetTransactionLogsByStatus(ctx, 2)
	if (len(allSuccessfulTransaction) == 0){
		return
	}
	if err != nil {
		return
	}
	for _, successfulTransaction := range allSuccessfulTransaction {
		logger.Log.Info().Msg("TRY FETCH ALL SUCCESSFUL TRANSACTIONS")
		ctx := context.TODO()
		err, notification := sentry.tx.DeletingSuccessfulTransactionLog(ctx, successfulTransaction)
		if err != nil || notification == "Failure" {
			continue
		}
		logger.Log.Info().Msg("TRY CLEAN SUCCESSFULLY")
	}
}
