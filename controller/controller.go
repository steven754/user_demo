package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test/models"
)

func CreateAccount(c *gin.Context) {
	// 1. 从请求中把数据拿出来
	var userInfo models.User
	err := c.BindJSON(&userInfo)
	if err != nil {
		if strings.Contains(err.Error(), "card") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1000,
				"msg":   "fail",
				"error": "card参数类型错误"})
		} else if strings.Contains(err.Error(), "iphone") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1000,
				"msg":   "fail",
				"error": "iphone参数类型错误"})
		} else if strings.Contains(err.Error(), "age") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1000,
				"msg":   "fail",
				"error": "age参数类型错误"})
		} else if strings.Contains(err.Error(), "sex") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1000,
				"msg":   "fail",
				"error": "sex参数类型错误"})
		} else if strings.Contains(err.Error(), "name") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1000,
				"msg":   "fail",
				"error": "name参数类型错误"})
		} else if strings.Contains(err.Error(), "account") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1000,
				"msg":   "fail",
				"error": "account参数类型错误"})
		} else if strings.Contains(err.Error(), "password") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1000,
				"msg":   "fail",
				"error": "password参数类型错误"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1000,
				"msg":   "fail",
				"error": "参数类型错误"})
		}
	} else if userInfo.Account == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1001,
			"msg":   "fail",
			"error": "Account参数无效",
		})
	} else if userInfo.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1002,
			"msg":   "fail",
			"error": "password参数无效"})
	} else {
		// 2. 存入数据库
		err := models.CreateUser(&userInfo)
		if err != nil {
			if strings.Contains(err.Error(), "1406") {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":  1000,
					"msg":   "fail",
					"error": "字段超长"})
			} else if strings.Contains(err.Error(), "1062") {
				c.JSON(http.StatusOK, gin.H{
					"code":  1000,
					"msg":   "fail",
					"error": "账号已存在"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":  1000,
					"msg":   "fail",
					"error": "未知错误",
				})
			}
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": userInfo,
			})
		}
	}

}

func GetUserList(c *gin.Context) {
	// 查询user这个表里的所有数据
	UserList, err := models.GetUserList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1000,
			"msg":   "fail",
			"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": UserList,
		})
	}
}

func GetUserInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1000,
			"msg":   "fail",
			"error": "无效的id"})
		return
	}
	userInfo, err := models.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1000,
			"msg":   "fail",
			"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": userInfo,
		})
	}
}

func UpdateUserInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1000,
			"msg":   "fail",
			"error": "无效的id"})
		return
	}
	UserInfo, err := models.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1000,
			"msg":   "fail",
			"error": err.Error()})
		return
	}
	err = c.BindJSON(&UserInfo)
	if err != nil {
		return
	}
	if err = models.UpdateAVisitInfo(UserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1000,
			"msg":   "fail",
			"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": UserInfo})
	}
}

func DeleteUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1000,
			"msg":   "fail",
			"error": "无效的id"})
		return
	}
	if err := models.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1000,
			"msg":   "fail",
			"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": id + "  deleted"})
	}
}
