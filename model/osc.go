package model

import "encoding/xml"

type OSMChange struct {
	XMLName        xml.Name          `xml:"osmChange" json:"-"`
	SequenceNumber SequenceNumber    `xml:"-" json:"sequenceNumber"`
	Timestamp      string            `xml:"-" json:"timestamp"`
	Version        string            `xml:"version,attr" json:"version"`
	Generator      string            `xml:"generator,attr" json:"generator"`
	Creates        []OSMChangeCreate `xml:"create" json:"created,omitempty"`
	Modifies       []OSMChangeModify `xml:"modify" json:"modified,omitempty"`
	Deletes        []OSMChangeDelete `xml:"delete" json:"deleted,omitempty"`
}

type OSMChangeCreate struct {
	XMLName   xml.Name            `xml:"create" json:"-"`
	Nodes     []OSMChangeNode     `xml:"node" json:"nodes,omitempty"`
	Ways      []OSMChangeWay      `xml:"way" json:"ways,omitempty"`
	Relations []OSMChangeRelation `xml:"relation" json:"relations,omitempty"`
}

type OSMChangeModify struct {
	XMLName   xml.Name            `xml:"modify" json:"-"`
	Nodes     []OSMChangeNode     `xml:"node" json:"nodes,omitempty"`
	Ways      []OSMChangeWay      `xml:"way" json:"ways,omitempty"`
	Relations []OSMChangeRelation `xml:"relation" json:"relations,omitempty"`
}

type OSMChangeDelete struct {
	XMLName   xml.Name            `xml:"delete" json:"-"`
	Nodes     []OSMChangeNode     `xml:"node" json:"nodes,omitempty"`
	Ways      []OSMChangeWay      `xml:"way" json:"ways,omitempty"`
	Relations []OSMChangeRelation `xml:"relation" json:"relations,omitempty"`
}

type OSMChangeNode struct {
	XMLName   xml.Name       `xml:"node" json:"-"`
	ID        string         `xml:"id,attr" json:"id"`
	Version   string         `xml:"version,attr" json:"version"`
	Timestamp string         `xml:"timestamp,attr" json:"timestamp"`
	UID       string         `xml:"uid,attr" json:"uid"`
	User      string         `xml:"user,attr" json:"user"`
	Changeset string         `xml:"changeset,attr" json:"changeset"`
	Latitude  string         `xml:"lat,attr" json:"lat,omitempty"`
	Longitude string         `xml:"lon,attr" json:"lon,omitempty"`
	Tags      []OSMChangeTag `xml:"tag" json:"tags,omitempty"`
}

type OSMChangeWay struct {
	XMLName   xml.Name           `xml:"way" json:"-"`
	ID        string             `xml:"id,attr" json:"id"`
	Version   string             `xml:"version,attr" json:"version"`
	Timestamp string             `xml:"timestamp,attr" json:"timestamp"`
	UID       string             `xml:"uid,attr" json:"uid"`
	User      string             `xml:"user,attr" json:"user"`
	Changeset string             `xml:"changeset,attr" json:"changeset"`
	Nodes     []OSMChangeWayNode `xml:"nd" json:"nodes,omitempty"`
	Tags      []OSMChangeTag     `xml:"tag" json:"tags,omitempty"`
}

type OSMChangeRelation struct {
	XMLName   xml.Name                  `xml:"relation" json:"-"`
	ID        string                    `xml:"id,attr" json:"id"`
	Version   string                    `xml:"version,attr" json:"version"`
	Timestamp string                    `xml:"timestamp,attr" json:"timestamp"`
	UID       string                    `xml:"uid,attr" json:"uid"`
	User      string                    `xml:"user,attr" json:"user"`
	Changeset string                    `xml:"changeset,attr" json:"changeset"`
	Members   []OSMChangeRelationMember `xml:"member" json:"members,omitempty"`
	Tags      []OSMChangeTag            `xml:"tag" json:"tags,omitempty"`
}

type OSMChangeRelationMember struct {
	XMLName xml.Name `xml:"member" json:"-"`
	Type    string   `xml:"type,attr" json:"type"`
	Ref     string   `xml:"ref,attr" json:"ref"`
	Role    string   `xml:"role,attr" json:"role"`
}

type OSMChangeWayNode struct {
	XMLName xml.Name `xml:"nd" json:"-"`
	Ref     string   `xml:"ref,attr" json:"ref"`
}

type OSMChangeTag struct {
	XMLName xml.Name `xml:"tag" json:"-"`
	Key     string   `xml:"k,attr" json:"key"`
	Value   string   `xml:"v,attr" json:"value"`
}
