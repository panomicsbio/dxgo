package dxgo

type ApiError struct {
	Type    string            `json:"type"`
	Message string            `json:"message"`
	Details map[string]string `json:"details"`
}

type ProjectDescribeInput struct {
	ID     string          `json:"id"`
	Fields map[string]bool `json:"fields"`
}

type ProjectDescribeOutput struct {
	Name    string   `json:"name"`
	Folders []string `json:"folders"`
}

type ProjectNewFolderInput struct {
	ID     string `json:"id"`
	Folder string `json:"folder"`
}

type DXAssetType string

type FindDataObjectsScope struct {
	Project *string `json:"project"`
	Folder  *string `json:"folder"`
	Recurse *bool   `json:"recurse"`
}

type FindDataObjectsInput struct {
	Name     string                 `json:"name"`
	Class    DXAssetType            `json:"class"`
	Scope    *FindDataObjectsScope  `json:"scope"`
	Describe bool                   `json:"describe"`
	Starting map[string]interface{} `json:"starting,omitempty"`
}

type FindDataObjectsResult struct {
	Project  string                 `json:"project"`
	ID       string                 `json:"id"`
	Describe map[string]interface{} `json:"describe"`
}

type FindDataObjectsOutput struct {
	Results []*FindDataObjectsResult `json:"results"`
	Next    map[string]interface{}   `json:"next"`
}

type SystemRequirements struct {
	InstanceType string `json:"instanceType"`
}

type AppletRunInput struct {
	ID                 string                 `json:"id"`
	Project            string                 `json:"project"`
	Input              map[string]interface{} `json:"input"`
	SystemRequirements *SystemRequirements    `json:"systemRequirements"`
	Properties         map[string]string      `json:"properties"`
	Error              ApiError               `json:"error"`
}

type AppletRunOutput struct {
	ID string `json:"id"`
}

type JobTerminateInput struct {
	ID string `json:"id"`
}

type JobDescribeInput struct {
	ID string `json:"id"`
}

type JobDescribeOutput struct {
	State string `json:"state"`
}

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
}

type RemoveObjectsInput struct {
	Project string   `json:"project"`
	Objects []string `json:"objects"`
	Force   bool     `json:"force"`
}

type RemoveFolderInput struct {
	Project string `json:"project"`
	Folder  string `json:"folder"`
	Force   bool   `json:"force"`
	Recurse bool   `json:"recurse"`
}
