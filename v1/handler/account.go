package handler

import (
	"account-web/initialize"
	"account-web/model"
	pb "github.com/Gentleelephant/proto-center/pb/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	GetHandlerMap    = make(map[string]gin.HandlerFunc)
	PostHandlerMap   = make(map[string]gin.HandlerFunc)
	PutHandlerMap    = make(map[string]gin.HandlerFunc)
	DeleteHandlerMap = make(map[string]gin.HandlerFunc)
)

// GetAccountList
// @Summary GetAccountList
// @GetAccountList get account list
// @ID GetAccountList
// @Tags GetAccountList
// @Accept json
// @Produce json
// @Param pageNo query uint true "The Request PageNo"
// @Param pageSize query uint true "The Request PageSize"
// @Success 200 {object} model.Response
// @Failure default {object} model.Response "Return if any error"
// @Header all {string} X-Request-Id "The unique id with this request"
// @Router /v1/account/list [post]
func GetAccountList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	GetHandlerMap["/account/list"] = GetAccountList
	pageNo := c.Query("pageNo")
	pageSize := c.Query("pageSize")
	uintPageNo, err := strconv.ParseUint(pageNo, 10, 64)
	if err != nil {
		return
	}
	uintPageSize, err := strconv.ParseUint(pageSize, 10, 64)
	if err != nil {
		return
	}
	list, err := initialize.AccountServiceClient.GetAccountList(c.Request.Context(), &pb.AccountPagingRequest{
		PageNo:   uint32(uintPageNo),
		PageSize: uint32(uintPageSize),
	})
	if err != nil {
		return
	}
	c.JSON(200, model.Response{
		Code:    200,
		Message: "success",
		Data:    list,
	})
}
