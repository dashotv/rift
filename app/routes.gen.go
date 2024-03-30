// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.infratographer.com/x/echox/echozap"
)

func init() {
	initializers = append(initializers, setupRoutes)
	healthchecks["routes"] = checkRoutes
	starters = append(starters, startRoutes)
}

func checkRoutes(app *Application) error {
	// TODO: check routes
	return nil
}

func startRoutes(ctx context.Context, app *Application) error {
	go func() {
		app.Routes()
		app.Log.Info("starting routes...")
		if err := app.Engine.Start(fmt.Sprintf(":%d", app.Config.Port)); err != nil {
			app.Log.Errorf("routes: %s", err)
		}
	}()
	return nil
}

func setupRoutes(app *Application) error {
	logger := app.Log.Named("routes").Desugar()
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(echozap.Middleware(logger))

	app.Engine = e
	// unauthenticated routes
	app.Default = app.Engine.Group("")
	// authenticated routes (if enabled, otherwise same as default)
	app.Router = app.Engine.Group("")

	// TODO: fix auth
	// if app.Config.Auth {
	// 	clerkSecret := app.Config.ClerkSecretKey
	// 	if clerkSecret == "" {
	// 		app.Log.Fatal("CLERK_SECRET_KEY is not set")
	// 	}
	//
	// 	clerkClient, err := clerk.NewClient(clerkSecret)
	// 	if err != nil {
	// 		app.Log.Fatalf("clerk: %s", err)
	// 	}
	//
	// 	app.Router.Use(requireSession(clerkClient))
	// }

	return nil
}

type Setting struct {
	Name  string `json:"name"`
	Value bool   `json:"value"`
}

func (a *Application) Routes() {
	a.Default.GET("/", a.indexHandler)
	a.Default.GET("/health", a.healthHandler)

	page := a.Router.Group("/page")
	page.GET("/", a.PageIndexHandler)
	page.POST("/", a.PageCreateHandler)
	page.GET("/:id", a.PageShowHandler)
	page.PUT("/:id", a.PageUpdateHandler)
	page.PATCH("/:id", a.PageSettingsHandler)
	page.DELETE("/:id", a.PageDeleteHandler)

	video := a.Router.Group("/video")
	video.GET("/", a.VideoIndexHandler)
	video.POST("/", a.VideoCreateHandler)
	video.GET("/:id", a.VideoShowHandler)
	video.PUT("/:id", a.VideoUpdateHandler)
	video.PATCH("/:id", a.VideoSettingsHandler)
	video.DELETE("/:id", a.VideoDeleteHandler)

	visit := a.Router.Group("/visit")
	visit.GET("/", a.VisitIndexHandler)
	visit.POST("/", a.VisitCreateHandler)
	visit.GET("/:id", a.VisitShowHandler)
	visit.PUT("/:id", a.VisitUpdateHandler)
	visit.PATCH("/:id", a.VisitSettingsHandler)
	visit.DELETE("/:id", a.VisitDeleteHandler)

}

func (a *Application) indexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, H{
		"name": "rift",
		"routes": H{
			"page":  "/page",
			"video": "/video",
			"visit": "/visit",
		},
	})
}

func (a *Application) healthHandler(c echo.Context) error {
	health, err := a.Health()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, H{"name": "rift", "health": health})
}

// Page (/page)
func (a *Application) PageIndexHandler(c echo.Context) error {
	page := QueryInt(c, "page")
	limit := QueryInt(c, "limit")
	return a.PageIndex(c, page, limit)
}
func (a *Application) PageCreateHandler(c echo.Context) error {
	var subject *Page
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.PageCreate(c, subject)
}
func (a *Application) PageShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.PageShow(c, id)
}
func (a *Application) PageUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	var subject *Page
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.PageUpdate(c, id, subject)
}
func (a *Application) PageSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	var setting *Setting
	if err := c.Bind(setting); err != nil {
		return err
	}
	return a.PageSettings(c, id, setting)
}
func (a *Application) PageDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.PageDelete(c, id)
}

// Video (/video)
func (a *Application) VideoIndexHandler(c echo.Context) error {
	page := QueryInt(c, "page")
	limit := QueryInt(c, "limit")
	return a.VideoIndex(c, page, limit)
}
func (a *Application) VideoCreateHandler(c echo.Context) error {
	var subject *Video
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.VideoCreate(c, subject)
}
func (a *Application) VideoShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.VideoShow(c, id)
}
func (a *Application) VideoUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	var subject *Video
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.VideoUpdate(c, id, subject)
}
func (a *Application) VideoSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	var setting *Setting
	if err := c.Bind(setting); err != nil {
		return err
	}
	return a.VideoSettings(c, id, setting)
}
func (a *Application) VideoDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.VideoDelete(c, id)
}

// Visit (/visit)
func (a *Application) VisitIndexHandler(c echo.Context) error {
	page := QueryInt(c, "page")
	limit := QueryInt(c, "limit")
	return a.VisitIndex(c, page, limit)
}
func (a *Application) VisitCreateHandler(c echo.Context) error {
	var subject *Visit
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.VisitCreate(c, subject)
}
func (a *Application) VisitShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.VisitShow(c, id)
}
func (a *Application) VisitUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	var subject *Visit
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.VisitUpdate(c, id, subject)
}
func (a *Application) VisitSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	var setting *Setting
	if err := c.Bind(setting); err != nil {
		return err
	}
	return a.VisitSettings(c, id, setting)
}
func (a *Application) VisitDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.VisitDelete(c, id)
}