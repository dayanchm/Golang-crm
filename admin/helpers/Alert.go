package helpers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("123123"))

func SetAlert(w http.ResponseWriter, r *http.Request, message string) error {
	sessions, err := store.Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return err
	}
	sessions.AddFlash(message)
	return sessions.Save(r, w)
}

func GetAlert(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	sessions, err := store.Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data := make(map[string]interface{})
	flashes := sessions.Flashes()

	if len(flashes) > 0 {
		data["is_alert"] = true
		data["message"] = flashes[0]
	} else {
		data["is_alert"] = false
		data["message"] = nil
	}

	sessions.Save(r, w)

	return data
}
