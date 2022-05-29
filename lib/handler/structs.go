package handler

type Statistics struct {
	Type									string								`json:"type,omitempty"`
	LoopNumber								int									`json:"loop-number,omitempty"`
	PagesLength								int									`json:"pages-length,omitempty"`
	CurrentPageNumber						int									`json:"current-page,omitempty"`
	CurrentPageCollectedMovies				int									`json:"page-collected-movies,omitempty"`
	TotalNumberOfMovies						int									`json:"total-movies,omitempty"`
	CurrentMovie							int									`json:"current-movie,omitempty"`
	NumberOfEpisodes						int									`json:"episodes-length,omitempty"`
	Available								int									`json:"available,omitempty"`
	UnAvailable								int									`json:"unavailable,omitempty"`
	Uploaded								int									`json:"uploaded,omitempty"`
	UnUploaded								int									`json:"unuploaded,omitempty"`
}