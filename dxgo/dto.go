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
