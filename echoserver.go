package echoserver

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
)

type Handler struct {
	Message string
}

func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	dump, err := httputil.DumpRequest(r, false)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	body, _ := io.ReadAll(r.Body)

	data := fmt.Sprintf(`CLIENT VALUE:
Remote Address: %s

HTTP VALUE:
%s

HTTP BODY:
%s

%s
`,
		r.RemoteAddr,
		strings.TrimSpace(string(dump)),
		string(body),
		h.Message)
	rw.Write([]byte(data))
}
