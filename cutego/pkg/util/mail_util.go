package util

import "cutego/pkg/config"
import "github.com/go-gomail/gomail"

// BatchSendCode 批量发送验证码
// Param: targetUserEmails 接收者邮箱地址
// Param: code 验证码
func BatchSendCode(targetUserEmails []string, code string) error {
	mailConfig := config.AppCoreConfig.CuteGoConfig.Mail
	sender := mailConfig.Username
	authCode := mailConfig.Password
	mailTitle := "CuteGo验证码"
	mailBody := "您的验证码为: " + code

	message := gomail.NewMessage()
	message.SetHeader("From", sender)
	message.SetHeader("To", targetUserEmails...)
	message.SetHeader("Subject", mailTitle)
	message.SetBody("text/html", mailBody)

	// 添加附件
	//zipPath := "./xxxx.zip"
	//message.Attach(zipPath)

	dialer := gomail.NewDialer(mailConfig.Host, mailConfig.Port, sender, authCode)
	return dialer.DialAndSend(message)
}

// SendCode 发送验证码
// Param: targetUserEmail 接收者邮箱地址
// Param: code 验证码
func SendCode(targetUserEmail string, code string) error {
	targetUserEmails := []string{targetUserEmail}
	return BatchSendCode(targetUserEmails, code)
}
