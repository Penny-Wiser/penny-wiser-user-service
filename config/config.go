package config

// GeneralConfig contains configurations required by application
// GeneraConfig is also required for dependency injection
type (
	MongoDbConfig struct {
		DatabaseName string
		DatabaseHost string
		DatabasePort string
		Username     string
		Password     string
	}

	LoggingConfig struct {
		LogLevel    string
		LogFilePath string
	}

	GeneralConfig struct {
		MongoDbConfig
		LoggingConfig
	}
)

func LoadConfig() *GeneralConfig {

	// Initialize config here
	mongoDbConfig := MongoDbConfig{
		"pennywiser",
		"localhost",
		"27017",
		"admin",
		"password123",
	}

	loggingConfig := LoggingConfig{
		"info",
		"log/pw_logs.log",
	}

	return &GeneralConfig{
		mongoDbConfig,
		loggingConfig,
	}
}
