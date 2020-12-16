package facade

import "fmt"

type Music struct {
}

func (m *Music) GetMusic() error {
	fmt.Println("get music material")
	return nil
}

type Video struct {
}

func (v *Video) GetVideo() error {
	fmt.Println("get video material")
	return nil
}

type Count struct {
	ZanCnt      int64
	CommentCnt  int64
	CollentCnt  int64
}

func (c *Count) GetCountByID(id int64) (*Count, error) {
	fmt.Println("get count for something")
	return c, nil
}

type Facade struct {
	m *Music
	v *Video
	c *Count
}

func (f *Facade) GetRecommendVideos() error {
	f.v.GetVideo()
	f.c.GetCountByID(111)
	return nil
}

func (f *Facade) GetRecommendMusics() error {
	f.m.GetMusic()
	f.c.GetCountByID(222)
	return nil
}
