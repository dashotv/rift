package app

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dashotv/fae"
)

// GET /page/
func (a *Application) PageIndex(c echo.Context, page int, limit int) error {
	skip := (page - 1) * limit
	if skip < 0 {
		skip = 0
	}

	count, err := a.DB.Page.Query().Count()
	if err != nil {
		return err
	}

	list, err := a.DB.Page.Query().Asc("name").Limit(limit).Skip(skip).Run()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, H{"error": false, "total": count, "result": list})
}

// POST /page/
func (a *Application) PageCreate(c echo.Context, req *Page) error {
	// if a.DB.PageExists(req.Name) {
	// 	return fae.New("page exists")
	// }

	if err := a.DB.Page.Save(req); err != nil {
		return fae.Wrap(err, "save failed")
	}

	// if err := a.Workers.Enqueue(&ScrapePage{Page: req}); err != nil {
	// 	return nil, err
	// }

	return c.JSON(http.StatusOK, H{"error": false, "result": req})
}

// GET /page/:id
func (a *Application) PageShow(c echo.Context, id string) error {
	page, err := a.DB.PageGet(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, H{"error": false, "result": page})
}

// PUT /page/:id
func (a *Application) PageUpdate(c echo.Context, id string, req *Page) error {
	if req.ID.Hex() != id {
		return fae.New("id mismatch")
	}

	if err := a.DB.Page.Save(req); err != nil {
		return fae.Wrap(err, "save failed")
	}

	return c.JSON(http.StatusOK, H{"error": false, "result": req})
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

	return c.JSON(http.StatusOK, H{"error": false, "result": subject})
}

// GET /page/:id/visits
func (a *Application) PageVisits(c echo.Context, id string, page int, limit int) error {
	p, err := a.DB.PageGet(id)
	if err != nil {
		return err
	}

	q := a.DB.Visit.Query().Where("page_id", p.ID)

	count, err := q.Count()
	if err != nil {
		return err
	}

	list, err := q.Desc("created_at").Limit(limit).Skip((page - 0) * limit).Run()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, H{"error": false, "total": count, "result": list})
}

// GET /page/:id/videos
func (a *Application) PageVideos(c echo.Context, id string, page int, limit int) error {
	p, err := a.DB.PageGet(id)
	if err != nil {
		return err
	}

	q := a.DB.Video.Query().Where("title", p.Name)

	count, err := q.Count()
	if err != nil {
		return err
	}

	list, err := q.Desc("created_at").Limit(limit).Skip((page - 1) * limit).Run()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, H{"error": false, "total": count, "result": list})
}

func (a *Application) PageRefresh(c echo.Context, id string) error {
	a.Log.Debugf("refreshing page %s", id)
	p, err := a.DB.PageGet(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, H{"error": true, "message": err.Error()})
	}

	if err := a.Workers.Enqueue(&ScrapePage{Page: p}); err != nil {
		return c.JSON(http.StatusInternalServerError, H{"error": true, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, H{"error": false, "result": p})
}
