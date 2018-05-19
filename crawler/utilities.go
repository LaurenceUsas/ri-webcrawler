package ricrawler

import (
	"net/http"
	"time"
)

func NewHTTPClient() *http.Client {
	nc := &http.Client{
		Timeout: time.Second * 30,
	}
	return nc
}
