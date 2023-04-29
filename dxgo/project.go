package dxgo

import (
	"encoding/json"
	"fmt"
)

func (c *DXClient) ProjectDescribe(input *ProjectDescribeInput) (*ProjectDescribeOutput, error) {
	data, err := c.retryableRequest(fmt.Sprintf("/%s/describe", input.ID), input)
	if err != nil {
		return nil, err
	}
	output := new(ProjectDescribeOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c *DXClient) ProjectNewFolder(input *ProjectNewFolderInput) (*ProjectNewFolderOutput, error) {
	data, err := c.retryableRequest(fmt.Sprintf("/%s/newFolder", input.ID), input)
	if err != nil {
		return nil, err
	}
	output := new(ProjectNewFolderOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c *DXClient) RemoveObjects(input *RemoveObjectsInput) (*RemoveObjectsOutput, error) {
	data, err := c.retryableRequest(fmt.Sprintf("/%s/removeObjects", input.Project), input)
	if err != nil {
		return nil, err
	}
	output := new(RemoveObjectsOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c *DXClient) RemoveFolder(input *RemoveFolderInput) (*RemoveFolderOutput, error) {
	data, err := c.retryableRequest(fmt.Sprintf("/%s/removeFolder", input.Project), input)
	if err != nil {
		return nil, err
	}
	output := new(RemoveFolderOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
