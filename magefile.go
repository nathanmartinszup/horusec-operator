// +build mage

package main

import (
	mage "github.com/ZupIT/horusec-devkit"
)

func Version(releaseType string) error {
	_, err := mage.UpVersions(releaseType)

	return err
}
