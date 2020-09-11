package loom

import (
	"testing"
	"time"
)

/********************************************************************
created:    2020-08-01
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type TestGo struct {
	wc WaitClose
}

func (my *TestGo) Init() {
	Go(my.goLoop)

	<-my.wc.C()
}

func (my *TestGo) goLoop(later Later) {
	var ticker = later.NewTicker(time.Second)
	var counter = 0
	for {
		select {
		case <-ticker.C:
			counter += 1
			println("goLoop")
			if counter == 3 {
				my.wc.Close(nil)
				return
			}
		}
	}
}

func TestLater_NewTicker(t *testing.T) {
	var test = &TestGo{}
	test.Init()
}
