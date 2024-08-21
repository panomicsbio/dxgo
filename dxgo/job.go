package dxgo

import (
	"fmt"
	"time"
)

func (c *DXClient) JobTerminate(input JobTerminateInput, timeout time.Duration) (JobTerminateOutput, error) {
	output := new(JobTerminateOutput)
	err := c.DoInto(fmt.Sprintf("/%s/terminate", input.ID), input, output, timeout)
	if err != nil {
		return JobTerminateOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) JobDescribe(input JobDescribeInput, timeout time.Duration) (JobDescribeOutput, error) {
	output := new(JobDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.ID), input, output, timeout)
	if err != nil {
		return JobDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
