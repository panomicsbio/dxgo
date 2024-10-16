package dxgo

import (
	"fmt"
	"time"
)

type AppletRunInput struct {
	ID                 string             `json:"id"`
	Project            *string            `json:"project,omitempty"`
	Input              map[string]any     `json:"input"`
	SystemRequirements SystemRequirements `json:"systemRequirements"`
	Properties         map[string]string  `json:"properties"`
	Detach             bool               `json:"detach"`
	HeadJobOnDemand    bool               `json:"headJobOnDemand"`
}

type AppletRunOutput struct {
	ID    string    `json:"id"`
	Error *ApiError `json:"error"`
}

func (c *DXClient) AppletRun(input AppletRunInput, timeout time.Duration) (AppletRunOutput, error) {
	output := new(AppletRunOutput)
	err := c.DoInto(fmt.Sprintf("/%s/run", input.ID), input, output, timeout)
	if err != nil {
		return AppletRunOutput{}, fmt.Errorf("doing request: %w", err)
	}

	return *output, nil
}
