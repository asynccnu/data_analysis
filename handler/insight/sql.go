package insight

import (
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
)

func sql2json(sql string) (json string) {
	var (
		hostname = viper.GetString("influxDB.addr") + "/query?"
		dbname   = "db=" + "insight" + "&"
	)
	sqlUrl := url.PathEscape(sql)
	encodeUrl := hostname + dbname + sqlUrl
	resp, err := http.Get(encodeUrl)
	if err != nil {
		log.Fatal("http get wrong", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json = string(body)
	return json
}
