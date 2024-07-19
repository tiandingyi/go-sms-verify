package api

import "github.com/gin-gonic/gin"

// gin.Engine 是 Gin 框架中的核心结构，用于创建和管理路由。
type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/otp", app.sendSMS())
	app.Router.POST("/verifyOTP", app.verifySMS())
}
