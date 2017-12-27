package hello

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	source := r.Header.Get("X-Appengine-Inbound-Appid")
	target := appengine.AppID(ctx)

	if source != target && appengine.IsDevAppServer() == false {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprint(w, "Hello")
}
