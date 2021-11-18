package main

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	os.Setenv(EnvPlatformVersion, "v1.0.0")
	os.Setenv(EnvActualVersion, "v2.3.8")
	os.Setenv(EnvReleaseVersion, "v2.3.9")

	err := UpdateVersioningFiles()
	fmt.Printf(err.Error())
}
