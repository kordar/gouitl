package goutil

import (
	"github.com/Unknwon/goconfig"
	"github.com/spf13/cast"
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

func GetSectionValue(section string, key string) string {
	if value, err := cfg.GetValue(section, key); err == nil {
		return value
	} else {
		return ""
	}
}

func GetSectionValueInt(section string, key string) int {
	value := GetSectionValue(section, key)
	return cast.ToInt(value)
}

func GetSection(section string) map[string]string {
	res, err := cfg.GetSection(section)
	if err != nil {
		return map[string]string{}
	}
	return res
}
