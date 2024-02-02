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

// GetArrivals initiates a request to retrieve arrival information for a locations at a specified time.
// It uses the VVS (Verkehrs- und Tarifverbund Stuttgart) API to calculate arrivals, taking into account various transport options and preferences.
//
// Parameters:
// - r ArrivalRequest: Struct containing the essential information for the journey request.
//    - StationId: The ID or code representing the location for the arrival search. This should be in the format recognized by the VVS API, e.g., "de:08111:2599".
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
// - *ArrivalResponse: A pointer to a ArrivalResponse struct containing the arrivals information found based on the request parameters. This includes detailed route information, timings, and any applicable fare information.
// - error: An error object that will be non-nil if there was an issue processing the request or communicating with the VVS API.
//
// Example usage:
// arrivalReq := ArrivalRequest{
//     StationId: "de:08111:2599",
//     TimeAt: time.Now(),
//     Limit: 10,
//     LangCode: "de",
// }
// arrivalResp, err := GetArrivals(arrivalReq, ReqParam{Name: ParamRouteType, Value: "leasttime"})
// if err != nil {
//     // Handle error
// }
// // Process arrivalResp
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

func GetDepartures(point string, checkTime time.Time) (*DepartureResponse, error) {
	// Assemble the query parameters
	params := url.Values{
		"locationServerActive":            []string{"1"},
		"lsShowTrainsExplicit":            []string{"1"},
		"stateless":                       []string{"1"},
		"language":                        []string{"de"},
		"SpEncId":                         []string{"0"},
		"anySigWhenPerfectNoOtherMatches": []string{"1"},
		"limit":                           []string{"100"},
		"depArr":                          []string{"departure"},
		"type_dm":                         []string{"any"},
		"anyObjFilter_dm":                 []string{"2"},
		"deleteAssignedStops":             []string{"1"},
		"name_dm":                         []string{point},
		"mode":                            []string{"direct"},
		"dmLineSelectionAll":              []string{"1"},
		"useRealtime":                     []string{"1"},
		"outputFormat":                    []string{"json"},
		"coordOutputFormat":               []string{"WGS84[DD.ddddd]"},
		"itdDateYear":                     []string{checkTime.Format("2006")},
		"itdDateMonth":                    []string{checkTime.Format("01")},
		"itdDateDay":                      []string{checkTime.Format("02")},
		"itdTimeHour":                     []string{checkTime.Format("15")},
		"itdTimeMinute":                   []string{checkTime.Format("04")},
	}

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

func setDefaultJourneyReqParams(urlParams url.Values) {
	for param, defaultValue := range defaultJourneyParams {
		urlParams.Set(param, defaultValue)
	}
}
