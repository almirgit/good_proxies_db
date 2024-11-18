package shared

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"sync"
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

// var Log *slog.Logger
// var Logfile *os.File
// var db *sql.DB // Global database connection

var (
	Log     *slog.Logger
	Logfile *os.File
	dbLock sync.Mutex
)

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

// getDBData takes a DatabaseConfig object, opens a connection to the PostgreSQL database,
// tests it, and returns the connection object.
func PgDBConnect(config DatabaseConfig, db *sql.DB) (*sql.DB, error) {

	var err error

	dbLock.Lock()
	defer dbLock.Unlock()


	// Check if the connection is already initialized
	if db != nil {
		// Verify the connection is still valid
		if err := db.Ping(); err == nil {
			Log.Info("Database connection is already initialized and valid.")
			return db, nil
		}
		// If invalid, close it
		Log.Warn("Database connection is invalid. Reinitializing...")
		_ = db.Close()
		db = nil
	}

	// Initialize the connection
	Log.Info("Initializing new database connection...")
	// Construct the PostgreSQL Data Source Name (DSN)
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
	//	config.Database_in.Host, config.Database_in.Username, config.Database_in.Password, config.Database_in.DBName, config.Database_in.SSLMode)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Username, config.Password,
		config.Host, config.Port,
		config.DBName, config.SSLMode)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		_ = db.Close()
		db = nil
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// func hi(text string) (string) {
// func Hi() {
// 	fmt.Println("Hi!")
// }
