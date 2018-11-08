package conf

import (
	"io/ioutil"
	"encoding/json"
)

type ConfigStruct struct {
	Debug bool `json:"debug"`
	ListenAddr string `json:"listen_addr"`

	LogDir    string `json:"log_dir"`
	LogFile   string `json:"log_file"`
	LogLevel  string `json:"log_level"`

	DBMyBlog MysqlConfig `json:"db_my_blog"`

}

type MysqlConfig struct {
	Conn     string `json:"conn"`
	PoolSize int    `json:"pool_size"`
}

var Config *ConfigStruct

func InitFileConfig(fileName string) (*ConfigStruct, error) {
	Config = new(ConfigStruct)
	if data, err := ioutil.ReadFile(fileName); err != nil {
		return nil, err
	} else {
		return Config, json.Unmarshal(data, Config)
	}
}
