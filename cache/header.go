package cache

const (
	CacheControlHeader = "Cache-Control"

	// Cache for up to a day if there is an error (barely supported by any browser)
	StaleOptions = "stale-if-error=86400"

	// Cache for up to an hour on the client
	StaticContentOptions = "max-age=3600, " + StaleOptions

	// Cache for 5 minutes
	DynamicContentOptions = "max-age=300, " + StaleOptions
)
