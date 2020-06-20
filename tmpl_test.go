package tmplt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	textText = Text("{{.V}}")
	testHTML = HTML("{{.V}}")
)

type v struct{ V int }

func Test(t *testing.T) {
	t.Parallel()

	got, err := textText.Execute(v{42})
	require.NoError(t, err)
	assert.Equal(t, "42", got)

	got, err = testHTML.Execute(v{42})
	require.NoError(t, err)
	assert.Equal(t, "42", got)
}

func TestFailure(t *testing.T) {
	t.Parallel()

	_, err := textText.Execute(42)
	assert.Error(t, err)

	_, err = testHTML.Execute(42)
	assert.Error(t, err)
}
