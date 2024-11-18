package proxy_data_input

import (
	"database/sql"
	"fmt"
	"good_proxies_db/shared"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

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
