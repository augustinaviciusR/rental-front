package app

import (
	"log"

	"github.com/Pervadinti/rental-front/app/config"
	"github.com/Pervadinti/rental-front/app/model"
	"github.com/Pervadinti/rental-front/app/store"
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

	cfg, err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	dataStore, err := store.NewGoogleSheetsStore(cfg.SpreadsheetId, cfg.GoogleServiceKey)
	if err != nil {
		log.Fatalf("failed to create datastore: %v", err)
	}

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return HTML(c, hellopage.LandingPage())
	})

	e.DELETE("/inquiry", func(c echo.Context) error {
		return HTML(c, templ.NopComponent)
	})
	e.GET("/inquiry", func(c echo.Context) error {
		return HTML(c, inquiryform.InquiryForm())
	})
	e.POST("/inquiry", func(c echo.Context) error {
		var inquiryform model.InquiryForm
		if err := c.Bind(&inquiryform); err != nil {
			return HTML(c, templ.NopComponent)
		}
		dataStore.StoreNewInquiry(c.Request().Context(), &inquiryform)
		return HTML(c, templ.NopComponent)
	})

	// e.GET("/gdpr", func(c echo.Context) error {
	// 	for _, cookie := range c.Cookies() {
	// 		if cookie.Name == "gdpr_consent" && cookie.Value == "true" {
	// 			return HTML(c, hellopage.LandingPage())
	// 		}
	// 	}
	// 	return HTML(c, gdpr.GdprPopupPage())
	// })

	// e.POST("/gdpr", func(c echo.Context) error {
	// 	cookie := &http.Cookie{
	// 		Name:     "gdpr_consent",
	// 		Value:    "true",
	// 		HttpOnly: true,
	// 		Expires:  time.Now().Add(30 * 24 * time.Hour),
	// 		Domain:   "http://localhost:8080",
	// 	}
	// 	c.SetCookie(cookie)

	// 	return HTML(c, hellopage.Landing())
	// })

	e.Logger.Fatal(e.Start(":8080"))
}
