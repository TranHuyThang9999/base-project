package configs

import (
	"encoding/json"
	"os"
)

type Configs struct {
	FileLc             string `json:"file_lc"`
	DataSource         string `json:"data_source"`
	Port               string `json:"port"`
	AccessSecret       string `json:"access_secret,omitempty"`
	ExpireAccess       string `json:"expire_access,omitempty"`
	AddressRedis       string `json:"addressRedis"`       // Địa chỉ và cổng Redis
	PasswordRedis      string `json:"passwordRedis"`      // Mật khẩu Redis (nếu có)
	DatabaseredisIndex string `json:"databaseredisIndex"` // Số DB trong Redis
}

var config *Configs

func Get() *Configs {
	return config
}
func LoadConfig(path string) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	byteValue, err := os.ReadFile(configFile.Name())
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}
}
