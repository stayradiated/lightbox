package xstream

type HotspotImage struct {
	ID       int    `json:"id"`
	Source   string `json:"source"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

type Hotspot struct {
	ID           int       `json:"id"`
	Titles       MapString `json:"titles"`
	Descriptions MapString `json:"descriptions"`
	// Link
	Location     string `json:"location"`
	Type         string `json:"type"`
	Availability struct {
		Domains []string `json:"domains"`
	} `json:"availability"`
	Order int   `json:"order"`
	Dates Dates `json:"dates"`
	// Categories
	Images []HotspotImage `json:"images"`
}

type HotspotsResponse struct {
	Count    int       `json:"count"`
	Hotspots []Hotspot `json:"hotspots"`
}
