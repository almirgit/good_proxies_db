package main

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

func check_proxies(dbdata []string) {
	logger, _ := shared.Loginit()
	//requestURL := "https://api-v2.capex.com/quotesv2?key=1&q=copper"
	requestURL := "https://static.kodera.hr"
	for _, proxy_server := range dbdata {
		//fmt.Println(proxy_server)
		proxy_server_url := "http://" + proxy_server
		logger.Info("Testing proxy", "proxy_server", proxy_server_url)
		proxyURL, _ := url.Parse(proxy_server_url)
		proxy := http.ProxyURL(proxyURL)
		transport := &http.Transport{Proxy: proxy}
		client := &http.Client{Transport: transport, Timeout: 10 * time.Second}
		//req, _ := http.NewRequest("POST", "http://server.name", somedata)
		//resp, err := client.Do(req)

		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			logger.Error("client: could not create request", "error", err)
			//os.Exit(1)
			continue
		}

		//res, err := http.DefaultClient.Do(req)
		res, err := client.Do(req)
		if err != nil {
			logger.Error("client: error making http request", "error", err)
			//os.Exit(1)
			continue
		}

		//fmt.Printf("client: got response!\n")
		logger.Debug("client: status code", "status code", res.StatusCode)
		err = save_good_proxy(proxy_server)
		if err != nil {
			//fmt.Print(err) // this works, but looks like Errorf doesn't work
			//fmt.Printf(os.Stderr, "Cannot write to sqlite DB file: %s\n", err)
			logger.Error("Cannot write to sqlite DB file", "error", err)
			//fmt.Errorf("Cannot write to sqlite DB file: %s\n", err)
		}
	}
}

func save_good_proxy(proxy_ip_port string) error {
	logger, _ := shared.Loginit()

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

	logger.Info("Added new good proxy", "proxy_name", proxy_ip_port)

	return nil
}
