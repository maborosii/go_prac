package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	// stat     int
	progress int
}

func (p *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 music", source)

	p.progress = 0
	for p.progress <= 100 {

		// ? 模仿播放
		time.Sleep(100 * time.Millisecond)
		fmt.Printf(".")
		p.progress += 10
	}

	fmt.Println("Finished playing", source)
}
