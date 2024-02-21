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
	fmt.Println("-- Normal Player Created")
	player.Play("one", output)
	player.Next(output)
	player.Next(output)
	fmt.Printf("-- Adding %s to user queue\n", track.title)
	player.AddToUserQueue(track)
	player.Next(output)
	player.Next(output)
	player.SetMode(shuffle)
	fmt.Println("-- Changed to Shuffle mode")
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
	player.SetMode(normal)
	fmt.Println("-- Changed to Normal mode")
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
	player.Next(output)
}
