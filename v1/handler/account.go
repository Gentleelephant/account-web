package handler

import (
	"account-web/initialize"
	"account-web/model"
	pb "github.com/Gentleelephant/proto-center/pb/model"
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Tags Account
// @Accept json
// @Produce json
// @Param pageNo query uint true "The Request PageNo"
// @Param pageSize query uint true "The Request PageSize"
// @Success 200 {object} model.Response
// @Failure default {object} model.Response "Return if any error"
// @Header all {string} X-Request-Id "The unique id with this request"
// @Router /v1/account/list [get]
func GetAccountList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
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
	c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    list,
	})
}

// GetAccountByMobile
// @Summary GetAccountByMobile
// @GetAccountByMobile get account
// @ID GetAccountByMobile
// @Tags Account
// @Accept json
// @Produce json
// @Param mobile query string true "The user mobile"
// @Success 200 {object} model.Response
// @Failure default {object} model.Response "Return if any error"
// @Header all {string} X-Request-Id "The unique id with this request"
// @Router /v1/account/mobile [get]
func GetAccountByMobile(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	mobile := c.Query("mobile")
	account, err := initialize.AccountServiceClient.GetAccountByMobile(c.Request.Context(), &pb.MobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    account,
	})
}

// GetAccountByID
// @Summary GetAccountByID
// @GetAccountByID get account
// @ID GetAccountByID
// @Tags Account
// @Accept json
// @Produce json
// @Param id query int32 true "The user id"
// @Success 200 {object} model.Response
// @Failure default {object} model.Response "Return if any error"
// @Header all {string} X-Request-Id "The unique id with this request"
// @Router /v1/account/id [get]
func GetAccountByID(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	id := c.Param("id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}
	account, err := initialize.AccountServiceClient.GetAccountByID(c.Request.Context(), &pb.IDRequest{
		Id: int32(idNum),
	})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    account,
	})
}

func AddAccount(c *gin.Context) {

}

func UpdateAccount(c *gin.Context) {

}

func CheckPassword(c *gin.Context) {

}
