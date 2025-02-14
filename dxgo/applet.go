package dxgo

import (
	"context"
	"fmt"
)

type AppletRunInput struct {
	ID                 string             `json:"id"`
	Project            *string            `json:"project,omitempty"`
	Input              map[string]any     `json:"input"`
	SystemRequirements SystemRequirements `json:"systemRequirements"`
	Properties         map[string]string  `json:"properties"`
	Detach             bool               `json:"detach"`
	HeadJobOnDemand    bool               `json:"headJobOnDemand"`
}

type AppletRunOutput struct {
	ID    string    `json:"id"`
	Error *ApiError `json:"error"`
}

func (c *DXClient) AppletRun(ctx context.Context, input AppletRunInput) (AppletRunOutput, error) {
	output := new(AppletRunOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/run", input.ID), input, output)
	if err != nil {
		return AppletRunOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

// AppletGetOutput represents the full specification of an applet
type AppletGetOutput struct {
	ID             string            `json:"id"`
	Project        string            `json:"project"`
	Class          string            `json:"class"`
	Types          []string          `json:"types"`
	Created        int64             `json:"created"`
	State          string            `json:"state"`
	Hidden         bool              `json:"hidden"`
	Links          []string          `json:"links"`
	Name           string            `json:"name"`
	Folder         string            `json:"folder"`
	Sponsored      bool              `json:"sponsored"`
	Tags           []string          `json:"tags"`
	Modified       int64             `json:"modified"`
	CreatedBy      CreatedBy         `json:"createdBy"`
	RunSpec        RunSpec           `json:"runSpec"`
	DxAPI          string            `json:"dxapi"`
	Access         Access            `json:"access"`
	Title          string            `json:"title"`
	Summary        string            `json:"summary"`
	Description    string            `json:"description"`
	DeveloperNotes string            `json:"developerNotes"`
	IgnoreReuse    bool              `json:"ignoreReuse"`
	InputSpec      []IOSpec          `json:"inputSpec,omitempty"`
	OutputSpec     []IOSpec          `json:"outputSpec,omitempty"`
	Properties     map[string]string `json:"properties,omitempty"`
	Details        interface{}       `json:"details,omitempty"`
}

func (c *DXClient) AppletGet(ctx context.Context, appletID string) (AppletGetOutput, error) {
	output := new(AppletGetOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/get", appletID), nil, output)
	if err != nil {
		return AppletGetOutput{}, fmt.Errorf("getting applet: %w", err)
	}

	return *output, nil
}
