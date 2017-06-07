package main

import (
	"time"
	"github.com/googollee/go-socket.io"
)

type WorldTimer struct {
	socket socketio.Socket
	whatFunc func(socket socketio.Socket)
	updateDuration time.Duration

	ticker *time.Ticker
	quit chan struct{}
}

func NewWorldTimer(socket socketio.Socket, whatFunc func(socket socketio.Socket), updateDuration time.Duration) *WorldTimer {
	return &WorldTimer{
		socket: socket,
		updateDuration: updateDuration,
		whatFunc: whatFunc,
	}
}

func (timer *WorldTimer) Start() {
	timer.ticker = time.NewTicker(40 * time.Millisecond)
	timer.quit = make(chan struct{})

	go func() {
		for {
			select {
			case <- timer.ticker.C:
				timer.whatFunc(timer.socket)
			case <- timer.quit:
				timer.ticker.Stop()
				return
			}
		}
	}()
}

func (timer *WorldTimer) Stop() {
	close(timer.quit)
}