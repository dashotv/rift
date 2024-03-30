package app

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dashotv/fae"
)

// GET /page/
func (a *Application) PageIndex(c echo.Context, page int, limit int) error {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 25
	}
	skip := (page - 1) * limit
	if skip < 0 {
		skip = 0
	}
	a.Log.Debugf("page: %d, limit: %d, skip: %d", page, limit, skip)

	count, err := a.DB.Page.Query().Count()
	if err != nil {
		return err
	}

	list, err := a.DB.Page.Query().Desc("created_at").Limit(limit).Skip(skip).Run()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, H{"total": count, "results": list})
}

// POST /page/
func (a *Application) PageCreate(c echo.Context, req *Page) error {
	if a.DB.PageExists(req.Name) {
		return fae.New("page exists")
	}

	if err := a.DB.Page.Save(req); err != nil {
		return fae.Wrap(err, "save failed")
	}

	// if err := a.Workers.Enqueue(&ScrapePage{Page: req}); err != nil {
	// 	return nil, err
	// }

	return c.JSON(http.StatusOK, H{"error": false, "page": req})
}

// GET /page/:id
func (a *Application) PageShow(c echo.Context, id string) error {
	page, err := a.DB.PageGet(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, H{"error": false, "page": page})
}

// PUT /page/:id
func (a *Application) PageUpdate(c echo.Context, id string, req *Page) error {
	p, err := a.DB.PageGet(id)
	if err != nil {
		return err
	}
	if p == nil {
		return fae.New("page not found")
	}

	if err := c.Bind(p); err != nil {
		return fae.Wrap(err, "bind failed")
	}

	if p.ID.Hex() != id {
		return fae.New("id mismatch")
	}

	if err := a.DB.Page.Save(p); err != nil {
		return fae.Wrap(err, "save failed")
	}

	return c.JSON(http.StatusOK, H{"error": false, "page": p})
}

// PATCH /page/:id
func (a *Application) PageSettings(c echo.Context, id string, setting *Setting) error {
	// subject, err := a.DB.Page.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// DELETE /page/:id
func (a *Application) PageDelete(c echo.Context, id string) error {
	subject, err := a.DB.PageGet(id)
	if err != nil {
		return fae.Wrap(err, "get failed")
	}

	if err := a.DB.Page.Delete(subject); err != nil {
		return fae.Wrap(err, "delete failed")
	}

	return c.JSON(http.StatusOK, H{"error": false, "page": subject})
}
