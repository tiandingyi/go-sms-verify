package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tiandingyi/go-sms-verify/data"
)

// 应用程序的超时时间为 10 秒。
const appTimeout = time.Second * 10

//定义了两个函数 sendSMS 和 verifySMS，它们都是 Config 结构体的方法，返回类型为 gin.HandlerFunc。
// 这意味着这两个函数可以被 Gin 框架注册为 HTTP 请求处理函数。

// context.WithTimeout() 函数用于创建一个带有超时的上下文（context）。
// 这个函数非常有用，因为它可以在执行一些可能会花费较长时间的操作时，设置一个最大的等待时间。
// 如果在指定的时间内操作没有完成，上下文将被取消，从而可以避免程序陷入无限等待的状态。
func (app *Config) sendSMS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validateBody(ctx, &payload)

		newData := data.OTPData{PhoneNumber: payload.PhoneNumber}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusAccepted, "OTP Sent Successfully")

	}
}
func (app *Config) verifySMS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(ctx, &payload)

		newData := data.VerifyData{User: payload.User, Code: payload.Code}

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusAccepted, "OTP Verify Successfully")

	}
}
