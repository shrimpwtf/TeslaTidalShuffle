package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"strconv"
	"time"
)

var TidalService *Service

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	color.Red("Tesla Tidal Shuffle")
	TidalService = NewService()
	playlists, err := TidalService.GetUserPlaylists(TidalService.UserID)
	if err != nil {
		panic(err)
	} else {
		color.Cyan("Playlists to Shuffle: ")
		for x, playlist := range playlists.Items {
			color.Magenta(strconv.Itoa(x) + " - " + playlist.Title)
		}
		color.Cyan("Choice: ")
		var choice string
		fmt.Scanln(&choice)
		choicei, err := strconv.Atoi(choice)
		if err != nil {
			panic(err)
		}
		toShuffle := playlists.Items[choicei]
		playlistTracks, err := TidalService.GetPlaylistTracks(toShuffle.UUID)
		if err != nil {
			panic(err)
		}
		songsToShuffle := playlistTracks.Items
		color.Green(strconv.Itoa(len(songsToShuffle)) + " Tracks Loaded!")
		color.Cyan("Shuffled Playlist Name: ")
		shuffledName := ""
		fmt.Scanln(&shuffledName)
		shuffledPlaylist, err := TidalService.CreatePlaylist(shuffledName, toShuffle.Description)
		if err != nil {
			panic(err)
		}

		rand.Shuffle(len(songsToShuffle), func(i, j int) { songsToShuffle[i], songsToShuffle[j] = songsToShuffle[j], songsToShuffle[i] })

		for _, song := range songsToShuffle {
			color.Yellow("Adding Song: " + song.Title)
			err := TidalService.AddTrackToPlaylist(shuffledPlaylist.UUID, song.ID)
			if err != nil {
				panic(err)
			}
			color.Green("Added Song: " + song.Title)
		}
	}
	color.Green("All Done :)")
	fmt.Scanln()
}
