package downloader

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/s3pweb/osm-changes-streamer/model"
)

func State(sn model.SequenceNumber) error {
	resp, err := httpClient.Get(fmt.Sprintf("https://planet.openstreetmap.org/replication/minute/%s.state.txt", sn))

	if err != nil {
		return fmt.Errorf("unable to download state %s %w", sn, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to download state %s status code %d", sn, resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("unable to download state %s %w", sn, err)
	}

	err = os.WriteFile("last-state.txt", data, 0644)

	return err
}

func StateFromTimestamp(timestamp string) {
	resp, err := httpClient.Get("https://replicate-sequences.osm.mazdermind.de/?" + timestamp)

	if err != nil {
		log.Fatalln("unable to download state", timestamp, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("unable to download state for", timestamp, "status code", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("unable to download state", timestamp, err)
	}

	os.WriteFile("last-state.txt", data, 0644)
}
