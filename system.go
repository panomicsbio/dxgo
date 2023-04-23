package dxgo

import "encoding/json"

func (c *DXClient) FindDataObjects(input *FindDataObjectsInput) (*FindDataObjectsOutput, error) {
	data, err := c.retryableRequest("/system/findDataObjects", input)
	if err != nil {
		return nil, err
	}
	output := new(FindDataObjectsOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
