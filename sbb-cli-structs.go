package main

import "time"
import "fmt"

type (
	// ConnectionsResponse represents the response from Response from http://transport.opendata.ch/v1/connections
	// Docs: http://transport.opendata.ch/docs.html#connections
	ConnectionsResponse struct {
		Connections []Connection `json:"connections"`
	}

	// Connection represents a single connection between two stations
	Connection struct {
		From      Location `json:"from"`
		To        Location `json:"to"`
		Duration  string   `json:"duration"`
		Transfers uint     `json:"transfers"`
	}

	// Location represents a time based location
	Location struct {
		Station       Station  `json:"station"`
		ArrivalTime   *SBBTime `json:"arrival"`
		DepartureTime *SBBTime `json:"departure"`
		Delay         uint     `json:"delay"`
		Platform      *string  `json:"platform"`
	}

	// Station represents a real world station
	Station struct {
		ID          string          `json:"id"`
		Name        string          `json:"name"`
		Coordinates WGS84Coordinate `json:"coordinate"`
	}

	// WGS84Coordinate is a WGS84 Coordinate
	WGS84Coordinate struct {
		Type      string  `json:"type"`
		Longitude float32 `json:"x"`
		Latitude  float32 `json:"y"`
	}

	// SBBTime Seems like Opendata.ch doesn't serialize the json correctly so we make a custom deserializer
	SBBTime struct {
		time.Time
	}
)

// UnmarshalJSON custom time unmarshaller because the time isn't in RFC3399 format.
func (s *SBBTime) UnmarshalJSON(b []byte) (err error) {
	str := string(b)

	// Get rid of the quotes "" around the value.
	// A second option would be to include them
	// in the date format string instead, like so below:
	//   time.Parse(`"`+time.RFC3339Nano+`"`, s)
	str = str[1 : len(str)-1]
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05-0700", str)
		if err != nil {
			fmt.Println(err)
		}
	}

	s.Time = t

	return
}
