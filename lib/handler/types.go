package handler

type Stats struct {
	Type									string								`json:"type,omitempty"`
	LoopNumber								int									`json:"loop-number,omitempty"`
	PagesPosition							int									`json:"pages-position,omitempty"`
	Position								int									`json:"position,omitempty"`
	Available								int									`json:"available,omitempty"`
	UnAvailable								int									`json:"unavailable,omitempty"`
	Uploaded								int									`json:"uploaded,omitempty"`
	UnUploaded								int									`json:"unuploaded,omitempty"`
}