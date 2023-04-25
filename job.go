package dxgo

import "fmt"

func (c *DXClient) JobTerminate(input *JobTerminateInput) error {
	_, err := c.retryableRequest(fmt.Sprintf("/%s/terminate", input.ID), input)
	if err != nil {
		return err
	}
	return nil
}
