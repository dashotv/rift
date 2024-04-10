package app

import "github.com/dashotv/fae"

func (c *Connector) PageExists(name string) bool {
	count, err := c.Page.Query().Where("name", name).Count()
	if err != nil {
		return true
	}
	if count > 0 {
		return true
	}

	return false
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
		return false, fae.Wrap(err, "is_visited: counting visit")
	}
	if count > 0 {
		return true, nil
	}

	if err := c.Visit.Save(&Visit{PageId: page.ID, Url: url}); err != nil {
		return false, fae.Wrap(err, "is_visited: saving visit")
	}

	return false, nil
}

func (c *Connector) VisitError(page *Page, url string, failure *fae.Error) error {
	visits, err := c.Visit.Query().Where("page_id", page.ID).Where("url", url).Run()
	if err != nil {
		return fae.Wrap(err, "visit_error: querying visit")
	}
	if len(visits) == 0 {
		return fae.Errorf("visit_error: no visit found")
	}
	if len(visits) > 1 {
		return fae.Errorf("visit_error: multiple visits found")
	}
	if failure == nil {
		return fae.Errorf("visit_error: failure is nil")
	}

	visits[0].Error = failure.Error()
	visits[0].Stacktrace = fae.StackTrace(failure)
	if err := c.Visit.Save(visits[0]); err != nil {
		return fae.Wrap(err, "visit_error: saving visit")
	}

	return nil
}
