package govvs

import (
	"testing"
	"time"

	"github.com/ownerofglory/govvs/station"
)

func TestGetJourney(t *testing.T) {
	testLimit := 4
	testCases := []struct {
		name        string
		orig        string
		dst         string
		limit       *int
		timeAt      *time.Time
		errExpected bool
	}{
		{
			name:        "When required params are given then journeys are returned",
			orig:        station.BOPSER_STUTTGART,
			dst:         station.HAUPTBAHNHOF_TIEF_STUTTGART,
			errExpected: false,
		},
		{
			name:        "When required params and limit are given then journeys no more than limit are returned",
			orig:        station.BOPSER_STUTTGART,
			dst:         station.HAUPTBAHNHOF_TIEF_STUTTGART,
			errExpected: false,
			limit:       &testLimit,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := JourneyRequest{
				OrigId: tc.orig,
				DstId:  tc.dst,
			}
			if tc.limit != nil {
				req.Limit = tc.limit
			}
			if tc.timeAt != nil {
				req.TimeAt = tc.timeAt
			}

			res, err := GetJourney(req)

			if !tc.errExpected && err != nil {
				t.Fatalf("Error occured: %v\n", err)
			}

			if tc.limit != nil && len(res.Journeys) > *tc.limit {
				t.Fatalf("Error occured: %v\n", err)
			}

			if res == nil {
				t.Fatalf("Response is nil")
			}
		})
	}
}

func TestGetArrivals(t *testing.T) {
	testLimit := 4
	testCases := []struct {
		name        string
		stationId   string
		limit       *int
		timeAt      *time.Time
		errExpected bool
	}{
		{
			name:        "When required params are given then arrivals are returned",
			stationId:   station.BOPSER_STUTTGART,
			errExpected: false,
		},
		{
			name:        "When required params and limit are given then arrivals no more than limit are returned",
			stationId:   station.BOPSER_STUTTGART,
			errExpected: false,
			limit:       &testLimit,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := ArrivalRequest{
				StationId: tc.stationId,
			}
			if tc.limit != nil {
				req.Limit = tc.limit
			}
			if tc.timeAt != nil {
				req.TimeAt = tc.timeAt
			}

			res, err := GetArrivals(req)

			if !tc.errExpected && err != nil {
				t.Fatalf("Error occured: %v\n", err)
			}

			if tc.limit != nil && len(res.ArrivalList) > *tc.limit {
				t.Fatalf("Error occured: %v\n", err)
			}

			if res == nil {
				t.Fatalf("Response is nil")
			}
		})
	}
}

func TestGetDepartures(t *testing.T) {
	testLimit := 4
	testCases := []struct {
		name        string
		stationId   string
		limit       *int
		timeAt      *time.Time
		errExpected bool
	}{
		{
			name:        "When required params are given then departures are returned",
			stationId:   station.BOPSER_STUTTGART,
			errExpected: false,
		},
		{
			name:        "When required params and limit are given then departures no more then limit are returned",
			stationId:   station.BOPSER_STUTTGART,
			errExpected: false,
			limit:       &testLimit,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := DepartureRequest{
				StationId: tc.stationId,
			}
			if tc.limit != nil {
				req.Limit = tc.limit
			}
			if tc.timeAt != nil {
				req.TimeAt = tc.timeAt
			}

			res, err := GetDepartures(req)

			if !tc.errExpected && err != nil {
				t.Fatalf("Error occured: %v\n", err)
			}

			if tc.limit != nil && len(res.DepartureList) > *tc.limit {
				t.Fatalf("Error occured: %v\n", err)
			}

			if res == nil {
				t.Fatalf("Response is nil")
			}
		})
	}
}

func TestGetStopFinder(t *testing.T) {
	testCases := []struct {
		name        string
		stopName    string
		stopType    string
		lang        *string
		useRealtime *bool
		errExpected bool
	}{
		{
			name:        "When required params are given then stop information is returned",
			stopName:    "Königstr. 1",
			stopType:    "any",
			errExpected: false,
		},
		{
			name:        "When real-time data is requested then real-time information is returned",
			stopName:    "Königstr. 1",
			stopType:    "any",
			useRealtime: boolPointer(true),
			errExpected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := StopFinderRequest{
				Name: tc.stopName,
				Type: tc.stopType,
			}

			if tc.lang != nil {
				req.Language = tc.lang
			}
			if tc.useRealtime != nil {
				req.UseRealtime = tc.useRealtime
			}

			res, err := GetStopFinder(req)

			// Check for errors
			if !tc.errExpected && err != nil {
				t.Fatalf("Error occurred: %v\n", err)
			}

			// Check that response is not nil
			if res == nil && !tc.errExpected {
				t.Fatalf("Expected response but got nil")
			}

			// Check that if an error is expected, no result should be returned
			if tc.errExpected && res != nil {
				t.Fatalf("Expected error but got valid response")
			}
		})
	}
}

func TestGetGeoObject(t *testing.T) {
	testCases := []struct {
		name        string
		lineID      string
		expLineName string
	}{
		{
			name:        "if valid line id is given geo object is returned",
			lineID:      "vvs:20007:+:H:j24:1",
			expLineName: "U7",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := GeoObjectRequest{
				LineID: tc.lineID,
			}
			geoObj, err := GetGeoObject(req)

			if err != nil {
				t.Fatalf("Error occured: %v\n", err)
			}

			if geoObj.Transportations[0].Number != tc.expLineName {
				t.Fatalf("Line doesn't match")
			}
		})
	}
}

// Helper function to create a pointer for boolean values
func boolPointer(b bool) *bool {
	return &b
}
