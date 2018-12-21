package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"strings"
)




// ============================================================
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/v1/mapbox-isochrone/{lng}/{lat}/{time}/{token}", v1MapboxIsochrone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8004", router))

}
// ============================================================


// ============================================================
func v1MapboxIsochrone (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var jsonResult map[string]string

	if isochrone, msg := v1DoMapboxIsochrone(params["lng"], params["lat"], params["time"], params["token"]); msg == "" {
		jsonResult = map[string]string{"mapbox": isochrone}
	} else {
		jsonResult = map[string]string{"mapbox": msg}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(jsonResult)
}
// ============================================================


// ============================================================
func v1DoMapboxIsochrone(sxLng string, syLat string, sTime string, sToken string) (geojson string, msg string) {

	mapbox_url := "https://api.mapbox.com/isochrone/v1/mapbox/driving/" + sxLng + "," + syLat + "?contours_minutes=" + sTime + "&polygons=true&access_token=" + sToken

	startSearchText := "\"geometry\":{\"coordinates\":"
	endSearchText   := ",\"type\":\"Polygon\""

	geojson = ""
	msg     = ""

	response, err := http.Get(mapbox_url)
	if err == nil {
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			geojson = ""
			msg     = err.Error()
		} 

		jsonText := string(body)

		nStart := strings.Index(jsonText, startSearchText) + len(startSearchText)
		nEnd   := strings.Index(jsonText, endSearchText)
		data   := jsonText[nStart:nEnd]

		x := strings.Split(data, ",")
		var s []string
		
		for n := 0; n < len(x); n+=2 {
			s = append(s, strings.Replace("[" + strings.Replace(x[n+1], "[", "", -1), "]", "", -1) + "," + strings.Replace(strings.Replace(x[n], "[", "", -1), "]", "", -1) + "]")
		}

		geojson = "[" + strings.Join(s, ",") + "]"

	} 

	return
}
// ============================================================
