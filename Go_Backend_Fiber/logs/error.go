package logs

import (
	// "fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
)
var buildInfo *debug.BuildInfo
var logger zerolog.Logger

func init() {

	
	buildInfo, _ = debug.ReadBuildInfo()

	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
	Level(zerolog.TraceLevel).
	With().
	Timestamp().
	Caller().
	Int("pid", os.Getpid()).
	Str("go_version", buildInfo.GoVersion).
	Logger()
}


func Error(s string,err error) {
    // log.Info().Msg("Info message")
	if err != nil {
		// fmt.Println("❌🔥",s,err)
		logger.Error().Msg("❌🔥Error message :"+s+err.Error())
	}
}