package dxgo

import (
	"encoding/json"
	"fmt"
)

func (c *DXClient) FileDownload(input *FileDownloadInput) (*FileDownloadOutput, error) {
	data, err := c.retryableRequest(fmt.Sprintf("/%s/download", input.ID), input)
	if err != nil {
		return nil, err
	}
	output := new(FileDownloadOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
