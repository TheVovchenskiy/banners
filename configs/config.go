package configs

import (
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

const (
	ServerPort  = 8080
	SwaggerPort = 8081
	LogsDir     = "./logs/"
	LogfileName = "server.log"
	LogLevel    = logrus.DebugLevel
)

type PgConfig struct {
	PgUser     string
	PgDBName   string
	PgPassword string
	PgPort     string
}

var (
	JwtKey   = []byte(os.Getenv("SECRET_KEY"))
	AdminKey = os.Getenv("ADMIN_REGISTRATION_KEY")

	PgConfigData = PgConfig{
		PgUser:     os.Getenv("PG_USER"),
		PgDBName:   os.Getenv("PG_DBNAME"),
		PgPassword: os.Getenv("PG_PASSWORD"),
		PgPort:     os.Getenv("PG_PORT"),
	}

	AccesTokenExpiresAt = time.Now().Add(24 * time.Hour).Unix()

	// RefreshJwtKey = []byte(os.Getenv("REFRESH_SECRET_KEY"))  TODO: add refresh jwt
)
