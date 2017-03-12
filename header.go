package swg

import "net/http"

func AddHeaders(resp http.ResponseWriter) {
	header := resp.Header()

	header.Add("Content-Security-Policy", "Content-Security-Policy: default-src 'none'; script-src 'self' 'unsafe-eval' https://cdnjs.cloudflare.com/; style-src 'self' https://cdnjs.cloudflare.com; img-src 'self'; font-src https://cdnjs.cloudflare.com/ https://fonts.googleapis.com/; connect-src 'self'; frame-src https://kiwiirc.com; frame-ancestors 'none'; upgrade-insecure-requests; block-all-mixed-content; ")
	header.Add("X-Content-Type-Options", "nosniff")
	header.Add("X-Frame-Options", "DENY")
	header.Add("X-XSS-Protection", "1; mode=block")
}
