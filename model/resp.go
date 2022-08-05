package model

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AddAccount struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
}

type UpdateAccount struct {
	Id       uint32 `json:"id"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
}
