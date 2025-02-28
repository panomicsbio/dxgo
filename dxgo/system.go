package dxgo

import (
	"context"
	"fmt"
)

type FindDataObjectsSortBy struct {
	Field    string           `json:"field"`
	Ordering DXSortByOrdering `json:"ordering"`
}

type FindDataObjectsInput struct {
	Name     any                    `json:"name,omitempty"`
	Class    DXAssetType            `json:"class"`
	Scope    *FindDataObjectsScope  `json:"scope,omitempty"`
	SortBy   *FindDataObjectsSortBy `json:"sortBy,omitempty"`
	Describe bool                   `json:"describe"`
	Starting map[string]any         `json:"starting,omitempty"`
}

type FindDataObjectsResult struct {
	Project  string         `json:"project"`
	ID       string         `json:"id"`
	Describe map[string]any `json:"describe"`
}

type FindDataObjectsOutput struct {
	Results []*FindDataObjectsResult `json:"results"`
	Next    map[string]any           `json:"next"`
	Error   *ApiError                `json:"error"`
}

func (c *DXClient) FindDataObjects(ctx context.Context, input FindDataObjectsInput) (FindDataObjectsOutput, error) {
	output := new(FindDataObjectsOutput)
	err := c.DoInto(ctx, "/system/findDataObjects", input, output)
	if err != nil {
		return FindDataObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}

type FindExecutionsInput struct {
	Project        string `json:"project,omitempty"`
	Class          string `json:"class,omitempty"`
	State          string `json:"state,omitempty"`
	OriginJob      string `json:"originJob,omitempty"`
	Starting       string `json:"starting,omitempty"`
	ParentAnalysis string `json:"parentAnalysis,omitempty"`
	ParentJob      string `json:"parentJob,omitempty"`
	IncludeSubjobs bool   `json:"includeSubjobs,omitempty"`
	LaunchedBy     string `json:"launchedBy,omitempty"`
}

type FindExecutionResult struct {
	ID       string         `json:"id"`
	Describe map[string]any `json:"describe"`
}

type FindExecutionsOutput struct {
	Results []*FindExecutionResult `json:"results"`
	Next    string                 `json:"string"`
	Error   *ApiError              `json:"error"`
}

func (c *DXClient) FindExecutions(ctx context.Context, input FindExecutionsInput) (FindExecutionsOutput, error) {
	output := new(FindExecutionsOutput)
	err := c.DoInto(ctx, "/system/findExecutions", input, output)
	if err != nil {
		return FindExecutionsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}

type FindProjectsInput struct {
	Name     any            `json:"name,omitempty"`
	Level    string         `json:"level,omitempty"`
	Starting string         `json:"starting,omitempty"`
	Describe map[string]any `json:"describe"`
}

type FindProjectsOutput struct {
	Results []*FindProjectsResult `json:"results"`
	Next    string                `json:"next"`
	Error   *ApiError             `json:"error"`
}

type FindProjectsResult struct {
	ID       string         `json:"id"`
	Level    string         `json:"level"`
	Public   bool           `json:"public"`
	Describe map[string]any `json:"describe"`
}

func (c *DXClient) FindProjects(ctx context.Context, input FindProjectsInput) (FindProjectsOutput, error) {
	output := new(FindProjectsOutput)
	err := c.DoInto(ctx, "/system/findProjects", input, output)
	if err != nil {
		return FindProjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}
