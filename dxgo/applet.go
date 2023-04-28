package dxgo

import (
	"encoding/json"
	"fmt"
)

func (c *DXClient) AppletRun(input *AppletRunInput) (*AppletRunOutput, error) {
	data, err := c.retryableRequest(fmt.Sprintf("/%s/run", input.ID), input)
	if err != nil {
		return nil, err
	}
	output := new(AppletRunOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
