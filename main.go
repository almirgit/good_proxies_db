package main

import (
	"good_proxies_go_ai/config"
	"good_proxies_go_ai/proxy_data_input"
	"good_proxies_go_ai/shared"
)

func main() {

	//LOG_FILE := os.Getenv("LOG_FILE")
	logger, logfile := shared.Loginit()

	defer logfile.Close() // executes on function exit

	logger.Info("Start")

	// Load configuration
	cfg, err := config.LoadConfig(".config.yml")
	if err != nil {
		logger.Error("Error loading config: %v", err)
	}

	// Get the database connection
	db, err := proxy_data_input.DBConnect(*cfg)
	if err != nil {
		logger.Error("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Connection successful
	//fmt.Println("Connected to the PostgreSQL database successfully!")

	proxy_list := proxy_data_input.GetProxies(db)

	// // Using a for loop with an index
	// for i := 0; i < len(proxy_list); i++ {
	// 	fmt.Println(proxy_list[i])
	// }

	check_proxies(proxy_list)

	//logger.Debug(cfg.CheckURLEndPoint)
}
