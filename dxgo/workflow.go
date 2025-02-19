package dxgo

import (
	"context"
	"fmt"
)

// WorkflowNewInput represents the input for creating a new workflow
// Project is required and must be of the form "project-xxxx"
// Name, if provided, may only contain the characters [a-zA-Z0-9._-]
// Types, if provided, must be from the list of supported workflow types
type WorkflowNewInput struct {
	Project      string            `json:"project"`                // Project in which to create the workflow
	Name         *string           `json:"name,omitempty"`         // Name of the workflow
	Title        *string           `json:"title,omitempty"`        // Title of the workflow
	Summary      *string           `json:"summary,omitempty"`      // A short description of the workflow
	Description  *string           `json:"description,omitempty"`  // A longer description of the workflow
	OutputFolder *string           `json:"outputFolder,omitempty"` // The default output folder for the workflow
	Tags         []string          `json:"tags,omitempty"`         // Tags to associate with the workflow
	Types        []string          `json:"types,omitempty"`        // Types of the workflow
	Hidden       *bool             `json:"hidden,omitempty"`       // Whether the workflow should be hidden
	Properties   map[string]string `json:"properties,omitempty"`   // Properties to associate with the workflow
	Details      map[string]any    `json:"details,omitempty"`      // Details about the workflow
	Inputs       []map[string]any  `json:"inputs,omitempty"`       // Input for the workflow described at https://documentation.dnanexus.com/developer/api/running-analyses/io-and-run-specifications#input-specification.
	Outputs      []map[string]any  `json:"outputs,omitempty"`      // Output for the workflow described at https://documentation.dnanexus.com/developer/api/running-analyses/io-and-run-specifications#output-specification.
	Folder       *string           `json:"folder,omitempty"`       // Folder path where the workflow should be created
	Parents      *bool             `json:"parents,omitempty"`      // Whether to create parent folders if they don't exist
	Stages       []WorkflowStage   `json:"stages,omitempty"`       // Initial stages of the workflow
	IgnoreReuse  []string          `json:"ignoreReuse,omitempty"`  // Stage IDs for which to ignore reuse
	Nonce        *string           `json:"nonce,omitempty"`        // Idempotency key
}

// WorkflowStage represents a stage in a workflow
// ID and Executable are required fields
type WorkflowStage struct {
	ID                 string             `json:"id"`                           // Unique identifier for the stage within the workflow
	Executable         string             `json:"executable"`                   // The app or applet to run in this stage
	Name               *string            `json:"name,omitempty"`               // Optional name for the stage
	Folder             *string            `json:"folder,omitempty"`             // Output folder for this stage
	Input              map[string]any     `json:"input,omitempty"`              // Input fields for the executable
	SystemRequirements SystemRequirements `json:"systemRequirements,omitempty"` // Computational resources required
	ExecutionPolicy    *ExecutionPolicy   `json:"executionPolicy,omitempty"`    // Policy for handling execution failures
}

// WorkflowNewOutput represents the output from creating a new workflow
type WorkflowNewOutput struct {
	ID          string    `json:"id"`
	EditVersion int       `json:"editVersion"`
	Error       *ApiError `json:"error"`
}

// WorkflowAddStageInput represents the input for adding a stage to a workflow
// EditVersion and Executable are required fields
type WorkflowAddStageInput struct {
	EditVersion        int                `json:"editVersion"`                  // Current version of the workflow
	ID                 *string            `json:"id,omitempty"`                 // Optional ID for the new stage
	Executable         string             `json:"executable"`                   // The app or applet to run in this stage
	Name               *string            `json:"name,omitempty"`               // Optional name for the stage
	Folder             *string            `json:"folder,omitempty"`             // Output folder for this stage
	Input              map[string]any     `json:"input,omitempty"`              // Input fields for the executable
	ExecutionPolicy    *ExecutionPolicy   `json:"executionPolicy,omitempty"`    // Policy for handling execution failures
	SystemRequirements SystemRequirements `json:"systemRequirements,omitempty"` // Computational resources required
}

// WorkflowAddStageOutput represents the output from adding a stage
type WorkflowAddStageOutput struct {
	ID          string    `json:"id"`
	EditVersion int       `json:"editVersion"`
	Stage       string    `json:"stage"`
	Error       *ApiError `json:"error"`
}

// WorkflowDescribeOutput represents the output from describing a workflow
type WorkflowDescribeOutput struct {
	ID          string            `json:"id"`
	Project     string            `json:"project"`
	Name        string            `json:"name"`
	State       string            `json:"state"`
	Stages      []WorkflowStage   `json:"stages"`
	EditVersion int               `json:"editVersion"`
	Created     int64             `json:"created"`
	Modified    int64             `json:"modified"`
	Properties  map[string]string `json:"properties"`
	Tags        []string          `json:"tags"`
	Error       *ApiError         `json:"error"`
}

// WorkflowRunInput represents the input for running a workflow
// Input is required and must contain all required inputs for each stage
type WorkflowRunInput struct {
	Name                      *string            `json:"name,omitempty"`                      // Name for the analysis
	Input                     map[string]any     `json:"input"`                               // Input for the analysis is launched with workflow run
	Project                   *string            `json:"project,omitempty"`                   // Project context for the run
	Folder                    *string            `json:"folder,omitempty"`                    // Default output folder
	StageFolders              map[string]string  `json:"stageFolders,omitempty"`              // Per-stage output folders
	Details                   map[string]any     `json:"details,omitempty"`                   // Details about the analysis
	SystemRequirements        SystemRequirements `json:"systemRequirements,omitempty"`        // Default system requirements
	ExecutionPolicy           *ExecutionPolicy   `json:"executionPolicy,omitempty"`           // Default execution policy
	DelayWorkspaceDestruction *bool              `json:"delayWorkspaceDestruction,omitempty"` // Whether to delay workspace destruction
	RerunStages               []string           `json:"rerunStages,omitempty"`               // Stages to force rerun
	IgnoreReuse               []string           `json:"ignoreReuse,omitempty"`               // Stage IDs for which to ignore reuse
	Detach                    *bool              `json:"detach,omitempty"`                    // Whether to detach the job after launching
}

// WorkflowRunOutput represents the output from running a workflow
type WorkflowRunOutput struct {
	ID     string    `json:"id"`
	Stages []string  `json:"stages"`
	Error  *ApiError `json:"error"`
}

// ExecutionPolicy represents the execution policy for a workflow stage
// All fields are optional and provide fine-grained control over job restart behavior
type ExecutionPolicy struct {
	RestartOn               map[string]int `json:"restartOn,omitempty"`               // Mapping of error types to number of restart attempts
	MaxRestarts             *int           `json:"maxRestarts,omitempty"`             // Maximum number of restarts regardless of error type
	OnNonRestartableFailure *string        `json:"onNonRestartableFailure,omitempty"` // Action to take on non-restartable failure
}

// WorkflowUpdateInput represents the input for updating a workflow
type WorkflowUpdateInput struct {
	EditVersion  int              `json:"editVersion"`            // Current version of the workflow
	Title        *string          `json:"title,omitempty"`        // Title of the workflow
	Summary      *string          `json:"summary,omitempty"`      // A short description of the workflow
	Description  *string          `json:"description,omitempty"`  // A longer description of the workflow
	OutputFolder *string          `json:"outputFolder,omitempty"` // The default output folder for the workflow
	Inputs       []map[string]any `json:"inputs,omitempty"`       // Input for the workflow described at https://documentation.dnanexus.com/developer/api/running-analyses/io-and-run-specifications#input-specification.
	Outputs      []map[string]any `json:"outputs,omitempty"`      // Output for the workflow described at https://documentation.dnanexus.com/developer/api/running-analyses/io-and-run-specifications#output-specification.
	Stages       []WorkflowStage  `json:"stages,omitempty"`       // Initial stages of the workflow
}

// WorkflowUpdateOutput represents the output from updating a workflow
type WorkflowUpdateOutput struct {
	ID          string    `json:"id"`
	EditVersion int       `json:"editVersion"`
	Error       *ApiError `json:"error"`
}

func (c *DXClient) WorkflowNew(ctx context.Context, input WorkflowNewInput) (WorkflowNewOutput, error) {
	output := new(WorkflowNewOutput)
	err := c.DoInto(ctx, "/workflow/new", input, output)
	if err != nil {
		return WorkflowNewOutput{}, fmt.Errorf("creating workflow: %w", err)
	}

	return *output, nil
}

func (c *DXClient) WorkflowAddStage(ctx context.Context, workflowID string, input WorkflowAddStageInput) (WorkflowAddStageOutput, error) {
	output := new(WorkflowAddStageOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/addStage", workflowID), input, output)
	if err != nil {
		return WorkflowAddStageOutput{}, fmt.Errorf("adding workflow stage: %w", err)
	}

	return *output, nil
}

func (c *DXClient) WorkflowDescribe(ctx context.Context, workflowID string) (WorkflowDescribeOutput, error) {
	output := new(WorkflowDescribeOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", workflowID), nil, output)
	if err != nil {
		return WorkflowDescribeOutput{}, fmt.Errorf("describing workflow: %w", err)
	}

	return *output, nil
}

func (c *DXClient) WorkflowRun(ctx context.Context, workflowID string, input WorkflowRunInput) (WorkflowRunOutput, error) {
	output := new(WorkflowRunOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/run", workflowID), input, output)
	if err != nil {
		return WorkflowRunOutput{}, fmt.Errorf("running workflow: %w", err)
	}

	return *output, nil
}

func (c *DXClient) WorkflowUpdate(ctx context.Context, workflowID string, input WorkflowUpdateInput) (WorkflowUpdateOutput, error) {
	output := new(WorkflowUpdateOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/update", workflowID), input, output)
	if err != nil {
		return WorkflowUpdateOutput{}, fmt.Errorf("updating workflow: %w", err)
	}

	return *output, nil
}
