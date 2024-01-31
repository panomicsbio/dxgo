package dxgo

import (
	"fmt"
)

func (c *DXClient) ProjectDescribe(input ProjectDescribeInput) (ProjectDescribeOutput, error) {
	output := new(ProjectDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.ID), input, output)
	if err != nil {
		return ProjectDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) ProjectNewFolder(input ProjectNewFolderInput) (ProjectNewFolderOutput, error) {
	output := new(ProjectNewFolderOutput)
	err := c.DoInto(fmt.Sprintf("/%s/newFolder", input.ID), input, output)
	if err != nil {
		return ProjectNewFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) RemoveObjects(input RemoveObjectsInput) (RemoveObjectsOutput, error) {
	output := new(RemoveObjectsOutput)
	err := c.DoInto(fmt.Sprintf("/%s/removeObjects", input.Project), input, output)
	if err != nil {
		return RemoveObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) RemoveFolder(input RemoveFolderInput) (RemoveFolderOutput, error) {
	output := new(RemoveFolderOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.Project), input, output)
	if err != nil {
		return RemoveFolderOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
