package msg

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "../../../app/models"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.String(http.StatusOK, "msg-controller is right")
}

func IndexList(c *gin.Context) {
	// var msgs Msg
	pageNum := c.DefaultQuery("pageNum", "1")
	num, _ := strconv.Atoi(pageNum)
	// lists, err := msgs.GetMsgList()
	// num := c.Query("pageNum")
	lists, err := GetMsgList1(num, 1)
	if err != nil {
		log.Fatalln("err")
	}
	fmt.Println(lists)
	c.HTML(http.StatusOK, "msg/index.html", gin.H{
		"title": "gin用户界面",
		"data":  lists,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "success",
	// 	"data": lists,
	// })
	// c.HTML(http.StatusOK, "index.html", gin.H{
	// 	"msg":  "success",
	// 	"data": lists,
	// })
}
