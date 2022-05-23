# osm-changes-streamer

Stream OpenStreetMap changes with optional filter.

## How it works

OSM changes are stored as XML file containing change sets (can be creation, modification and deletion) associated with a state file (having a sequence number).

Each minute an [OSM changes set file](https://planet.openstreetmap.org/replication/minute/) is being generated.

This project do the following:

* if no last-state.txt file exists use the start sequence number provider or retrieve it from the start date provided
* download .osc.gz associated with sequence number, decode it and generate a JSON representation of the changes set
* apply an optional filter
* output on standard output
* increment sequence number and loop
* wait until sequence number is available

All is being done in memory.

If for some reason program is restarted, it will start from last known sequence number (being stored in last-state.txt).
