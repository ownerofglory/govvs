package govvs

import (
	"testing"
	"time"

	"github.com/ownerofglory/govvs/station"
)

func TestGetJourney(t *testing.T) {
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

			if res == nil {
				t.Fatalf("Response is nil")
			}
		})
	}
}

func TestGetArrivals(t *testing.T) {
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

			if res == nil {
				t.Fatalf("Response is nil")
			}
		})
	}
}

func TestGetDepartures(t *testing.T) {
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

			if res == nil {
				t.Fatalf("Response is nil")
			}
		})
	}
}
