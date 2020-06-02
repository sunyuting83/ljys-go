package apis

// Movielist index movie list data
type Movielist struct {
	ID         int64        `json:"id"`
	CName      string       `json:"c_name"`
	Smallclass []MovieSmall `json:"smallclass"`
	Movie      []MovieLs    `json:"movie"`
}

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
