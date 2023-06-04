package models

type Client struct {
	UserName     string `form:"username" binding:"required"`
	Password     string `form:"password" binding:"required"`
	GrantType    string `form:"grant_type" binding:"required"`
	ClientId     string `form:"client_id" binding:"required"`
	ClientSecret string `form:"client_secret" binding:"required"`
}
