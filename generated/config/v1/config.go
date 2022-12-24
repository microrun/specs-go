// Code generated from github.com/microrun/specs/config/v1. DO NOT EDIT.
package v1

import (
	"encoding/json"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// App defines model for App.
type App struct {
	Build App_Build `json:"build"`
}

// App_Build defines model for App.Build.
type App_Build struct {
	union json.RawMessage
}

// BuildDockerfile defines model for BuildDockerfile.
type BuildDockerfile struct {
	Dockerfile string `json:"dockerfile"`
}

// BuildImage defines model for BuildImage.
type BuildImage struct {
	Image string `json:"image"`
}

// Config defines model for Config.
type Config struct {
	Apps      map[string]App      `json:"apps"`
	Providers map[string]Provider `json:"providers"`
}

// Provider defines model for Provider.
type Provider struct {
	// Provider URI pointing to the provider implementation
	// Supported protocols: file:// and git over https://
	Provider string `json:"provider"`
}

// AsBuildImage returns the union data inside the App_Build as a BuildImage
func (t App_Build) AsBuildImage() (BuildImage, error) {
	var body BuildImage
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBuildImage overwrites any union data inside the App_Build as the provided BuildImage
func (t *App_Build) FromBuildImage(v BuildImage) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBuildImage performs a merge with any union data inside the App_Build, using the provided BuildImage
func (t *App_Build) MergeBuildImage(v BuildImage) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsBuildDockerfile returns the union data inside the App_Build as a BuildDockerfile
func (t App_Build) AsBuildDockerfile() (BuildDockerfile, error) {
	var body BuildDockerfile
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBuildDockerfile overwrites any union data inside the App_Build as the provided BuildDockerfile
func (t *App_Build) FromBuildDockerfile(v BuildDockerfile) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBuildDockerfile performs a merge with any union data inside the App_Build, using the provided BuildDockerfile
func (t *App_Build) MergeBuildDockerfile(v BuildDockerfile) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t App_Build) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *App_Build) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
