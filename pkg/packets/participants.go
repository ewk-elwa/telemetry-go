package packets

import (
	"github.com/ewk-elwa/telemetry-go/pkg/env/driver"
	"github.com/ewk-elwa/telemetry-go/pkg/env/team"
)

// This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name.
// If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.

//N.B. on Xbox One, the names will always be the driver name, on PS4 the name will be the LAN name if playing a LAN game,
// otherwise it will be the driver name.

// The array should be indexed by vehicle index.

// Frequency: Every 5 seconds
// Size: 1257 bytes
// Version: 1

type ParticipantData struct {
	AIControlled  uint8         // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID      driver.Driver // Driver id - see appendix
	NetworkID     uint8         // Network id – unique identifier for network players
	TeamID        team.Team     // Team id - see appendix
	MyTeam        uint8         // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber    uint8         // Race number of the car
	Nationality   uint8         // Nationality of the driver
	Name          [48]byte      // Name of participant in UTF-8 format – null terminated, Will be truncated with … (U+2026) if too long
	YourTelemetry uint8         // The player's UDP setting, 0 = restricted, 1 = public
}

type PacketParticipantsData struct {
	Header        PacketHeader
	NumActiveCars uint8               // Number of active cars in the data – should match number of
	Participants  [22]ParticipantData // cars on HUD
}

func (p *PacketParticipantsData) Self() ParticipantData {
	return p.Participants[p.Header.PlayerCarIndex]
}

func (p *ParticipantData) NameToString() string {
	return string(p.Name[:])
}
