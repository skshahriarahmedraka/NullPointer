package config

import (
	// "fmt"
	"os"

	"app/logs"
	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	box := packr.New("mybox", "../")
	s, err := box.FindString(".env")
	logs.Error("Error in loading env file", err)
	// fmt.Println("🚀", s)
	myEnv, err := godotenv.Unmarshal(s)
	logs.Error("Error in unmarshalling env file", err)

	for i, j := range myEnv {
		os.Setenv(i, j)
	}

}
