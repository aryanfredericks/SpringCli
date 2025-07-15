package utils

import (
	"fmt"
	"strings"
)

func GenerateFinalURL(
	selectedProjectType, selectedLanguage, selectedVersion, projectName, groupId, artifactId, packageName, selectedPackaging, javaVersion string, dependencies []string,
) string {
	return fmt.Sprintf(
		"https://start.spring.io/starter.zip?type=%s&language=%s&bootVersion=%s&baseDir=%s&groupId=%s&artifactId=%s&name=%s&packageName=%s&packaging=%s&javaVersion=%s&dependencies=%s",
		selectedProjectType, selectedLanguage, selectedVersion, projectName, groupId, artifactId, projectName, packageName, selectedPackaging, javaVersion, strings.Join(dependencies, ","),
	)
}
