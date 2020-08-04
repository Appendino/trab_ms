package movie

type Movie struct {
	MovieID     int    `json:"movieId,omitempty"`
	Name        string `json:"movieName,omitempty"`
	Description string `json:"movieDesc,omitempty"`
	Genre       string `json:"genre,omitempty"`
	TotalViews  int    `json:"totalViews,omitempty"`
}
