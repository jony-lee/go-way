package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
	"time"
)

// define email interface, and implemented auth and send method
type Mail interface {
	Auth()
	Send(message Message) error
}

type SendMail struct {
	User     string
	Password string
	Host     string
	Port     string
	auth     smtp.Auth
}

type Attachment struct {
	name        string
	contentType string
	withFile    bool
}

type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	attachment  Attachment
}

func loadJson(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
func (mail *SendMail) Auth() {
	mail.auth = smtp.PlainAuth("", mail.User, mail.Password, mail.Host)
}

func (mail SendMail) Send(message Message) error {
	mail.Auth()
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.from
	Header["To"] = strings.Join(message.to, ";")
	Header["Cc"] = strings.Join(message.cc, ";")
	Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)

	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.contentType + "\r\n"
	body += "\r\n" + message.body + "\r\n"
	buffer.WriteString(body)

	if message.attachment.withFile {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + message.attachment.name + "\"\r\n"
		buffer.WriteString(attachment)
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			}
		}()
		mail.writeFile(buffer, message.attachment.name)
	}

	buffer.WriteString("\r\n--" + boundary + "--")
	err := smtp.SendMail(mail.Host+":"+mail.Port, mail.auth, message.from, message.to, buffer.Bytes())
	if err != nil {
		panic(err)
	}
	return nil
}

func (mail SendMail) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

// read and write the file to buffer
func (mail SendMail) writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}

func main() {
	mail := SendMail{}
	loadJson("src/github.com/jony-lee/go-way/sci-hub/paper/config.json", &mail)
	//mail = &SendMail{user: "chunyunzeng@hotmail.com", password: "password", host: "smtp.mxhichina.com", port: "25"}
	message := Message{from: "jonylee_cn@163.com",
		to:          []string{"1528315273@qq.com"},
		cc:          []string{},
		bcc:         []string{},
		subject:     "HELLO WORLD",
		body:        "",
		contentType: "text/plain;charset=utf-8",
		attachment: Attachment{
			name:        "./pdf/1.pdf",
			contentType: "Content-Disposition",
			withFile:    true,
		},
	}
	//TODO 邮箱登陆认证出问题了
	if err := mail.Send(message); err != nil {
		panic(err)
	}

}
