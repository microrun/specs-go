package v1

import (
	"encoding/json"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

type App struct {
	Build App_Build `json:"build"`
}

type App_Build struct {
	union json.RawMessage
}

type BuildDockerfile struct {
	Dockerfile string `json:"dockerfile"`
}

type BuildImage struct {
	Image string `json:"image"`
}

type Config struct {
	Apps      map[string]App      `json:"apps"`
	Providers map[string]Provider `json:"providers"`
}

type Provider struct {
	Provider string `json:"provider"`
}

func (t App_Build) AsBuildImage() (BuildImage, error) {
	var body BuildImage
	err := json.Unmarshal(t.union, &body)
	return body, err
}

func (t *App_Build) FromBuildImage(v BuildImage) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

func (t *App_Build) MergeBuildImage(v BuildImage) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t App_Build) AsBuildDockerfile() (BuildDockerfile, error) {
	var body BuildDockerfile
	err := json.Unmarshal(t.union, &body)
	return body, err
}

func (t *App_Build) FromBuildDockerfile(v BuildDockerfile) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

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
