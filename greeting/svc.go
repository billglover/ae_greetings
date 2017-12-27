package greeting

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/en/", handlerEN)
	http.HandleFunc("/zn/", handlerZN)
}

func handlerEN(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	g, err := getGreeting(ctx, "hello")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(g)
}

func handlerZN(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	g, err := getGreeting(ctx, "nihao")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(g)
}

func getGreeting(ctx context.Context, svcName string) ([]byte, error) {
	client := urlfetch.Client(ctx)

	hostname, err := appengine.ModuleHostname(ctx, svcName, "", "")
	if err != nil {
		return nil, fmt.Errorf("unable to find service %s", svcName)
	}

	scheme := "https"
	if appengine.IsDevAppServer() {
		scheme = "http"
	}

	req, _ := http.NewRequest("GET", scheme+"://"+hostname, nil)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to query internal service")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response from internal service")
	}

	return body, nil
}
