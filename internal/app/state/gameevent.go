package state

import (
	"reflect"
)

// ByTeam extracts the `ByTeam` attribute from the game event details
func (e GameEvent) ByTeam() Team {
	v := reflect.ValueOf(e.Event)
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsNil() {
			byTeamValue := v.Field(i).Elem().FieldByName("ByTeam")
			if byTeamValue.IsValid() && !byTeamValue.IsNil() {
				return Team(Team(byTeamValue.Elem().Int()))
			}
		}
	}
	return Team_UNKNOWN
}
