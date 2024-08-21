package dxgo

import (
	"fmt"
	"time"
)

func (c *DXClient) ProjectDescribe(input ProjectDescribeInput, timeout time.Duration) (ProjectDescribeOutput, error) {
	output := new(ProjectDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.ID), input, output, timeout)
	if err != nil {
		return ProjectDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) ProjectNewFolder(input ProjectNewFolderInput, timeout time.Duration) (ProjectNewFolderOutput, error) {
	output := new(ProjectNewFolderOutput)
	err := c.DoInto(fmt.Sprintf("/%s/newFolder", input.ID), input, output, timeout)
	if err != nil {
		return ProjectNewFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) RemoveObjects(input RemoveObjectsInput, timeout time.Duration) (RemoveObjectsOutput, error) {
	output := new(RemoveObjectsOutput)
	err := c.DoInto(fmt.Sprintf("/%s/removeObjects", input.Project), input, output, timeout)
	if err != nil {
		return RemoveObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) RemoveFolder(input RemoveFolderInput, timeout time.Duration) (RemoveFolderOutput, error) {
	output := new(RemoveFolderOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.Project), input, output, timeout)
	if err != nil {
		return RemoveFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
