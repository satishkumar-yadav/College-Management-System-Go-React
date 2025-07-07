package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

// from test site not gpt
func SendGmail(receiver string) {
	from := "satishkumar.yadav.7549@gmail.com"
	password := "ltcn wzqx ztov hycb" // smtp app password
	to := []string{"recipient@example.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" // 465 for ssl and 587 for tls

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		fmt.Println("TLS Dial Error:", err)
		return
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		fmt.Println("SMTP Client Error:", err)
		return
	}
	defer c.Close()

	if err = c.Auth(auth); err != nil {
		fmt.Println("SMTP Auth Error:", err)
		return
	}

	if err = c.Mail(from); err != nil {
		fmt.Println("SMTP Mail Error:", err)
		return
	}

	if err = c.Rcpt(to[0]); err != nil {
		fmt.Println("SMTP Rcpt Error:", err)
		return
	}

	wc, err := c.Data()
	if err != nil {
		fmt.Println("SMTP Data Error:", err)
		return
	}

	msg := []byte("Subject: Test email from Go!\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		"This is a test email sent from a Go program using Gmail SMTP.")

	_, err = wc.Write(msg)
	if err != nil {
		fmt.Println("SMTP Write Error:", err)
		return
	}

	err = wc.Close()
	if err != nil {
		fmt.Println("SMTP Close Error:", err)
		return
	}

	if err = c.Quit(); err != nil {
		fmt.Println("SMTP Quit Error:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
