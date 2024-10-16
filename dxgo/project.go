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
	ID     string          `json:"id"`
	Fields map[string]bool `json:"fields,omitempty"`
}

type ProjectDescribeOutput struct {
	Name                 string         `json:"name"`
	Folders              []string       `json:"folders"`
	FileUploadParameters map[string]any `json:"fileUploadParameters"`
	Error                *ApiError      `json:"error"`
}

func (c *DXClient) ProjectDescribe(input ProjectDescribeInput, timeout time.Duration) (ProjectDescribeOutput, error) {
	output := new(ProjectDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.ID), input, output, timeout)
	if err != nil {
		return ProjectDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type ProjectNewFolderInput struct {
	ID     string `json:"id"`
	Folder string `json:"folder"`
}

type ProjectNewFolderOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) ProjectNewFolder(input ProjectNewFolderInput, timeout time.Duration) (ProjectNewFolderOutput, error) {
	output := new(ProjectNewFolderOutput)
	err := c.DoInto(fmt.Sprintf("/%s/newFolder", input.ID), input, output, timeout)
	if err != nil {
		return ProjectNewFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type RemoveObjectsInput struct {
	Project string   `json:"project"`
	Objects []string `json:"objects"`
	Force   bool     `json:"force"`
}

type RemoveObjectsOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) RemoveObjects(input RemoveObjectsInput, timeout time.Duration) (RemoveObjectsOutput, error) {
	output := new(RemoveObjectsOutput)
	err := c.DoInto(fmt.Sprintf("/%s/removeObjects", input.Project), input, output, timeout)
	if err != nil {
		return RemoveObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type RemoveFolderInput struct {
	Project string `json:"project"`
	Folder  string `json:"folder"`
	Force   bool   `json:"force"`
	Recurse bool   `json:"recurse"`
}

type RemoveFolderOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) RemoveFolder(input RemoveFolderInput, timeout time.Duration) (RemoveFolderOutput, error) {
	output := new(RemoveFolderOutput)
	err := c.DoInto(fmt.Sprintf("/%s/removeFolder", input.Project), input, output, timeout)
	if err != nil {
		return RemoveFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
