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
	replacePathAnalytic            = "'this.components.analytic.container.image.tag=\"%s\"'"
	replacePathApi                 = "'this.components.api.container.image.tag=\"%s\"'"
	replacePathAuth                = "'this.components.auth.container.image.tag=\"%s\"'"
	replacePathCore                = "'this.components.core.container.image.tag=\"%s\"'"
	replacePathManager             = "'this.components.manager.container.image.tag=\"%s\"'"
	replacePathMessages            = "'this.components.messages.container.image.tag=\"%s\"'"
	replacePathVulnerability       = "'this.components.vulnerability.container.image.tag=\"%s\"'"
	replacePathWebhook             = "'this.components.webhook.container.image.tag=\"%s\"'"
	replacePathDatabaseMigration   = "'this.global.database.migration.image.tag=\"%s\"'"
	replacePathAnalyticDatabase    = "'this.components.analytic.database.migration.image.tag=\"%s\"'"
	pathToReplaceSeedKustomization = "config/manager/kustomization.yaml"
	pathToReplaceSeedReadme        = "README.md"
	defaultJsonPath                = "api/v2alpha1/horusec_platform_defaults.json"
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

	//for _, valueToReplace := range replaceValues() {
	//	if err := replacePlatformVersion(valueToReplace); err != nil {
	//		return err
	//	}
	//}

	return updateOperatorVersion()
}

func replacePlatformVersion(valueToReplace string) error {
	valueReplaced := fmt.Sprintf(valueToReplace, getPlatformVersion())

	fmt.Printf("json -I -f %s -e %s", defaultJsonPath, valueReplaced)
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
	seedValue := fmt.Sprintf("'s/%s/%s/g'", getActualVersion(), getReleaseVersion())

	return sh.Run("sed", "-i", seedValue, pathToReplaceSeedReadme)
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
