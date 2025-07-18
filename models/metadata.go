package models

type Metadata struct {
	Dependencies struct {
		Type   string `json:"type"`
		Values []struct {
			Name   string `json:"name"`
			Values []struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"values"`
		} `json:"values"`
	} `json:"dependencies"`

	Type struct {
		Type    string `json:"type"`
		Default string `json:"default"`
		Values  []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Action      string `json:"action"`
			Tags        struct {
				Build   string `json:"build"`
				Dialect string `json:"dialect"`
				Format  string `json:"format"`
			} `json:"tags"`
		} `json:"values"`
	} `json:"type"`

	Language struct {
		Type    string `json:"type"`
		Default string `json:"default"`
		Values  []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"language"`

	BootVersion struct {
		Default string `json:"default"`
		Values  []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"bootVersion"`

	GroupId struct {
		Type    string `json:"type"`
		Default string `json:"default"`
	} `json:"groupId"`

	ArtifactId struct {
		Type    string `json:"type"`
		Default string `json:"default"`
	} `json:"artifiactId"`

	Name struct {
		Type    string `json:"type"`
		Default string `json:"default"`
	} `json:"name"`

	PackageName struct {
		Type    string `json:"type"`
		Default string `json:"default"`
	} `json:"packageName"`

	Packaging struct {
		Type    string `json:"type"`
		Default string `json:"default"`
		Values  []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"packaging"`

	JavaVersion struct {
		Type    string `json:"type"`
		Default string `json:"default"`
		Values  []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"javaVersion"`
}
