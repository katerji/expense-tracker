package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

type env struct {
	jWTToken        string
	jWTRefreshToken string
	dbHost          string
	dbPassword      string
	dbUser          string
	dbPort          string
	dbName          string
	redisURL        string
	webServerPort   string
	openAISecret    string
}

func newEnv() *env {
	InitEnv()
	return &env{
		jWTToken:        os.Getenv("JWT_SECRET"),
		jWTRefreshToken: os.Getenv("JWT_REFRESH_SECRET"),
		dbHost:          os.Getenv("DB_HOST"),
		dbPassword:      os.Getenv("DB_PASSWORD"),
		dbUser:          os.Getenv("DB_USERNAME"),
		dbPort:          os.Getenv("DB_PORT"),
		dbName:          os.Getenv("DB_DATABASE"),
		redisURL:        os.Getenv("UPSTASH_REDIS_URL"),
		webServerPort:   os.Getenv("WEB_SERVER_PORT"),
		openAISecret:    os.Getenv("OPEN_AI_SECRET"),
	}
}

func JWTToken() string {
	return getInstance().jWTToken
}

func JWTRefreshToken() string {
	return getInstance().jWTRefreshToken
}

func DbHost() string {
	return getInstance().dbHost
}

func DbPassword() string {
	return getInstance().dbPassword
}

func DbUser() string {
	return getInstance().dbUser
}

func DbPort() string {
	return getInstance().dbPort
}

func DbName() string {
	return getInstance().dbName
}

func RedisURL() string {
	return getInstance().redisURL
}

func WebServerPort() string {
	return getInstance().webServerPort
}

func OpenAISecret() string {
	return getInstance().openAISecret
}

var instance *env

func getInstance() *env {
	if instance == nil {
		instance = newEnv()
	}
	return instance
}

func InitEnv() {
	const projectDirName = "expense-tracker"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
