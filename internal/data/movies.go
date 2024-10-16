package data

import (
	"encoding/json"
	"fmt"
	"time"
)

// omitempty - hide fields in the output if and only if they are empty
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // - directive
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"runtime,omitempty,string"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

// Alternative #1
// less clever more clear
func (m Movie) MarshalJSON() ([]byte, error) {
	var runtime string

	if m.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}

	aux := struct {
		ID      int64    `json:"id"`
		Title   string   `json:"title"`
		Year    int32    `json:"year,omitempty"`
		Runtime string   `json:"runtime,omitempty,string"`
		Genres  []string `json:"genres,omitempty"`
		Version int32    `json:"version"`
	}{
		ID:      m.ID,
		Title:   m.Title,
		Year:    m.Year,
		Runtime: runtime,
		Genres:  m.Genres,
		Version: m.Version,
	}

	return json.Marshal(aux)
}

// Alternative #2
// more clever less clear
// func (m Movie) MarshalJSON() ([]byte, error) {
// 	var runtime string

// 	if m.Runtime != 0 {
// 		runtime = fmt.Sprintf("%d mins", m.Runtime)
// 	}

// 	type MovieAlias Movie

// 	aux := struct {
// 		MovieAlias
// 		Runtime string `json:"runtime,omitempty"`
// 	}{
// 		MovieAlias: MovieAlias(m),
// 		Runtime:    runtime,
// 	}

// 	return json.Marshal(aux)
// }
