package main

import "fmt"

type Song struct {
	Name     string
	Artist   string
	Previous *Song
	Next     *Song
}

type Playlist struct {
	Name       string
	Head       *Song
	Tail       *Song
	NowPlaying *Song
}

func CreatePlaylist(name string) *Playlist {
	return &Playlist{
		Name: name,
	}
}

func (p *Playlist) AddSong(name, author string) error {
	s := &Song{
		Name:   name,
		Artist: author,
	}
	if p.Head == nil {
		p.Head = s
	} else {
		currentNode := p.Tail
		currentNode.Next = s
		s.Previous = p.Tail
	}
	p.Tail = s
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

func (p *Playlist) PreviousSong() *Song {
	p.NowPlaying = p.NowPlaying.Previous
	return p.NowPlaying
}

func main() {
	playlist := CreatePlaylist("playlistName")
	fmt.Println("Created Playlist")
	fmt.Println()

	fmt.Print("Adding songs to the Playlist...\n\n")
	playlist.AddSong("Ophelia", "The Lumineers")
	playlist.AddSong("Shape of you", "Ed Sheeran")
	playlist.AddSong("Stubborn Love", "The Lumineers")
	playlist.AddSong("Feels", "Calvin Harris")
	fmt.Println("Showing all songs in Playlist...")
	playlist.ShowAllSongs()
	fmt.Println()

	playlist.StartPlaying()
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println()

	playlist.NextSong()
	fmt.Println("Changing Next Song...")
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println()
	playlist.NextSong()
	fmt.Println("Changing Next Song...")
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println()

	playlist.PreviousSong()
	fmt.Println("Changing Previous Song...")
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println()
	playlist.PreviousSong()
	fmt.Println("Changing Previous Song...")
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
}
