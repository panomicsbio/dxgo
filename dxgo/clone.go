package dxgo

import (
	"context"
	"fmt"
)

type CloneInput struct {
	// (optional, required if folders is not provided) List of object IDs (strings of the form "class-xxxx") in the source container to be cloned
	Objects []string `json:"objects,omitempty"`
	// (optional, required if objects is not provided) List of folders in the source container to be cloned
	Folders         []string `json:"folders,omitempty"`
	Project         string   `json:"project"`
	DestinationPath string   `json:"destination"`
	//  boolean (optional, default false) Whether the destination folder and/or parent folders should be created if they do not exist
	Parents bool `json:"parents"`
}

type CloneOutput struct {
	ID      string    `json:"id"`
	Project string    `json:"project"`
	Exists  []string  `json:"exists"`
	Error   *ApiError `json:"error"`
}

func (c *DXClient) Clone(ctx context.Context, class string, input CloneInput) (CloneOutput, error) {
	output := new(CloneOutput)
	err := c.DoInto(ctx, fmt.Sprintf("/%s/clone", class), input, output)
	if err != nil {
		return CloneOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
