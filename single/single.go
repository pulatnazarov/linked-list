package main

import (
	"fmt"
)

// Song ...
type Song struct {
	Name   string
	Artist string
	Next   *Song
}

// Playlist ...
type Playlist struct {
	Name       string
	Head       *Song
	NowPlaying *Song
}

// NewPlaylist ...
func NewPlaylist(name string) *Playlist {
	return &Playlist{
		Name: name,
	}
}

// AddSong ...
func (p *Playlist) AddSong(name, artist string) error {
	song := &Song{
		Name:   name,
		Artist: artist,
	}

	if p.Head == nil {
		p.Head = song
	} else {
		currentNode := p.Head
		if currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = song
	}

	return nil
}

func (p *Playlist) ShowAllSongs() error {
	currentNode := p.Head
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}

func (p *Playlist) StartPlaying() *Song {
	p.NowPlaying = p.Head
	return p.NowPlaying
}

func (p *Playlist) NextSong() *Song {
	p.NowPlaying = p.NowPlaying.Next
	return p.NowPlaying
}

func main() {
	playlist := NewPlaylist("my single")
	fmt.Println("createdPlaylist")
	fmt.Println()

	fmt.Println("adding songs")
	playlist.AddSong("A", "A")
	playlist.AddSong("B", "B")
	playlist.AddSong("C", "C")
	playlist.AddSong("D", "D")

	playlist.ShowAllSongs()
	fmt.Println()

	playlist.StartPlaying()
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println()

	playlist.NextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println()
	playlist.NextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)

}
