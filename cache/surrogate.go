package cache

const (
	// Serve stale content for up to 5 minutes when updating/a week when an error occurs
	SurrogateStaleOptions = "stale-while-revalidate=300, stale-if-error=604800"

	// Cache for up to 1 month
	SurrogateStaticContentOptions = "max-age=2628000, " + SurrogateStaleOptions

	// Cache for up to an hour
	SurrogateDynamicContentOptions = "max-age=3600, " + SurrogateStaleOptions
)
