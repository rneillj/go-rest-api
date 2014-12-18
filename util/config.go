package config

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

type Config struct {
    Auth struct {
        Rackspace struct {
            Username string
            ApiKey string
            Endpoint string
        }
    }
}

var config *Config = nil

func GetConfig() *Config {

    if config == nil {
        config = new(Config)

        file, err := ioutil.ReadFile("config/config.json")
        if err != nil {
            log.Fatal(err)
        }

        err = json.Unmarshal(file, &config)
        if err != nil {
            log.Fatal(err)
        }
    }

    return config
}
