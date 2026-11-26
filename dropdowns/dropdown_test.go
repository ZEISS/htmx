package dropdowns_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	htmx "github.com/zeiss/htmx"
	"github.com/zeiss/htmx/dropdowns"
)

func TestDropdown(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<details class=\"dropdown\"></details>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<details class=\"custom-class dropdown\"></details>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := dropdowns.Dropdown(
				dropdowns.Props{
					ClassNames: test.classes,
				},
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}
