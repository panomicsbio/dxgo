package dxgo

import (
	"context"
	"fmt"
)

func (c *DXClient) AppletRun(ctx context.Context, input AppletRunInput) (AppletRunOutput, error) {
	output := new(AppletRunOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/run", input.ID), input, output)
	if err != nil {
		return AppletRunOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
