package app

import (
	inquiryform "github.com/Pervadinti/rental-front/templates/inquiry_form"
	hellopage "github.com/Pervadinti/rental-front/templates/landing"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response())
}

func App() {
	e := echo.New()

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return HTML(c, hellopage.LandingPage())
	})

	//e.GET("/menu/mobile/toggle", func(c echo.Context) error {
	//	showStr := c.QueryParam("show")
	//
	//	// Default to false if 'show' query parameter is empty.
	//	if showStr == "" {
	//		showStr = "false"
	//	}
	//
	//	show, err := strconv.ParseBool(showStr)
	//	if err != nil {
	//		// Handle invalid 'show' query parameter. Return a bad request error.
	//		return c.String(http.StatusBadRequest, "Invalid 'show' query parameter.")
	//	}
	//
	//	if show {
	//		return HTML(c, common.MobileMenu())
	//	}
	//
	//	return HTML(c, common.MobileMenuToggle())
	//})

	e.DELETE("/inquiry", func(c echo.Context) error {
		return HTML(c, templ.NopComponent)
	})
	e.GET("/inquiry", func(c echo.Context) error {
		return HTML(c, inquiryform.InquiryForm())
	})
	e.POST("/inquiry", func(c echo.Context) error {
		return HTML(c, templ.NopComponent)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
