package smtp

import (
	config "CRUD/internal/app"
	logger "CRUD/internal/app/logs"
	auth2 "CRUD/internal/auth/services"
	"net/smtp"
)

type mailer struct {
	login, password, server, port string
}

func NewMailer() *mailer {
	configuration := config.GetConfig()
	return &mailer{
		login:    (*configuration).SMTP.Login,
		password: (*configuration).SMTP.Password,
		server:   (*configuration).SMTP.Server,
		port:     (*configuration).SMTP.Port,
	}
}

func (m *mailer) SendVerificationMail(email string) {
	configuration := config.GetConfig()
	address := m.server + ":" + m.port

	subject := "Subject: Youdomen email verification\n"
	body := "Чтобы подтвердить почту перейдите по ссылке: "
	hash := auth2.GetMD5Hash(email)
	link := "http://" + (*configuration).Server.Host + "/verification?email=" + email + "&sign=" + hash
	message := []byte(subject + body + link)

	auth := smtp.PlainAuth("", m.login, m.password, m.server)
	err := smtp.SendMail(address, auth, m.login, []string{email}, message)
	if err != nil {
		logger.Error.Println("Send mail: ", err)
		panic(err)
	}
}
