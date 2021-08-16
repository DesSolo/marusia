### marusia
Библиотека для создания скиллов голосового помощника [Маруся](https://marusia.mail.ru/)

### Пример
Простейший скилл - повторение реплик пользователя.
```golang
package main

import (
	"log"
	"github.com/DesSolo/marusia"
)

func echoHandler(resp *marusia.Response, req *marusia.Request) *marusia.Response {
	message := req.OriginalUtterance()
	resp.Text(message)
	resp.TTS(message)
	return resp
}

func main() {
	dr := marusia.NewDialogRouter(true)
	dr.RegisterDefault(echoHandler)

	config := marusia.NewConfig(
		false,
		"",
		"",
		":9000",
		"/webhook",
	)
	skill := marusia.NewSkill(config, dr)
	if err := skill.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

```