package dxgo

type ProjectDescribeInput struct {
	ID     string          `json:"id"`
	Fields map[string]bool `json:"fields"`
}

type ProjectDescribeOutput struct {
	Name    string   `json:"name"`
	Folders []string `json:"folders"`
}

type DXAssetType string

type FindDataObjectsScope struct {
	Project *string `json:"project"`
	Folder  *string `json:"folder"`
	Recurse *bool   `json:"recurse"`
}

type FindDataObjectsInput struct {
	Class    DXAssetType           `json:"class"`
	Scope    *FindDataObjectsScope `json:"scope"`
	Describe bool                  `json:"describe"`
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
