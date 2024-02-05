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
	Name                 string         `json:"name"`
	Folders              []string       `json:"folders"`
	FileUploadParameters map[string]any `json:"fileUploadParameters"`
	Error                *ApiError      `json:"error"`
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

type DXSortByOrdering string

const Ascending DXSortByOrdering = "ascending"
const Descending DXSortByOrdering = "descending"

type FindDataObjectsSortBy struct {
	Field    string           `json:"field"`
	Ordering DXSortByOrdering `json:"ordering"`
}

type FindDataObjectsInput struct {
	Name     string                 `json:"name,omitempty"`
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

type SystemRequirementsValue struct {
	InstanceType string `json:"instanceType"`
}

type SystemRequirements map[string]*SystemRequirementsValue

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
	ID    string    `json:"id"`
	Error *ApiError `json:"error"`
}

type FileUploadInput struct {
	ID    string `json:"id"`
	Size  int    `json:"size"`
	MD5   string `json:"md5"`
	Index int    `json:"index"`
}

type FileUploadOutput struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Error   *ApiError         `json:"error"`
}

type FindExecutionsInput struct {
	Project   string `json:"project,omitempty"`
	Class     string `json:"class,omitempty"`
	State     string `json:"state,omitempty"`
	OriginJob string `json:"originJob,omitempty"`
	Starting  string `json:"starting,omitempty"`
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

type FileCloseInput struct {
	ID string `json:"id"`
}

type FileCloseOutput struct {
	Error *ApiError `json:"error"`
}

type FileDescribeInput struct {
	ID string `json:"id"`
}

type FileDescribeOutput struct {
	Folder string    `json:"folder"`
	State  string    `json:"state"`
	Name   string    `json:"name"`
	Error  *ApiError `json:"error"`
}

type FindProjectsInput struct {
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
