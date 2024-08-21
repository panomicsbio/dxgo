package dxgo

import (
	"fmt"
	"time"
)

func (c *DXClient) FindDataObjects(input FindDataObjectsInput, timeout time.Duration) (FindDataObjectsOutput, error) {
	output := new(FindDataObjectsOutput)
	err := c.DoInto("/system/findDataObjects", input, output, timeout)
	if err != nil {
		return FindDataObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}

func (c *DXClient) FindExecutions(input FindExecutionsInput, timeout time.Duration) (FindExecutionsOutput, error) {
	output := new(FindExecutionsOutput)
	err := c.DoInto("/system/findExecutions", input, output, timeout)
	if err != nil {
		return FindExecutionsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}

func (c *DXClient) FindProjects(input FindProjectsInput, timeout time.Duration) (FindProjectsOutput, error) {
	output := new(FindProjectsOutput)
	err := c.DoInto("/system/findProjects", input, output, timeout)
	if err != nil {
		return FindProjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}
