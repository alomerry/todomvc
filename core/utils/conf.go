package utils

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
	"reflect"
	"sync"
)

func Config(once *sync.Once, cfg interface{}, location string) {
	once.Do(func() {
		filePath, err := filepath.Abs(location)
		if err != nil {
			panic(err)
		}
		if _, err := toml.DecodeFile(filePath, cfg); err != nil {
			panic(err)
		}
	})
}

func GetConfItem(cfg interface{}, field string) reflect.Value {
	return reflect.ValueOf(cfg).Elem().FieldByName(field)
}
