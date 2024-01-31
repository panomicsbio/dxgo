package dxgo

import (
	"fmt"
)

func (c *DXClient) FileDownload(input FileDownloadInput) (FileDownloadOutput, error) {
	output := new(FileDownloadOutput)
	err := c.DoInto(fmt.Sprintf("/%s/download", input.ID), input, output)
	if err != nil {
		return FileDownloadOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileNew(input FileNewInput) (FileNewOutput, error) {
	output := new(FileNewOutput)
	err := c.DoInto("/file/new", input, output)
	if err != nil {
		return FileNewOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileUpload(input FileUploadInput) (FileUploadOutput, error) {
	output := new(FileUploadOutput)
	err := c.DoInto(fmt.Sprintf("/%s/upload", input.ID), input, output)
	if err != nil {
		return FileUploadOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileClose(input FileCloseInput) (FileCloseOutput, error) {
	output := new(FileCloseOutput)
	err := c.DoInto(fmt.Sprintf("/%s/close", input.ID), input, output)
	if err != nil {
		return FileCloseOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileDescribe(input FileDescribeInput) (FileDescribeOutput, error) {
	output := new(FileDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.ID), input, output)
	if err != nil {
		return FileDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
