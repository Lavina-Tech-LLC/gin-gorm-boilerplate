package auth

import (
	"crypto/md5"
	"encoding/hex"
	"gin-gorm-boilerplate/internal/dbCon"
	"gin-gorm-boilerplate/internal/models"
	"io/ioutil"
	"strings"

	lvn "github.com/Lavina-Tech-LLC/lavina-utils"
	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
	lavina "github.com/Lavina-Tech-LLC/lavinagopackage"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

var (
	noAuthResourses = []string{
		"/",
		"/network/ip",
	}
)

func AuthenticateUser(c *gin.Context) {

	// no Auth
	if slices.Index(noAuthResourses, c.FullPath()) >= 0 {
		c.Set("User", "noAuth")
		return
	}

	usr := models.Users{}
	usr.Key = c.Request.Header.Get("k")
	err := usr.GetUserByKey()
	if err != nil {
		c.JSON(401, lavina.Response("", "No user match", false))
		c.Abort()
		return
	}

	c.Set("User", usr.Name)

	signStr := ""

	if c.Request.Method == "POST" {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			llog.Error("Unable to get body")
		}
		signStr += string(body)
		c.Set("body", body)
	} else if c.Request.Method == "GET" {
		signStr += c.Request.URL.Path + lvn.Ternary(c.Request.URL.RawQuery == "", "", "?"+c.Request.URL.RawQuery).(string)
	} else if c.Request.Method == "OPTIONS" {
		return
	} else {
		c.JSON(400, lavina.Response("", "Bad Request", false))
		c.Abort()
		return
	}

	signStr += usr.Secret

	sign := md5.Sum([]byte(signStr))
	if dbCon.Dm {
		llog.Info("Signing string is: " + signStr)
		llog.Info("Generated: " + hex.EncodeToString(sign[:]))
		llog.Info("Recieved: " + strings.ToLower(c.Request.Header.Get("s")))
	}

	if strings.ToLower(c.Request.Header.Get("s")) != hex.EncodeToString(sign[:]) {
		c.JSON(401, lavina.Response("", "Bad credentials", false))
		c.Abort()
		return
	}

}

func AuthorizeUser(c *gin.Context) {
	enforcer := dbCon.GetCasbin

	// Load stored policy from database
	enforcer.LoadPolicy()

	// Do permission checking
	isOk, _ := enforcer.Enforce(c.MustGet("User").(string), c.FullPath(), c.Request.Method)

	if !isOk {
		c.JSON(401, lavina.Response("", "Unauthorized", false))
		c.Abort()
		return
	}

}
