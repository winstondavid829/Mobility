package controllers

import (
	"context"
	"entertainment/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func AuthenticateSpotify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		// if auth.ValidateUserTokenInHeader(c.Request) == false {
		// 	c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": fmt.Sprintf("%v", "Unauthorized Login Attempt / Token Expired")})
		// 	return

		// }
		// ctx := context.Background()
		config := &clientcredentials.Config{
			ClientID:     os.Getenv("SPOTIFY_CLIENTID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			TokenURL:     spotifyauth.TokenURL,
		}

		token, err := config.Token(ctx)
		if err != nil {
			log.Println("couldn't get token: %v", err)

			c.JSON(http.StatusBadRequest, gin.H{"Status": false, "Result": err.Error()})
			return
		}
		httpClient := spotifyauth.New().Client(ctx, token)
		client := spotify.New(httpClient)
		// search for playlists and albums containing "holiday"
		results, err := client.Search(ctx, "gym", spotify.SearchTypePlaylist|spotify.SearchTypeAlbum)
		if err != nil {
			log.Println(err)
		}

		// handle album results
		var albums []models.Audio_Albums
		if results.Albums != nil {
			fmt.Println("Albums:")
			for _, item := range results.Albums.Albums {
				fmt.Println("   ", item.Name, item.ExternalURLs, item.Images, item.AlbumGroup, item.AlbumType, item.URI)
				new := models.Audio_Albums{
					Name:         item.Name,
					ExternalURLs: item.ExternalURLs,
					Images:       item.Images,
					AlbumGroup:   item.AlbumGroup,
					AlbumType:    item.AlbumType,
					URI:          item.URI,
				}
				albums = append(albums, new)
			}
		}
		// handle playlist results
		var playLists []models.Audio_Playlists
		if results.Playlists != nil {
			fmt.Println("Playlists:")
			for _, item := range results.Playlists.Playlists {
				fmt.Println("   ", item.Name, item.ExternalURLs, item.Images, item.Tracks, item.URI)
				new_list := models.Audio_Playlists{
					Name:         item.Name,
					ExternalURLs: item.ExternalURLs,
					Images:       item.Images,
					Tracks:       item.Tracks,
					URI:          item.URI,
				}

				playLists = append(playLists, new_list)
			}
		}

		sptify_res := models.SpotifyResult{
			Playlists: playLists,
			Albums:    albums,
		}

		c.JSON(http.StatusOK, gin.H{"Status": true, "Result": sptify_res})
	}
}
