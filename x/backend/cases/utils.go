package cases

import (
	"os"
)

// GetBackendDBDir return the path "$GOPATH/src/github.com/omexapp/omexchain/x/backend/cases"
func GetBackendDBDir() string {
	gopath := os.Getenv("GOPATH")
	dir := gopath + "/src/github.com/omexapp/omexchain/x/backend/cases"
	return dir
}
