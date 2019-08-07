package main

import (
	"fmt"
	"log"
	"net/http"
)

type MatomoVisitLog struct {
	IdVisit              string `json:"idvisit"`
	IdSite               string `json:"idsite"`
	IdVisitor            string `json:"idvisitor"`
	VisitFirstActionTime string `json:"visit_first_action_time"`
	VisitLastActionTime  string `json:"visit_last_action_time"`
	RefererUrl           string `json:"referer_url"`
	Url                  string `json:"url"`
	Longitude            string `json:"longitude"`
	Latitude             string `json:"latitude"`
	Language             string `json:"language"`
}

func main() {
	http.HandleFunc("/piwik.php", piwik)
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func piwik(w http.ResponseWriter, r *http.Request) {
	var buf []byte
	n, _ := r.Body.Read(buf)
	defer r.Body.Close()
	log.Printf("Body : %v\n", string(n))
	log.Printf("Query : %v\n", r.URL.Query().Encode())
	log.Printf("Header : %+#v\n", r.Header)
	log.Printf("RemoteAddr : %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "-")
}

type Sinker interface {
	Sink()
}
