package govvs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/ownerofglory/govvs/station"
)

const (
	baseURLJourney   = "https://www3.vvs.de/mngvvs/XML_TRIP_REQUEST2"
	baseURLArrival   = "http://www3.vvs.de/vvs/widget/XML_DM_REQUEST"
	baseURLDeparture = "http://www3.vvs.de/vvs/widget/XML_DM_REQUEST"
)

func GetJourneyBtwStations(startStation, destinationStation string, checkTime time.Time) (*JourneyResponse, error) {
	startSt, _ := station.StationNameToGlobalId(startStation)
	dstSt, _ := station.StationNameToGlobalId(destinationStation)

	return GetJourney(startSt, dstSt, checkTime)
}

func GetJourney(start, destination string, checkTime time.Time) (*JourneyResponse, error) {
	// Prepare query parameters
	params := url.Values{
		"SpEncId":                []string{"0"},
		"calcOneDirection":       []string{"1"},
		"changeSpeed":            []string{"normal"},
		"computationType":        []string{"sequence"},
		"coordOutputFormat":      []string{"EPSG:4326"},
		"cycleSpeed":             []string{"14"},
		"deleteAssignedStops":    []string{"0"},
		"deleteITPTWalk":         []string{"0"},
		"descWithElev":           []string{"1"},
		"illumTransfer":          []string{"on"},
		"imparedOptionsActive":   []string{"1"},
		"itOptionsActive":        []string{"1"},
		"itdDate":                []string{checkTime.Format("20060102")},
		"itdTime":                []string{checkTime.Format("1504")},
		"language":               []string{"de"},
		"locationServerActive":   []string{"1"},
		"macroWebTrip":           []string{"true"},
		"name_destination":       []string{destination},
		"name_origin":            []string{start},
		"noElevationProfile":     []string{"1"},
		"noElevationSummary":     []string{"1"},
		"outputFormat":           []string{"rapidJSON"},
		"outputOptionsActive":    []string{"1"},
		"ptOptionsActive":        []string{"1"},
		"routeType":              []string{"leasttime"},
		"searchLimitMinutes":     []string{"360"},
		"securityOptionsActive":  []string{"1"},
		"serverInfo":             []string{"1"},
		"showInterchanges":       []string{"1"},
		"trITArrMOT":             []string{"100"},
		"trITArrMOTvalue":        []string{"15"},
		"trITDepMOT":             []string{"100"},
		"trITDepMOTvalue":        []string{"15"},
		"tryToFindLocalityStops": []string{"1"},
		"type_destination":       []string{"any"},
		"type_origin":            []string{"any"},
		"useElevationData":       []string{"1"},
		"useLocalityMainStop":    []string{"0"},
		"useRealtime":            []string{"1"},
		"useUT":                  []string{"1"},
		"version":                []string{"10.2.10.139"},
		"w_objPrefAl":            []string{"12"},
		"w_regPrefAm":            []string{"1"},
	}

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

func GetArrivalsForStation(point string, checkTime time.Time) (*ArrivalResponse, error) {
	trgStation, _ := station.StationNameToGlobalId(point)

	return GetArrivals(trgStation, checkTime)
}

func GetArrivals(point string, checkTime time.Time) (*ArrivalResponse, error) {
	// Prepare query parameters
	params := url.Values{
		"locationServerActive":            []string{"1"},
		"lsShowTrainsExplicit":            []string{"1"},
		"stateless":                       []string{"1"},
		"language":                        []string{"de"},
		"SpEncId":                         []string{"0"},
		"anySigWhenPerfectNoOtherMatches": []string{"1"},
		"limit":                           []string{"10"},
		"depArr":                          []string{"arrival"},
		"type_dm":                         []string{"any"},
		"anyObjFilter_dm":                 []string{"2"},
		"deleteAssignedStops":             []string{"1"},
		"name_dm":                         []string{point},
		"mode":                            []string{"direct"},
		"dmLineSelectionAll":              []string{"1"},
		"useRealtime":                     []string{"1"},
		"outputFormat":                    []string{"json"},
		"coordOutputFormat":               []string{"WGS84[DD.ddddd]"},
		"itdDateTimeDepArr":               []string{"arr"},
		"itdDateYear":                     []string{checkTime.Format("2006")},
		"itdDateMonth":                    []string{checkTime.Format("01")},
		"itdDateDay":                      []string{checkTime.Format("02")},
		"itdTimeHour":                     []string{checkTime.Format("15")},
		"itdTimeMinute":                   []string{checkTime.Format("04")},
		"itdTripDateTimeDepArr":           []string{"arr"},
	}

	// Build the URL with encoded query parameters
	url := baseURLArrival + "?" + params.Encode()

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

func GetDeparturesForStation(point string, checkTime time.Time) (*DepartureResponse, error) {
	trgStation, _ := station.StationNameToGlobalId(point)

	return GetDepartures(trgStation, checkTime)
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
	fullURL := baseURLDeparture + "?" + params.Encode()

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
