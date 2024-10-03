package main

import (
	"flag"
	"fmt"
	"good_proxies_db/config"
	"good_proxies_db/proxy_data_input"
	"good_proxies_db/proxy_data_output"
	"good_proxies_db/shared"
	"os"

	//"golang.org/x/text/language/display"
)

var (
	// These variables will be set at build time using -ldflags
	version   string
	commitSHA string
)

func display_error(errmsg string, err error) {
	shared.Log.Error(errmsg, "error", err)
	fmt.Fprintln(os.Stderr, errmsg)
}

func main() {

	//LOG_FILE := os.Getenv("LOG_FILE")
	shared.Log, shared.Logfile = shared.Loginit()
	defer shared.Logfile.Close() // executes on function exit

	configf := flag.String("config", ".config.yml", "Configuration file")
	disp_version := flag.Bool("version", false, "Display version of the program")
	flag.Parse()	

	if *disp_version {
		fmt.Printf("%s%s\n", version, commitSHA)
		return
	}
	
	for {
		shared.Log.Info("Start processing cycle")

		// Load configuration
		cfg, err := config.LoadConfig(*configf)
		if err != nil {
			display_error("Error reading .config.yml", err)
			//logger.Error("Error reading .config.yml", "error", err)
			return
		}

		// Get the database connection
		db, err := proxy_data_input.PgDBConnect(*cfg)
		if err != nil {
			//logger.Error("Error connecting to database: %v", "error", err)
			display_error("Error connecting to database", err)
			return
		}
		defer db.Close()

		proxy_list := proxy_data_input.GetProxies(db)

		// // Using a for loop with an index
		// for i := 0; i < len(proxy_list); i++ {
		// 	fmt.Println(proxy_list[i])
		// }

		proxy_data_output.Check_proxies(db, *cfg, proxy_list)
		proxy_data_output.Check_stored_proxies(db, *cfg)
	}

}
