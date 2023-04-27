package dxgo

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
	ID                 string              `json:"id"`
	Project            string              `json:"project"`
	SystemRequirements *SystemRequirements `json:"systemRequirements"`
	Properties         map[string]string   `json:"properties"`
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
