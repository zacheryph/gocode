package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	echo "gopkg.in/labstack/echo.v1"
	"gopkg.in/labstack/echo.v1/middleware"
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
	admin.Get("/backup", requestBoltBackup)

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
	res := c.Form("response")
	if name == "" || email == "" {
		c.String(http.StatusBadRequest, "Name and Email must be given. Please go back and try again.")
		return nil
	}

	rsvp := Rsvp{0, name, email, res, time.Now()}
	if err := addRsvp(rsvp); err != nil {
		c.String(http.StatusInternalServerError,
			"Failed to add RSVP at this time. Please try again later")
	}

	c.HTML(http.StatusOK, saveSuccess)
	return nil
}

func requestListRsvp(c *echo.Context) error {
	w := c.Response()
	w.Header().Add("Content-Type", "text/plain")
	listRsvp(w)
	return nil
}

func requestBoltBackup(c *echo.Context) error {
	w := c.Response()

	err := db.View(func(tx *bolt.Tx) error {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", `attachment; filename="rsvp.boltdb"`)
		w.Header().Set("Content-Length", strconv.Itoa(int(tx.Size())))
		_, err := tx.WriteTo(w)
		return err
	})

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return err
}
