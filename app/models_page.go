package app

import "github.com/dashotv/fae"

func (c *Connector) PageExists(name string) bool {
	count, err := c.Page.Query().Where("name", name).Count()
	if err != nil {
		return false
	}
	if count > 0 {
		return false
	}

	return true
}

func (c *Connector) PageGet(id string) (*Page, error) {
	m := &Page{}
	err := c.Page.Find(id, m)
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) PageList() ([]*Page, error) {
	list, err := c.Page.Query().Limit(10).Run()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (c *Connector) IsVisited(page *Page, url string) (bool, error) {
	count, err := c.Visit.Query().Where("page_id", page.ID).Where("url", url).Count()
	if err != nil {
		return false, fae.Errorf("is_visited: counting visit: %w", err)
	}
	if count > 0 {
		return true, nil
	}

	if err := c.Visit.Save(&Visit{PageId: page.ID, Url: url}); err != nil {
		return false, fae.Errorf("is_visited: saving visit: %w", err)
	}

	return false, nil
}
