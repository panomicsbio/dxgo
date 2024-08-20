package dxgo

import (
	"context"
	"fmt"
)

func (c *DXClient) FindDataObjects(ctx context.Context, input FindDataObjectsInput) (FindDataObjectsOutput, error) {
	output := new(FindDataObjectsOutput)
	err := c.DoInto(ctx, "/system/findDataObjects", input, output)
	if err != nil {
		return FindDataObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}

func (c *DXClient) FindExecutions(ctx context.Context, input FindExecutionsInput) (FindExecutionsOutput, error) {
	output := new(FindExecutionsOutput)
	err := c.DoInto(ctx, "/system/findExecutions", input, output)
	if err != nil {
		return FindExecutionsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}

func (c *DXClient) FindProjects(ctx context.Context, input FindProjectsInput) (FindProjectsOutput, error) {
	output := new(FindProjectsOutput)
	err := c.DoInto(ctx, "/system/findProjects", input, output)
	if err != nil {
		return FindProjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}
