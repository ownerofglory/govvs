package govvs

import "time"

// Request params common for Journey and Arrival
const (
	// ParamSpEncId specifies the special encryption ID for the request, usually set to "0" for public requests.
	ParamSpEncId = "SpEncId"

	// ParamCoordOutputFormat sets the coordinate output format, with "EPSG:4326" being a common value for latitude and longitude.
	ParamCoordOutputFormat = "coordOutputFormat"

	// ParamDeleteAssignedStops indicates whether to exclude specific stops from the result. "1" to exclude, "0" to include.
	ParamDeleteAssignedStops = "deleteAssignedStops"
	// ParamLanguage specifies the language of the response. Common values include "de" for German.
	ParamLanguage = "language"

	// ParamLocationServerActive indicates if the location server is active. "1" for active, "0" for inactive.
	ParamLocationServerActive = "locationServerActive"

	// ParamOutputFormat specifies the format of the output. "rapidJSON" is an example for JSON formatted output.
	ParamOutputFormat = "outputFormat"

	// ParamUseRealtime specifies whether to use real-time data. "1" for yes, "0" for no.
	ParamUseRealtime = "useRealtime"
)

// Constants for VVS (Verkehrs- und Tarifverbund Stuttgart) API request parameters.
const (

	// ParamCalcOneDirection indicates whether to calculate routes in one direction only. Accepts "1" (true) or "0" (false).
	ParamCalcOneDirection = "calcOneDirection"

	// ParamChangeSpeed denotes the speed at which a user can change between different transport modes. Common values include "normal" and "fast".
	ParamChangeSpeed = "changeSpeed"

	// ParamComputationType defines the computation strategy for the trip, e.g., "sequence" for sequential trip planning.
	ParamComputationType = "computationType"

	// ParamCycleSpeed represents the cycling speed in km/h when a bicycle mode is selected. Example values might include "14" for 14 km/h.
	ParamCycleSpeed = "cycleSpeed"

	// ParamDeleteITPTWalk specifies whether to remove walking parts from the itinerary. "1" to remove, "0" to include.
	ParamDeleteITPTWalk = "deleteITPTWalk"

	// ParamDescWithElev indicates if the description should include elevation data. "1" for yes, "0" for no.
	ParamDescWithElev = "descWithElev"

	// ParamIllumTransfer specifies whether the transfer should be illuminated. "on" for illuminated transfers.
	ParamIllumTransfer = "illumTransfer"

	// ParamImparedOptionsActive activates options for impaired individuals. "1" for active, "0" for inactive.
	ParamImparedOptionsActive = "imparedOptionsActive"

	// ParamItOptionsActive activates itinerary options. "1" for active, "0" for inactive.
	ParamItOptionsActive = "itOptionsActive"

	// ParamItdDate sets the date for the trip request in format "YYYYMMDD".
	ParamItdDate = "itdDate"

	// ParamItdTime sets the time for the trip request in format "HHMM".
	ParamItdTime = "itdTime"

	// ParamMacroWebTrip specifies whether the web trip macro is used. "true" for yes, "false" for no.
	ParamMacroWebTrip = "macroWebTrip"

	// ParamNameDestination specifies the destination name or code in the request.
	ParamNameDestination = "name_destination"

	// ParamNameOrigin specifies the origin name or code in the request.
	ParamNameOrigin = "name_origin"

	// ParamNoElevationProfile indicates if the elevation profile should be omitted. "1" for yes, "0" for no.
	ParamNoElevationProfile = "noElevationProfile"

	// ParamNoElevationSummary indicates if the elevation summary should be omitted. "1" for yes, "0" for no.
	ParamNoElevationSummary = "noElevationSummary"

	// ParamOutputOptionsActive activates output options. "1" for active, "0" for inactive.
	ParamOutputOptionsActive = "outputOptionsActive"

	// ParamPtOptionsActive activates public transport options. "1" for active, "0" for inactive.
	ParamPtOptionsActive = "ptOptionsActive"

	// ParamRouteType specifies the type of route requested, e.g., "leasttime" for the fastest route.
	ParamRouteType = "routeType"

	// ParamSearchLimitMinutes sets the maximum number of minutes to search for a route.
	ParamSearchLimitMinutes = "searchLimitMinutes"

	// ParamSecurityOptionsActive activates security options for the trip. "1" for active, "0" for inactive.
	ParamSecurityOptionsActive = "securityOptionsActive"

	// ParamServerInfo includes server information in the response. "1" for yes, "0" for no.
	ParamServerInfo = "serverInfo"

	// ParamShowInterchanges specifies whether to show interchange options in the trip. "1" for yes, "0" for no.
	ParamShowInterchanges = "showInterchanges"

	// ParamTrITArrMOT specifies the mode of transport for arrival. "100" might represent a specific mode.
	ParamTrITArrMOT = "trITArrMOT"

	// ParamTrITArrMOTvalue specifies the value associated with the mode of transport for arrival.
	ParamTrITArrMOTvalue = "trITArrMOTvalue"

	// ParamTrITDepMOT specifies the mode of transport for departure. "100" might represent a specific mode.
	ParamTrITDepMOT = "trITDepMOT"

	// ParamTrITDepMOTvalue specifies the value associated with the mode of transport for departure.
	ParamTrITDepMOTvalue = "trITDepMOTvalue"

	// ParamTryToFindLocalityStops indicates whether to try finding locality stops. "1" for yes, "0" for no.
	ParamTryToFindLocalityStops = "tryToFindLocalityStops"

	// ParamTypeDestination specifies the type of the destination, e.g., "any" for any type.
	ParamTypeDestination = "type_destination"

	// ParamTypeOrigin specifies the type of the origin, e.g., "any" for any type.
	ParamTypeOrigin = "type_origin"

	// ParamUseElevationData specifies whether to use elevation data. "1" for yes, "0" for no.
	ParamUseElevationData = "useElevationData"

	// ParamUseLocalityMainStop specifies whether to use the main stop for a locality. "1" for yes, "0" for no.
	ParamUseLocalityMainStop = "useLocalityMainStop"

	// ParamUseUT specifies whether to use the universal time. "1" for yes, "0" for no.
	ParamUseUT = "useUT"

	// ParamVersion specifies the version of the API being used.
	ParamVersion = "version"

	// ParamWObjPrefAl specifies the preference for object allocation. Values are typically numeric.
	ParamWObjPrefAl = "w_objPrefAl"

	// ParamWRegPrefAm specifies the preference for regional amplitude. Values are typically numeric.
	ParamWRegPrefAm = "w_regPrefAm"
)

// Constants for VVS API request parameters.
const (
	// ParamLsShowTrainsExplicit indicates if trains should be explicitly shown.
	ParamLsShowTrainsExplicit = "lsShowTrainsExplicit"

	// ParamStateless indicates if the request should be stateless.
	ParamStateless = "stateless"

	// ParamAnySigWhenPerfectNoOtherMatches indicates if any significant matches should be shown when there's a perfect match.
	ParamAnySigWhenPerfectNoOtherMatches = "anySigWhenPerfectNoOtherMatches"

	// ParamLimit limits the number of results returned.
	ParamLimit = "limit"

	// ParamDepArr indicates the type of time reference (departure or arrival).
	ParamDepArr = "depArr"

	// ParamTypeDm specifies the type of demand (e.g., any).
	ParamTypeDm = "type_dm"

	// ParamAnyObjFilterDm is a filter used for objects in the request.
	ParamAnyObjFilterDm = "anyObjFilter_dm"

	// ParamNameDm is the name of the stop or location for which information is requested.
	ParamNameDm = "name_dm"

	// ParamMode specifies the mode of transportation.
	ParamMode = "mode"

	// ParamDmLineSelectionAll indicates if all lines should be selected.
	ParamDmLineSelectionAll = "dmLineSelectionAll"

	// ParamItdDateTimeDepArr specifies the time and type of the request (departure or arrival).
	ParamItdDateTimeDepArr = "itdDateTimeDepArr"

	// ParamItdDateYear specifies the year of the request date.
	ParamItdDateYear = "itdDateYear"

	// ParamItdDateMonth specifies the month of the request date.
	ParamItdDateMonth = "itdDateMonth"

	// ParamItdDateDay specifies the day of the request date.
	ParamItdDateDay = "itdDateDay"

	// ParamItdTimeHour specifies the hour of the request time.
	ParamItdTimeHour = "itdTimeHour"

	// ParamItdTimeMinute specifies the minute of the request time.
	ParamItdTimeMinute = "itdTimeMinute"

	// ParamItdTripDateTimeDepArr specifies the detailed time and date for departure or arrival.
	ParamItdTripDateTimeDepArr = "itdTripDateTimeDepArr"
)

// JourneyRequest represents the parameters required to request journey information between two points.
//
// Fields:
// - OrigId (string): Identifier for the origin station or location. This should be a unique identifier recognized by the API.
// - DstId (string): Identifier for the destination station or location. Similar to OrigId, this should be a unique identifier recognized by the API.
// - TimeAt (*time.Time): Specifies the desired departure time for the journey. If nil, the current time is assumed.
// - Limit (*int): Limits the number of journeys returned in the response. If nil, a default limit defined by the API is used.
// - LangCode (*string): Language code for the response (e.g., "en" for English, "de" for German). If nil, the default language set by the API is used.
//
// This struct is used to encapsulate the request parameters for fetching journey planning information between two locations.
type JourneyRequest struct {
	OrigId   string
	DstId    string
	TimeAt   *time.Time
	Limit    *int
	LangCode *string
}

// ArrivalRequest represents the parameters required to request arrival information for a specific station or location.
//
// Fields:
// - StationId (string): The unique identifier of the station for which arrival information is requested. This identifier should be recognized by the API.
// - TimeAt (*time.Time): Specifies the time at which arrival information is desired. If nil, the current time is assumed.
// - Limit (*int): Limits the number of arrival records returned in the response. If nil, a default limit set by the API is used.
// - LangCode (*string): Specifies the language of the response (e.g., "en" for English, "de" for German). If nil, the API's default language setting is used.
//
// This struct is utilized to gather arrival times and related information for trains, buses, or other modes of transport at a specified station.
type ArrivalRequest struct {
	StationId string
	TimeAt    *time.Time
	Limit     *int
	LangCode  *string
}

// DepartureRequest represents the parameters required to request departure information for a specific station or location.
//
// Fields:
// - StationId (string): The unique identifier for the station from which departure information is sought. This should correspond to an identifier known to the API.
// - TimeAt (*time.Time): Indicates the specific time for which departure information is requested. Defaults to the current time if nil.
// - Limit (*int): Controls the maximum number of departure entries to be included in the response. Uses a default value if nil.
// - LangCode (*string): Language code for the API response, such as "en" for English or "de" for German. Defaults to the API's standard language if nil.
//
// Employed in obtaining information about scheduled departures from a particular station, including times, lines, and destinations.
type DepartureRequest struct {
	StationId string
	TimeAt    *time.Time
	Limit     *int
	LangCode  *string
}

// ReqParam represents an individual request parameter that can be included in API requests.
//
// Fields:
// - Name (string): The name of the parameter, as expected by the API. This should match the exact name of the parameter in the API documentation.
// - Value (string): The value assigned to the parameter. The type and format of this value should comply with what is expected by the API for the given parameter name.
//
// ReqParam is utilized to dynamically add additional parameters to requests, allowing for more customized and precise API queries.
//
// Example usage:
// To specify a limit on the number of results returned by an API, you might create a ReqParam like so:
// limitParam := ReqParam{Name: "limit", Value: "10"}
// This ReqParam can then be passed to functions that accept variable request parameters, allowing for
// dynamic construction of API queries.
type ReqParam struct {
	Name  string
	Value string
}

var defaultJourneyParams = map[string]string{
	ParamSpEncId:                "0",
	ParamCalcOneDirection:       "1",
	ParamChangeSpeed:            "normal",
	ParamComputationType:        "sequence",
	ParamCoordOutputFormat:      "EPSG:4326",
	ParamCycleSpeed:             "14",
	ParamDeleteAssignedStops:    "0",
	ParamDeleteITPTWalk:         "0",
	ParamDescWithElev:           "1",
	ParamIllumTransfer:          "on",
	ParamImparedOptionsActive:   "1",
	ParamItOptionsActive:        "1",
	ParamLanguage:               "de",
	ParamLocationServerActive:   "1",
	ParamMacroWebTrip:           "true",
	ParamNoElevationProfile:     "1",
	ParamNoElevationSummary:     "1",
	ParamOutputFormat:           "rapidJSON",
	ParamOutputOptionsActive:    "1",
	ParamPtOptionsActive:        "1",
	ParamRouteType:              "leasttime",
	ParamSearchLimitMinutes:     "360",
	ParamSecurityOptionsActive:  "1",
	ParamServerInfo:             "1",
	ParamShowInterchanges:       "1",
	ParamTrITArrMOT:             "100",
	ParamTrITArrMOTvalue:        "15",
	ParamTrITDepMOT:             "100",
	ParamTrITDepMOTvalue:        "15",
	ParamTryToFindLocalityStops: "1",
	ParamTypeDestination:        "any",
	ParamTypeOrigin:             "any",
	ParamUseElevationData:       "1",
	ParamUseLocalityMainStop:    "0",
	ParamUseRealtime:            "1",
	ParamUseUT:                  "1",
	ParamVersion:                "10.2.10.139",
	ParamWObjPrefAl:             "12",
	ParamWRegPrefAm:             "1",
}

var defaultArrivalParams = map[string]string{
	ParamLocationServerActive:            "1",
	ParamLsShowTrainsExplicit:            "1",
	ParamStateless:                       "1",
	ParamLanguage:                        "de",
	ParamSpEncId:                         "0",
	ParamAnySigWhenPerfectNoOtherMatches: "1",
	ParamLimit:                           "10",
	ParamDepArr:                          "arrival",
	ParamTypeDm:                          "any",
	ParamAnyObjFilterDm:                  "2",
	ParamDeleteAssignedStops:             "1",
	ParamMode:                            "direct",
	ParamDmLineSelectionAll:              "1",
	ParamUseRealtime:                     "1",
	ParamOutputFormat:                    "json",
	ParamCoordOutputFormat:               "WGS84[DD.ddddd]",
	ParamItdDateTimeDepArr:               "arr",
	ParamItdTripDateTimeDepArr:           "arr",
}

var defaultDepartureParams = map[string]string{
	ParamLocationServerActive:            "1",
	ParamLsShowTrainsExplicit:            "1",
	ParamStateless:                       "1",
	ParamLanguage:                        "de",
	ParamSpEncId:                         "0",
	ParamAnySigWhenPerfectNoOtherMatches: "1",
	ParamLimit:                           "100",
	ParamDepArr:                          "departure",
	ParamTypeDm:                          "any",
	ParamAnyObjFilterDm:                  "2",
	ParamDeleteAssignedStops:             "1",
	ParamMode:                            "direct",
	ParamDmLineSelectionAll:              "1",
	ParamUseRealtime:                     "1",
	ParamOutputFormat:                    "json",
	ParamCoordOutputFormat:               "WGS84[DD.ddddd]",
}
