package govvs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	baseURLJourney    = "https://www3.vvs.de/mngvvs/XML_TRIP_REQUEST2"
	baseURLWidgetDM   = "https://www3.vvs.de/vvs/widget/XML_DM_REQUEST"
	baseURLStopFinder = "https://www3.vvs.de/mngvvs/XML_STOPFINDER_REQUEST"
	baseURLGeoObject  = "https://www3.vvs.de/mngvvs/XML_GEOOBJECT_REQUEST"
)

const (
	defaultDepartureLimit = 100
	defaultArrivalLimit   = 100
	defaultJourneyLimit   = 100
)

// GetJourney initiates a request to retrieve journey information between two locations at a specified time.
// It uses the VVS (Verkehrs- und Tarifverbund Stuttgart) API to calculate routes, taking into account various transport options and preferences.
//
// Parameters:
// - r JourneyRequest: Struct containing the essential information for the journey request.
//   - OrigId: The ID or code representing the origin location for the journey. This should be in the format recognized by the VVS API, e.g., "de:08111:2599".
//   - DstId: The ID or code representing the destination location for the journey, formatted similarly to OrigId.
//   - TimeAt: Pointer to a time.Time object specifying the desired departure or arrival time for the journey.
//   - Limit: Pointer to an integer specifying the maximum number of journey options to return. If nil, a default value defined by the API will be used.
//   - LangCode: Pointer to a string specifying the language code for the response, e.g., "de" for German. If nil, the API's default language will be used.
//
// - reqParams ...ReqParam: An optional slice of ReqParam structs allowing for additional parameters to be specified for the journey request. These can include:
//   - Name: The name of the parameter, matching the constant names defined for VVS API parameters, e.g., ParamCalcOneDirection.
//   - Value: The value of the parameter, as a string, e.g., "1" for true or specific values like "EPSG:4326" for coordinate formats.
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
//
//	journeyReq := JourneyRequest{
//	    OrigId: "de:08111:2599",
//	    DstId: "de:08111:6075",
//	    TimeAt: time.Now(),
//	    Limit: 10,
//	    LangCode: "de",
//	}
//
// journeyResp, err := GetJourney(journeyReq, ReqParam{Name: ParamRouteType, Value: "leasttime"})
//
//	if err != nil {
//	    // Handle error
//	}
//
// // Process journeyResp
func GetJourney(r JourneyRequest, reqParams ...ReqParam) (*JourneyResponse, error) {
	// Prepare query parameters
	params := url.Values{}

	// set default params
	setDefaultJourneyReqParams(params)

	// Set dynamic values separately
	limit := defaultJourneyLimit
	if r.Limit != nil {
		limit = *r.Limit
	}
	limitParam := strconv.Itoa(limit)
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
	params.Set(ParamLimit, limitParam)

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
	body, err := io.ReadAll(resp.Body)
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
//
//   - OrigId (string): The origin station identifier. Not used in this context but can be set for consistency.
//
//   - DstId (string): The destination station identifier. This is typically the station ID for which arrivals are requested.
//
//   - TimeAt (*time.Time): The specific date and time for which arrival information is requested. If nil, current time is assumed.
//
//   - Limit (*int): The maximum number of arrival entries to retrieve. If nil, a default limit is applied.
//
//   - LangCode (*string): The language code for the response (e.g., "de" for German). If nil, a default language is used.
//
//   - reqParams ...ReqParam: Optional parameters to further customize the request. Each ReqParam consists of a Name and Value.
//     Supported optional parameters (refer to constants for parameter names):
//
//   - ParamLocationServerActive: Activates the location server for processing the request.
//
//   - ParamLsShowTrainsExplicit: Specifies whether trains should be shown explicitly in the response.
//
//   - ParamStateless: Indicates if the request should be processed without maintaining state.
//
//   - ParamLanguage: Sets the language of the response.
//
//   - ParamSpEncId: Security parameter, typically "0".
//
//   - ParamAnySigWhenPerfectNoOtherMatches: Used to refine search results.
//
//   - ParamLimit: Controls the maximum number of results returned.
//
//   - ParamDepArr: Indicates whether departure or arrival times are requested ("arrival" for arrivals).
//
//   - ParamTypeDm: Specifies the type of data management request.
//
//   - ParamAnyObjFilterDm: Filters objects based on specified criteria.
//
//   - ParamDeleteAssignedStops: Indicates whether assigned stops should be removed from the response.
//
//   - ParamNameDm: The name or ID of the station for which arrivals are being requested.
//
//   - ParamMode: Specifies the mode of transport.
//
//   - ParamDmLineSelectionAll: Determines if all lines should be selected in the request.
//
//   - ParamUseRealtime: Specifies whether real-time data should be used.
//
//   - ParamOutputFormat: Defines the output format of the response.
//
//   - ParamCoordOutputFormat: Sets the coordinate format for the response.
//
//   - ParamItdDateTimeDepArr: Specifies the desired time for departure or arrival.
//
//   - Additional parameters related to date and time of the request: ParamItdDateYear, ParamItdDateMonth, ParamItdDateDay, ParamItdTimeHour, ParamItdTimeMinute.
//
// Returns:
// - (*ArrivalResponse, error): A pointer to an ArrivalResponse struct containing the requested arrival information, or an error if the request fails.
//
// Example usage:
// response, err := GetArrivals(ArrivalRequest{DstId: "stationID", TimeAt: &time.Now()}, ReqParam{Name: ParamLimit, Value: "5"})
//
//	if err != nil {
//	    // handle error
//	}
//
// // process response
func GetArrivals(r ArrivalRequest, reqParams ...ReqParam) (*ArrivalResponse, error) {
	// Prepare query parameters
	params := url.Values{}

	setDefaultArrivalReqParams(params)

	limit := defaultArrivalLimit
	if r.Limit != nil {
		limit = *r.Limit
	}
	limitParam := strconv.Itoa(limit)
	timeAt := time.Now()
	if r.TimeAt != nil {
		timeAt = *r.TimeAt
	}
	year := timeAt.Format("2006")
	month := timeAt.Format("01")
	day := timeAt.Format("02")
	hour := timeAt.Format("15")
	minute := timeAt.Format("04")
	stationId := r.StationId

	params.Set(ParamItdDateYear, year)
	params.Set(ParamItdDateMonth, month)
	params.Set(ParamItdDateDay, day)
	params.Set(ParamItdTimeHour, hour)
	params.Set(ParamItdTimeMinute, minute)
	params.Set(ParamNameDm, stationId)
	params.Set(ParamLimit, limitParam)

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
	body, err := io.ReadAll(resp.Body)
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
//
//   - OrigId (string): The origin station identifier. In the context of departures, this is the station ID from which departures are requested.
//
//   - DstId (string): The destination station identifier. Not typically used for departure requests but can be set for API consistency.
//
//   - TimeAt (*time.Time): The specific date and time for which departure information is requested. If nil, the current time is assumed.
//
//   - Limit (*int): The maximum number of departure entries to retrieve. If nil, a default limit is applied.
//
//   - LangCode (*string): The language code for the response (e.g., "de" for German). If nil, a default language is used.
//
//   - reqParams ...ReqParam: Optional parameters to further customize the request. Each ReqParam consists of a Name and Value.
//     Supported optional parameters (refer to constants for parameter names):
//
//   - ParamLocationServerActive: Activates the location server for processing the request.
//
//   - ParamLsShowTrainsExplicit: Specifies whether trains should be shown explicitly in the response.
//
//   - ParamStateless: Indicates if the request should be processed without maintaining state.
//
//   - ParamLanguage: Sets the language of the response.
//
//   - ParamSpEncId: Security parameter, typically "0".
//
//   - ParamAnySigWhenPerfectNoOtherMatches: Used to refine search results.
//
//   - ParamLimit: Controls the maximum number of results returned.
//
//   - ParamDepArr: Indicates whether departure or arrival times are requested ("departure" for departures).
//
//   - ParamTypeDm: Specifies the type of data management request.
//
//   - ParamAnyObjFilterDm: Filters objects based on specified criteria.
//
//   - ParamDeleteAssignedStops: Indicates whether assigned stops should be removed from the response.
//
//   - ParamNameDm: The name or ID of the station for which departures are being requested.
//
//   - ParamMode: Specifies the mode of transport.
//
//   - ParamDmLineSelectionAll: Determines if all lines should be selected in the request.
//
//   - ParamUseRealtime: Specifies whether real-time data should be used.
//
//   - ParamOutputFormat: Defines the output format of the response.
//
//   - ParamCoordOutputFormat: Sets the coordinate format for the response.
//
//   - Additional parameters related to date and time of the request: ParamItdDateYear, ParamItdDateMonth, ParamItdDateDay, ParamItdTimeHour, ParamItdTimeMinute.
//
// Returns:
// - (*DepartureResponse, error): A pointer to a DepartureResponse struct containing the requested departure information, or an error if the request fails.
//
// Example usage:
// response, err := GetDepartures(DepartureRequest{OrigId: "stationID", TimeAt: &time.Now()}, ReqParam{Name: ParamLimit, Value: "10"})
//
//	if err != nil {
//	    // handle error
//	}
//
// // process response
func GetDepartures(r DepartureRequest, reqParams ...ReqParam) (*DepartureResponse, error) {
	// Assemble the query parameters
	params := url.Values{}

	setDefaultDepartureReqParams(params)

	limit := defaultDepartureLimit
	if r.Limit != nil {
		limit = *r.Limit
	}
	limitParam := strconv.Itoa(limit)
	timeAt := time.Now()
	if r.TimeAt != nil {
		timeAt = *r.TimeAt
	}
	year := timeAt.Format("2006")
	month := timeAt.Format("01")
	day := timeAt.Format("02")
	hour := timeAt.Format("15")
	minute := timeAt.Format("04")
	stationId := r.StationId

	params.Set(ParamItdDateYear, year)
	params.Set(ParamItdDateMonth, month)
	params.Set(ParamItdDateDay, day)
	params.Set(ParamItdTimeHour, hour)
	params.Set(ParamItdTimeMinute, minute)
	params.Set(ParamNameDm, stationId)
	params.Set(ParamLimit, limitParam)

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
	body, err := io.ReadAll(resp.Body)
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

// GetStopFinder sends a request to retrieve stop information for a specified location.
// The function takes a StopFinderRequest struct containing essential request parameters and
// a variable number of ReqParam structs to specify additional request options.
//
// Parameters:
// - r StopFinderRequest: A struct containing the base request parameters such as stop name and type.
//
//   - Name (string): The name of the stop or location to search for.
//
//   - Type (string): The type of location being queried (e.g., "any" for any type).
//
//   - CoordOutputFormat (string): The format in which coordinates should be returned (e.g., "EPSG:4326").
//
//   - OutputFormat (string): Specifies the format of the response (e.g., "rapidJSON").
//
//   - ServerInfo (*bool): Whether to include server information in the response. If nil, the default value is used.
//
//   - Language (*string): The language for the response (e.g., "de" for German). If nil, the default language is used.
//
//   - UseRealtime (*bool): Whether to use real-time data in the response. If nil, the default value is used.
//
//   - reqParams ...ReqParam: Optional parameters to further customize the request.
//
//     Supported optional parameters (refer to constants for parameter names):
//
//   - ParamCoordOutputFormat: Specifies the coordinate format (e.g., "EPSG:4326").
//
//   - ParamOutputFormat: Specifies the response format (e.g., "rapidJSON").
//
//   - ParamServerInfo: Indicates if server info should be included in the response.
//
//   - ParamUseRealtime: Specifies whether to use real-time data.
//
//   - ParamLanguage: The language of the response.
//
// Returns:
// - (*StopFinderResponse, error): A pointer to a StopFinderResponse struct containing the stop information, or an error if the request fails.
//
// Example usage:
// response, err := GetStopFinder(StopFinderRequest{Name: "Fred-Uhlman-Str 11", Type: "any"}, ReqParam{Name: ParamLimit, Value: "10"})
//
//	if err != nil {
//	    // handle error
//	}
//
// // process response
func GetStopFinder(r StopFinderRequest, reqParams ...ReqParam) (*StopFinderResponse, error) {
	// Assemble the query parameters
	params := url.Values{}

	// Set default StopFinder parameters
	setDefaultStopFinderReqParams(params)

	// Use provided StopFinderRequest fields
	params.Set(ParamNameSF, r.Name)
	params.Set(ParamTypeSF, r.Type)

	// Optional parameters handling
	if r.CoordOutputFormat != "" {
		params.Set(ParamCoordOutputFormat, r.CoordOutputFormat)
	}
	if r.OutputFormat != "" {
		params.Set(ParamOutputFormat, r.OutputFormat)
	}
	if r.ServerInfo != nil {
		params.Set(ParamServerInfo, boolToString(*r.ServerInfo))
	}
	if r.Language != nil {
		params.Set(ParamLanguage, *r.Language)
	}
	if r.UseRealtime != nil {
		params.Set(ParamUseRealtime, boolToString(*r.UseRealtime))
	}

	// Apply any additional optional request parameters
	overrideReqParams(params, reqParams...)

	// Build the complete URL with the encoded query parameters
	fullURL := baseURLStopFinder + "?" + params.Encode()

	// Perform the GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// Unmarshal the JSON response into the StopFinderResponse struct
	var stopFinderResponse StopFinderResponse
	err = json.Unmarshal(body, &stopFinderResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return &stopFinderResponse, nil
}

// GetGeoObject fetches geographical and transportation data for a specified line or object from the VVS API.
//
// Parameters:
// - r GeoObjectRequest: Struct containing the parameters for the geo-object request.
//   - LineID: The ID of the line or transportation object to fetch data for. For example, "vvs:20007:+:H:j24:1".
//   - CoordOutputFormat: The format of coordinates in the response, e.g., "EPSG:4326".
//   - OutputFormat: The desired response format, e.g., "rapidJSON".
//   - Version: The API version to use in the request.
//
// - reqParams ...ReqParam: Additional optional parameters that can be included in the request.
//
// Returns:
// - *GeoObjectResponse: A pointer to the GeoObjectResponse struct containing the transportation data and associated details.
// - error: An error object that will be non-nil if there was an issue processing the request or communicating with the VVS API.
//
// Example usage:
//
//	geoReq := GeoObjectRequest{
//	    LineID:           "vvs:20007:+:H:j24:1",
//	    CoordOutputFormat: "EPSG:4326",
//	    OutputFormat:      "rapidJSON",
//	    Version:           "10.2.10.139",
//	}
//
// geoResp, err := GetGeoObject(geoReq, ReqParam{Name: ParamServerInfo, Value: "1"})
//
//	if err != nil {
//	    // Handle error
//	}
//
// // Process geoResp
func GetGeoObject(r GeoObjectRequest, reqParams ...ReqParam) (*GeoObjectResponse, error) {
	// Prepare query parameters
	params := url.Values{}

	// Set mandatory parameters
	//params.Set(ParamLineID, r.LineID)
	params.Set(ParamOutputFormat, "rapidJSON")
	params.Set(ParamCoordOutputFormat, "EPSG:4326")
	params.Set(ParamVersion, "10.2.10.139")
	params.Set(ParamSpEncId, "0")
	params.Set(ParamServerInfo, "1")
	params.Set("stFaZon", "1")
	params.Set("vSL", "1")

	// Override optional parameters if provided
	overrideReqParams(params, reqParams...)

	// Build the request URL
	url := baseURLGeoObject + "?" + params.Encode() + "&" + fmt.Sprintf("%s=%s", ParamLineID, r.LineID)

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	// Unmarshal the JSON response into the GeoObjectResponse struct
	var geoResponse GeoObjectResponse
	err = json.Unmarshal(body, &geoResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return nil, err
	}

	// Return the populated GeoObjectResponse struct
	return &geoResponse, nil
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

func setDefaultStopFinderReqParams(params url.Values) {
	for key, value := range defaultStopFinderParams {
		params.Set(key, value)
	}
}

// boolToString converts a boolean value to "1" or "0" for use in query parameters.
func boolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
