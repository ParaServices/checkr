package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_ListPackages(t *testing.T) {
	client := newClient(t)
	packages, err := client.ListPackages()
	require.NoError(t, err)
	require.NotNil(t, packages)
	require.NotEmpty(t, packages.Data)
}
