package alanube

import "net/url"

// queryWith creates url.Values with a single key-value pair.
func queryWith(key, value string) url.Values {
	q := url.Values{}
	q.Set(key, value)
	return q
}
