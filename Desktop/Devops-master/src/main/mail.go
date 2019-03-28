/*package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// 메일서버 로그인 정보 설정
	auth := smtp.PlainAuth("", "sender@live.com", "pwd", "smtp.live.com")

	from := "sender@live.com"
	to := []string{"jh3859025@gmail.com"} // 복수 수신자 가능

	// 메시지 작성
	headerSubject := "Subject: 테스트\r\n"
	headerBlank := "\r\n"
	body := "안녕하세요 GOLang 메일 테스트 입니다 \r\n"
	msg := []byte(headerSubject + headerBlank + body)

	// 메일 보내기
	err := smtp.SendMail("smtp.live.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}

	fmt.Println("complete to send mail")
}

*/

package main

import (
	"net/smtp"
)

func main() {

	auth := smtp.PlainAuth("", "jhyoon@sender.com", "call", "smtp.live.com")
	from := "jhyoon@sender.com"
	to := []string{"jh3859025@gmail.com"}

	headerSubject := "test1\r\n"
	headerBlank := "\r\n"
	body := "안녕하세요 윤지훈입니다. 새해복 많이받으세요 !!"
	msg := []byte(headerBlank + headerSubject + body)

	err := smtp.SendMail("smtp.live.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}
}
