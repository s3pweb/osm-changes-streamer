package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/s3pweb/osm-changes-streamer/downloader"
	"github.com/s3pweb/osm-changes-streamer/filter"
	"github.com/s3pweb/osm-changes-streamer/model"
)

var (
	sequenceNumber model.SequenceNumber
	timestamp      string
)

func main() {

	now := time.Now().UTC()

	resetSequenceNumber := flag.Int("reset-sequence-number", -1, "reset to sequence number")
	resetTimestamp := flag.String("reset-timestamp", "", "reset to timestamp (format: "+now.Format("2006-01-02T15:04:05Z")+")")

	f := flag.String("filter", "", "filter to be used (same syntax as osmium)")

	flag.Parse()

	if _, err := os.Stat("last-state.txt"); errors.Is(err, os.ErrNotExist) {

		// defaults to today at midnight
		if *resetTimestamp == "" {
			*resetTimestamp = now.Format("2006-01-02T00:00:00Z")
		}

	}

	if *resetSequenceNumber != -1 {
		downloader.State(model.SequenceNumber(*resetSequenceNumber))
	}

	if *resetTimestamp != "" {
		_, err := time.Parse("2006-01-02T15:04:05Z", *resetTimestamp)

		if err != nil {
			flag.PrintDefaults()
			log.Fatalln("bad timestamp format")
		}

		downloader.StateFromTimestamp(*resetTimestamp)
	}

	lastState, err := os.Open("last-state.txt")

	if err != nil {
		log.Fatalln("unable to open last-state", err)
	}

	scanner := bufio.NewScanner(lastState)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "sequenceNumber=") {
			number, err := strconv.Atoi(strings.TrimPrefix(line, "sequenceNumber="))

			if err != nil {
				log.Fatalln("expected sequenceNumber to be a number", err)
			}

			sequenceNumber = model.SequenceNumber(number)
		}

		if strings.HasPrefix(line, "timestamp=") {
			timestamp = strings.TrimPrefix(line, "timestamp=")
		}
	}

	log.Println(sequenceNumber)

	for {

		sequenceNumber++
		log.Println("processing", sequenceNumber)

		var changes *model.OSMChange

		for {
			changes, err = downloader.Changes(sequenceNumber)

			if err == nil {
				break
			}

			log.Printf("%v: sleep for some time\n", err)

			// TODO: compute next timestamp diff
			time.Sleep(15 * time.Second)
		}

		filter := filter.Filter(*f)

		filtered := filter.Filter(changes)

		filtered.SequenceNumber = sequenceNumber
		filtered.Timestamp = timestamp

		if len(filtered.Creates) > 0 || len(filtered.Modifies) > 0 || len(filtered.Deletes) > 0 {

			data, err := json.MarshalIndent(filtered, "", "  ")

			if err != nil {
				log.Fatalln("unable to generate JSON", err)
			}

			fmt.Println(string(data))

		}

		for {
			err = downloader.State(sequenceNumber)

			if err == nil {
				break
			}

			log.Printf("%v: sleep for some time\n", err)

			// TODO: compute next timestamp diff
			time.Sleep(15 * time.Second)
		}

	}

}
