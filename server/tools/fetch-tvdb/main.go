package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"

	"bitbucket.org/stayradiated/lightbox/server/tvdb"
)

var key string = "71D4D0F8D1336E0D"
var mirror string = "http://thetvdb.com"
var language string = "en"

func main() {

	var id, tvdbid int
	var name string

	flag.IntVar(&id, "id", -1, "lightbox id")
	flag.StringVar(&name, "name", "", "series name")
	flag.IntVar(&tvdbid, "tvdbid", -1, "tvdb id")
	flag.Parse()

	if id < 0 || name == "" {
		fmt.Println("WARNING: Must specify -id and -name")
		return
	}

	time, _ := ServerTime()
	fmt.Println("Server Time:", time)

	fp := filepath.Join("series", fmt.Sprintf("%d - %s", id, name))

	if tvdbid < 0 {
		results, err := SearchSeries(name)
		if err != nil {
			fmt.Println("ERROR: Could not find match")
			return
		}
		result := results[0]
		tvdbid = result.ID
		fmt.Println(result.SeriesName, result.FirstAired, result.Genre)
	}

	if err := DownloadInfoForSeries(tvdbid, name, fp); err != nil {
		panic(err)
	}
}

func DownloadInfoForSeries(seriesID int, name, fp string) error {
	fn := "package.zip"

	_, err := os.Stat(filepath.Join(fp, "en.xml"))
	if err == nil {
		fmt.Println("Already exists. Skipping...")
		return nil
	}

	if err := os.MkdirAll(fp, 0755); err != nil {
		return err
	}

	src, err := GetSeriesPackage(seriesID)
	if err != nil {
		return err
	}
	defer src.Close()

	f, err := os.Create(filepath.Join(fp, fn))
	if err != nil {
		return err
	}

	if _, err = io.Copy(f, src); err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	cmd := exec.Command("unzip", fn)
	cmd.Dir = fp
	if err := cmd.Run(); err != nil {
		return err
	}

	return err
}

func GetSeriesPackage(seriesID int) (io.ReadCloser, error) {
	url := fmt.Sprintf("%s/api/%s/series/%d/all/%s.zip",
		mirror, key, seriesID, language)

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return r.Body, nil
}

func SearchSeries(name string) ([]tvdb.Series, error) {

	v := url.Values{}
	v.Set("seriesname", name)
	params := v.Encode()

	url := fmt.Sprintf("%s/api/GetSeries.php?%s", mirror, params)

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	response := GetSeriesResponse{}
	x := xml.NewDecoder(r.Body)
	err = x.Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Series, nil
}

func ServerTime() (int, error) {
	url := fmt.Sprintf("%s/api/Updates.php?type=none", mirror)

	r, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer r.Body.Close()

	response := UpdateResponse{}
	x := xml.NewDecoder(r.Body)
	err = x.Decode(&response)
	if err != nil {
		return 0, err
	}

	return response.Time, nil
}

type UpdateResponse struct {
	Time int
}
type GetSeriesResponse struct {
	Series []tvdb.Series
}
