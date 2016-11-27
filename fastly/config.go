package fastly

import (
	"errors"
	"log"
	"strings"
)

const (
	configKeySeparator    = '/'
	configOptionSeparator = ";"
)

var errInvalidFormat = errors.New("Fastly config: API_KEY/SERVICE_ID[;healthcheck]")

type Cache struct {
	Log *log.Logger

	APIKey    string
	ServiceID string

	healthCheck bool
}

func ParseConfig(logger *log.Logger, config string) (*Cache, error) {
	options := strings.Split(config, configOptionSeparator)
	config = options[0]

	pos := strings.IndexByte(config, configKeySeparator)
	if pos == -1 {
		return nil, errInvalidFormat
	}

	cache := &Cache{
		Log:       logger,
		APIKey:    config[:pos],
		ServiceID: config[pos+1:],
	}

	for i := 1; i < len(options); i++ {
		switch options[i] {
		case "healthcheck":
			go cache.verifyHealthCheck()
		}
	}

	return cache, nil
}
