package filter

import (
	"testing"

	"github.com/s3pweb/osm-changes-streamer/model"
)

func TestMatch(t *testing.T) {

	oneTestFailed := false

	for i, tc := range []struct {
		filter string
		typ    string
		tags   []model.OSMChangeTag
		match  bool
	}{
		{
			filter: "",
			typ:    "n",
			match:  true,
		},
		{
			filter: "key",
			typ:    "n",
			tags: []model.OSMChangeTag{
				{
					Key:   "key",
					Value: "value",
				},
			},
			match: true,
		},
		{
			filter: "key=value",
			typ:    "n",
			tags: []model.OSMChangeTag{
				{
					Key:   "other_key",
					Value: "other_value",
				},
				{
					Key:   "key",
					Value: "value",
				},
			},
			match: true,
		},
		{
			filter: "key=value",
			typ:    "n",
			tags: []model.OSMChangeTag{
				{
					Key:   "other_key",
					Value: "other_value",
				},
				{
					Key:   "key",
					Value: "bad_value",
				},
			},
			match: false,
		},
	} {
		f := Filter(tc.filter)
		m := f.Match(tc.typ, tc.tags...)
		if m != tc.match {
			t.Logf("test %d failed: expected %t for %s %s %+v", i, tc.match, tc.filter, tc.typ, tc.tags)
			oneTestFailed = true
		}
	}

	if oneTestFailed {
		t.FailNow()
	}
}
