package mp

import "fmt"

type player interface {
	Play(source string)
}

func PlayMusic(source string, mtype string) {
	var p player
	switch mtype {
	case "mp3":
		p = &MP3Player{}
	case "wav":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)
}
