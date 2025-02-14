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

type CreatedBy struct {
	User       string `json:"user"`
	Job        string `json:"job,omitempty"`
	Executable string `json:"executable,omitempty"`
}

// RunSpec represents the execution specification for an applet or app
type RunSpec struct {
	Code                   string             `json:"code"`
	Interpreter            string             `json:"interpreter"`
	Distribution           string             `json:"distribution,omitempty"`
	Release                string             `json:"release,omitempty"`
	Version                string             `json:"version,omitempty"`
	RestartableEntryPoints bool               `json:"restartableEntryPoints,omitempty"`
	SystemRequirements     SystemRequirements `json:"systemRequirements,omitempty"`
	TimeoutPolicy          *TimeoutPolicy     `json:"timeoutPolicy,omitempty"`
}

// TimeoutPolicy represents the timeout configuration for an execution
type TimeoutPolicy struct {
	Hours   *int `json:"*,omitempty"`
	Main    *int `json:"main,omitempty"`
	Default *int `json:"default,omitempty"`
}

// IOSpec represents input/output specifications for an applet or app
type IOSpec struct {
	Name        string   `json:"name"`
	Class       string   `json:"class"`
	Optional    bool     `json:"optional,omitempty"`
	Types       []string `json:"types,omitempty"`
	Patterns    []string `json:"patterns,omitempty"`
	Help        string   `json:"help,omitempty"`
	Label       string   `json:"label,omitempty"`
	Default     any      `json:"default,omitempty"`
	Choices     []any    `json:"choices,omitempty"`
	Group       string   `json:"group,omitempty"`
	Suggestions []any    `json:"suggestions,omitempty"`
}

// Access represents access control settings for a resource
type Access struct {
	Network       []string          `json:"network"`
	Project       string            `json:"project"`
	AllProjects   string            `json:"allProjects"`
	Developer     bool              `json:"developer"`
	ProjectAccess string            `json:"projectAccess"`
	Permissions   map[string]string `json:"permissions,omitempty"`
}

// SystemRequirements represents specific computational resource requirements
type SystemRequirements struct {
	InstanceType string                                  `json:"instanceType,omitempty"`
	ClusterSpec  string                                  `json:"clusterSpec,omitempty"`
	EntryPoints  map[string]SystemRequirementsEntryPoint `json:"entryPoints,omitempty"`
}

// SystemRequirementEntryPoint represents requirements for a specific entry point
type SystemRequirementsEntryPoint struct {
	InstanceType string `json:"instanceType,omitempty"`
	ClusterSpec  string `json:"clusterSpec,omitempty"`
}
