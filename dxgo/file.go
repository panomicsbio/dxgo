package dxgo

import (
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

func (c *DXClient) FileDownload(input FileDownloadInput, timeout time.Duration) (FileDownloadOutput, error) {
	output := new(FileDownloadOutput)
	err := c.DoInto(fmt.Sprintf("/%s/download", input.ID), input, output, timeout)
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
}

type FileNewOutput struct {
	ID    string    `json:"id"`
	Error *ApiError `json:"error"`
}

func (c *DXClient) FileNew(input FileNewInput, timeout time.Duration) (FileNewOutput, error) {
	output := new(FileNewOutput)
	err := c.DoInto("/file/new", input, output, timeout)
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

func (c *DXClient) FileUpload(input FileUploadInput, timeout time.Duration) (FileUploadOutput, error) {
	output := new(FileUploadOutput)
	err := c.DoInto(fmt.Sprintf("/%s/upload", input.ID), input, output, timeout)
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

func (c *DXClient) FileClose(input FileCloseInput, timeout time.Duration) (FileCloseOutput, error) {
	output := new(FileCloseOutput)
	err := c.DoInto(fmt.Sprintf("/%s/close", input.ID), input, output, timeout)
	if err != nil {
		return FileCloseOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}

type FileDescribeOutput struct {
	Folder string    `json:"folder"`
	State  string    `json:"state"`
	Name   string    `json:"name"`
	Error  *ApiError `json:"error"`
}

type FindProjectsInput struct {
	Name     any            `json:"name,omitempty"`
	Level    string         `json:"level,omitempty"`
	Starting string         `json:"starting,omitempty"`
	Describe map[string]any `json:"describe"`
}

func (c *DXClient) FileDescribe(input FileDescribeInput, timeout time.Duration) (FileDescribeOutput, error) {
	output := new(FileDescribeOutput)
	err := c.DoInto(fmt.Sprintf("/%s/describe", input.ID), input, output, timeout)
	if err != nil {
		return FileDescribeOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
