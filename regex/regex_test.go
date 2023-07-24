package regex_test

import (
	"regexp"
	"strings"
	"testing"
)

func BenchmarkStringNoRegex(b *testing.B) {
	myblogurl := "https://omarghader.github.io/"
	for n := 0; n < b.N; n++ {
		// check if the string starts with http or https
		if !strings.HasPrefix(myblogurl, "http") && !strings.HasPrefix(myblogurl, "https") {
			b.Errorf("string doesn't start with http or https")
		}
	}
}

func BenchmarkStringRegex(b *testing.B) {
	myblogurl := "https://omarghader.github.io/"
	regex := regexp.MustCompile("http[s]?")
	for n := 0; n < b.N; n++ {
		// check if the string starts with http or https
		if !regex.MatchString(myblogurl) {
			b.Errorf("regex doesn't match the url")
		}
	}
}

func BenchmarkStringRegexBegin(b *testing.B) {
	myblogurl := "https://omarghader.github.io/"
	regex := regexp.MustCompile("^http")
	for n := 0; n < b.N; n++ {
		// check if the string starts with http or https
		if !regex.MatchString(myblogurl) {
			b.Errorf("regex doesn't match the url")
		}
	}
}
