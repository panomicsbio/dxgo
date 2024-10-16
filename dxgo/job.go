package dxgo

import (
	"fmt"
	"time"
)

type JobTerminateInput struct {
	ID string `json:"id"`
}

type JobTerminateOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) JobTerminate(input JobTerminateInput, timeout time.Duration) (JobTerminateOutput, error) {
	output := new(JobTerminateOutput)
	err := c.DoInto(fmt.Sprintf("/%s/terminate", input.ID), input, output, timeout)
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

func (c *DXClient) JobDescribe(input JobDescribeInput, timeout time.Duration) (JobDescribeOutput, error) {
	output := new(JobDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.ID), input, output, timeout)
	if err != nil {
		return JobDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
