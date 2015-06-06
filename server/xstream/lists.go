package xstream

type List struct {
	ID         int    `json:"id"`
	SectionID  int    `json:"section_id"`
	Type       string `json:"type"`
	ViewType   string `json:"view_type"`
	ExternalID string `json:"external_id"`
	Titles     struct {
		Default string `json:"default"`
	} `json:"titles"`
	Order    int        `json:"order"`
	Visible  string     `json:"visible"`
	Elements SeriesList `json:"elements"`
}

type ListElements struct {
	ID       int        `json:"id"`
	Elements SeriesList `json:"elements"`
}
