package flash

import "net/http"

func ClearFlashMessage(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clearCookie(w, "alert")
		clearCookie(w, "notice")
		h(w, r)
	})
}

func clearCookie(w http.ResponseWriter, name string) {
	c := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	}
	http.SetCookie(w, c)
}

type Flash struct {
	Alert     string
	HasAlert  bool
	Notice    string
	HasNotice bool
}

func NewFlashFromCookie(r *http.Request) *Flash {
	var alertMsg string
	alert, _ := r.Cookie("alert")
	if alert != nil {
		alertMsg = alert.Value
	}
	var noticeMsg string
	notice, _ := r.Cookie("notice")
	if notice != nil {
		noticeMsg = notice.Value
	}
	return newFlash(alertMsg, noticeMsg)
}

func newFlash(alert, notice string) *Flash {
	return &Flash{
		Alert:     alert,
		HasAlert:  len(alert) > 0,
		Notice:    notice,
		HasNotice: len(notice) > 0,
	}
}

func SetAlert(w http.ResponseWriter, alert string) {
	setFlash(w, alert, "alert")
}

func SetNotice(w http.ResponseWriter, notice string) {
	setFlash(w, notice, "notice")
}

func setFlash(w http.ResponseWriter, msg string, name string) {
	c := &http.Cookie{
		Name:     name,
		Value:    msg,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   100,
	}
	http.SetCookie(w, c)
}
