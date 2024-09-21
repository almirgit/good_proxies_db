package proxy_data_input

import (
	"database/sql"
	"fmt"
	"good_proxies_db/shared"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// getDBData takes a DatabaseConfig object, opens a connection to the PostgreSQL database,
// tests it, and returns the connection object.
func PgDBConnect(config shared.Config) (*sql.DB, error) {
	// Construct the PostgreSQL Data Source Name (DSN)
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
	//	config.Database_in.Host, config.Database_in.Username, config.Database_in.Password, config.Database_in.DBName, config.Database_in.SSLMode)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Database_in.Username, config.Database_in.Password,
		config.Database_in.Host, config.Database_in.Port,
		config.Database_in.DBName, config.Database_in.SSLMode)

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Return the database connection
	return db, nil
}

func GetProxies(db *sql.DB) []string {	

	//rows, err := db.Query("select proxy_ip, proxy_port from data.proxy_list pl order by moddate desc limit 1000")
	rows, _ := db.Query("select proxy_ip, proxy_port from data.proxy_list pl order by moddate desc limit 1000")
	//defer rows.Close()
	// if err != nil {
	// 	logger.Error("Error: %v", err)
	// }

	//var res Proxy
	col := shared.Proxy{}
	set := []string{}
	for rows.Next() {
		rows.Scan(&col.Proxy_ip, &col.Proxy_port)
		res := fmt.Sprintf("%v:%v", col.Proxy_ip, col.Proxy_port)
		set = append(set, res)
	}

	//fmt.Printf("type of set: %T\n", set)
	//type of set: []string
	return set
}
