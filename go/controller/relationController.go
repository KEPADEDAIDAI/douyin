package controller

import (
	"douyin/go/model"
	"douyin/go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FollowListResponse struct {
	model.Response
	UserList []model.Follower `json:"user_list"`
}

type FollowerListResponse struct {
	model.Response
	UserList []model.Follower `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {

	//1.取数据
	//token := c.Query("token")
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	to_user_id, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	action_type_, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	action_type := int32(action_type_)
	//2.token鉴权

	//3.service层处理
	err := service.RelationAction(user_id, to_user_id, action_type)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "操作成功！",
		})
	}

}

// FollowList
func FollowList(c *gin.Context) {

	//1.数据预处理
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	//2.token鉴权

	//3.service层处理
	followlist, err := service.FollowList(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, FollowListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "查找列表失败！",
			},
			UserList: nil,
		})
	} else {
		c.JSON(http.StatusOK, FollowListResponse{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  "已找到列表！",
			},
			UserList: followlist,
		})
	}
}

// FollowerList
func FollowerList(c *gin.Context) {
	//1.数据预处理
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	//2.token鉴权

	//3.service层处理
	followlist, err := service.FollowerList(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, FollowListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "查找列表失败！",
			},
			UserList: nil,
		})
	} else {
		c.JSON(http.StatusOK, FollowListResponse{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  "已找到列表！",
			},
			UserList: followlist,
		})
	}
}
