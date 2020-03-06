package model

// Movie represents movie domain model
type Movie struct {
	Base
	Seen       string    `json:"seen" pg:"default:false,notnull"`
	Votes      int       `json:"votes" pg:"default:0"`
	Title      string    `json:"title"`
	Year       string    `json:"year,omitempty"`
	Rated      string    `json:"rated,omitempty"`
	Released   string    `json:"released,omitempty"`
	Runtime    string    `json:"runtime,omitempty"`
	Genre      string    `json:"genre,omitempty"`
	Director   string    `json:"director,omitempty"`
	Writer     string    `json:"writer,omitempty"`
	Actors     string    `json:"actors,omitempty"`
	Plot       string    `json:"plot,omitempty"`
	Language   string    `json:"language,omitempty"`
	Country    string    `json:"country,omitempty"`
	Awards     string    `json:"awards,omitempty"`
	Poster     string    `json:"poster,omitempty" `
	Ratings    []Ratings `json:"ratings,omitempty"`
	Metascore  string    `json:"metascore,omitempty"`
	ImdbRating string    `json:"imdb_rating,omitempty"`
	ImdbVotes  string    `json:"imdb_votes,omitempty"`
	ImdbID     string    `json:"imdb_id,omitempty"`
	Type       string    `json:"type,omitempty"`
	DVD        string    `json:"dvd,omitempty"`
	BoxOffice  string    `json:"box_office,omitempty"`
	Production string    `json:"production,omitempty"`
	Website    string    `json:"website,omitempty"`
	Response   string    `json:"response,omitempty"`
}

// IncrementVote adds 1 to the field votes
func (m *Movie) IncrementVote() {
	m.Votes = m.Votes + 1
}

// Ratings represents rating domain model
type Ratings struct {
	Base
	Source string `json:"source,omitempty"`
	Value  string `json:"value,omitempty"`
}

// MovieNight represents MovieNight domain model
type MovieNight struct {
	Base
	WeekNumber      int     `json:"week_number"`
	SelectedMovie   *Movie  `json:"selected_movie,omitempty"`
	SuggestedMovies []Movie `json:"suggested_movies,omitempty"`
	Date            string  `json:"date"`
	Host            *User   `json:"host,omitempty"`
	RSVP            []*User `json:"rsvp,omitempty"`
}

func (m *MovieNight) UpdateHost(host *User) {
	m.Host = host
}

func (m *MovieNight) UpdateRSVP(user *User) {
	if m.RSVP == nil {
		m.RSVP = make([]*User, 0)
	}

	m.RSVP = append(m.RSVP, user) //////////////////////////////////////////////////////////////////////////////////////////////
}
