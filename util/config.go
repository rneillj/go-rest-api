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

func GetConfig() (map[string]interface{}) {
    //c := &Config{}
    config := make(map[string]interface{})

    file, err := ioutil.ReadFile("config/config.json")
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(file, &config)
    if err != nil {
        log.Fatal(err)
    }

    //config["auth"] = *c

    return config
}
