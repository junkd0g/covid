package sortcon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	applogger "github.com/junkd0g/covid/lib/applogger"
	mcountry "github.com/junkd0g/covid/lib/model/country"
	stats "github.com/junkd0g/covid/lib/stats"
	merror "github.com/junkd0g/neji"
)

//SortRequest used for the https request's body
type SortRequest struct {
	Type string `json:"type"`
}

/*
	POST request to /api/sort endpoint

	Request:

	{
		"type" : "deaths"
	}

	Response

	{
    	"data": [{
        	"country": "Italy",
            "cases": 124632,
            "todayCases": 4805,
            "deaths": 15362,
            "todayDeaths": 681,
            "recovered": 20996,
            "active": 88274,
            "critical": 3994,
			"casesPerOneMillion": 2061,
			"tests": 21298974,
            "testsPerOneMillion": 64371
        },
        {
            "country": "Spain",
            "cases": 124736,
            "todayCases": 5537,
            "deaths": 11744,
            "todayDeaths": 546,
            "recovered": 34219,
            "active": 78773,
            "critical": 6416,
			"casesPerOneMillion": 2668,
			"tests": 21298974,
            "testsPerOneMillion": 64371
		}]
	}

*/
func Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "sortcon", "Handle",
		"Endpoint /api/sort called with response JSON body "+string(jsonBody), status, elapsed)
}

//Perform used in the /sort endpoint's handle to return
//	the structs.Countries struct as a json response by calling
//	stats.SortByDeaths() or tats.GetAllCountries() or stats.GetAllCountries()
//  or stats.SortByCasesPerOneMillion() or stats.SortByCritical() or
//  stats.SortByActive() or stats.SortByRecovered() or stats.SortByTodayDeaths()
//  or stats.SortByTodayCases() or stats.SortByCases()
//  which get and return sorted by field data: array
//
//	CompareRequest used as the struct for the request
//		example:
//			{
//				"type" : "deaths"
//			}
//
//	In this JSON format
//	{
//		"data": [{
//			"country": "Italy",
//			"cases": 124632,
//			"todayCases": 4805,
//			"deaths": 15362,
//			"todayDeaths": 681,
//			"recovered": 20996,
//			"active": 88274,
//			"critical": 3994,
//			"casesPerOneMillion": 2061
//		},
//		{
//			"country": "Spain",
//			"cases": 124736,
//			"todayCases": 5537,
//			"deaths": 11744,
//			"todayDeaths": 546,
//			"recovered": 34219,
//			"active": 78773,
//			"critical": 6416,
//			"casesPerOneMillion": 2668
//		}]
//	}
//
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func perform(r *http.Request) ([]byte, int) {
	var sortRequest SortRequest

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "sortcon", "perform", errIoutilReadAll.Error())
		statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(400, errIoutilReadAll)
		return statsErrJSONBody, 400
	}

	unmarshallError := json.Unmarshal(b, &sortRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "sortcon", "perform", unmarshallError.Error())
		statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(400, unmarshallError)
		return statsErrJSONBody, 400
	}

	sortType := sortRequest.Type
	var countries mcountry.Countries
	var countriesError error

	switch sortType {
	case "deaths":
		countries, countriesError = stats.SortByDeaths()
		if countriesError != nil {
			applogger.Log("ERROR", "sortcon", "perform", "Deaths sorting error: "+countriesError.Error())
			statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, countriesError)
			return statsErrJSONBody, 500
		}
	case "cases":
		countries, countriesError = stats.SortByCases()
		if countriesError != nil {
			applogger.Log("ERROR", "sortcon", "perform", "Cases sorting error: "+countriesError.Error())
			statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, countriesError)
			return statsErrJSONBody, 500
		}
	case "todayCases":
		countries, countriesError = stats.SortByTodayCases()
		if countriesError != nil {
			applogger.Log("ERROR", "sortcon", "perform", "Today cases sorting error: "+countriesError.Error())
			statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, countriesError)
			return statsErrJSONBody, 500
		}
	case "todayDeaths":
		countries, countriesError = stats.SortByTodayDeaths()
		if countriesError != nil {
			applogger.Log("ERROR", "sortcon", "perform", "Today deaths sorting error: "+countriesError.Error())
			statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, countriesError)
			return statsErrJSONBody, 500
		}
	case "recovered":
		countries, countriesError = stats.SortByRecovered()
		if countriesError != nil {
			applogger.Log("ERROR", "sortcon", "perform", "Recovered sorting error: "+countriesError.Error())
			statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, countriesError)
			return statsErrJSONBody, 500
		}
	case "active":
		countries, countriesError = stats.SortByActive()
		if countriesError != nil {
			applogger.Log("ERROR", "sortcon", "perform", "Active sorting error: "+countriesError.Error())
			statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, countriesError)
			return statsErrJSONBody, 500
		}
	case "critical":
		countries, countriesError = stats.SortByCritical()
		if countriesError != nil {
			applogger.Log("ERROR", "sortcon", "perform", "Critical sorting error: "+countriesError.Error())
			statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, countriesError)
			return statsErrJSONBody, 500
		}
	case "casesPerOneMillion":
		countries, countriesError = stats.SortByCasesPerOneMillion()
		if countriesError != nil {
			applogger.Log("ERROR", "sortcon", "perform", "Cases per one million sorting error: "+countriesError.Error())
			statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, countriesError)
			return statsErrJSONBody, 500
		}
	default:
		countries, countriesError = stats.GetAllCountries()
		if countriesError != nil {
			applogger.Log("ERROR", "sortcon", "perform", "Default sorting error: "+countriesError.Error())
			statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, countriesError)
			return statsErrJSONBody, 500
		}
	}

	jsonBody, err := json.Marshal(countries)
	if err != nil {
		applogger.Log("ERROR", "sortcon", "perform", err.Error())
		errorJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, err)
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
