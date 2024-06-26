package query_eviction_config

import "time"

// QueryEvictionConfig is a struct that contains the configuration options for the QueryCache instance.
//
// Fields:
//   - TTL: The Time-to-Live (TTL) duration for the QueryCache instance.
//   - Now: A function that returns the current time.
type QueryEvictionConfig struct {
	// TTL is the Time-to-Live (TTL) duration for the QueryCache instance.
	TTL time.Duration
	// Now is a function that returns the current time, which is used for calculating the expiration time of cache entries.
	Now func() time.Time
}
