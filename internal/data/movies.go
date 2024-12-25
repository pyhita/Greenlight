package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type Movie struct {
	ID        int64     `json:"id,omitempty"`      // Unique integer ID for the movie
	CreatedAt time.Time `json:"created_at"`        // Timestamp for when the movie is added to our database
	Title     string    `json:"title,omitempty"`   // Movie title
	Year      int32     `json:"year,omitempty"`    // Movie release year
	Runtime   int32     `json:"-"`                 // Movie runtime (in minutes)
	Genres    []string  `json:"genres,omitempty"`  // Slice of genres for the movie (romance, comedy, etc.)
	Version   int32     `json:"version,omitempty"` // The version number starts at 1 and will be incremented each
	// time the movie information is updated
}

func (m Movie) MarshalJSON() ([]byte, error) {
	type MovieAlias Movie

	runtime := fmt.Sprintf("%d mins", m.Runtime)
	aux := struct {
		MovieAlias
		Runtime string `json:"runtime,omitempty"`
	}{
		MovieAlias: (MovieAlias)(m),
		Runtime:    runtime,
	}

	return json.Marshal(aux)
}
