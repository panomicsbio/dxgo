package dxgo

import (
	"context"
	"fmt"
)

type JobTerminateInput struct {
	ID string `json:"id"`
}

type JobTerminateOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) JobTerminate(ctx context.Context, input JobTerminateInput) (JobTerminateOutput, error) {
	output := new(JobTerminateOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/terminate", input.ID), input, output)
	if err != nil {
		return JobTerminateOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type JobDescribeInput struct {
	ID     string          `json:"id"`
	Fields map[string]bool `json:"fields"`
}

type JobDescribeOutput struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	State  string         `json:"state"`
	Input  map[string]any `json:"input"`
	Output map[string]any `json:"output"`
	Error  *ApiError      `json:"error"`
}

func (c *DXClient) JobDescribe(ctx context.Context, input JobDescribeInput) (JobDescribeOutput, error) {
	output := new(JobDescribeOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", input.ID), input, output)
	if err != nil {
		return JobDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
