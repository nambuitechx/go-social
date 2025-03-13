package configs

import (
	"os"
	"log"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Settings struct {
	ServerHost 			string
	ServerPort 			int

	DatabaseHost 		string
	DatabasePort 		int
	DatabaseName 		string
	DatabaseUser 		string
	DatabasePassword	string

	RedisHost			string
	RedisPort			int
	RedisDB				int
	RedisUser			string
	RedisPassword		string
	RedisProtocol		int

	KafkaHost			string
	KafkaPort			int
	KafkaProtocol		string

	JwtSecretKey		string
}

func NewSettings() *Settings {
	// Load env
	dotenvErr := godotenv.Load()
	
	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}

	settings := &Settings{}

	// Server
	serverHost, ok := os.LookupEnv("SERVER_HOST")
	if ok {
		settings.ServerHost = serverHost
	} else {
		settings.ServerHost = "localhost"
	}

	serverPort, ok := os.LookupEnv("SERVER_PORT")
	if ok {
		port, err := strconv.ParseInt(serverPort, 10, 64)
		if err != nil {
			log.Fatal("Invalid server port")
		}
		settings.ServerPort = int(port)
	} else {
		settings.ServerPort = 8000
	}

	// Database
	databaseHost, ok := os.LookupEnv("DATABASE_HOST")
	if ok {
		settings.DatabaseHost = databaseHost
	} else {
		settings.DatabaseHost = "localhost"
	}

	databasePort, ok := os.LookupEnv("DATABASE_PORT")
	if ok {
		port, err := strconv.ParseInt(databasePort, 10, 64)
		if err != nil {
			log.Fatal("Invalid database port")
		}
		settings.DatabasePort = int(port)
	} else {
		settings.DatabasePort = 5432
	}

	databaseName, ok := os.LookupEnv("DATABASE_NAME")
	if ok {
		settings.DatabaseName = databaseName
	} else {
		settings.DatabaseName = "go_social"
	}

	databaseUser, ok := os.LookupEnv("DATABASE_USER")
	if ok {
		settings.DatabaseUser = databaseUser
	} else {
		settings.DatabaseUser = "admin"
	}

	databasePassword, ok := os.LookupEnv("DATABASE_PASSWORD")
	if ok {
		settings.DatabasePassword = databasePassword
	} else {
		settings.DatabasePassword = "admin"
	}

	// Redis
	redisHost, ok := os.LookupEnv("REDIS_HOST")
	if ok {
		settings.RedisHost = redisHost
	} else {
		settings.RedisHost = "localhost"
	}

	redisPort, ok := os.LookupEnv("REDIS_PORT")
	if ok {
		port, err := strconv.ParseInt(redisPort, 10, 64)
		if err != nil {
			log.Fatal("Invalid redis port")
		}
		settings.RedisPort = int(port)
	} else {
		settings.RedisPort = 6379
	}

	redisDB, ok := os.LookupEnv("REDIS_DB")
	if ok {
		port, err := strconv.ParseInt(redisDB, 10, 64)
		if err != nil {
			log.Fatal("Invalid redis db")
		}
		settings.RedisDB = int(port)
	} else {
		settings.RedisDB = 0
	}

	redisUser, ok := os.LookupEnv("REDIS_USER")
	if ok {
		settings.RedisUser = redisUser
	} else {
		settings.RedisUser = "admin"
	}

	redisPassword, ok := os.LookupEnv("REDIS_PASSWORD")
	if ok {
		settings.RedisPassword = redisPassword
	} else {
		settings.RedisPassword = "admin"
	}

	redisProtocol, ok := os.LookupEnv("REDIS_PROTOCOL")
	if ok {
		port, err := strconv.ParseInt(redisProtocol, 10, 64)
		if err != nil {
			log.Fatal("Invalid redis protocol")
		}
		settings.RedisProtocol = int(port)
	} else {
		settings.RedisProtocol = 3
	}

	// Kafka
	kafkaHost, ok := os.LookupEnv("KAFKA_HOST")
	if ok {
		settings.KafkaHost = kafkaHost
	} else {
		settings.KafkaHost = "localhost"
	}

	kafkaPort, ok := os.LookupEnv("KAFKA_PORT")
	if ok {
		port, err := strconv.ParseInt(kafkaPort, 10, 64)
		if err != nil {
			log.Fatal("Invalid kafka port")
		}
		settings.KafkaPort = int(port)
	} else {
		settings.KafkaPort = 9092
	}

	kafkaProtocol, ok := os.LookupEnv("KAFKA_PROTOCOL")
	if ok {
		settings.KafkaProtocol = kafkaProtocol
	} else {
		settings.KafkaProtocol = "tcp"
	}

	return settings
}

func GetJwtSecret() string {
	jwtSecretKey, ok := os.LookupEnv("JWT_SECRET_KEY")

	if ok {
		return jwtSecretKey
	} else {
		return "my-secret"
	}
}
