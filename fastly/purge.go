package fastly

import (
	"fmt"
	"net/http"
	"time"
)

const (
	apiURL    = "https://api.fastly.com/service/%s/%s"
	keyHeader = "Fastly-Key"

	purgeAll = "purge_all"
	purgeKey = "purge/"

	softPurgeHeader = "Fastly-Soft-Purge"
)

func (c *Cache) doPurge(path string) error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(apiURL, c.ServiceID, path), nil)
	if err != nil {
		return err
	}

	req.Header.Add(keyHeader, c.APIKey)
	req.Header.Add(softPurgeHeader, "1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s %s -> %s", req.Method, req.URL, resp.Status)
	}

	return nil
}

func (c *Cache) purge(path string, description string) bool {
	t := time.Now()

	if err := c.doPurge(path); err != nil {
		c.Log.Println("Failed to purge", description, err)
		return false
	}

	c.Log.Println("Successfully purged", description, "in", time.Since(t))
	return true
}

func (c *Cache) PurgeAll() bool {
	return c.purge(purgeAll, "all")
}

func (c *Cache) PurgeKey(key string) bool {
	return c.purge(purgeKey+key, "key '"+key+"'")
}
