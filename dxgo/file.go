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

func (c *DXClient) FileNew(input *FileNewInput) (*FileNewOutput, error) {
	data, err := c.retryableRequest("/file/new", input)
	if err != nil {
		return nil, err
	}
	output := new(FileNewOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c *DXClient) FileUpload(input *FileUploadInput) (*FileUploadOutput, error) {
	data, err := c.retryableRequest(fmt.Sprintf("/%s/upload", input.ID), input)
	if err != nil {
		return nil, err
	}
	output := new(FileUploadOutput)
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
