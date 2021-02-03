package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen-1312/faucet/config"
	"github.com/hoangnguyen-1312/faucet/config/database"
	"github.com/hoangnguyen-1312/faucet/internal/incognito"
	"github.com/hoangnguyen-1312/faucet/modules/transaction/entity"
	"github.com/hoangnguyen-1312/faucet/modules/transaction/handler"
	"github.com/hoangnguyen-1312/faucet/modules/transaction/repository"
)
func main() {

	conf := config.Init()
	db := database.InitDatabase(conf.DatabaseConnection)
	defer db.Close()
	database.Migration(db)

	router := gin.Default()
	var routerConfig cors.Config
	routerConfig = cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(routerConfig))
	client := incognito.NewClient(conf.IncHost, conf.IncPort, conf.IncWs, conf.Mode, conf.IncHttps)
	routerGroupV1 := router.Group("/v1")

	repo := repository.NewLogTransactionRepository(db)
	entity := entity.NewLogTransactionUsecase(repo, *client)
	handler.NewHandlerTransaction(routerGroupV1, entity)
	handler.NewDefaultSentry(entity)
	router.Run(conf.AppPort)
}