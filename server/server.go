package server

import (
	"gost/gist"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Addr = ""
)

func Start() {
	router := gin.Default()

	// get all
	router.GET("/gist", func(c *gin.Context) {
		ret := gist.GetAllKV()

		c.IndentedJSON(200, ret)
	})

	// get one
	router.GET("/gist/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		ret, err := gist.Get(uid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.String(200, ret.String()+"\n")
		}
	})

	// set one
	router.POST("/gist", func(c *gin.Context) {
		msg, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		uid, err := gist.Post(msg)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.String(200, uid.String()+"\n")
		}
	})

	// describe one
	router.GET("/gist/:uid/*action", func(c *gin.Context) {
		uid := c.Param("uid")
		action := c.Param("action")
		switch action {
		case "/describe":
			ret, err := gist.Describe(uid)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			} else {
				c.String(200, ret)
			}
		default:
			c.String(http.StatusInternalServerError, "action : "+action+" not supported")
		}
	})

	// delete one
	router.DELETE("/gist/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		err := gist.Remove(uid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.String(200, "deleted")
		}
	})

	// health check
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.Run(Addr)
}
