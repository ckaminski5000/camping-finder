package geocoding

import (
	"camping-finder/pkg/errors"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type UrlParams = struct {
	City    string
	State   string
	Country string
}

type LatLng = struct {
	Lat float64
	Lng float64
}

type Response interface{}

type Location struct {
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	State     string  `json:"state"`
}

// this function will send an api request to recreation.gov and find out if there are available campsites
func GetLngLat(city string, state string) (LatLng, error) {
	baseUrl := "https://api.api-ninjas.com/v1/geocoding"

	values := url.Values{}
	values.Add("city", city)
	values.Add("state", state)
	values.Add("country", "United States")

	encodedURL := baseUrl + "?" + values.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", encodedURL, nil)
	if err != nil {
		err := errors.CreateError(http.StatusInternalServerError, "Error in creating request to get longitude and latitude")
		return LatLng{}, err
	}
	req.Header.Add("X-Api-Key", os.Getenv("GEOCODING_API_KEY"))

	// Send GET request
	response, err := client.Do(req)
	if err != nil {
		err := errors.CreateError(http.StatusInternalServerError, "Error making the GET request:")
		return LatLng{}, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		err := errors.CreateError(http.StatusInternalServerError, "Error reading the response body while getting Long/Lat data")
		return LatLng{}, err
	}

	var jsonResponse []Location

	err = json.Unmarshal(responseBody, &jsonResponse)
	if err != nil {
		err := errors.CreateError(http.StatusInternalServerError, "Error unmarshaling JSON while getting Long/Lat data")
		return LatLng{}, err
	}

	fmt.Println("length: ", responseBody)
	fmt.Println("length: ", string(responseBody))

	if len(string(responseBody)) < 3 {
		err := errors.CreateError(http.StatusBadRequest, "We did not find that Location.  Try another location")
		return LatLng{}, err
	}
	return transformGeocodingResponse(jsonResponse), nil

}

func transformGeocodingResponse(json []Location) LatLng {
	location := LatLng{
		Lat: json[0].Latitude,
		Lng: json[0].Longitude,
	}
	return location
}
