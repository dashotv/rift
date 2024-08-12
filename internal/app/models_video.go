package app

import "github.com/dashotv/fae"

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

func (c *Connector) VideoFindOrCreate(displayID string) (*Video, error) {
	list, err := app.DB.Video.Query().Where("display_id", displayID).Run()
	if err != nil {
		return nil, fae.Wrap(err, "finding video")
	}
	if len(list) > 0 {
		return list[0], nil
	}

	v := &Video{}
	v.DisplayID = displayID

	return v, nil
}
