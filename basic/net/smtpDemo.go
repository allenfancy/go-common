package main

import (
	"strings"
	"net/smtp"
	"fmt"
)

func SendToMail(user, pwd, host, to, subject, body, mailType string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, pwd, hp[0])
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type:text/" + mailType + ";charset=UTF-8"
	} else {
		contentType = "Content-Type:text/plain" + ";charset=UTF-8"
	}
	msg := []byte("To:" + to + "\r\n From: " + user + ">\r\n Subject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func main() {
	user := "verify@.com"
	password := ""
	host := "smtpdm.aliyun.com:25"
	to := "59871794@qq.com"
	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
