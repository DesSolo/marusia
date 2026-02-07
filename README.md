### marusia
Библиотека для создания скиллов голосового помощника [Маруся](https://marusia.mail.ru/)

### Пример
Простейший скилл - повторение реплик пользователя.

```golang
package main

import (
	"log"
    "context"
	
	"github.com/DesSolo/marusia"
)

func echoHandler(_ context.Context, req *marusia.Request, r *marusia.Response) {
	message := req.OriginalUtterance()
	r.Text(message)
	r.TTS(message)
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
