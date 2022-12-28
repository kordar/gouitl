package goutil

import (
	"github.com/spf13/viper"
	"log"
)

var cfg = viper.New()
var errCfg error

func ConfigInit(filepath string) {
	ConfigInit2(filepath, "ini")
}

func ConfigInit2(filepath string, in string) {
	cfg.SetConfigFile(filepath)
	cfg.SetConfigType(in)
	errCfg = cfg.ReadInConfig()
	if errCfg != nil {
		log.Fatal(errCfg)
	}
}

func GetSystemValue(key string) string {
	return cfg.GetString("system." + key)
}

func GetSettingValue(key string) string {
	return cfg.GetString("setting." + key)
}

func GetSectionValue(section string, key string) string {
	return cfg.GetString(section + "." + key)
}

func GetSectionValueInt(section string, key string) int {
	return cfg.GetInt(section + "." + key)
}

func GetSection(section string) map[string]string {
	return cfg.GetStringMapString(section)
}

func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return cfg.UnmarshalKey(key, rawVal, opts...)
}

func GetCfg() *viper.Viper {
	return cfg
}
