package conf

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type GbeConfig struct {
	DataSource DataSourceCondig `json:dataSource`
	RestServer RestServerConfig `json:restServer`
}

type DataSourceCondig struct {
	DriverName        string `json:"driverName"`
	Addr              string `json:"addr"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}

var config GbeConfig
var configOnce sync.Once

func GetConfig() *GbeConfig {
	configOnce.Do(func() {
		bytes, err := ioutil.ReadFile("conf.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})

	return &config
}
