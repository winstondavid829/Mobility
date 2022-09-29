package models

import (
	"github.com/zmb3/spotify/v2"
)

type Audio_Albums struct {
	Name         string            `json:"Name" bson:"Name"`
	ExternalURLs map[string]string `json:"ExternalURLs" bson:"ExternalURLs"`
	Images       []spotify.Image   `json:"Images" bson:"Images"`
	AlbumGroup   string            `json:"AlbumGroup" bson:"AlbumGroup"`
	AlbumType    string            `json:"AlbumType" bson:"AlbumType"`
	URI          spotify.URI       `json:"URI" bson:"URI"`
}

type Audio_Playlists struct {
	Name         string                 `json:"Name" bson:"Name"`
	ExternalURLs map[string]string      `json:"ExternalURLs" bson:"ExternalURLs"`
	Images       []spotify.Image        `json:"Images" bson:"Images"`
	Tracks       spotify.PlaylistTracks `json:"Tracks" bson:"Tracks"`
	URI          spotify.URI            `json:"URI" bson:"URI"`
}

type SpotifyResult struct {
	Playlists []Audio_Playlists `json:"Playlists" bson:"Playlists"`
	Albums    []Audio_Albums    `json:"Albums" bson:"Albums"`
}

type SpotifyResponses struct {
	Status bool          `json:"Status" bson:"Status"`
	Result SpotifyResult `json:"Result" bson:"Result"`
}

type Responses struct {
	Status bool   `json:"Status" bson:"Status"`
	Result string `json:"Result" bson:"Result"`
}
