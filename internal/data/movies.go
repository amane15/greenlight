package data

import (
	"greenlight/internal/validator"
	"time"
)

// omitempty - hide fields in the output if and only if they are empty
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // - directive
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty,string"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain atleast 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")

}

// Alternative #1
// less clever more clear
// func (m Movie) MarshalJSON() ([]byte, error) {
// 	var runtime string

// 	if m.Runtime != 0 {
// 		runtime = fmt.Sprintf("%d mins", m.Runtime)
// 	}

// 	aux := struct {
// 		ID      int64    `json:"id"`
// 		Title   string   `json:"title"`
// 		Year    int32    `json:"year,omitempty"`
// 		Runtime string   `json:"runtime,omitempty,string"`
// 		Genres  []string `json:"genres,omitempty"`
// 		Version int32    `json:"version"`
// 	}{
// 		ID:      m.ID,
// 		Title:   m.Title,
// 		Year:    m.Year,
// 		Runtime: runtime,
// 		Genres:  m.Genres,
// 		Version: m.Version,
// 	}

// 	return json.Marshal(aux)
// }

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
