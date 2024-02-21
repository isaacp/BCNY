package main

import (
	"fmt"
	"math/rand"
)

type (
	AudioOutput struct{}

	Player interface {
		Next(output AudioOutput)
		Play(playlist string, output AudioOutput)
		AddToUserQueue(track Track)
	}

	SequencedPlayer struct {
		Lists           map[string]Playlist
		CurrentPlaylist string
		CurrentTrack    int
		Output          AudioOutput
		UserQueue       []Track
	}

	RandomPlayer struct {
		SequencedPlayer
		played map[string]bool
	}

	Playlist struct {
		name   string
		tracks []Track
	}

	Track struct {
		title string
	}
)

func (a *AudioOutput) startAudioOutput(trackName string) {
	fmt.Printf("Playing track: %s\n", trackName)
}

func addToUserQueue(queue *[]Track, track Track) {
	if *queue == nil {
		*queue = make([]Track, 0)
	}
	*queue = append(*queue, track)
}

func (p *SequencedPlayer) AddToUserQueue(track Track) {
	addToUserQueue(&p.UserQueue, track)
}

func (p *RandomPlayer) AddToUserQueue(track Track) {
	addToUserQueue(&p.UserQueue, track)
}

func (p *RandomPlayer) Play(playlist string, output AudioOutput) {
	list := p.Lists[playlist]
	if len(list.tracks) > 0 {
		p.CurrentPlaylist = playlist
		p.CurrentTrack = 0
		p.Output.startAudioOutput(p.Lists[p.CurrentPlaylist].tracks[p.CurrentTrack].title)
	}
}

func (p *SequencedPlayer) Play(playlist string, output AudioOutput) {
	list := p.Lists[playlist]
	if len(list.tracks) > 0 {
		p.CurrentPlaylist = playlist
		p.CurrentTrack = 0
		p.Output.startAudioOutput(p.Lists[p.CurrentPlaylist].tracks[p.CurrentTrack].title)
	}
}

func (p *SequencedPlayer) Next(output AudioOutput) {
	if len(p.UserQueue) > 0 {
		track := p.UserQueue[0]
		p.UserQueue = p.UserQueue[1:]
		p.Output.startAudioOutput(track.title)
		return
	}

	if p.CurrentTrack == len(p.Lists[p.CurrentPlaylist].tracks)-1 {
		p.CurrentTrack = 0
	} else {
		p.CurrentTrack++
	}
	p.Output.startAudioOutput(p.Lists[p.CurrentPlaylist].tracks[p.CurrentTrack].title)
}

func (p *RandomPlayer) Next(output AudioOutput) {
	if len(p.UserQueue) > 0 {
		track := p.UserQueue[0]
		p.UserQueue = p.UserQueue[1:]
		p.Output.startAudioOutput(track.title)
		return
	}

	p.CurrentTrack = randomize(len(p.Lists[p.CurrentPlaylist].tracks))
	count := 0
	for p.played[p.Lists[p.CurrentPlaylist].tracks[p.CurrentTrack].title] {
		count++

		if count >= len(p.Lists[p.CurrentPlaylist].tracks) {
			fmt.Println("-- All tracks have been played randomly.")
			return
		}

		p.CurrentTrack = randomize(len(p.Lists[p.CurrentPlaylist].tracks))

	}
	p.played[p.Lists[p.CurrentPlaylist].tracks[p.CurrentTrack].title] = true
	p.Output.startAudioOutput(p.Lists[p.CurrentPlaylist].tracks[p.CurrentTrack].title)
}

func (p *SequencedPlayer) Shuffle() Player {
	return &RandomPlayer{
		SequencedPlayer: SequencedPlayer{
			Lists:           p.Lists,
			CurrentPlaylist: p.CurrentPlaylist,
			CurrentTrack:    p.CurrentTrack,
			Output:          p.Output,
			UserQueue:       p.UserQueue,
		},
		played: make(map[string]bool),
	}
}

func (p *RandomPlayer) Sequence() Player {
	return &SequencedPlayer{
		Lists:           p.Lists,
		CurrentPlaylist: p.CurrentPlaylist,
		CurrentTrack:    p.CurrentTrack,
		Output:          p.Output,
		UserQueue:       p.UserQueue,
	}
}

func NewPlayer(lists []Playlist, random bool) Player {
	playlists := make(map[string]Playlist)
	for _, list := range lists {
		playlists[list.name] = list
	}

	if random {
		return &RandomPlayer{
			SequencedPlayer: SequencedPlayer{
				Lists:           playlists,
				CurrentPlaylist: "",
				CurrentTrack:    0,
				Output:          AudioOutput{},
				UserQueue:       make([]Track, 0),
			},
			played: make(map[string]bool),
		}
	} else {
		return &SequencedPlayer{
			Lists:           playlists,
			CurrentPlaylist: "",
			CurrentTrack:    0,
			Output:          AudioOutput{},
			UserQueue:       make([]Track, 0),
		}
	}
}

func randomize(count int) int {
	ndx := rand.Int() % count
	return (ndx)
}
