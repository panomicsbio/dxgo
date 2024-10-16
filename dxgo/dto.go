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

func (e *ApiError) Error() string {
	return e.String()
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

type SystemRequirementsValue struct {
	InstanceType string `json:"instanceType"`
}

type SystemRequirements map[string]*SystemRequirementsValue
