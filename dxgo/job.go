package dxgo

import (
	"fmt"
)

func (c *DXClient) JobTerminate(input JobTerminateInput) (JobTerminateOutput, error) {
	output := new(JobTerminateOutput)
	err := c.DoInto(fmt.Sprintf("/%s/terminate", input.ID), input, output)
	if err != nil {
		return JobTerminateOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) JobDescribe(input JobDescribeInput) (JobDescribeOutput, error) {
	output := new(JobDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.ID), input, output)
	if err != nil {
		return JobDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
