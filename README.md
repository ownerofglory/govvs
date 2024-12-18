# Go VVS - VVS (Verkehrsverbund Stuttgart) API library
Simple Go library wrapping VVS API
![Test Pipeline status](https://github.com/ownerofglory/govvs/actions/workflows/test-pipeline.yml/badge.svg)

## Features
- Journey search 
- Station finder
- Departure information
- Arrival information
- Serving lines search
- Geo information for lines

## Prerequisites
- Go version >= 1.19

## Usage
### Example: station finder
```go
package main

import (
    "github.com/ownerofglory/govvs"
    "github.com/ownerofglory/govvs/station"
)

func main() {
	req := govvs.StopFinderRequest{
		Name: "Augustinum",
		Type: "any",
	}
	res, err := govvs.GetStopFinder(req)

	if err != nil {
		// handle error ...
	}

	// process response
	// ...
}
```

### Example: departure monitor
```go
package main

import (
    "github.com/ownerofglory/govvs"
    "github.com/ownerofglory/govvs/station"
)

func main() {
    stationName := station.HAUPTBAHNHOF_TIEF_STUTTGART

    // convert station name to station id (de:XXXXX:YYYY)
    stationId, _ := station.StationNameToGlobalId(stationName)

    req := govvs.DepartureRequest{
        StationId: stationId,
    }
    resp, err := govvs.GetDepartures(req)
    if err != nil {
        // handle error ...
    }
    // process response
    // ...
}
```


### Example: journey search
```go
package main

import (
    "github.com/ownerofglory/govvs"
	"github.com/ownerofglory/govvs/station"
)

func main() {
    // convert station name to station id (de:XXXXX:YYYY)
    origId, _ := station.StationNameToGlobalId(station.HAUPTBAHNHOF_TIEF_STUTTGART)
    destId, _ := station.StationNameToGlobalId(station.FLUGHAFENMESSE_ECHTERDINGEN)

    req := govvs.JourneyRequest{
		OrigId: origId,
		DstId:  destId,
	}

	resp, err := govvs.GetJourney(req)
    if err != nil {
        // handle error ...
    }
    // process response
    // ...
}
```