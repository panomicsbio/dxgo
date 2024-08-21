package dxgo

import (
	"fmt"
	"time"
)

func (c *DXClient) FileDownload(input FileDownloadInput, timeout time.Duration) (FileDownloadOutput, error) {
	output := new(FileDownloadOutput)
	err := c.DoInto(fmt.Sprintf("/%s/download", input.ID), input, output, timeout)
	if err != nil {
		return FileDownloadOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileNew(input FileNewInput, timeout time.Duration) (FileNewOutput, error) {
	output := new(FileNewOutput)
	err := c.DoInto("/file/new", input, output, timeout)
	if err != nil {
		return FileNewOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileUpload(input FileUploadInput, timeout time.Duration) (FileUploadOutput, error) {
	output := new(FileUploadOutput)
	err := c.DoInto(fmt.Sprintf("/%s/upload", input.ID), input, output, timeout)
	if err != nil {
		return FileUploadOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileClose(input FileCloseInput, timeout time.Duration) (FileCloseOutput, error) {
	output := new(FileCloseOutput)
	err := c.DoInto(fmt.Sprintf("/%s/close", input.ID), input, output, timeout)
	if err != nil {
		return FileCloseOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileDescribe(input FileDescribeInput, timeout time.Duration) (FileDescribeOutput, error) {
	output := new(FileDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.ID), input, output, timeout)
	if err != nil {
		return FileDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
