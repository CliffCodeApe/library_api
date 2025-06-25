package config

import (
	"fmt"
	"library_api/utils"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfiguration struct {
	Port                 int
	IsProduction         bool
	DbURI                string
	AccessTokenLifeTime  uint
	RefreshTokenLifeTime uint
	BaseURL              string
	StoragePath          string
}

var config *AppConfiguration

func Get() *AppConfiguration {
	return config
}

func Load() {
	log.Println("load config from environment")
	// Get the configuration
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading environment variables, try to get from environtment OS")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	// set default value port if env doesn't have PORT config
	if err != nil {
		port = 8080
	}

	isProduction := utils.SafeCompareString(os.Getenv("IS_PRODUCTION"), "true")

	ATLifeTime, err := strconv.Atoi(os.Getenv("JWT_LIFETIME"))

	if err != nil {
		ATLifeTime = 7200
	}

	RFLifeTime, err := strconv.Atoi(os.Getenv("REFRESH_JWT_LIFETIME"))

	if err != nil {
		RFLifeTime = 172800
	}

	baseURL := os.Getenv("BASE_URL")

	storagePath := os.Getenv("STORAGE_PATH")

	// set global variable config
	config = &AppConfiguration{
		Port:                 port,
		IsProduction:         isProduction,
		DbURI:                loadDatabaseConfig(),
		AccessTokenLifeTime:  uint(ATLifeTime),
		RefreshTokenLifeTime: uint(RFLifeTime),
		BaseURL:              baseURL,
		StoragePath:          storagePath,
	}
}

func loadDatabaseConfig() string {
	user := getFromEnv("DB_USER")
	pass := getFromEnv("DB_PASS")
	name := getFromEnv("DB_NAME")
	host := getFromEnv("DB_HOST")
	port := getFromEnv("DB_PORT")
	sslMode := getFromEnv("DB_SSL_MODE") == "true"
	timeZone := getFromEnv("DB_TIME_ZONE")

	mode := "enable"

	if !sslMode {
		mode = "disable"
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, pass, name, port, mode, timeZone)
}

func getFromEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable is not set", value)
	}

	return value
}
