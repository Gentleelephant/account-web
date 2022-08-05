package handler

import (
	"account-web/initialize"
	"account-web/model"
	pb "github.com/Gentleelephant/proto-center/pb/service/account"
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Router /v1/account/page [get]
func GetAccountList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	req := &pb.AccountPagingRequest{}
	err = c.ShouldBindQuery(req)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	list, err := initialize.AccountServiceClient.GetAccountList(c.Request.Context(), req)
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
// @Router /v1/account [get]
func GetAccountByMobile(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	req := &pb.MobileRequest{}
	err = c.ShouldBindQuery(req)
	if err != nil {
		return
	}
	account, err := initialize.AccountServiceClient.GetAccountByMobile(c.Request.Context(), req)
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
// @Param id path int32 true "The user id"
// @Success 200 {object} model.Response
// @Failure default {object} model.Response "Return if any error"
// @Header all {string} X-Request-Id "The unique id with this request"
// @Router /v1/account/[id] [get]
func GetAccountByID(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	req := &pb.IDRequest{}
	err = c.ShouldBindJSON(req)
	if err != nil {
		return
	}
	account, err := initialize.AccountServiceClient.GetAccountByID(c.Request.Context(), req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    account,
	})
}

// AddAccount
// @Summary AddAccount
// @AddAccount add account
// @ID AddAccount
// @Tags Account
// @Accept json
// @Produce json
// @Param body  account pb.AddAccountRequest true "The account"
// @Success 200 {object} model.Response
// @Failure default {object} model.Response "Return if any error"
// @Header all {string} X-Request-Id "The unique id with this request"
// @Router /v1/account [post]
func AddAccount(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	var account *pb.AddAccountRequest
	if err := c.ShouldBindJSON(&account); err != nil {
		return
	}
	resp, err := initialize.AccountServiceClient.AddAccount(c.Request.Context(), &pb.AddAccountRequest{
		Mobile:   account.Mobile,
		Password: account.Password,
		Nickname: account.Nickname,
		Gender:   account.Gender,
	})
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	})

}

// UpdateAccount
// @Summary UpdateAccount
// @AddAccount update account
// @ID UpdateAccount
// @Tags Account
// @Accept json
// @Produce json
// @Param body account pb.UpdateAccountRequest true "The account"
// @Success 200 {object} model.Response
// @Failure default {object} model.Response "Return if any error"
// @Header all {string} X-Request-Id "The unique id with this request"
// @Router /v1/account [put]
func UpdateAccount(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	var account *pb.UpdateAccountRequest
	if err = c.ShouldBindJSON(&account); err != nil {
		return
	}
	resp, err := initialize.AccountServiceClient.UpdateAccount(c.Request.Context(), &pb.UpdateAccountRequest{
		Id:       account.Id,
		Mobile:   account.Mobile,
		Password: account.Password,
		Nickname: account.Nickname,
		Gender:   account.Gender,
	})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

// CheckPassword
// @Summary CheckPassword
// @CheckPassword Check password
// @ID CheckPassword
// @Tags Account
// @Accept json
// @Produce json
// @Param body info pb.CheckPasswordRequest true "The password info"
// @Success 200 {object} model.Response
// @Failure default {object} model.Response "Return if any error"
// @Header all {string} X-Request-Id "The unique id with this request"
// @Router /v1/account [post]
func CheckPassword(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	var account *pb.CheckPasswordRequest
	if err = c.ShouldBindJSON(&account); err != nil {
		return
	}
	resp, err := initialize.AccountServiceClient.CheckPassword(c.Request.Context(), &pb.CheckPasswordRequest{
		Password:     account.Password,
		HashPassword: account.HashPassword,
		AccountId:    account.AccountId,
	})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}
