package config

import (
    "os"
    "github.com/BurntSushi/toml"
)


type Config struct{
    Owner    string     `toml:"owner"`
    Repo     string     `toml:"repo"`
    Token    string     `toml:"token"`
    BaseURL  string     `toml:"baseURL"`
}

var DefaultFile = "config.toml"

func (conf *Config) ReadConfig(fn string) (err error) {
    var (
        cBytes  []byte
    )
    if _, err = os.Stat(fn); err != nil {
        fn = DefaultFile
    }
    if cBytes, err = os.ReadFile(fn); err != nil {
        return
    }
    if err = toml.Unmarshal(cBytes, conf); err != nil {
        return
    }
    return
}


func ReadConfig(fn string) (ac *Config, err error){

    ac = &Config{}

    var (
        cBytes  []byte
    )

    if _, err = os.Stat(fn); err != nil {
        fn = DefaultFile
    }

    if cBytes, err = os.ReadFile(fn); err != nil {
        return
    }

    if err = toml.Unmarshal(cBytes, ac); err != nil {
        return
    }

    return
}

func ReadConfigMap(fn string) (m map[string]interface{}, err error){
    var (
        cBytes  []byte
    )
    if _, err = os.Stat(fn); err != nil {
        fn = DefaultFile
    }
    if cBytes, err = os.ReadFile(fn); err != nil {
        return
    }

    if err = toml.Unmarshal(cBytes, &m); err != nil {
        return
    }
    return
}

