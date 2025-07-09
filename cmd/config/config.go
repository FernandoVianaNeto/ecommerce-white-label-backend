package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApplicationCfg *ApplicationConfig
	MongoCfg       *MongoConfig
	GoogleAuthCfg  *GoogleAuthConfig
	SendGridCfg    *SendGridConfig
	MinIoCfg       *MinIOConfig
	NatsCfg        *NatsConfig
)

const (
	AppName     = "ecommerce-white-label-backend"
	AppVersion  = "1.0.0"
	Development = "development"
	Staging     = "stage"
	Production  = "production"
)

type ApplicationConfig struct {
	Env         string
	AppVersion  string
	AppPort     int
	JwtSecret   string
	Environment string
}

type GoogleAuthConfig struct {
	ClientId string
}

type SendGridConfig struct {
	ApiKey string
}

type MinIOConfig struct {
	Host                   string
	User                   string
	Password               string
	ProfileBucket          string
	ProductBucket          string
	PresignedURLExpiration string
}

type NatsConfig struct {
	Host         string
	User         string
	Password     string
	UserTopic    string
	ProductTopic string
}

type MongoConfig struct {
	UserCollection              string
	ProductCollection           string
	ResetPasswordCodeCollection string
	Dsn                         string
	Database                    string
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func InitializeConfigs() {
	initialize()
	initializeApplicationConfigs()
	initializeMongoConfigs()
	initializeGoogleAuthConfigs()
	initializeSendGridConfigs()
	initializeMinIOConfigs()
	initializeNatsConfigs()
}

func getEnv(key string, defaultVal string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func initializeApplicationConfigs() {
	if ApplicationCfg == nil {
		ApplicationCfg = &ApplicationConfig{
			Env:        getEnv("APP_ENV", "local"),
			AppVersion: AppVersion,
			AppPort:    getEnvAsInt("APP_PORT", 3001),
			JwtSecret:  getEnv("JWT_SECRET", "hash_aleatorio"),
		}
	}
}

func initializeMongoConfigs() {
	if MongoCfg == nil {
		MongoCfg = &MongoConfig{
			UserCollection:              getEnv("MONGO_USER_COLLECTION", "user"),
			ProductCollection:           getEnv("MONGO_PRODUCT_COLLECTION", "product"),
			ResetPasswordCodeCollection: getEnv("MONGO_RESET_PASSWORD_CODE_COLLECTION", "reset_password_code"),
			Dsn:                         getEnv("MONGO_DSN", "mongodb://localhost:27017"),
			Database:                    getEnv("MONGO_DB", "ecommerce_white_label_backend"),
		}
	}
}

func initializeGoogleAuthConfigs() {
	if GoogleAuthCfg == nil {
		GoogleAuthCfg = &GoogleAuthConfig{
			ClientId: getEnv("GOOGLE_CLIENT_ID", "your-google-client-id"),
		}
	}
}

func initializeSendGridConfigs() {
	if SendGridCfg == nil {
		SendGridCfg = &SendGridConfig{
			ApiKey: getEnv("SEND_GRID_API_KEY", "your-sendgrid-api-key"),
		}
	}
}

func initializeMinIOConfigs() {
	if MinIoCfg == nil {
		MinIoCfg = &MinIOConfig{
			Host:                   getEnv("MINIO_HOST", "localhost:9000"),
			User:                   getEnv("MINIO_USER", "root"),
			Password:               getEnv("MINIO_PASSWORD", "password"),
			ProfileBucket:          getEnv("MINIO_PROFILE_BUCKET", "profile"),
			ProductBucket:          getEnv("MINIO_PRODUCT_BUCKET", "product"),
			PresignedURLExpiration: getEnv("PRESIGNED_URL_EXPIRATION", "60"), // Default to 60 minutes
		}
	}
}

func initializeNatsConfigs() {
	if NatsCfg == nil {
		NatsCfg = &NatsConfig{
			Host:         getEnv("NATS_HOST", "nats://localhost:4222"),
			User:         getEnv("NATS_USER", "root"),
			Password:     getEnv("NATS_PASSWORD", "password"),
			UserTopic:    getEnv("NATS_USER_TOPIC", "user.events"),
			ProductTopic: getEnv("NATS_PRODUCT_TOPIC", "product.events"),
		}
	}
}
