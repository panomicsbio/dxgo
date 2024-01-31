package dxgo

import (
	"fmt"
)

func (c *DXClient) AppletRun(input AppletRunInput) (AppletRunOutput, error) {
	output := new(AppletRunOutput)
	err := c.DoInto(fmt.Sprintf("/%s/run", input.ID), input, output)
	if err != nil {
		return AppletRunOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
