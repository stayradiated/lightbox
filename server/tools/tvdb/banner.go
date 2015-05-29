package main

type Banner struct {
	ID            int    `xml:"id"`
	BannerPath    string `xml:"BannerPath"`
	BannerType    string `xml:"BannerType"`
	BannerSize    string `xml:"BannerType2"`
	Colors        string `xml:"Colors"`
	Language      string `xml:"Language"`
	Rating        string `xml:"Rating"`
	RatingCount   string `xml:"RatingCount"`
	Season        string `xml:"Season"`
	SeriesName    string `xml:"SeriesName"`
	ThumbnailPath string `xml:"ThumbnailPath"`
	VignettePath  string `xml:"VignettePath'`
}

type Banners struct {
	Banners []Banner `xml:"Banner"`
}
