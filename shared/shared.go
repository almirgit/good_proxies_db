package shared

import (
	"fmt"
	"log/slog"
	"os"
)

// Config holds the overall configuration for the application
type Config struct {
	CheckURLEndPoint string         `yaml:"CheckURLEndPoint"`
	Database_in      DatabaseConfig `yaml:"database_in"`
	Database_out     DatabaseConfig `yaml:"database_out"`	
}

// DatabaseConfig holds the database-related configuration
type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
	DBName   string `yaml:"dbname"`
}

type Proxy struct {
	Proxy_ip   string `db:"proxy_ip"`
	Proxy_port string `db:"proxy_port"`
}

var Log *slog.Logger
var Logfile *os.File
const LOG_FILE = "output.log"
const FILE string = "good_proxies.db"

// func Loginit(logfile *os.File) (*slog.Logger) {
func Loginit() (*slog.Logger, *os.File) {
	logfile, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(err)
	}

	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	//logger := slog.New(slog.NewJSONHandler(os.Stderr, handlerOpts))
	logger := slog.New(slog.NewTextHandler(logfile, handlerOpts))
	slog.SetDefault(logger)

	return logger, logfile
}

// func hi(text string) (string) {
func Hi() {
	fmt.Println("Hi!")
}
