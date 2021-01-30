package cases

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetBackendDBDir(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	dir := gopath + "/src/github.com/omexapp/omexchain/x/backend/cases"
	require.Equal(t, dir, GetBackendDBDir())
}
