package dxgo

import (
	"encoding/json"
	"fmt"
)

func (c *DXClient) JobTerminate(input *JobTerminateInput) error {
	_, err := c.retryableRequest(fmt.Sprintf("/%s/terminate", input.ID), input)
	if err != nil {
		return err
	}
	return nil
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
