package main

import (
	"good_proxies_go_ai/config"
	"good_proxies_go_ai/proxy_data_input"
	"good_proxies_go_ai/proxy_data_output"
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
		return
	}

	// Get the database connection
	db, err := proxy_data_input.DBInConnect(*cfg)
	if err != nil {
		logger.Error("Error connecting to database: %v", err)
	}
	defer db.Close()

	proxy_list := proxy_data_input.GetProxies(db)

	// // Using a for loop with an index
	// for i := 0; i < len(proxy_list); i++ {
	// 	fmt.Println(proxy_list[i])
	// }

	proxy_data_output.Check_proxies(*cfg, proxy_list)
	//TODO: check_stored_proxies

	//logger.Debug(cfg.CheckURLEndPoint)
}
