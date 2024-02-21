package main

import (
	"fmt"
	"math/rand"
)

const (
	normal mode = iota
	shuffle
)

type (
	mode        = int
	AudioOutput struct{}

	Player interface {
		Next(output AudioOutput)
		Play(playlist string, output AudioOutput)
		AddToUserQueue(track Track)
		SetMode(mode mode)
	}

	PlayerMode interface {
		next(output AudioOutput, currentTrack int, plylist Playlist) int
	}

	NormalMode struct{}

	RandomMode struct {
		played map[string]bool
	}

	ConcretePlayer struct {
		Lists           map[string]Playlist
		CurrentPlaylist string
		CurrentTrack    int
		Output          AudioOutput
		UserQueue       []Track
		Mode            PlayerMode
	}

	Playlist struct {
		name   string
		tracks []Track
	}

	Track struct {
		title string
	}
)

func NewPlayer(lists []Playlist, random bool) Player {
	playlists := make(map[string]Playlist)
	for _, list := range lists {
		playlists[list.name] = list
	}

	player := &ConcretePlayer{
		Lists:           playlists,
		CurrentPlaylist: "",
		CurrentTrack:    0,
		Output:          AudioOutput{},
		UserQueue:       make([]Track, 0),
		Mode:            &NormalMode{},
	}

	if random {
		player.SetMode(shuffle)
	}

	return player
}

func (a *AudioOutput) startAudioOutput(trackName string) {
	fmt.Printf("Playing track: %s\n", trackName)
}

func (m *NormalMode) next(output AudioOutput, currentTrack int, playlist Playlist) int {
	if currentTrack == len(playlist.tracks)-1 {
		currentTrack = 0
	} else {
		currentTrack++
	}
	output.startAudioOutput(playlist.tracks[currentTrack].title)
	return currentTrack
}

func (m *RandomMode) next(output AudioOutput, currentTrack int, playlist Playlist) int {
	hold := currentTrack
	currentTrack = randomize(len(playlist.tracks))
	for m.played[playlist.tracks[currentTrack].title] {

		if len(m.played) >= len(playlist.tracks) {
			fmt.Println("-- All tracks have been played randomly.")
			return hold
		}

		currentTrack = randomize(len(playlist.tracks))

	}
	m.played[playlist.tracks[currentTrack].title] = true
	output.startAudioOutput(playlist.tracks[currentTrack].title)
	return currentTrack
}

func (p *ConcretePlayer) Next(output AudioOutput) {
	if len(p.UserQueue) <= 0 {
		p.CurrentTrack = p.Mode.next(output, p.CurrentTrack, p.Lists[p.CurrentPlaylist])
		return
	}

	track := p.UserQueue[0]
	p.UserQueue = p.UserQueue[1:]
	p.Output.startAudioOutput(track.title)
}

func (p *ConcretePlayer) Play(playlist string, output AudioOutput) {
	if len(p.Lists) == 0 {
		return
	}

	list, ok := p.Lists[playlist]
	if !ok {
		return
	}
	if len(list.tracks) > 0 {
		p.CurrentPlaylist = playlist
		p.CurrentTrack = 0
		p.Output.startAudioOutput(p.Lists[p.CurrentPlaylist].tracks[p.CurrentTrack].title)
	}
}

func (p *ConcretePlayer) AddToUserQueue(track Track) {
	if p.UserQueue == nil {
		p.UserQueue = make([]Track, 0)
	}
	p.UserQueue = append(p.UserQueue, track)
}

func (p *ConcretePlayer) SetMode(mode mode) {
	switch mode {
	case normal:
		p.Mode = &NormalMode{}
	case shuffle:
		p.Mode = &RandomMode{
			played: make(map[string]bool),
		}
	}
}

func randomize(count int) int {
	ndx := rand.Int() % count
	return (ndx)
}
