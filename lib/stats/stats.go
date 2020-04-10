package stats

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"

	caching "../caching"
	pconf "../config"
	structs "../structs"
)

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

// requestData does an HTTP GET request to the third party API that 
// contains covid-9 stats
// It returns []structs.Country and any write error encountered.
func requestData() ([]structs.Country, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", serverConf.API.URL, nil)
	if err != nil {
		return []structs.Country{}, err
	}

	res, resError := client.Do(req)
	if resError != nil {
		return []structs.Country{}, resError
	}
	defer res.Body.Close()

	b, readBoyError := ioutil.ReadAll(res.Body)
	if readBoyError != nil {
		return []structs.Country{}, readBoyError
	}

	keys := make([]structs.Country, 0)
	if errUnmarshal := json.Unmarshal(b, &keys); err != nil {
		return []structs.Country{}, errUnmarshal
	}

	return keys, nil
}

// GetAllCountries get n array of all countries that have
// Covid-19 stats (data starts from date 22/01/2020)
// Check if there are cached data if not does a HTTP
// request to the 3rd party API (check requestData())
// It returns structs.Countries ([] Country) and any write error encountered.
func GetAllCountries() (structs.Countries, error) {
	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()
	
	cachedData, cacheGetError := caching.Get(conn, "total")
	if cacheGetError != nil {
		return structs.Countries{}, cacheGetError
	}

	var s structs.Countries

	if len(cachedData.Data) == 0 {
		response, responseError := requestData()

		if responseError != nil {
			return structs.Countries{}, responseError
		}

		s = structs.Countries{Data: response}
		caching.Set(conn, s, "total")

	} else {
		return cachedData, nil
	}

	return s, nil
}

// GetCountry seach through an array of structs.Country and
// gets COVID-19 stats for that specific country
// It returns structs.Country and any write error encountered.
func GetCountry(name string) (structs.Country, error) {
	allCountries, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Country{}, allCountriesError
	}

	for _, v := range allCountries.Data {
		if v.Country == name {
			return v, nil
		}
	}

	return structs.Country{}, nil
}

// SortByCases sorts an array of Country structs by Country.Cases
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByCases() (structs.Countries, error) {

	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Cases != allCountries[j].Cases {
			return allCountries[i].Cases > allCountries[j].Cases
		}
		return allCountries[i].Deaths > allCountries[j].Deaths
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

// SortByDeaths sorts an array of Country structs by Country.Deaths
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByDeaths() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Deaths != allCountries[j].Deaths {
			return allCountries[i].Deaths > allCountries[j].Deaths
		}
		return allCountries[i].Cases > allCountries[j].Cases
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

// SortByTodayCases sorts an array of Country structs by Country.TodayCases
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByTodayCases() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayCases > allCountries[j].TodayCases
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

// SortByTodayDeaths sorts an array of Country structs by Country.TodayDeaths
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByTodayDeaths() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayDeaths > allCountries[j].TodayDeaths
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

// SortByRecovered sorts an array of Country structs by Country.Recovered
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByRecovered() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Recovered > allCountries[j].Recovered
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

// SortByActive sorts an array of Country structs by Country.Active
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByActive() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Active > allCountries[j].Active
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

// SortByCritical sorts an array of Country structs by Country.Critical
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByCritical() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Critical > allCountries[j].Critical
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

// SortByCasesPerOneMillion sorts an array of Country structs by Country.CasesPerOneMillion
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByCasesPerOneMillion() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].CasesPerOneMillion > allCountries[j].CasesPerOneMillion
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

// PercentancePerCountry gets a country's COVID-19 stats (getting the from GetCountry)
// and calculate today's total cases percentance and today's death percentance
// It returns structs.CountryStats and any write error encountered.
func PercentancePerCountry(name string) (structs.CountryStats, error) {
	country, countryError := GetCountry(name)
	if countryError != nil {
		return structs.CountryStats{}, nil
	}

	var todayPerCentOfTotalCases = country.TodayCases * 100 / country.Cases
	var todayPerCentOfTotalDeaths = country.TodayDeaths * 100 / country.Deaths

	countryStats := structs.CountryStats{Country: country.Country,
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths}

	return countryStats, nil
}

// GetTotalStats gets worlds COVID-19 total statistics.
// The statistics are total cases, total deaths today's total deaths
// totltoal cases, percentace totay increase in deaths and cases
// It returns structs.TotalStats and any write error encountered.
func GetTotalStats() (structs.TotalStats, error) {
	var totalDeaths = 0
	var totalCases = 0
	var todayTotalDeaths = 0
	var todayTotalCases = 0

	allCountriesArr, errorAllCountries := GetAllCountries()
	if errorAllCountries != nil {
		return structs.TotalStats{}, nil
	}

	allCountries := allCountriesArr.Data

	for _, v := range allCountries {
		if v.Country == "World" {
			continue
		}
		totalDeaths = totalDeaths + v.Deaths
		totalCases = totalCases + v.Cases
		todayTotalDeaths = todayTotalDeaths + v.TodayDeaths
		todayTotalCases = todayTotalCases + v.TodayCases
	}

	var todayPerCentOfTotalCases = todayTotalDeaths * 100 / totalDeaths
	var todayPerCentOfTotalDeaths = todayTotalCases * 100 / totalCases

	totalStatsStuct := structs.TotalStats{
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths,
		TotalCases:                totalCases,
		TotalDeaths:               totalDeaths,
		TodayTotalCases:           todayTotalCases,
		TodayTotalDeaths:          todayTotalDeaths,
	}

	return totalStatsStuct, nil
}

// GetAllCountriesName get names of the countries that we have Covid-19 stats
// It returns structs.AllCountriesName and any write error encountered.
func GetAllCountriesName() (structs.AllCountriesName, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.AllCountriesName{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	var counties []string

	for _, v := range allCountries {
		counties = append(counties, v.Country)
	}

	return structs.AllCountriesName{Countries: counties}, nil
}
