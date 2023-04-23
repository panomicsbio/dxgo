package dxgo

func (c *DXClient) findDataObjects() {
	data, err := c.retryableRequest("/system/findDataObjects", struct{}{})
	if err != nil {
		println(err)
		return
	}
	println(string(data))
}
