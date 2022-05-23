package downloader

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/s3pweb/osm-changes-streamer/model"
)

func Changes(sn model.SequenceNumber) (*model.OSMChange, error) {
	resp, err := httpClient.Get(fmt.Sprintf("https://planet.openstreetmap.org/replication/minute/%s.osc.gz", sn))

	if err != nil {
		log.Fatalln("unable to download changes", sn, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("changes %s does not exists", sn)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to download changes %s status code %d", sn, resp.StatusCode)
	}

	reader, err := gzip.NewReader(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("unable to unzip %s %w", sn, err)
	}

	data, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, fmt.Errorf("unable to unzip %s %w", sn, err)
	}

	var changes model.OSMChange

	err = xml.Unmarshal(data, &changes)

	if err != nil {
		return nil, fmt.Errorf("unable to decode xml %s %w", sn, err)
	}

	return &changes, nil
}
