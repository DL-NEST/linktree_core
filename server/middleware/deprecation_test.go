package middleware

import (
	"regexp"
	"testing"
)

func BenchmarkDeprecation(b *testing.B) {
	var deprecation = []string{
		"/v1/*", "/ss/*", "/ssc/*", "/sds/*",
	}
	for i := 0; i < b.N; i++ {
		for _, i2 := range deprecation {
			_, err := regexp.MatchString(i2, "/v1/sd")
			if err != nil {
				return
			}
		}
	}
}
