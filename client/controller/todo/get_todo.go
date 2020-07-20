package todo

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/proto/todo"
)

func GetTodoServiceClient(ctx *gin.Context) {
	conn := utils.GetConnection()
	defer conn.Close()
	getTodoClient := todo.NewTodoServiceClient(conn)
	res := getParameters(ctx)
	request := &todo.GetTodoRequest{
		SortBy:      res["sortBy"].(string),
		IsAscend:    res["isAscend"].(bool),
		Page:        int32(res["page"].(int64)),
		PageSize:    int32(res["pageSize"].(int64)),
		AccessToken: res["accessToken"].(string),
	}
	if _, ok := res["startedAt"]; ok {
		request.StartedAt = res["startedAt"].(int64)
	}
	if _, ok := res["endedAt"]; ok {
		request.EndedAt = res["endedAt"].(int64)
	}
	if _, ok := res["status"]; ok {
		if res["status"].(bool) {
			request.Status = "true"
		} else {
			request.Status = "false"
		}
	}
	if _, ok := res["keyword"]; ok {
		request.KeyWord = res["keyword"].(string)
	}
	reply, err := getTodoClient.GetTodo(ctx, request)
	if err != nil {
		panic(nil)
	}
	ctx.JSON(codes.OK, gin.H{
		"todos": reply.Todos,
		"total": reply.Total,
	})
}
func getParameters(ctx *gin.Context) map[string]interface{} {
	res := make(map[string]interface{})
	var err error
	res["isAscend"], err = strconv.ParseBool(ctx.DefaultQuery("isAscend", "true"))
	if err != nil {
		panic(nil)
	}
	res["page"], err = strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 32)
	if err != nil {
		panic(nil)
	}
	res["pageSize"], err = strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		panic(nil)
	}
	res["sortBy"], res["accessToken"] =
		ctx.DefaultQuery("sortBy", "createdAt"),
		ctx.GetHeader("AccessToken")
	if ctx.Query("startedAt") != "" {
		res["startedAt"], err = strconv.ParseInt(ctx.Query("startedAt"), 10, 64)
		if err != nil {
			panic(nil)
		}
	}
	if ctx.Query("endedAt") != "" {
		res["endedAt"], err = strconv.ParseInt(ctx.Query("endedAt"), 10, 64)
		if err != nil {
			panic(nil)
		}
	}
	if ctx.Query("status") != "" {
		res["status"], err = strconv.ParseBool(ctx.Query("status"))
		if err != nil {
			panic(nil)
		}
	}
	if ctx.Query("keyword") != "" {
		res["keyword"] = ctx.Query("keyword")
	}
	return res
}
