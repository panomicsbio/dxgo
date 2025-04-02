package dxgo

import (
	"context"
	"fmt"
	"time"
)

type FileDownloadInput struct {
	ID               string  `json:"id"`
	Duration         *int    `json:"duration,omitempty"`
	Filename         *string `json:"filename,omitempty"`
	Project          string  `json:"project"`
	Preauthenticated *bool   `json:"preauthenticated,omitempty"`
	StickyIP         *bool   `json:"stickyIP,omitempty"`
}

type FileDownloadOutput struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Error   *ApiError         `json:"error"`
}

func (c *DXClient) FileDownload(ctx context.Context, input FileDownloadInput) (FileDownloadOutput, error) {
	output := new(FileDownloadOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/download", input.ID), input, output)
	if err != nil {
		return FileDownloadOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type FileNewInput struct {
	Project string `json:"project"`
	Folder  string `json:"folder"`
	Parent  bool   `json:"parents"`
	Name    string `json:"name"`
	Media   string `json:"media"`
}

type FileNewOutput struct {
	ID    string    `json:"id"`
	Error *ApiError `json:"error"`
}

func (c *DXClient) FileNew(ctx context.Context, input FileNewInput) (FileNewOutput, error) {
	output := new(FileNewOutput)
	err := c.DoInto(ctx, "/file/new", input, output)
	if err != nil {
		return FileNewOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type FileUploadInput struct {
	ID    string `json:"id"`
	Size  int    `json:"size"`
	MD5   string `json:"md5"`
	Index int    `json:"index"`
}

type FileUploadOutput struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Error   *ApiError         `json:"error"`
}

func (c *DXClient) FileUpload(ctx context.Context, input FileUploadInput) (FileUploadOutput, error) {
	output := new(FileUploadOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/upload", input.ID), input, output)
	if err != nil {
		return FileUploadOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type FileCloseInput struct {
	ID string `json:"id"`
}

type FileCloseOutput struct {
	Error *ApiError `json:"error"`
}

func (c *DXClient) FileClose(ctx context.Context, input FileCloseInput) (FileCloseOutput, error) {
	output := new(FileCloseOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/close", input.ID), input, output)
	if err != nil {
		return FileCloseOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type FileDescribeInput struct {
	ID string `json:"id"`
}

type FileDescribeOutput struct {
	Folder  string    `json:"folder"`
	State   string    `json:"state"`
	Name    string    `json:"name"`
	Media   string    `json:"media"`
	Size    uint64    `json:"size"`
	Created time.Time `json:"created"`
	Error   *ApiError `json:"error"`
}

func (c *DXClient) FileDescribe(ctx context.Context, input FileDescribeInput) (FileDescribeOutput, error) {
	output := new(FileDescribeOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/describe", input.ID), input, output)
	if err != nil {
		return FileDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

func (c *DXClient) FileClone(ctx context.Context, sourceProject string, input CloneInput) (CloneOutput, error) {
	return c.Clone(ctx, sourceProject, input)
}
