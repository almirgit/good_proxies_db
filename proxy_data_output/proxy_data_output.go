//package main
package proxy_data_output

import (
	"database/sql"
	//"fmt"
	//"good_proxies_go_ai/input"
	"good_proxies_go_ai/shared"
	//"log/slog"
	"net/http"
	"net/url"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func http_request(requestURL string, proxy_server_url string) (int) {
	//requestURL := "https://static.kodera.hr"	
	
	//logger, _ := shared.Loginit()
	proxyURL, _ := url.Parse(proxy_server_url)
	proxy := http.ProxyURL(proxyURL)
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport, Timeout: 10 * time.Second}
	//req, _ := http.NewRequest("POST", "http://server.name", somedata)
	//resp, err := client.Do(req)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		shared.Log.Error("client: could not create request", "error", err)
		//os.Exit(1)
		//continue
		return -1
	}
	
	res, err := client.Do(req)
	if err != nil {
		shared.Log.Error("client: error making http request", "error", err)
		return -1
	}

	//fmt.Printf("client: got response!\n")
	shared.Log.Debug("client: status code", "status code", res.StatusCode)
	return res.StatusCode
}

func Check_proxies(config shared.Config, dbdata []string) {
	//logger, _ := shared.Loginit()	
	
	for _, proxy_server := range dbdata {
		//fmt.Println(proxy_server)
		proxy_server_url := "http://" + proxy_server
		shared.Log.Info("Testing proxy", "proxy_server", proxy_server_url)
		
		http_result := http_request(config.CheckURLEndPoint, proxy_server_url)

		if (http_result == 200) {
			err := add_good_proxy(proxy_server)
			if err != nil {
				//fmt.Print(err) // this works, but looks like Errorf doesn't work
				//fmt.Printf(os.Stderr, "Cannot write to sqlite DB file: %s\n", err)
				shared.Log.Error("Cannot write to sqlite DB file", "error", err)
				//fmt.Errorf("Cannot write to sqlite DB file: %s\n", err)
			}
		}
	}
}

func add_good_proxy(proxy_ip_port string) error {
	//logger, _ := shared.Loginit()

	db, err := sql.Open("sqlite3", shared.FILE)
	if err != nil {
		return err
	}

	const create string = `
		CREATE TABLE IF NOT EXISTS proxies (
		id INTEGER NOT NULL PRIMARY KEY autoincrement,
		proxy_ip_port TEXT not null unique,
		entdate text not null DEFAULT (CURRENT_TIMESTAMP)
		);`

	if _, err := db.Exec(create); err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO proxies (proxy_ip_port) VALUES(?);", proxy_ip_port)
	if err != nil {
		return err
	}

	shared.Log.Info("Added new good proxy", "proxy_name", proxy_ip_port)

	return nil
}



func Check_stored_proxies(config shared.Config, dbdata []string) {
	//logger, _ := shared.Loginit()
	
	for _, proxy_server := range dbdata {
		proxy_server_url := "http://" + proxy_server
		shared.Log.Info("Testing proxy", "proxy_server", proxy_server_url)
		
		http_result := http_request(config.CheckURLEndPoint, proxy_server_url)

		if (http_result == 200) {
			err := remove_good_proxy(proxy_server)
			if err != nil {
				shared.Log.Error("Cannot write to sqlite DB file", "error", err)
			}
		}
	}
}


func remove_good_proxy(proxy_ip_port string) error {
    // No operation performed
	return nil
}
