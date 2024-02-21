package main

import (
	"fmt"
)

func main() {
	output := AudioOutput{}
	lists := []Playlist{
		{
			name: "one",
			tracks: []Track{
				{
					title: "one1",
				},
				{
					title: "one2",
				},
				{
					title: "one3",
				},
				{
					title: "one4",
				},
				{
					title: "one5",
				},
			},
		},
	}

	track := Track{
		title: "sillyTrack",
	}

	track2 := Track{
		title: "sillyTrack2",
	}

	player := NewPlayer(lists, false)
	fmt.Println("-- Sequenced Player Created")
	player.Play("one", output)
	player.Next(output)
	player.Next(output)
	fmt.Printf("-- Adding %s to user queue\n", track.title)
	player.AddToUserQueue(track)
	player.Next(output)
	player.Next(output)
	player = player.(*SequencedPlayer).Shuffle()
	fmt.Println("-- Randomized Player Created")
	fmt.Println("-- Tracks should start to be Randomized")
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	fmt.Printf("-- Adding %s to user queue\n", track2.title)
	player.AddToUserQueue(track2)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player = player.(*RandomPlayer).Sequence()
	fmt.Println("-- Sequenced Player Created")
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
}
