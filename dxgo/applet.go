package dxgo

import (
	"fmt"
	"time"
)

func (c *DXClient) AppletRun(input AppletRunInput, timeout time.Duration) (AppletRunOutput, error) {
	output := new(AppletRunOutput)
	err := c.DoInto(fmt.Sprintf("/%s/run", input.ID), input, output, timeout)
	if err != nil {
		return AppletRunOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
