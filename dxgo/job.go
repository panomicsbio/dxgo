package dxgo

import (
	"context"
	"fmt"
)

func (c *DXClient) JobTerminate(ctx context.Context, input JobTerminateInput) (JobTerminateOutput, error) {
	output := new(JobTerminateOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/terminate", input.ID), input, output)
	if err != nil {
		return JobTerminateOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) JobDescribe(ctx context.Context, input JobDescribeInput) (JobDescribeOutput, error) {
	output := new(JobDescribeOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", input.ID), input, output)
	if err != nil {
		return JobDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
