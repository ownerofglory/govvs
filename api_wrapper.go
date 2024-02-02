package govvs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURLJourney  = "https://www3.vvs.de/mngvvs/XML_TRIP_REQUEST2"
	baseURLWidgetDM = "http://www3.vvs.de/vvs/widget/XML_DM_REQUEST"
)

// GetJourney initiates a request to retrieve journey information between two locations at a specified time.
// It uses the VVS (Verkehrs- und Tarifverbund Stuttgart) API to calculate routes, taking into account various transport options and preferences.
//
// Parameters:
// - r JourneyRequest: Struct containing the essential information for the journey request.
//    - OrigId: The ID or code representing the origin location for the journey. This should be in the format recognized by the VVS API, e.g., "de:08111:2599".
//    - DstId: The ID or code representing the destination location for the journey, formatted similarly to OrigId.
//    - TimeAt: Pointer to a time.Time object specifying the desired departure or arrival time for the journey.
//    - Limit: Pointer to an integer specifying the maximum number of journey options to return. If nil, a default value defined by the API will be used.
//    - LangCode: Pointer to a string specifying the language code for the response, e.g., "de" for German. If nil, the API's default language will be used.
//
// - reqParams ...ReqParam: An optional slice of ReqParam structs allowing for additional parameters to be specified for the journey request. These can include:
//    - Name: The name of the parameter, matching the constant names defined for VVS API parameters, e.g., ParamCalcOneDirection.
//    - Value: The value of the parameter, as a string, e.g., "1" for true or specific values like "EPSG:4326" for coordinate formats.
//
// The function supports various optional parameters that can be used to refine the journey search, including but not limited to:
// - ChangeSpeed: Adjusts the assumed speed of changing between transport modes.
// - RouteType: Determines the criteria for route selection, such as "LEASTTIME" for the shortest overall journey time.
// - UseRealtime: Specifies whether to include real-time data in the journey calculation.
//
// Returns:
// - *JourneyResponse: A pointer to a JourneyResponse struct containing the journey options found based on the request parameters. This includes detailed route information, timings, and any applicable fare information.
// - error: An error object that will be non-nil if there was an issue processing the request or communicating with the VVS API.
//
// Example usage:
// journeyReq := JourneyRequest{
//     OrigId: "de:08111:2599",
//     DstId: "de:08111:6075",
//     TimeAt: time.Now(),
//     Limit: 10,
//     LangCode: "de",
// }
// journeyResp, err := GetJourney(journeyReq, ReqParam{Name: ParamRouteType, Value: "leasttime"})
// if err != nil {
//     // Handle error
// }
// // Process journeyResp
func GetJourney(r JourneyRequest, reqParams ...ReqParam) (*JourneyResponse, error) {
	// Prepare query parameters
	params := url.Values{}

	// set default params
	setDefaultJourneyReqParams(params)

	// Set dynamic values separately
	timeAt := time.Now()
	if r.TimeAt != nil {
		timeAt = *r.TimeAt
	}
	orig := r.OrigId
	dst := r.DstId

	params.Set(ParamItdDate, timeAt.Format("20060102"))
	params.Set(ParamItdTime, timeAt.Format("1504"))
	params.Set(ParamNameDestination, dst)
	params.Set(ParamNameOrigin, orig)

	// override optional params6
	overrideReqParams(params, reqParams...)

	// Build the URL with encoded query parameters
	url := baseURLJourney + "?" + params.Encode()

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	// Unmarshal the JSON response into the JourneyResponse struct
	var journeyResponse JourneyResponse
	err = json.Unmarshal(body, &journeyResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return nil, err
	}

	// Return the populated JourneyResponse struct
	return &journeyResponse, nil
}

// GetArrivals sends a request to retrieve arrival information for a specified station or location.
// The function takes an ArrivalRequest struct containing essential request parameters and
// a variable number of ReqParam structs to specify additional request options.
//
// Parameters:
// - r ArrivalRequest: A struct containing the base request parameters such as station ID and request time.
//   - OrigId (string): The origin station identifier. Not used in this context but can be set for consistency.
//   - DstId (string): The destination station identifier. This is typically the station ID for which arrivals are requested.
//   - TimeAt (*time.Time): The specific date and time for which arrival information is requested. If nil, current time is assumed.
//   - Limit (*int): The maximum number of arrival entries to retrieve. If nil, a default limit is applied.
//   - LangCode (*string): The language code for the response (e.g., "de" for German). If nil, a default language is used.
//
// - reqParams ...ReqParam: Optional parameters to further customize the request. Each ReqParam consists of a Name and Value.
//   Supported optional parameters (refer to constants for parameter names):
//   - ParamLocationServerActive: Activates the location server for processing the request.
//   - ParamLsShowTrainsExplicit: Specifies whether trains should be shown explicitly in the response.
//   - ParamStateless: Indicates if the request should be processed without maintaining state.
//   - ParamLanguage: Sets the language of the response.
//   - ParamSpEncId: Security parameter, typically "0".
//   - ParamAnySigWhenPerfectNoOtherMatches: Used to refine search results.
//   - ParamLimit: Controls the maximum number of results returned.
//   - ParamDepArr: Indicates whether departure or arrival times are requested ("arrival" for arrivals).
//   - ParamTypeDm: Specifies the type of data management request.
//   - ParamAnyObjFilterDm: Filters objects based on specified criteria.
//   - ParamDeleteAssignedStops: Indicates whether assigned stops should be removed from the response.
//   - ParamNameDm: The name or ID of the station for which arrivals are being requested.
//   - ParamMode: Specifies the mode of transport.
//   - ParamDmLineSelectionAll: Determines if all lines should be selected in the request.
//   - ParamUseRealtime: Specifies whether real-time data should be used.
//   - ParamOutputFormat: Defines the output format of the response.
//   - ParamCoordOutputFormat: Sets the coordinate format for the response.
//   - ParamItdDateTimeDepArr: Specifies the desired time for departure or arrival.
//   - Additional parameters related to date and time of the request: ParamItdDateYear, ParamItdDateMonth, ParamItdDateDay, ParamItdTimeHour, ParamItdTimeMinute.
//
// Returns:
// - (*ArrivalResponse, error): A pointer to an ArrivalResponse struct containing the requested arrival information, or an error if the request fails.
//
// Example usage:
// response, err := GetArrivals(ArrivalRequest{DstId: "stationID", TimeAt: &time.Now()}, ReqParam{Name: ParamLimit, Value: "5"})
// if err != nil {
//     // handle error
// }
// // process response
func GetArrivals(r ArrivalRequest, reqParams ...ReqParam) (*ArrivalResponse, error) {
	// Prepare query parameters
	params := url.Values{}

	setDefaultArrivalReqParams(params)

	year := r.TimeAt.Format("2006")
	month := r.TimeAt.Format("01")
	day := r.TimeAt.Format("02")
	hour := r.TimeAt.Format("15")
	minute := r.TimeAt.Format("04")
	stationId := r.StationId

	params.Set(ParamItdDateYear, year)
	params.Set(ParamItdDateMonth, month)
	params.Set(ParamItdDateDay, day)
	params.Set(ParamItdTimeHour, hour)
	params.Set(ParamItdTimeMinute, minute)
	params.Set(ParamNameDm, stationId)

	overrideReqParams(params, reqParams...)

	// Build the URL with encoded query parameters
	url := baseURLWidgetDM + "?" + params.Encode()

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	// Unmarshal the JSON response into the ArrivalResponse struct
	var arrivalResponse ArrivalResponse
	err = json.Unmarshal(body, &arrivalResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return nil, err
	}

	// Return the populated ArrivalResponse struct
	return &arrivalResponse, nil
}

// GetDepartures sends a request to retrieve departure information for a specified station or location.
// The function takes a DepartureRequest struct containing essential request parameters and
// a variable number of ReqParam structs to specify additional request options.
//
// Parameters:
// - r DepartureRequest: A struct containing the base request parameters such as station ID and request time.
//   - OrigId (string): The origin station identifier. In the context of departures, this is the station ID from which departures are requested.
//   - DstId (string): The destination station identifier. Not typically used for departure requests but can be set for API consistency.
//   - TimeAt (*time.Time): The specific date and time for which departure information is requested. If nil, the current time is assumed.
//   - Limit (*int): The maximum number of departure entries to retrieve. If nil, a default limit is applied.
//   - LangCode (*string): The language code for the response (e.g., "de" for German). If nil, a default language is used.
//
// - reqParams ...ReqParam: Optional parameters to further customize the request. Each ReqParam consists of a Name and Value.
//   Supported optional parameters (refer to constants for parameter names):
//   - ParamLocationServerActive: Activates the location server for processing the request.
//   - ParamLsShowTrainsExplicit: Specifies whether trains should be shown explicitly in the response.
//   - ParamStateless: Indicates if the request should be processed without maintaining state.
//   - ParamLanguage: Sets the language of the response.
//   - ParamSpEncId: Security parameter, typically "0".
//   - ParamAnySigWhenPerfectNoOtherMatches: Used to refine search results.
//   - ParamLimit: Controls the maximum number of results returned.
//   - ParamDepArr: Indicates whether departure or arrival times are requested ("departure" for departures).
//   - ParamTypeDm: Specifies the type of data management request.
//   - ParamAnyObjFilterDm: Filters objects based on specified criteria.
//   - ParamDeleteAssignedStops: Indicates whether assigned stops should be removed from the response.
//   - ParamNameDm: The name or ID of the station for which departures are being requested.
//   - ParamMode: Specifies the mode of transport.
//   - ParamDmLineSelectionAll: Determines if all lines should be selected in the request.
//   - ParamUseRealtime: Specifies whether real-time data should be used.
//   - ParamOutputFormat: Defines the output format of the response.
//   - ParamCoordOutputFormat: Sets the coordinate format for the response.
//   - Additional parameters related to date and time of the request: ParamItdDateYear, ParamItdDateMonth, ParamItdDateDay, ParamItdTimeHour, ParamItdTimeMinute.
//
// Returns:
// - (*DepartureResponse, error): A pointer to a DepartureResponse struct containing the requested departure information, or an error if the request fails.
//
// Example usage:
// response, err := GetDepartures(DepartureRequest{OrigId: "stationID", TimeAt: &time.Now()}, ReqParam{Name: ParamLimit, Value: "10"})
// if err != nil {
//     // handle error
// }
// // process response
func GetDepartures(r DepartureRequest, reqParams ...ReqParam) (*DepartureResponse, error) {
	// Assemble the query parameters
	params := url.Values{}

	setDefaultDepartureReqParams(params)

	year := r.TimeAt.Format("2006")
	month := r.TimeAt.Format("01")
	day := r.TimeAt.Format("02")
	hour := r.TimeAt.Format("15")
	minute := r.TimeAt.Format("04")
	stationId := r.StationId

	params.Set(ParamItdDateYear, year)
	params.Set(ParamItdDateMonth, month)
	params.Set(ParamItdDateDay, day)
	params.Set(ParamItdTimeHour, hour)
	params.Set(ParamItdTimeMinute, minute)
	params.Set(ParamNameDm, stationId)

	overrideReqParams(params, reqParams...)

	// Build the complete URL with the encoded query parameters
	fullURL := baseURLWidgetDM + "?" + params.Encode()

	// Perform the GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// Unmarshal the JSON response into the DepartureResponse struct
	var departureResponse DepartureResponse
	err = json.Unmarshal(body, &departureResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return &departureResponse, nil
}

func overrideReqParams(urlParams url.Values, params ...ReqParam) {
	for _, param := range params {
		urlParams.Set(param.Name, param.Value)
	}
}

func setDefaultArrivalReqParams(urlParams url.Values) {
	for param, defaultValue := range defaultArrivalParams {
		urlParams.Set(param, defaultValue)
	}
}

func setDefaultDepartureReqParams(urlParams url.Values) {
	for param, defaultValue := range defaultDepartureParams {
		urlParams.Set(param, defaultValue)
	}
}

func setDefaultJourneyReqParams(urlParams url.Values) {
	for param, defaultValue := range defaultJourneyParams {
		urlParams.Set(param, defaultValue)
	}
}
