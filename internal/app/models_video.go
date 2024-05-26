package app

func (c *Connector) VideoGet(id string) (*Video, error) {
	m := &Video{}
	err := c.Video.Find(id, m)
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) VideoList() ([]*Video, error) {
	list, err := c.Video.Query().Limit(10).Run()
	if err != nil {
		return nil, err
	}

	return list, nil
}
