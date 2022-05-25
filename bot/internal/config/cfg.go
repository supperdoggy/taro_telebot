package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var config appConfig

type appConfig struct {
	DBUrl                          string `json:"db_url"`
	DBName                         string `json:"db_name"`
	WarningCollectionName          string `json:"warning_collection_name"`
	AdviceCollectionName           string `json:"advice_collection_name"`
	PicCollectionName              string `json:"pic_collection_name"`
	RuLocCollectionName            string `json:"ru_loc_collection_name"`
	DailyTaroCollectionName        string `json:"daily_taro_collection_name"`
	DailyTaroHistoryCollectionName string `json:"daily_taro_history_collection_name"`

	IsProd bool   `json:"is_prod"`
	Token  string `json:"token"`
}

func init() {
	path := flag.String("config", "", "for config path")
	flag.Parse()
	if path == nil || *path == "" {
		panic("no config path specified")
	}

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err.Error())
	}

	if config.Token == "" {
		panic("token empty")
	}

}

func GetConfig() appConfig {
	return config
}
