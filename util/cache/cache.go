package cache

import (
    "fmt"

    "github.com/censhin/go-rest-api/util"

    "github.com/bradfitz/gomemcache/memcache"
)

var (
    mc *memcache.Client= memcache.New("127.0.0.1:11211")
    user string = config.GetConfig().Auth.Rackspace.Username
)

func CacheValue(value []byte, action string) error {
    key := fmt.Sprintf("%s%s", user, action)

    err := mc.Set(&memcache.Item{Key: key, Value: value})

    return err
}

func GetValue(action string) ([]byte, error) {
    key := fmt.Sprintf("%s%s", user, action)

    item, err := mc.Get(key)

    if item != nil {
        return item.Value, nil
    } else {
        return nil, err
    }
}