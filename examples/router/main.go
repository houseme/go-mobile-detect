package main

import (
	"fmt"
	"net/http"
	"regexp"

	md "github.com/houseme/go-mobiledetect"
)

// route manager
type route struct {
	re      *regexp.Regexp
	handler func(http.ResponseWriter, *http.Request, []string, *md.MobileDetect)
}

// RouterHandler .
type RouterHandler struct {
	routes []*route
	detect *md.MobileDetect
}

// AddRoute .
func (h *RouterHandler) AddRoute(re string, handler func(http.ResponseWriter, *http.Request, []string, *md.MobileDetect)) {
	r := &route{regexp.MustCompile(re), handler}
	h.routes = append(h.routes, r)
}

func (h *RouterHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.detect = md.NewMobileDetect(r, nil)
	for _, route := range h.routes {
		matches := route.re.FindStringSubmatch(r.URL.String())
		if matches != nil {
			route.handler(rw, r, matches, h.detect)
			break
		}
	}
}

func homepageHandler(w http.ResponseWriter, r *http.Request, matches []string, detect *md.MobileDetect) {
	fmt.Fprint(w, "Hello World\n")
	fmt.Fprintf(w, "Matches %+v\n", matches)
	fmt.Fprintf(w, "Is Mobile? %+v\n", detect.IsMobile())
	fmt.Fprintf(w, "Is Tablet? %+v\n", detect.IsTablet())
}

func main() {
	reHandler := new(RouterHandler)
	reHandler.AddRoute("/device/[mobile|desktop]", homepageHandler)
	http.ListenAndServe(":9999", reHandler)
}
