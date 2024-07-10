package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sanitize_Do(t *testing.T) {
	tests := []struct {
		name string
		xss string
		expect string
	}{
		{
			name: "query with script",
			xss: "<style>.xss{background-image:url(\\\\\\\"javascript:alert('XSS')\\\\\\\");}</style>",
			expect: `{"query":"mutation {foo("silver")}}`,
		},
		{
			name: "query with script",
			xss: `<image src="javascript:alert('XSS')">`,
			expect: `{"query":"mutation {foo("silver")}}`,
		},
		{
			name: "query with unicode",
			xss: "\\u0022\\u003e\\u003c\\u0069\\u006d\\u0067\\u0020\\u0073\\u0072\\u0063\\u003d\\u0022\\u0022\\u0069",
			expect: `{"query":"mutation {foo("silver")}}`,
		},
		{
			name: "query with encoded html tags",
			xss: "&lt;_is&gt;_awesome&lt;&gt;",
			expect: `{"query":"mutation {foo("silver_is_awesome")}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := []byte(fmt.Sprintf(`{"query":"mutation {foo("silver%s")}}`, tt.xss))
			NewSanitize().Do(&body)
			assert.Equalf(t, tt.expect, string(body), "SanitizeBody(%v)", tt.xss)
		})
	}
}
