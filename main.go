package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/idna"
)

var htmlt = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Test A tag with unicode DNS</title>
  </head>
  <body>
    <a href="http://{{.IDNAName}}:8090" target="_blank">Link to ßan-jöśè.com IDNA</a>
	<br />
	<a href="http://{{.Name}}:8090" target="_blank">Link to ßan-jöśè.com</a>
  </body>
</html>
`

var t = template.New("index.html")

func init() {
	t.Parse(htmlt)
}

func getASCII(s string) string {
	id := idna.New(idna.Transitional(false))
	ss, _ := id.ToASCII(s)
	return ss
}

func getASCIITransitional(s string) string {
	id := idna.New(idna.Transitional(true))
	ss, _ := id.ToASCII(s)
	return ss
}
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "text/html")
	type A struct {
		Name     string
		IDNAName string
	}
	unicodeDNSName := "ßan-jöśè.com"
	uDNSNameIDNAASCII, _ := idna.ToASCII(unicodeDNSName)
	fmt.Println(uDNSNameIDNAASCII)
	t.Execute(w, A{
		IDNAName: uDNSNameIDNAASCII,
		Name:     unicodeDNSName,
	})
	fmt.Println("Got domain name", r.Host)
	fmt.Println(idna.ToUnicode(uDNSNameIDNAASCII))
	fmt.Println(getASCII(unicodeDNSName), getASCIITransitional(unicodeDNSName))
	h := strings.Split(r.Host, ":")[0]
	fmt.Println(idna.ToUnicode(h))
}
func main() {
	r := httprouter.New()
	r.GET("/", Index)
	http.ListenAndServe("0.0.0.0:8090", r)
}
