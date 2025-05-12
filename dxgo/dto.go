package dxgo

import (
	"fmt"
	"strings"
)

type ApiError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Details any    `json:"details"`
}

func (e *ApiError) String() string {
	var details = make([]string, 0)

	switch v := e.Details.(type) {
	case map[string]any:
		for k, v := range v {
			details = append(details, fmt.Sprintf("%s=%s", k, v))
		}
	case []any:
		for _, v := range v {
			details = append(details, fmt.Sprintf("%v", v))
		}
	}

	return fmt.Sprintf("%s - %s, Details: %s", e.Type, e.Message, strings.Join(details, ", "))
}

func (e *ApiError) Error() string {
	return e.String()
}

type DXAssetType string

const (
	DXAssetTypeApplet   DXAssetType = "applet"
	DXAssetTypeWorkflow DXAssetType = "workflow"
	DXAssetTypeFile     DXAssetType = "file"
)

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
	NotSpecified *TimeoutPolicyValue `json:"*,omitempty"`
	Main         *TimeoutPolicyValue `json:"main,omitempty"`
	Default      *TimeoutPolicyValue `json:"default,omitempty"`
}

type TimeoutPolicyValue struct {
	Days    *int `json:"days,omitempty"`
	Hours   *int `json:"hours,omitempty"`
	Minutes *int `json:"minutes,omitempty"`
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

const SystemRequirementsAllEntryPoints = "*"

// SystemRequirements represents specific computational resource requirements
type SystemRequirements map[string]SystemRequirementsEntry

type SystemRequirementsEntry struct {
	InstanceType string       `json:"instanceType,omitempty"`
	ClusterSpec  *ClusterSpec `json:"clusterSpec,omitempty"`
	FpgaDriver   *string      `json:"fpgaDriver,omitempty"`
	NvidiaDriver *string      `json:"nvidiaDriver,omitempty"`
}

// ClusterSpec indicates that this job requires a cluster of instances rather than just a single worker node
type ClusterSpec struct {
	// The type of cluster, supported values are: dxspark, apachespark, and generic
	Type *string `json:"type,omitempty"`
	// Requested version for dxspark or apachespark clusters. Supported values are: 2.4.4, 3.2.3
	Version              *string `json:"version,omitempty"`
	InitialInstanceCount *string `json:"initialInstanceCount,omitempty"`
	Ports                *string `json:"ports,omitempty"`
	BootstrapScript      *string `json:"bootstrapScript,omitempty"`
}
