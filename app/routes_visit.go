package app

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dashotv/fae"
)

// GET /visit/
func (a *Application) VisitIndex(c echo.Context, page int, limit int) error {
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

	pageId := c.QueryParam("page_id")

	q := a.DB.Visit.Query()
	if pageId != "" {
		q = q.Where("page_id", pageId)
	}

	count, err := q.Count()
	if err != nil {
		return fae.Wrap(err, "count")
	}
	list, err := q.Desc("created_at").Limit(limit).Skip(skip).Run()
	if err != nil {
		return fae.Wrap(err, "query")
	}

	return c.JSON(http.StatusOK, H{"error": false, "total": count, "results": list})
}

// POST /visit/
func (a *Application) VisitCreate(c echo.Context, visit *Visit) error {
	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// GET /visit/:id
func (a *Application) VisitShow(c echo.Context, id string) error {
	// subject, err := a.DB.Visit.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// PUT /visit/:id
func (a *Application) VisitUpdate(c echo.Context, id string, visit *Visit) error {
	// subject, err := a.DB.Visit.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// PATCH /visit/:id
func (a *Application) VisitSettings(c echo.Context, id string, setting *Setting) error {
	// subject, err := a.DB.Visit.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// DELETE /visit/:id
func (a *Application) VisitDelete(c echo.Context, id string) error {
	// subject, err := a.DB.Visit.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}
