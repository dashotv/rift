package app

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/dashotv/fae"
)

// GET /video/
func (a *Application) VideoIndex(c echo.Context, page int, limit int) error {
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
	a.Log.Debugf("VideoIndex: page=%d limit=%d skip=%d", page, limit, skip)
	count, err := a.DB.Video.Query().Count()
	if err != nil {
		return fae.Wrap(err, "count")
	}

	list, err := a.DB.Video.Query().Desc("created_at").Limit(limit).Skip(skip).Run()
	if err != nil {
		return fae.Wrap(err, "query")
	}

	return c.JSON(http.StatusOK, H{"error": false, "total": count, "result": list})
}

// POST /video/
func (a *Application) VideoCreate(c echo.Context, video *Video) error {
	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// GET /video/:id
func (a *Application) VideoShow(c echo.Context, id string) error {
	// subject, err := a.DB.Video.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// PUT /video/:id
func (a *Application) VideoUpdate(c echo.Context, id string, video *Video) error {
	subject, err := a.DB.VideoGet(id)
	if err != nil {
		return fae.Errorf("video not found: %s", id)
	}

	subject.UpdatedAt = time.Now()

	if err := a.DB.Video.Save(subject); err != nil {
		return fae.Wrap(err, "save")
	}

	return c.JSON(http.StatusOK, H{"error": false, "result": subject})
}

// PATCH /video/:id
func (a *Application) VideoSettings(c echo.Context, id string, setting *Setting) error {
	// subject, err := a.DB.Video.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// DELETE /video/:id
func (a *Application) VideoDelete(c echo.Context, id string) error {
	// subject, err := a.DB.Video.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}
