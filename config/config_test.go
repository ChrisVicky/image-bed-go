package config

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
    fn := "./config.toml"
    c, err := ReadConfig(fn)
    if err != nil {
        t.Errorf("Error: %v", err)
        return
    }
    t.Logf("%+v", c)
}

func TestReadConfigMap(t *testing.T){
    
    fn := "./config.toml"
    _, err := ReadConfigMap(fn)
    if err != nil {
        t.Errorf("Error: %v", err)
        return
    }
}

