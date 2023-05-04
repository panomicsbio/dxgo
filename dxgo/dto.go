package dxgo

import (
	"fmt"
	"strings"
)

type ApiError struct {
	Type    string            `json:"type"`
	Message string            `json:"message"`
	Details map[string]string `json:"details"`
}

func (e *ApiError) String() string {
	var details = make([]string, len(e.Details))
	idx := 0
	for k, v := range e.Details {
		details[idx] = fmt.Sprintf("%s=%s", k, v)
		idx++
	}
	return fmt.Sprintf("%s - %s, Details: %s", e.Type, e.Message, strings.Join(details, ", "))
}

type ProjectDescribeInput struct {
	ID     string          `json:"id"`
	Fields map[string]bool `json:"fields,omitempty"`
}

type ProjectDescribeOutput struct {
	Name                 string                 `json:"name"`
	Folders              []string               `json:"folders"`
	FileUploadParameters map[string]interface{} `json:"fileUploadParameters"`
	Error                *ApiError              `json:"error"`
}

type ProjectNewFolderInput struct {
	ID     string `json:"id"`
	Folder string `json:"folder"`
}

type ProjectNewFolderOutput struct {
	Error *ApiError `json:"error"`
}

type DXAssetType string

type FindDataObjectsScope struct {
	Project *string `json:"project"`
	Folder  *string `json:"folder"`
	Recurse *bool   `json:"recurse"`
}

type FindDataObjectsInput struct {
	Name     string                 `json:"name,omitempty"`
	Class    DXAssetType            `json:"class"`
	Scope    *FindDataObjectsScope  `json:"scope"`
	Describe bool                   `json:"describe"`
	Starting map[string]interface{} `json:"starting,omitempty"`
}

type FindDataObjectsResult struct {
	Project  string                 `json:"project"`
	ID       string                 `json:"id"`
	Describe map[string]interface{} `json:"describe"`
}

type FindDataObjectsOutput struct {
	Results []*FindDataObjectsResult `json:"results"`
	Next    map[string]interface{}   `json:"next"`
	Error   *ApiError                `json:"error"`
}

type SystemRequirementsValue struct {
	InstanceType string `json:"instanceType"`
}

type SystemRequirements map[string]*SystemRequirementsValue

type AppletRunInput struct {
	ID                 string                 `json:"id"`
	Project            *string                `json:"project,omitempty"`
	Input              map[string]interface{} `json:"input"`
	SystemRequirements *SystemRequirements    `json:"systemRequirements"`
	Properties         map[string]string      `json:"properties"`
	Detach             bool                   `json:"detach"`
	HeadJobOnDemand    bool                   `json:"headJobOnDemand"`
}

type AppletRunOutput struct {
	ID    string    `json:"id"`
	Error *ApiError `json:"error"`
}

type JobTerminateInput struct {
	ID string `json:"id"`
}

type JobTerminateOutput struct {
	Error *ApiError `json:"error"`
}

type JobDescribeInput struct {
	ID string `json:"id"`
}

type JobDescribeOutput struct {
	State string    `json:"state"`
	Error *ApiError `json:"error"`
}

type FileDownloadInput struct {
	ID               string  `json:"id"`
	Duration         *int    `json:"duration,omitempty"`
	Filename         *string `json:"filename,omitempty"`
	Project          string  `json:"project"`
	Preauthenticated *bool   `json:"preauthenticated,omitempty"`
	StickyIP         *bool   `json:"stickyIP,omitempty"`
}

type FileDownloadOutput struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Error   *ApiError         `json:"error"`
}

type RemoveObjectsInput struct {
	Project string   `json:"project"`
	Objects []string `json:"objects"`
	Force   bool     `json:"force"`
}

type RemoveObjectsOutput struct {
	Error *ApiError `json:"error"`
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

type FileNewInput struct {
	Project string `json:"project"`
	Folder  string `json:"folder"`
	Parent  bool   `json:"parents"`
	Name    string `json:"name"`
}

type FileNewOutput struct {
	ID string `json:"id"`
}

type FileUploadInput struct {
	ID    string `json:"id"`
	Size  uint   `json:"size"`
	MD5   string `json:"md5"`
	Index uint   `json:"index"`
}

type FileUploadOutput struct {
	URL string `json:"url"`
}
