// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Astronaut struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IsPilot bool   `json:"isPilot"`
}

type Mission struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Crew        []*Astronaut `json:"crew"`
}

type NewAstronautInput struct {
	Name    string `json:"name"`
	IsPilot bool   `json:"isPilot"`
}

type NewMissionInput struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	CrewMembersID []string `json:"crewMembersId"`
}
