package swg

import "net/http"

func AddHeaders(resp http.ResponseWriter) {
	header := resp.Header()

	header.Add("X-Content-Type-Options", "nosniff")
	header.Add("X-Frame-Options", "DENY")
	header.Add("X-XSS-Protection", "1; mode=block")
}
