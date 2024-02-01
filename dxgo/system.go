package dxgo

import "fmt"

func (c *DXClient) FindDataObjects(input FindDataObjectsInput) (FindDataObjectsOutput, error) {
	output := new(FindDataObjectsOutput)
	err := c.DoInto("/system/findDataObjects", input, output)
	if err != nil {
		return FindDataObjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}

func (c *DXClient) FindExecutions(input FindExecutionsInput) (FindExecutionsOutput, error) {
	output := new(FindExecutionsOutput)
	err := c.DoInto("/system/findExecutions", input, output)
	if err != nil {
		return FindExecutionsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}

func (c *DXClient) FindProjects(input FindProjectsInput) (FindProjectsOutput, error) {
	output := new(FindProjectsOutput)
	err := c.DoInto("/system/findProjects", input, output)
	if err != nil {
		return FindProjectsOutput{}, fmt.Errorf("doing request: %w", err)
	}
	return *output, nil
}
