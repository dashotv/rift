package app

func (c *Connector) VisitGet(id string) (*Visit, error) {
	m := &Visit{}
	err := c.Visit.Find(id, m)
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) VisitList() ([]*Visit, error) {
	list, err := c.Visit.Query().Limit(10).Run()
	if err != nil {
		return nil, err
	}

	return list, nil
}
