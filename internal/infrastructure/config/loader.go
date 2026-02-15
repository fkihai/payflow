package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App          AppConfig            `yaml:"app"`
	Server       ServerConfig         `yaml:"server"`
	Database     DatabaseConfig       `yaml:"database"`
	Peyment      PaymentGatewayConfig `yaml:"payment_gateway"`
	MessageQueue MessageQueueConfig   `yaml:"message_queue"`
}

type AppConfig struct {
	Name     string `yaml:"name"`
	Env      string `yaml:"environment"`
	TimeZone string `yaml:"timezone"`
}

type ServerConfig struct {
	Http HttpServerConfig `yaml:"http"`
}

type HttpServerConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type DatabaseConfig struct {
	Primary PrimaryDatabaseConfig `yaml:"primary"`
}

type PrimaryDatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	SSLMode  string `yaml:"ssl_mode"`
	Port     int
	Name     string
	User     string
	Password string
}

type PaymentGatewayConfig struct {
	WebhookPath   string `yaml:"webhook_path"`
	Env           string `yaml:"environment"`
	SandBoxUrl    string `yaml:"sandbox_url"`
	ProductionUrl string `yaml:"production_url"`
	ServerKey     string
}

type MessageQueueConfig struct {
	Type     string                     `yaml:"type"`
	RabbitMQ RabbitMQMessageQueueConfig `yaml:"rabbitmq"`
}

type RabbitMQMessageQueueConfig struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	VHost string `yaml:"vhost"`
}

func getString(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func getInt(k string, d int) int {
	v := os.Getenv(k)
	if v == "" {
		return d
	}

	var i int
	_, err := fmt.Sscanf(v, "%d", &i)
	if err != nil {
		panic("invalid int env: " + k)
	}

	return i
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(yamlFile, cfg); err != nil {
		return nil, err
	}

	// load env
	godotenv.Load()

	cfg.Database.Primary.Port = getInt("DB_PORT", 5432)
	cfg.Database.Primary.Name = getString("DB_NAME", "payflow")
	cfg.Database.Primary.User = getString("DB_USER", "dev")
	cfg.Database.Primary.Password = getString("DB_PASSWORD", "root")

	cfg.Peyment.ServerKey = getString("PAYMENT_SERVER_KEY", "")

	return cfg, nil
}
