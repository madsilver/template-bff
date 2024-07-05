package utils

import (
	"html"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type SanitizeFunc func(*string)

type Sanitize interface {
	Apply(v *string, fns ...SanitizeFunc)
	Do(v *[]byte)
}

type sanitize struct{}

func NewSanitize() Sanitize {
	return &sanitize{}
}

func (s *sanitize) Apply(v *string, fns ...SanitizeFunc) {
	for _, fn := range fns {
		fn(v)
	}
}

func (s *sanitize) Do(b *[]byte) {
	if len(*b) > 0 {
		v := string(*b)
		s.Apply(&v, BlueMonday, UnicodeChar, HtmlScape)
		*b = []byte(v)
	}
}

func BlueMonday(v *string) {
	s := bluemonday.NewPolicy().Sanitize(*v)
	s = strings.ReplaceAll(s, "&#34;", "\"")
	*v = html.UnescapeString(s)
}

func UnicodeChar(v *string) {
	RemovePattern(v, `\\u[\dA-Fa-f]{4}`)
}

func HtmlScape(v *string) {
	RemovePattern(v, `[<>=]`)
}

func RemovePattern(v *string, pattern string) {
	re := regexp.MustCompile(pattern)
	*v = re.ReplaceAllString(*v, "")
}
