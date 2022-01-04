package util

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var cfg *goconfig.ConfigFile
var errCfg error

func ConfigInit(filepath string) {
	cfg, errCfg = goconfig.LoadConfigFile(filepath)
	if errCfg != nil {
		log.Fatal(errCfg)
	}
}

func GetSystemValue(key string) string {
	if value, err := cfg.GetValue("system", key); err == nil {
		return value
	} else {
		return ""
	}
}

func GetSettingValue(key string) string {
	if value, err := cfg.GetValue("setting", key); err == nil {
		return value
	} else {
		return ""
	}
}

func GetSectionValue(key string, field string) string {
	if value, err := cfg.GetValue(key, field); err == nil {
		return value
	} else {
		return ""
	}
}
