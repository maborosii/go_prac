package library

import "errors"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (example *MusicManager) Len() int {
	return len(example.musics)
}

func (example *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(example.musics) {
		return nil, errors.New("index out of range")
	}
	return &example.musics[index], nil
}

func (example *MusicManager) Find(name string) *MusicEntry {
	if len(example.musics) == 0 {
		return nil
	}
	for _, m := range example.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (example *MusicManager) Add(music *MusicEntry) {
	example.musics = append(example.musics, *music)
}

func (example *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(example.musics) || len(example.musics) == 0 {
		// ? index不合法
		return nil
	}
	removedMusic := &example.musics[index]
	if index < len(example.musics)-1 && index != 0 {
		// ? 切片长度大于等于3的中间元素
		example.musics = append(example.musics[:index], example.musics[index+1:]...)
	} else if len(example.musics) == 1 {
		// ? 长度为一的切片
		example.musics = make([]MusicEntry, 0)
	} else if index == 0 {
		// ? 头元素
		example.musics = example.musics[:index+1]
	} else {
		// ? 尾元素
		example.musics = example.musics[:index]
	}

	return removedMusic
}
