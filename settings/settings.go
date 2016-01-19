package settings

import (
	"fmt"
	"os"
)

var env = "develop"

// Init - Loads the Settings up for the environment
func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
}
