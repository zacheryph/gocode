package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var saveSuccess = `<html><head>
<meta http-equiv="Refresh" content="5; url=/" />
</head><body>
<pre>You have Successfully RSVP'd.  Thank you.  You will now be Redirected.</pre>
</body></html>`

func echoServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Post("/rsvp", requestAddRsvp)

	admin := e.Group("/rsvp")
	admin.Use(middleware.BasicAuth(checkAuth))
	admin.Get("/list", requestListRsvp)

	e.Static("/", *rootDir)

	fmt.Println("Starting Server:", *httpServ)
	e.Run(*httpServ)
}

func checkAuth(user, passwd string) bool {
	return true
}

func requestAddRsvp(c *echo.Context) error {
	name := c.Form("name")
	email := c.Form("email")
	if name == "" || email == "" {
		c.String(http.StatusBadRequest, "Name and Email must be given. Please go back and try again.")
		return nil
	}

	if err := addRsvp(name, email); err != nil {
		c.String(http.StatusInternalServerError,
			"Failed to add RSVP at this time. Please try again later")
	}

	c.HTML(http.StatusOK, saveSuccess)
	return nil
}

func requestListRsvp(c *echo.Context) error {
	res := c.Response()
	res.Header().Add("Content-Type", "text/plain")
	listRsvp(res)
	return nil
}
