package dxgo

import (
	"context"
	"fmt"
)

func (c *DXClient) ProjectDescribe(ctx context.Context, input ProjectDescribeInput) (ProjectDescribeOutput, error) {
	output := new(ProjectDescribeOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", input.ID), input, output)
	if err != nil {
		return ProjectDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) ProjectNewFolder(ctx context.Context, input ProjectNewFolderInput) (ProjectNewFolderOutput, error) {
	output := new(ProjectNewFolderOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/newFolder", input.ID), input, output)
	if err != nil {
		return ProjectNewFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) RemoveObjects(ctx context.Context, input RemoveObjectsInput) (RemoveObjectsOutput, error) {
	output := new(RemoveObjectsOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/removeObjects", input.Project), input, output)
	if err != nil {
		return RemoveObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) RemoveFolder(ctx context.Context, input RemoveFolderInput) (RemoveFolderOutput, error) {
	output := new(RemoveFolderOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", input.Project), input, output)
	if err != nil {
		return RemoveFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
