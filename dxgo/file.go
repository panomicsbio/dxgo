package dxgo

import (
	"context"
	"fmt"
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

func (c *DXClient) FileNew(ctx context.Context, input FileNewInput, public bool) (FileNewOutput, error) {
	// Dnanexus provides file upload URLS differently depending on whether the request if from the platform or not, from a job or not.
	// In order to allow the client browser to multipart upload file parts directly we must impersonate a browser platform request.
	// This is a major hack but the official API does not provide a way to do this.
	headers := map[string]string{}
	if public {
		headers["Host"] = "api.dnanexus.com"
		headers["Origin"] = "https://platform.dnanexus.com"
	}

	output := new(FileNewOutput)
	err := c.DoIntoWithHeaders(ctx, "/file/new", input, output, headers)
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

func (c *DXClient) FileUpload(ctx context.Context, input FileUploadInput, public bool) (FileUploadOutput, error) {
	output := new(FileUploadOutput)

	// Dnanexus provides file upload URLS differently depending on whether the request if from the platform or not, from a job or not.
	// In order to allow the client browser to multipart upload file parts directly we must impersonate a browser platform request.
	// This is a major hack but the official API does not provide a way to do this.
	headers := map[string]string{}
	if public {
		headers["Host"] = "api.dnanexus.com"
		headers["Origin"] = "https://platform.dnanexus.com"
	}
	err := c.DoIntoWithHeaders(ctx, fmt.Sprintf("/%s/upload", input.ID), input, output, headers)
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

func (c *DXClient) FileClose(ctx context.Context, input FileCloseInput, public bool) (FileCloseOutput, error) {
	output := new(FileCloseOutput)

	// Dnanexus provides file upload URLS differently depending on whether the request if from the platform or not, from a job or not.
	// In order to allow the client browser to multipart upload file parts directly we must impersonate a browser platform request.
	// This is a major hack but the official API does not provide a way to do this.
	headers := map[string]string{}
	if public {
		headers["Host"] = "api.dnanexus.com"
		headers["Origin"] = "https://platform.dnanexus.com"
	}
	err := c.DoIntoWithHeaders(ctx, fmt.Sprintf("/%s/close", input.ID), input, output, headers)
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
	Created uint64    `json:"created"`
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
