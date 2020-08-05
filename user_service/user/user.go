package user

type UserMovie struct {
	UserID  int `json:"userId,omitempty"`
	MovieID int `json:"movieId,omitempty"`
}

type Movie struct {
	MovieID     int    `json:"movieId,omitempty"`
	Name        string `json:"movieName,omitempty"`
	Description string `json:"movieDesc,omitempty"`
	Genre       string `json:"genre,omitempty"`
	TotalViews  int    `json:"totalViews,omitempty"`
}

type User struct {
	UserID int     `json:"userId,omitempty"`
	Movies []Movie `json:"movies,omitempty"`
}
