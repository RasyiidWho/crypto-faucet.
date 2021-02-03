package config
import (
	"github.com/hoangnguyen-1312/faucet/logger"
	"strings"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseConnection DatabaseConnection
	AppPort            string
	CORS               []string
	Mode               string
	IncHost            string
	IncPort            string
	IncWs              string
	IncHttps           string
}

type DatabaseConnection struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
}

func NewDatabaseConnection(host string, port string, user string, dbname string, password string) DatabaseConnection {
	return DatabaseConnection{
		Host:     host,
		Port:     port,
		User:     user,
		Dbname:   dbname,
		Password: password,
	}
}

func Init() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	logLevel := os.Getenv("LOGLEVEL")
	logger.InitLog(logLevel)
	logger.Log.Info().Str("level", logLevel).Msg("Initialized Logger Successfully")

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}
	logger.Log.Info().Str("mode", mode).Msg("Running Mode")

	appPort := os.Getenv("PORT")
	logger.Log.Info().Str("port", appPort).Msg("Application Port")

	var cors []string
	tempCORS := os.Getenv("CORS")
	if tempCORS == "" {
		cors = []string{"*"}
	} else {
		cors = strings.Split(tempCORS, ",")
		if len(cors) == 0 {
			cors = []string{"*"}
		}
	}
	logger.Log.Info().Strs("cors", cors).Msg("Cross Origin Resource Sharing")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_DATABASE")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbConn := NewDatabaseConnection(dbHost, dbPort, dbUser, dbName, dbPassword)
	logger.Log.Info().Str("host", dbHost).Str("port", dbPort).Str("user", dbUser).Str("password", dbPassword).Str("database", dbName).Msg("Database Connection Parameters")

	incHost := os.Getenv("INC_HOST")
	incPort := os.Getenv("INC_PORT")
	incWs := os.Getenv("INC_WS")
	incHttps := os.Getenv("INC_HTTPS")
	
	logger.Log.Info().
		Str("inc_host", incHost).
		Str("inc_port", incPort).
		Str("inc_ws", incWs).
		Str("inc_https", incHttps).
		Msg("Incognito Client Info")

	return Config{
		DatabaseConnection: dbConn,
		AppPort:            appPort,
		CORS:               cors,
		Mode:               mode,
		IncHost:            incHost,
		IncPort:            incPort,
		IncHttps:           incHttps,
		IncWs:              incWs,
	}
}
