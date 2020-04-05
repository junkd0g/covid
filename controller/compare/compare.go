package compare

/*
	Controller used for the endpoints:
		/compare
		/compare/firstdeath
		/compare/perday
*/

import (
	"encoding/json"

	curve "../../lib/curve"
	structs "../../lib/structs"

	"io/ioutil"
	"net/http"
)

//CompareRequest used for the https request's body
type CompareRequest struct {
	NameOne string `json:"countryOne"`
	NameTwo string `json:"countryTwo"`
}

//Perform used in the /compare endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.CompareDeathsCountries which get and return grobal statistics
//
//	CompareRequest used as the struct fro the request
//		example:
//			{
//				"countryOne" : "Italy",
//				"countryTwo" : "Spain"
//			}
//
//	Data structure that returns for two countries the names
//  and an array that contains deaths per day. It is sorted
//  and the first element is for the date 22/01/2020
//
//	In this JSON format
//  {
//    "countryOne": {
//        "country": "Spain",
//        "data": [
//            5,
//            10,
//		   	  17
//		   ]
//		},
//		"countryTwo": {
//      	"country": "Italy",
//       	"data": [
//            	197,
//            	233,
//				366
//			]
//		}
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func Perform(r *http.Request) ([]byte, int) {
	var compareRequest CompareRequest
	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	json.Unmarshal(b, &compareRequest)

	country := curve.CompareDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo)
	jsonBody, _ := json.Marshal(country)

	return jsonBody, 200
}

//PerformFromFirstDeath used in the /compare/firstdeath endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.CompareDeathsFromFirstDeathCountries which get
//	and return grobal statistics
//
//	CompareRequest used as the struct fro the request
//		example:
//			{
//				"countryOne" : "Italy",
//				"countryTwo" : "Spain"
//			}
//
//	Data structure that returns for two countries the names
//  and an array that contains total deaths per day. It is sorted
//  and the first element is for the date when the country
//  had their first death
//
//	In this JSON format
//  {
//    "countryOne": {
//        "country": "Spain",
//        "data": [
//            5,
//            10,
//		   	  17
//		   ]
//		},
//		"countryTwo": {
//      	"country": "Italy",
//       	"data": [
//            	197,
//            	233,
//				366
//			]
//		}
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func PerformFromFirstDeath(r *http.Request) ([]byte, int) {
	var compareRequest CompareRequest
	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	json.Unmarshal(b, &compareRequest)

	country := curve.CompareDeathsFromFirstDeathCountries(compareRequest.NameOne, compareRequest.NameTwo)
	jsonBody, _ := json.Marshal(country)

	return jsonBody, 200
}

//PerformPerDayDeath used in the /compare/firstdeath endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.ComparePerDayDeathsCountries which get
//	and return grobal statistics
//
//	CompareRequest used as the struct fro the request
//		example:
//			{
//				"countryOne" : "Italy",
//				"countryTwo" : "Spain"
//			}
//
//	Data structure that returns for two countries the names
//  and an array that contains deaths per day. It is sorted
//  and the first element is for the date when the country
//  had their first death
//
//	In this JSON format
//  {
//    "countryOne": {
//        "country": "Spain",
//        "data": [
//            5,
//            10,
//		   	  17
//		   ]
//		},
//		"countryTwo": {
//      	"country": "Italy",
//       	"data": [
//            	197,
//            	233,
//				366
//			]
//		}
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func PerformPerDayDeath(r *http.Request) ([]byte, int) {
	var compareRequest CompareRequest
	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	json.Unmarshal(b, &compareRequest)

	country := curve.ComparePerDayDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo)
	jsonBody, err := json.Marshal(country)
	if err != nil {
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
