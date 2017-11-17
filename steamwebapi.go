package steamwebapi

import (
	"net/http"
	"time"
)

const (
	ENDPOINT = "https://api.steampowered.com"
)

func httpGetWithTimeout(url string, timeout time.Duration) (*http.Response, error) {
	cli := http.Client{
		Timeout: timeout,
	}
	return cli.Get(url)
}
