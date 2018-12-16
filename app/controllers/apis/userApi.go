package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "../../../app/models"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It Works!")
}

func AddUserApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	u := User{FirstName: firstName, LastName: lastName}

	rs, err := u.AddUser()
	if err != nil {
		log.Fatalln(err)
		return
	}

	msg := fmt.Sprintf("insert success %d", rs)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

func UpdateUserApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	u := User{Id: id}
	u.GetUserInfoById()
	if u.Id > 0 {
		u.FirstName = firstName
		u.LastName = lastName
		ra, err := u.UpdateUser()
		if err != nil {
			log.Fatalln(err)
		}

		msg := fmt.Sprintf("update success %d", ra)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "not found",
		})
	}
}

func DelUserApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	u := User{Id: id}
	u.GetUserInfoById()
	if u.Id > 0 {
		rs, _err := u.DelUser()
		if _err != nil {
			log.Fatalln(_err)
		}

		msg := fmt.Sprintf("delete success %d", rs)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	}

}

//查询所有
func GetUserApi(c *gin.Context) {
	var u User
	users, err := u.GetUserInfo()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"msg":  "success",
	})

}

//根据id查询
func GetUserByIdApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	u := User{Id: id}
	u.GetUserInfoById()
	if u.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": u,
			"msg":  "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "user not found",
		})
	}
}
