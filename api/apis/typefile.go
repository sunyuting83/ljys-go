package apis

// Movielist index movie list data
type Movielist struct {
	ID         int64        `json:"id"`
	CName      string       `json:"c_name"`
	Smallclass []MovieSmall `json:"smallclass"`
	Movie      []MovieLs    `json:"movie"`
}

// Classlist struct
type Classlist struct {
	ID    int64     `json:"id"`
	CName string    `json:"c_name"`
	Movie []MovieLs `json:"movie"`
}

// MovieSmall small classify
type MovieSmall struct {
	ID    int64  `json:"id"`
	CName string `json:"c_name"`
}

// MovieLs movie data
type MovieLs struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Img     string `json:"img"`
	Remarks string `json:"remarks"`
	Score   string `json:"score"`
}

// MovieData movie data
type MovieData struct {
	Img       string   `json:"img"`
	Remarks   string   `json:"reamarks"`
	PlayPath  PlayPath `json:"paly_path"`
	Languarge string   `json:"languarge"`
	Profiles  string   `json:"profiles"`
	Year      string   `json:"year"`
	Soure     string   `json:"soure"`
}

// PlayPath paly_path
type PlayPath struct {
	Hls       []Playlist `json:"hls"`
	Player    []Playlist `json:"player"`
	More      bool       `json:"more"`
	PlayCount int64      `json:"play_count"`
}

// Playlist play list
type Playlist struct {
	Name string `json:"name"`
	Path string `json:"page"`
}

// HotKey hot key
type HotKey struct {
	Key   string `json:"key"`
	Click int64  `json:"click"`
}
