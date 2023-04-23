package dxgo

func (c *DXClient) FindDataObjects() {
	data, err := c.retryableRequest("/system/findDataObjects", struct{}{})
	if err != nil {
		println(err.Error())
		return
	}
	println("----------")
	println(string(data))
	println("----------")
}
