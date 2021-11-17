//go:build mage
// +build mage

package main

import (
	mage "github.com/ZupIT/horusec-devkit"
)

func Version(releaseType string) error {
	err := mage.UpVersions(releaseType)

	return err
}
