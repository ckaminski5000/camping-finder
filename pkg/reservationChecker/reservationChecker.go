package reservationchecker

// example url: https://www.recreation.gov/api/search?fq=campsite_type_of_use%3AOvernight&fq=campsite_type_of_use%3Ana&fq=entity_type%3Acampground&sort=available&start=0&size=19&exact=false&lat=34.176&lng=-118.9316&location=Newbury%20Park%2C%20California&radius=300&start_date=2023-09-08T00%3A00%3A00Z&end_date=2023-09-10T00%3A00%3A00Z&include_partially_available=false&include_notreservable=false&include_unavailable=false

import (
	"camping-finder/pkg/errors"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type UrlParams = struct {
	Lat                         string
	Lng                         string
	Location                    string
	State                       string
	Radius                      string
	Start_date                  string
	End_date                    string
	Include_partially_available string
	Include_notreservable       string
	Include_unavailable         string
}

// this function will send an api request to recreation.gov and find out if there are available campsites

var testParams = UrlParams{
	Lat:                         "34.176",
	Lng:                         "-118.9316",
	Location:                    "Newbury Park",
	State:                       "California",
	Radius:                      "300",
	Start_date:                  "2023-09-08T00:00:00Z",
	End_date:                    "2023-09-10T00:00:00Z",
	Include_partially_available: "false",
	Include_notreservable:       "false",
	Include_unavailable:         "false",
}

func GetRecAreaSuggestions(input string) interface{} {
	baseUrl := "https://www.recreation.gov/api/search/suggest"

	values := url.Values{}
	values.Add("q", input)
	values.Add("geocoder", "true")

	encodedURL := baseUrl + "?" + values.Encode()

	// Send GET request
	client := &http.Client{}
	req, err := http.NewRequest("GET", encodedURL, nil)
	if err != nil {
		err := errors.CreateError(http.StatusInternalServerError, "Error in creating request to get longitude and latitude")
		return err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Referer", "https://www.recreation.gov/")
	// Send GET request
	response, err := client.Do(req)
	if err != nil {
		err := errors.CreateError(http.StatusInternalServerError, "Error making the GET request:")
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return fmt.Sprintf("Error reading the response body: %s", err)
	}

	recreationGovResponse := RecSuggestionList{}

	err = json.Unmarshal(body, &recreationGovResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return fmt.Sprintf("Error decoding JSON: %s", err)
	}

	return transformSuggestionList(recreationGovResponse)
}

func transformSuggestionList(s RecSuggestionList) InventorySuggestions {

	inventorySuggestions := s.InventorySuggestions
	fmt.Println("suggestions1: ", inventorySuggestions[0])
	var recAreaSuggestions InventorySuggestions

	for i, _ := range inventorySuggestions {
		if inventorySuggestions[i].EntityType == "campground" {
			recAreaSuggestions = append(recAreaSuggestions, inventorySuggestions[i])
		}
	}

	return recAreaSuggestions
}

func GetLocationListByCity(params UrlParams) interface{} {
	baseUrl := "https://www.recreation.gov/api/search"

	values := url.Values{}
	values.Add("fq", "campsite_type_of_use:Overnight")
	values.Add("fq", "campsite_type_of_use:na")
	values.Add("fq", "entity_type:campground")
	values.Add("sort", "available")
	values.Add("start", "0")
	values.Add("size", "19")
	values.Add("exact", "false")
	values.Add("lat", params.Lat)
	values.Add("lng", params.Lng)
	values.Add("location", params.Location)
	values.Add("radius", params.Radius)
	values.Add("start_date", params.Start_date)
	values.Add("end_date", params.End_date)
	values.Add("include_partially_available", "false")
	values.Add("include_notreservable", "false")
	values.Add("include_unavailable", "false")

	encodedURL := baseUrl + "?" + values.Encode()
	fmt.Printf("encodedUrl: %s", encodedURL)

	// Send GET request
	client := &http.Client{}
	req, err := http.NewRequest("GET", encodedURL, nil)
	if err != nil {
		err := errors.CreateError(http.StatusInternalServerError, "Error in creating request to get longitude and latitude")
		return err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Referer", "https://www.recreation.gov/")
	// Send GET request
	response, err := client.Do(req)
	if err != nil {
		err := errors.CreateError(http.StatusInternalServerError, "Error making the GET request:")
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return fmt.Sprintf("Error reading the response body: %s", err)
	}

	recreationGovResponse := RecreationResponse{}

	err = json.Unmarshal(body, &recreationGovResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return fmt.Sprintf("Error decoding JSON: %s", err)
	}

	// return response to client
	return recreationGovResponse

}
