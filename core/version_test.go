package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	v, err := NewVersion(1).MarshalJSON()
	require.NoError(t, err)
	t.Log(v)
	a := NewVersion(2)
	require.NoError(t, a.UnmarshalJSON(v))
	if !a.Eq(1) {
		t.Fatal("version not equal")
	}
}
