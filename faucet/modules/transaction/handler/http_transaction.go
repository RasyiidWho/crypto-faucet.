package handler

import (
	_ "fmt"
	"math/rand"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen-1312/faucet/domain"
	"github.com/hoangnguyen-1312/faucet/logger"
	"github.com/hoangnguyen-1312/faucet/serializer"
)

type HandlerTransaction struct {
	tx domain.IUsecaseTransactionLog
}
func NewHandlerTransaction(r *gin.RouterGroup, tx domain.IUsecaseTransactionLog) {
	handler := &HandlerTransaction{
		tx: tx,
	}
	r.GET("/donate/:paymentAddress", handler.CreateAndSendTransaction)
	r.GET("/all", handler.GetAllTransactions)
	r.GET("/balance", handler.GetPRVBalance)
	r.GET("/info", handler.GetInfoNumberOfTransactionLogByAddress)

}


func (h *HandlerTransaction) CreateAndSendTransaction (c *gin.Context) {
	requestID := rand.Int63()
	paymentAddress := c.Param("paymentAddress")
	logger.Log.Info().
		Int64(logger.LogRequestID, requestID).
		Str(logger.LogMethod, "GET").
		Str(logger.LogURLPath, "/donate/").
		Str(logger.LogAddress, paymentAddress).
		Msg("ACCESS")
	transaction, err := h.tx.CreateTransactionLog(c, paymentAddress)
	if err != nil {
		logger.Log.Err(err.Err).
			Int64(logger.LogRequestID, requestID).
			Str(logger.LogMethod, "GET").
			Str(logger.LogURLPath, "/donate/").
			Str(logger.LogAddress, paymentAddress).
			Msg("FAILURE")
		c.JSON(http.StatusBadRequest, serializer.NewResponseError(err))
		return
	}

	if transaction.PaymentAddress == "" {
		logger.Log.Info().
				Int64(logger.LogRequestID, requestID).
				Str(logger.LogMethod, "GET").
				Str(logger.LogURLPath, "/donate/").
				Str(logger.LogAddress, paymentAddress).
				Msg("FAILURE")
		c.JSON(http.StatusBadRequest, serializer.NewResponseTransactionDetail(&transaction))
		return
	}

	logger.Log.Info().
		Int64(logger.LogRequestID, requestID).
		Str(logger.LogMethod, "GET").
		Str(logger.LogURLPath, "/donate/").
		Str(logger.LogAddress, paymentAddress).
		Msg("SUCCESS")
	c.JSON(http.StatusOK, serializer.NewResponseSuccess(serializer.NewResponseTransactionDetail(&transaction)))
	return	
}

func (h *HandlerTransaction) GetAllTransactions(c *gin.Context) {
	requestID := rand.Int63()
	logger.Log.Info().
		Int64(logger.LogRequestID, requestID).
		Str(logger.LogMethod, "GET").
		Str(logger.LogURLPath, "/all/").
		Msg("ACCESS")
	
	transactions := domain.TxLogs{}
	transactions, err := h.tx.GetAllTransaction(c)
	if err != nil {
		logger.Log.Err(err.Err).
			Int64(logger.LogRequestID, requestID).
			Str(logger.LogMethod, "GET").
			Str(logger.LogURLPath, "/all/").
			Msg("FAILURE")
		c.JSON(http.StatusBadRequest, serializer.NewResponseError(err))
		return
	}
	logger.Log.Info().
		Int64(logger.LogRequestID, requestID).
		Str(logger.LogMethod, "GET").
		Str(logger.LogURLPath, "/all/").
		Msg("SUCCESS")
	c.JSON(http.StatusOK, serializer.NewResponseSuccess(serializer.NewResponseTransactions(transactions)))
	return
}

func (h *HandlerTransaction) GetPRVBalance(c *gin.Context) {
	requestID := rand.Int63()
	logger.Log.Info().
		Int64(logger.LogRequestID, requestID).
		Str(logger.LogMethod, "GET").
		Str(logger.LogURLPath, "/balance/").
		Msg("ACCESS")

	balance, err := h.tx.GetPRVBalance(c)
	if err != nil {
		logger.Log.Err(err.Err).
			Int64(logger.LogRequestID, requestID).
			Str(logger.LogMethod, "GET").
			Str(logger.LogURLPath, "/balance/").
			Msg("FAILURE")
		c.JSON(http.StatusBadRequest, serializer.NewResponseError(err))
		return
	}
	logger.Log.Info().
		Int64(logger.LogRequestID, requestID).
		Str(logger.LogMethod, "GET").
		Str(logger.LogURLPath, "/balance/").
		Msg("SUCCESS")
	c.JSON(http.StatusOK, balance)
	return
}

func (h *HandlerTransaction) GetInfoNumberOfTransactionLogByAddress (c *gin.Context) {
	requestID := rand.Int63()
	logger.Log.Info().
		Int64(logger.LogRequestID, requestID).
		Str(logger.LogMethod, "GET").
		Str(logger.LogURLPath, "/info/").
		Msg("ACCESS")
	
	transactions, err1 := h.tx.GetAllTransaction(c)
	if err1 != nil {
		logger.Log.Err(err1.Err).
			Int64(logger.LogRequestID, requestID).
			Str(logger.LogMethod, "GET").
			Str(logger.LogURLPath, "/info/").
			Msg("FAILURE")
		c.JSON(http.StatusBadRequest, serializer.NewResponseError(err1))
		return
	}
	infoNumberOfTransactionLogs := make(map[string]int64)
	for _, transaction := range transactions {
		count, err2 := h.tx.GetTotalTransactionLogsByAddress(c, transaction.PaymentAddress)
		_, found := infoNumberOfTransactionLogs[transaction.PaymentAddress]
		if found == false {
			infoNumberOfTransactionLogs[transaction.PaymentAddress] = count
		}
		if err2 != nil {
			logger.Log.Err(err2.Err).
				Int64(logger.LogRequestID, requestID).
				Str(logger.LogMethod, "GET").
				Str(logger.LogURLPath, "/info/").
				Msg("FAILURE")
			c.JSON(http.StatusBadRequest, serializer.NewResponseError(err2))
			return
		}
	}
	var err4 *domain.Error
	infoNumberOfTransactionLogs["all"], err4 = h.tx.CountAllTransactions(c)
	if err4 != nil {
		logger.Log.Err(err4.Err).
			Int64(logger.LogRequestID, requestID).
			Str(logger.LogMethod, "GET").
			Str(logger.LogURLPath, "/info/").
			Msg("FAILURE")
		c.JSON(http.StatusBadRequest, serializer.NewResponseError(err4))
		return
	}
	
	logger.Log.Info().
		Int64(logger.LogRequestID, requestID).
		Str(logger.LogMethod, "GET").
		Str(logger.LogURLPath, "/info/").
		Msg("SUCCESS")
	c.JSON(http.StatusOK, infoNumberOfTransactionLogs)
	return	
}