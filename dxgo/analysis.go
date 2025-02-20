package dxgo

import (
	"context"
	"fmt"
)

// AnalysisDescribeInput represents the input parameters for describing an analysis
type AnalysisDescribeInput struct {
	DefaultFields *bool           `json:"defaultFields,omitempty"` // Whether to include default fields
	Fields        map[string]bool `json:"fields,omitempty"`        // Specific fields to include/exclude
}

// AnalysisDescribeOutput represents the output from describing an analysis
type AnalysisDescribeOutput struct {
	ID                        string            `json:"id"`                        // The object ID (analysis-xxxx)
	Class                     string            `json:"class"`                     // Always "analysis"
	Name                      string            `json:"name"`                      // Analysis name
	Executable                string            `json:"executable"`                // ID of workflow that was run
	ExecutableName            string            `json:"executableName"`            // Name of workflow that was run
	Created                   int64             `json:"created"`                   // Creation timestamp
	Modified                  int64             `json:"modified"`                  // Last update timestamp
	BillTo                    string            `json:"billTo"`                    // Billing account ID
	Project                   string            `json:"project"`                   // Project ID
	Folder                    string            `json:"folder"`                    // Output folder path
	RootExecution             string            `json:"rootExecution"`             // Root job/analysis ID
	ParentJob                 *string           `json:"parentJob"`                 // Parent job ID if applicable
	ParentJobTry              *int              `json:"parentJobTry"`              // Parent job try number
	ParentAnalysis            *string           `json:"parentAnalysis"`            // Parent analysis ID if applicable
	DetachedFrom              *string           `json:"detachedFrom"`              // Job ID this was detached from
	DetachedFromTry           *int              `json:"detachedFromTry"`           // Detached job try number
	Analysis                  *string           `json:"analysis"`                  // Analysis ID if part of stage
	Stage                     *string           `json:"stage"`                     // Stage ID if part of analysis
	Workflow                  WorkflowMetadata  `json:"workflow"`                  // Workflow metadata
	Stages                    []AnalysisStage   `json:"stages"`                    // Stage execution metadata
	State                     string            `json:"state"`                     // Analysis state
	Workspace                 string            `json:"workspace"`                 // Temporary workspace ID
	LaunchedBy                string            `json:"launchedBy"`                // User ID who launched root execution
	Tags                      []string          `json:"tags"`                      // Associated tags
	Properties                map[string]string `json:"properties"`                // Associated properties
	Details                   map[string]any    `json:"details"`                   // Stored JSON details
	RunInput                  map[string]any    `json:"runInput"`                  // Original API call input
	OriginalInput             map[string]any    `json:"originalInput"`             // Effective analysis input
	Input                     map[string]any    `json:"input"`                     // Same as OriginalInput
	Output                    map[string]any    `json:"output"`                    // Available outputs
	DelayWorkspaceDestruction bool              `json:"delayWorkspaceDestruction"` // Keep workspace for 3 days
	IgnoreReuse               []string          `json:"ignoreReuse"`               // Stage IDs to ignore reuse
	PreserveJobOutputs        map[string]string `json:"preserveJobOutputs"`        // Job output preservation settings
	DetailedJobMetrics        bool              `json:"detailedJobMetrics"`        // Whether detailed metrics enabled
	CostLimit                 *float64          `json:"costLimit"`                 // Root execution cost limit
	Rank                      int               `json:"rank"`                      // Analysis rank
	Error                     *ApiError         `json:"error"`                     // Error information if request failed
}

// WorkflowMetadata represents metadata about the workflow that was run
type WorkflowMetadata struct {
	ID              string           `json:"id"`                        // Workflow ID
	Name            string           `json:"name"`                      // Workflow name
	Inputs          []map[string]any `json:"inputs"`                    // Input specification
	Outputs         []map[string]any `json:"outputs"`                   // Output specification
	Stages          []WorkflowStage  `json:"stages"`                    // Stage metadata
	EditVersion     int              `json:"editVersion"`               // Edit version at run time
	InitializedFrom map[string]any   `json:"initializedFrom,omitempty"` // Initialization source
}

// AnalysisStage represents metadata about a stage execution
type AnalysisStage struct {
	ID        string        `json:"id"`        // Stage ID
	Execution ExecutionInfo `json:"execution"` // Execution information
}

// ExecutionInfo represents information about stage execution
type ExecutionInfo struct {
	ID      string         `json:"id"`      // Execution ID
	Details map[string]any `json:"details"` // Additional execution details
}

// AnalysisTerminateOutput represents the output from terminating an analysis
type AnalysisTerminateOutput struct {
	ID    string    `json:"id"`    // ID of the terminated analysis
	Error *ApiError `json:"error"` // Error information if request failed
}

func (c *DXClient) AnalysisDescribe(ctx context.Context, analysisID string, input *AnalysisDescribeInput) (AnalysisDescribeOutput, error) {
	output := new(AnalysisDescribeOutput)

	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", analysisID), input, output)
	if err != nil {
		return AnalysisDescribeOutput{}, fmt.Errorf("describing analysis: %w", err)
	}

	return *output, nil
}

func (c *DXClient) AnalysisTerminate(ctx context.Context, analysisID string) (AnalysisTerminateOutput, error) {
	output := new(AnalysisTerminateOutput)

	err := c.DoInto(ctx, fmt.Sprintf("/%s/terminate", analysisID), nil, output)
	if err != nil {
		return AnalysisTerminateOutput{}, fmt.Errorf("terminating analysis: %w", err)
	}

	return *output, nil
}
