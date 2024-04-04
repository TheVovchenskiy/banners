package configs

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	ServerPort  = 8080
	SwaggerPort = 8081
	LogsDir     = "./logs/"
	LogfileName = "server.log"
	LogLevel    = logrus.DebugLevel
)

var (
	JwtKey = []byte(os.Getenv("SECRET_KEY"))
	// RefreshJwtKey = []byte(os.Getenv("REFRESH_SECRET_KEY"))  TODO: add refresh jwt
)
