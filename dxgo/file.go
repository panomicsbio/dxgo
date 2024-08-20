package dxgo

import (
	"context"
	"fmt"
)

func (c *DXClient) FileDownload(ctx context.Context, input FileDownloadInput) (FileDownloadOutput, error) {
	output := new(FileDownloadOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/download", input.ID), input, output)
	if err != nil {
		return FileDownloadOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileNew(ctx context.Context, input FileNewInput) (FileNewOutput, error) {
	output := new(FileNewOutput)
	err := c.DoInto(ctx, "/file/new", input, output)
	if err != nil {
		return FileNewOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileUpload(ctx context.Context, input FileUploadInput) (FileUploadOutput, error) {
	output := new(FileUploadOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/upload", input.ID), input, output)
	if err != nil {
		return FileUploadOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileClose(ctx context.Context, input FileCloseInput) (FileCloseOutput, error) {
	output := new(FileCloseOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/close", input.ID), input, output)
	if err != nil {
		return FileCloseOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileDescribe(ctx context.Context, input FileDescribeInput) (FileDescribeOutput, error) {
	output := new(FileDescribeOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", input.ID), input, output)
	if err != nil {
		return FileDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
