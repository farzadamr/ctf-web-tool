package httpclient

import (
	"net/http"
	"time"
)

var Client = &http.Client{Timeout: 7 * time.Second}
