package dxgo

import (
	"context"
	"fmt"
)

type NewProjectInput struct {
	Name                     string            `json:"name"`
	Summary                  string            `json:"summary,omitempty"`
	Description              string            `json:"description,omitempty"`
	Protected                bool              `json:"protected,omitempty"`
	Restricted               bool              `json:"restricted,omitempty"`
	DownloadRestricted       bool              `json:"downloadRestricted,omitempty"`
	ExternalUploadRestricted bool              `json:"externalUploadRestricted,omitempty"`
	DatabaseUIViewOnly       bool              `json:"databaseUIViewOnly,omitempty"`
	ContainsPHI              bool              `json:"containsPHI,omitempty"`
	Tags                     []string          `json:"tags,omitempty"`
	Properties               map[string]string `json:"properties,omitempty"`
	BillTo                   string            `json:"billTo,omitempty"`
	Region                   string            `json:"region,omitempty"`
	MonthlyComputeLimit      *int              `json:"monthlyComputeLimit,omitempty"`
	MonthlyEgressBytesLimit  *int              `json:"monthlyEgressBytesLimit,omitempty"`
}
type NewProjectOutput struct {
	ID    string    `json:"id"`
	Error *ApiError `json:"error"`
}

func (c *DXClient) NewProject(ctx context.Context, input NewProjectInput) (NewProjectOutput, error) {
	output := new(NewProjectOutput)
	err := c.DoInto(ctx, "/project/new", input, output)
	if err != nil {
		return NewProjectOutput{}, err
	}

	return *output, nil
}

type ProjectDescribeInput struct {
	Fields map[string]bool `json:"fields,omitempty"`
}

type ProjectDescribeOutput struct {
	Name                 string         `json:"name"`
	BillTo               string         `json:"billTo"`
	Folders              []string       `json:"folders"`
	FileUploadParameters map[string]any `json:"fileUploadParameters"`
	Error                *ApiError      `json:"error"`
}

func (c *DXClient) ProjectDescribe(ctx context.Context, projectID string, input ProjectDescribeInput) (ProjectDescribeOutput, error) {
	output := new(ProjectDescribeOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", projectID), input, output)
	if err != nil {
		return ProjectDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type ProjectNewFolderInput struct {
	Folder string `json:"folder"`
	// Parents (optional, default false) Whether the parent folders should be created if they do not exist.
	Parents bool `json:"parents,omitempty"`
}

type ProjectNewFolderOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) ProjectNewFolder(ctx context.Context, projectID string, input ProjectNewFolderInput) (ProjectNewFolderOutput, error) {
	output := new(ProjectNewFolderOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/newFolder", projectID), input, output)
	if err != nil {
		return ProjectNewFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type RemoveObjectsInput struct {
	Objects []string `json:"objects"`
	Force   bool     `json:"force"`
}

type RemoveObjectsOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) RemoveObjects(ctx context.Context, projectID string, input RemoveObjectsInput) (RemoveObjectsOutput, error) {
	output := new(RemoveObjectsOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/removeObjects", projectID), input, output)
	if err != nil {
		return RemoveObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type RemoveFolderInput struct {
	Folder  string `json:"folder"`
	Force   bool   `json:"force"`
	Recurse bool   `json:"recurse"`
}

type RemoveFolderOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) RemoveFolder(ctx context.Context, projectID string, input RemoveFolderInput) (RemoveFolderOutput, error) {
	output := new(RemoveFolderOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/removeFolder", projectID), input, output)
	if err != nil {
		return RemoveFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type MoveObjectsInput struct {
	Objects     []string `json:"objects,omitempty"`
	Folders     []string `json:"folders,omitempty"`
	Destination string   `json:"destination"`
}

type MoveObjectsOutput struct {
	ID    string    `json:"id"`
	Error *ApiError `json:"error"`
}

func (c *DXClient) MoveObjects(ctx context.Context, projectID string, input MoveObjectsInput) (MoveObjectsOutput, error) {
	output := new(MoveObjectsOutput)

	err := c.DoInto(ctx, fmt.Sprintf("/%s/move", projectID), input, output)
	if err != nil {
		return MoveObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type DestroyProjectInput struct {
	TerminateJobs bool `json:"terminateJobs"`
}

type DestroyProjectOutput struct {
	ID    string    `json:"id"`
	Error *ApiError `json:"error"`
}

func (c *DXClient) DestroyProject(ctx context.Context, projectID string, input DestroyProjectInput) (DestroyProjectOutput, error) {
	output := new(DestroyProjectOutput)

	err := c.DoInto(ctx, fmt.Sprintf("/%s/destroy", projectID), input, output)
	if err != nil {
		return DestroyProjectOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type SetProjectProperties struct {
	ID         string            `json:"id"`
	Properties map[string]string `json:"properties"`
}

func (c *DXClient) SetProjectProperties(ctx context.Context, projectID string, input SetProjectProperties) error {
	var output any
	err := c.DoInto(ctx, fmt.Sprintf("/%s/setProperties", projectID), input, output)
	if err != nil {
		return fmt.Errorf("doing request: %w", err)
	}
	return nil
}
