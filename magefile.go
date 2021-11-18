//go:build mage
// +build mage

package main

import (
	"github.com/ZupIT/horusec-devkit/pkg/utils/mageutils"
)

func UpVersions(releaseType string) error {
	return mageutils.UpVersions(releaseType)
}

func CherryPick() error {
	return mageutils.CherryPick()
}

func UpdateVersioningFiles() {

}
