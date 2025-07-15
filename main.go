package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/aryanredericks/SpringCli/models"
	"github.com/aryanredericks/SpringCli/utils"
)

func main() {

	done := make(chan bool)
	go utils.Spinner("Fetching Spring Initializr metadata...", done)
	META_DATA_URL := "https://start.spring.io/metadata/client"

	resp, err := http.Get(META_DATA_URL)
	if err != nil {
		panic("Failed to fetch metadata: " + err.Error())
	}

	defer resp.Body.Close()

	done <- true
	fmt.Print("\r✅ Metadata fetched successfully!        		\n")
	var meta models.Metadata
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		panic("Failed to decode metadata: " + err.Error())
	}

	// Project Type
	var selectedProjectType string
	projectType := []string{}
	for _, pT := range meta.Type.Values {
		projectType = append(projectType, pT.ID)
	}
	survey.AskOne(&survey.Select{
		Options: projectType,
		Message: "Select Project Type:",
		Default: meta.Type.Default,
	}, &selectedProjectType)

	// PROJECT Language
	var selectedLanguage string
	projectLanguage := []string{}
	for _, pL := range meta.Language.Values {
		projectLanguage = append(projectLanguage, pL.ID)
	}
	survey.AskOne(&survey.Select{
		Options: projectLanguage,
		Message: "Select Project Language:",
		Default: meta.Language.Default,
	}, &selectedLanguage)

	// Spring Boot Build Version
	var selectedVersion string
	versions := []string{}
	for _, v := range meta.BootVersion.Values {
		versions = append(versions, v.ID)
	}
	survey.AskOne(&survey.Select{
		Options: versions,
		Message: "Select Boot Version: ",
		Default: meta.BootVersion.Default,
	}, &selectedVersion)

	// Group ID
	var groupId string
	survey.AskOne(&survey.Input{
		Message: "Group ID:",
		Default: meta.GroupId.Default,
	}, &groupId)

	// Artifact ID
	var artifactId string
	survey.AskOne(&survey.Input{
		Message: "Artifact ID:",
		Default: meta.ArtifactId.Default,
	}, &artifactId)

	// Name
	var projectName string
	survey.AskOne(&survey.Input{
		Message: "Project Name:",
		Default: artifactId,
	}, &projectName)

	// Package Name
	var packageName string
	survey.AskOne(&survey.Input{
		Message: "Package Name:",
		Default: groupId + "." + artifactId,
	}, &packageName)

	// Packaging
	var selectedPackaging string
	packagingOptions := []string{}
	for _, p := range meta.Packaging.Values {
		packagingOptions = append(packagingOptions, p.ID)
	}
	survey.AskOne(&survey.Select{
		Message: "Select Packaging:",
		Options: packagingOptions,
		Default: meta.Packaging.Default,
	}, &selectedPackaging)

	// Java Version
	var javaVersion string
	javaOptions := []string{}
	for _, j := range meta.JavaVersion.Values {
		javaOptions = append(javaOptions, j.ID)
	}
	survey.AskOne(&survey.Select{
		Message: "Select Java Version:",
		Options: javaOptions,
		Default: meta.JavaVersion.Default,
	}, &javaVersion)

	var dependencies []string
	survey.AskOne(&survey.MultiSelect{
		Message: "Select dependencies:",
		Options: []string{"web", "data-jpa", "security", "devtools", "lombok", "websocket", "data-mongodb", "graphql", "mysql"},
		Default: []string{"web"},
	}, &dependencies)

	url := utils.GenerateFinalURL(
		selectedProjectType, selectedLanguage, selectedVersion, projectName, groupId, artifactId, packageName, selectedPackaging, javaVersion, dependencies,
	)

	fmt.Println("\nDownloading project from:")
	fmt.Println(url)

	resp, err = http.Get(url)
	if err != nil {
		panic("Failed to download project: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Failed to download project: " + resp.Status)
	}

	// SAVE THE ZIP
	zipFileName := fmt.Sprintf("%s.zip", projectName)
	utils.SaveZipFile(zipFileName, resp)

	// UNZIP FILE
	if err := utils.Unzip(zipFileName, projectName); err != nil {
		panic("❌ Failed to unzip project: " + err.Error())
	}

	// CLEANUP
	if err := os.Remove(zipFileName); err != nil {
		fmt.Println("⚠️  Warning: Failed to delete zip file:", err)
	}

	fmt.Println("\033[32m✅ Project downloaded successfully!\033[0m")
	fmt.Println("\033[32m🚀 Happy coding!\033[0m")
}
