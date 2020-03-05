package model

// Movie represents movie domain model
type Movie struct {
	Base
	Seen       bool
	Votes      int       `json:"Votes"`
	Title      string    `json:"Title"`
	Year       string    `json:"Year,omitempty"`
	Rated      string    `json:"Rated,omitempty"`
	Released   string    `json:"Released,omitempty"`
	Runtime    string    `json:"Runtime,omitempty"`
	Genre      string    `json:"Genre,omitempty"`
	Director   string    `json:"Director,omitempty"`
	Writer     string    `json:"Writer,omitempty"`
	Actors     string    `json:"Actors,omitempty"`
	Plot       string    `json:"Plot,omitempty"`
	Language   string    `json:"Language,omitempty"`
	Country    string    `json:"Country,omitempty"`
	Awards     string    `json:"Awards,omitempty"`
	Poster     string    `json:"Poster,omitempty" `
	Ratings    []Ratings `json:"Ratings,omitempty"`
	Metascore  string    `json:"Metascore,omitempty"`
	ImdbRating string    `json:"imdbRating,omitempty"`
	ImdbVotes  string    `json:"imdbVotes,omitempty"`
	ImdbID     string    `json:"imdbID,omitempty"`
	Type       string    `json:"Type,omitempty"`
	DVD        string    `json:"DVD,omitempty"`
	BoxOffice  string    `json:"BoxOffice,omitempty"`
	Production string    `json:"Production,omitempty"`
	Website    string    `json:"Website,omitempty"`
	Response   string    `json:"Response,omitempty"`
}

// Ratings represents rating domain model
type Ratings struct {
	Base
	Source string `json:"Source,omitempty"`
	Value  string `json:"Value,omitempty"`
}

// MovieNight represents MovieNight domain model
type MovieNight struct {
	Base
	WeekNumber      int
	SelectedMovie   Movie   `json:"Selected_Movie,omitempty"`
	SuggestedMovies []Movie `json:"Suggested_Movies,omitempty"`
	Date            string  `json:"date"`
	Host            User    `json:"Host"`
}
