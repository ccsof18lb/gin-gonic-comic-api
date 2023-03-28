package model

type Comic struct {
	ComicId string 	`json:"comicId,omitempty" validate:"required"`
	Title string `json:"title,omitempty" validate:"required"`
 	ReleaseDate string	`json:"releaseDate,omitempty" validate:"required"`
	TrailerLink string	`json:"trailerLink,omitempty" validate:"required"`
	Genres []string	`json:"genres,omitempty" validate:"required"`
	Poster string	`json:"poster,omitempty" validate:"required"`
	Backdrops []string	`json:"backdrops,omitempty" validate:"required"`
	ReviewIds []string	`json:"reviewIds,omitempty" validate:"required"`
}