package dxgo

import (
	"encoding/json"
	"fmt"
)

func (c *DXClient) JobTerminate(input *JobTerminateInput) (*JobTerminateOutput, error) {
	data, err := c.retryableRequest(fmt.Sprintf("/%s/terminate", input.ID), input)
	if err != nil {
		return nil, err
	}
	output := new(JobTerminateOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c *DXClient) JobDescribe(input *JobDescribeInput) (*JobDescribeOutput, error) {
	data, err := c.retryableRequest(fmt.Sprintf("/%s/describe", input.ID), input)
	if err != nil {
		return nil, err
	}
	output := new(JobDescribeOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
