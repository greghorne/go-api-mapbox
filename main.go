package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	// "strconv"
	"strings"
	"fmt"
)




// ============================================================
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/v1/mapbox-isochrone/{lng}/{lat}/{time}/{appid}/{appcode}", v1HereIsochrone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8003", router))

}
// ============================================================


// ============================================================
func v1HereIsochrone (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var jsonResult map[string]string

	if isochrone, msg := v1DoHereIsochrone(params["lng"], params["lat"], params["time"], params["appid"], params["appcode"]); msg == "" {
		jsonResult = map[string]string{"here": isochrone}
	} else {
		jsonResult = map[string]string{"intersects": ""}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(jsonResult)
}
// ============================================================


// ============================================================
func v1DoHereIsochrone(sxLng string, syLat string, sTime string, sAppID string, sAppCode string) (geojson string, msg string) {

	https://api.mapbox.com/isochrone/v1/mapbox/driving/-97,36?contours_minutes=3&polygons=true&access_token=

	here_url := "https://api.mapbox.com/isochrone/v1/mapbox/driving/" + sxLng + "," + syLat + "?contours_minutes=" + sTime + "&polygons=true&access_token=" sToken
fmt.Println(here_url
)
	startSearchText := "[{id:0,shape:"
	endSearchText   := "}]}],start:"

	geojson = ""
	msg     = ""

	response, err := http.Get(here_url)
	if err == nil {
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			geojson = ""
			msg     = err.Error()
		} 

		jsonText := strings.Replace(string(body), "\"", "", -1)

		nStart   := strings.Index(jsonText, startSearchText) + len(startSearchText)
		nEnd     := strings.Index(jsonText, endSearchText)

		x := strings.Split(jsonText[nStart:nEnd], ",")

		var s []string
		for n := 0; n < len(x); n+=2 {
			
			switch num := n; num {
				case 0:
					s = append(s, (x[n] + "," + x[n+1] +"],"))
				case len(x) - 2:
					s = append(s, ("[" + x[n] + "," + x[n+1]))
				default:
					s = append(s, "[" + (x[n] + "," + x[n+1]) + "],")
			}
		}

		geojson = strings.Join(s, "")
	} 

	return
}
// ============================================================
