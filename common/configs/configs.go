package configs

import (
	"encoding/json"
	"os"
)

type Configs struct {
	FileLc             string `json:"file_lc"`
	DataSource         string `json:"data_source"`
	Port               string `json:"port"`
	AccessSecret       string `json:"access_secret"`
	ExpireAccess       string `json:"expire_access"`
	AddressRedis       string `json:"address_redis"`
	PasswordRedis      string `json:"password_redis"`
	DatabaseRedisIndex int    `json:"database_redis_index"`
	KeyAes             string `json:"key_aes"`
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
