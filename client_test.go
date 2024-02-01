package govvs

import (
	"testing"
	"time"

	"github.com/ownerofglory/govvs/station"
)

func TestGetJourneyBtwStations(t *testing.T) {
	testCases := []struct {
		name    string
		orig    string
		dst     string
		success bool
	}{
		{
			name:    "Success",
			orig:    station.SILLENBUCH_STUTTGART,
			dst:     station.CHARLOTTENPLATZ_STUTTGART,
			success: true,
		},
		{
			name:    "When station doesn exist error is thrown",
			orig:    "NONEXISTINGORIG",
			dst:     "NONEXISTINGDST",
			success: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := GetJourneyBtwStations(tc.orig, tc.dst, time.Now())

			if tc.success && err != nil {
				t.Fatal(err)
			}

			if resp == nil {
				t.Fatal("Response is nil")
			}
		})
	}
}

func TestGetArrivalsForStation(t *testing.T) {
	testCases := []struct {
		name    string
		station string
		success bool
	}{
		{
			name:    "Success",
			station: station.SILLENBUCH_STUTTGART,
			success: true,
		},
		{
			name:    "When station doesn exist error is thrown",
			station: "NONEXISTINGORIG",
			success: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := GetArrivalsForStation(tc.station, time.Now())

			if tc.success && err != nil {
				t.Fatal(err)
			}

			if resp == nil {
				t.Fatal("Response is nil")
			}
		})
	}
}

func TestGetDeparturesForStation(t *testing.T) {
	testCases := []struct {
		name    string
		station string
		success bool
	}{
		{
			name:    "Success",
			station: station.SILLENBUCH_STUTTGART,
			success: true,
		},
		{
			name:    "When station doesn exist error is thrown",
			station: "NONEXISTINGORIG",
			success: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := GetDeparturesForStation(tc.station, time.Now())

			if tc.success && err != nil {
				t.Fatal(err)
			}

			if resp == nil {
				t.Fatal("Response is nil")
			}
		})
	}
}
