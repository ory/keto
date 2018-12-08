package cmd

import (
	"fmt"
	"os"

	"github.com/akutz/gotil"
)

var port int

func init() {
	var osArgs = make([]string, len(os.Args))
	port = gotil.RandomTCPPort()
	os.Setenv("DATABASE_URL", "memory")
	os.Setenv("PORT", fmt.Sprintf("%d", port))
	os.Setenv("KETO_URL", fmt.Sprintf("http://127.0.0.1:%d", port))
	copy(osArgs, os.Args)
}
