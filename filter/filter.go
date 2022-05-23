package filter

import (
	"log"
	"strings"

	"github.com/s3pweb/osm-changes-streamer/model"
)

type Filter string

func (f Filter) Match(typ string, tags ...model.OSMChangeTag) bool {

	parts := strings.Split(string(f), " ")

	for _, each := range parts {

		match := true

		eachParts := strings.Split(each, "/")

		// if type not specified
		if len(eachParts) == 1 {
			eachParts = []string{typ, each}
		}

		if eachParts[0] != "a" && eachParts[0] != typ {
			match = false
			continue
		}

		found := false

		if eachParts[1] == "" {
			found = true
		}

		if !found {
			for _, eachTag := range tags {
				differParts := strings.Split(eachParts[1], "!=")

				if len(differParts) == 2 {
					// TODO
					log.Fatalln("differ operator not supported")
				}

				equalsParts := strings.Split(eachParts[1], "=")

				if len(equalsParts) == 2 {
					if eachTag.Key == equalsParts[0] {
						for _, multi := range strings.Split(equalsParts[1], ",") {
							if eachTag.Value == multi {
								found = true
								break
							}
						}
					}
				}

				if eachTag.Key == eachParts[1] {
					found = true
					break
				}
			}
		}

		if !found {
			match = false
		}

		if match {
			return true
		}
	}

	return false
}

func (filter Filter) Filter(input *model.OSMChange) *model.OSMChange {

	ret := &model.OSMChange{
		Version:   input.Version,
		Generator: strings.Join([]string{input.Generator, "filtered with", string(filter)}, " "),
	}

	// creates
	for _, each := range input.Creates {
		create := model.OSMChangeCreate{}

		// for each node
		for _, eachNode := range each.Nodes {
			if filter.Match("n", eachNode.Tags...) {
				create.Nodes = append(create.Nodes, eachNode)
			}
		}

		// for each way
		for _, eachWay := range each.Ways {
			if filter.Match("w", eachWay.Tags...) {
				create.Ways = append(create.Ways, eachWay)
			}
		}

		// for each relation
		for _, eachRelation := range each.Relations {
			if filter.Match("r", eachRelation.Tags...) {
				create.Relations = append(create.Relations, eachRelation)
			}
		}

		if len(create.Nodes) > 0 || len(create.Ways) > 0 || len(create.Relations) > 0 {
			ret.Creates = append(ret.Creates, create)
		}
	}

	// modifies
	for _, each := range input.Modifies {
		modify := model.OSMChangeModify{}

		// for each node
		for _, eachNode := range each.Nodes {
			if filter.Match("n", eachNode.Tags...) {
				modify.Nodes = append(modify.Nodes, eachNode)
			}
		}

		// for each way
		for _, eachWay := range each.Ways {
			if filter.Match("w", eachWay.Tags...) {
				modify.Ways = append(modify.Ways, eachWay)
			}
		}

		// for each relation
		for _, eachRelation := range each.Relations {
			if filter.Match("r", eachRelation.Tags...) {
				modify.Relations = append(modify.Relations, eachRelation)
			}
		}

		if len(modify.Nodes) > 0 || len(modify.Ways) > 0 || len(modify.Relations) > 0 {
			ret.Modifies = append(ret.Modifies, modify)
		}
	}

	// deletes
	ret.Deletes = input.Deletes

	return ret
}
