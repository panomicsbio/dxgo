package dxgo

import (
	"fmt"
	"time"
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

func (c *DXClient) NewProject(input NewProjectInput, timeout time.Duration) (NewProjectOutput, error) {
	output := new(NewProjectOutput)
	err := c.DoInto("/project/new", input, output, timeout)
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
	Folders              []string       `json:"folders"`
	FileUploadParameters map[string]any `json:"fileUploadParameters"`
	Error                *ApiError      `json:"error"`
}

func (c *DXClient) ProjectDescribe(projectID string, input ProjectDescribeInput, timeout time.Duration) (ProjectDescribeOutput, error) {
	output := new(ProjectDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", projectID), input, output, timeout)
	if err != nil {
		return ProjectDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type ProjectNewFolderInput struct {
	Folder string `json:"folder"`
}

type ProjectNewFolderOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) ProjectNewFolder(projectID string, input ProjectNewFolderInput, timeout time.Duration) (ProjectNewFolderOutput, error) {
	output := new(ProjectNewFolderOutput)
	err := c.DoInto(fmt.Sprintf("/%s/newFolder", projectID), input, output, timeout)
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

func (c *DXClient) RemoveObjects(projectID string, input RemoveObjectsInput, timeout time.Duration) (RemoveObjectsOutput, error) {
	output := new(RemoveObjectsOutput)
	err := c.DoInto(fmt.Sprintf("/%s/removeObjects", projectID), input, output, timeout)
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

func (c *DXClient) RemoveFolder(projectID string, input RemoveFolderInput, timeout time.Duration) (RemoveFolderOutput, error) {
	output := new(RemoveFolderOutput)
	err := c.DoInto(fmt.Sprintf("/%s/removeFolder", projectID), input, output, timeout)
	if err != nil {
		return RemoveFolderOutput{}, fmt.Errorf("doing request: %w", err)
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

func (c *DXClient) DestroyProject(projectID string, input DestroyProjectInput, timeout time.Duration) (DestroyProjectOutput, error) {
	output := new(DestroyProjectOutput)

	err := c.DoInto(fmt.Sprintf("/%s/destroy", projectID), input, output, timeout)
	if err != nil {
		return DestroyProjectOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
