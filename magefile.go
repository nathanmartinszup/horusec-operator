//go:build mage
// +build mage

package main

import (
	"fmt"
	"github.com/ZupIT/horusec-devkit/pkg/utils/mageutils"
	"github.com/magefile/mage/sh"
	"os"
)

const (
	replacePathAnalytic          = "'this.components.analytic.container.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	replacePathApi               = "'this.components.api.container.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	replacePathAuth              = "'this.components.auth.container.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	replacePathCore              = "'this.components.core.container.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	replacePathManager           = "'this.components.manager.container.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	replacePathMessages          = "'this.components.messages.container.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	replacePathVulnerability     = "'this.components.vulnerability.container.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	replacePathWebhook           = "'this.components.webhook.container.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	replacePathDatabaseMigration = "'this.global.database.migration.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	replacePathAnalyticDatabase  = "'this.components.analytic.database.migration.image.tag=\"${{ github.event.inputs.horusecPlatformVersion }}\"'"
	defaultJsonPath              = "api/v2alpha1/horusec_platform_defaults.json"
	seedOperatorVersion          = "\"s/%s/%s/g\""
)

const (
	EnvPlatformVersion = "HORUSEC_PLATFORM_VERSION"
	EnvActualVersion   = "HORUSEC_ACTUAL_VERSION"
	EnvReleaseVersion  = "HORUSEC_RELEASE_VERSION"
)

func UpVersions(releaseType string) error {
	return mageutils.UpVersions(releaseType)
}

func CherryPick() error {
	return mageutils.CherryPick()
}

func UpdateVersioningFiles() error {
	if err := sh.RunV("npm", "install", "-g", "json"); err != nil {
		return err
	}

	for _, valueToReplace := range replaceValues() {
		if err := replacePlatformVersion(valueToReplace); err != nil {
			return err
		}
	}

	return updateOperatorVersion()
}

func replacePlatformVersion(valueToReplace string) error {
	valueReplaced := fmt.Sprintf(valueToReplace, getPlatformVersion())

	return sh.RunV("json", "-I", "-f", defaultJsonPath, "-e", valueReplaced)
}

func replaceValues() []string {
	return []string{
		replacePathAnalytic,
		replacePathApi,
		replacePathAuth,
		replacePathCore,
		replacePathManager,
		replacePathMessages,
		replacePathVulnerability,
		replacePathWebhook,
		replacePathDatabaseMigration,
		replacePathAnalyticDatabase,
	}
}

func updateOperatorVersion() error {
	seedValue := fmt.Sprintf(seedOperatorVersion, getActualVersion(), getReleaseVersion())

	return sh.RunV("find", ".", "-type", "f", "-not", "-path", "\"./.git/*\"", "-not", "-path",
		"\"./go.mod\"", "-not", "-path", "\"./go.sum\"", "|", "xargs", "sed", "-i", seedValue)
}

func getActualVersion() string {
	return os.Getenv(EnvActualVersion)
}

func getReleaseVersion() string {
	return os.Getenv(EnvReleaseVersion)
}

func getPlatformVersion() string {
	return os.Getenv(EnvPlatformVersion)
}
