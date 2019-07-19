package observer

import "fmt"

// 被观察者
type ISubject interface {
	AddObServer(observers ...IObserver)
	NotifyAll()
}

// 观察者
type IObserver interface {
	Notify()
}

type SubjectA struct {
	obs []IObserver
}

func (s *SubjectA) AddObServer(observers ...IObserver) {
	s.obs = append(s.obs, observers...)
}

func (s *SubjectA) NotifyAll() {
	for _, obs := range s.obs {
		obs.Notify()
	}
}

type ObserverX struct {
}

func (ox *ObserverX) Notify() {
	fmt.Println("ox do something")
}

type ObserverY struct {
}

func (oy *ObserverY) Notify() {
	fmt.Println("oy do something")
}
