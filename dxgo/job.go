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
	ID string `json:"id"`
}

type JobDescribeOutput struct {
	State string    `json:"state"`
	Error *ApiError `json:"error"`
}

func (c *DXClient) JobDescribe(ctx context.Context, input JobDescribeInput) (JobDescribeOutput, error) {
	output := new(JobDescribeOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", input.ID), input, output)
	if err != nil {
		return JobDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
