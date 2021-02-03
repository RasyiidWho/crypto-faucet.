package logger
import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLog(logLevel string) {
	switch logLevel {
	case LogLevelDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("|  %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s=", i)
	}
	output.FormatFieldValue = func(value interface{}) string {
		return fmt.Sprintf("%s", value)
	}
	Log = zerolog.New(output).With().Timestamp().Logger()
}