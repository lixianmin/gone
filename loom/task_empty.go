package loom

/********************************************************************
created:    2020-08-25
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type taskEmpty struct{}

func (task taskEmpty) Do() error {
	return nil
}

func (task taskEmpty) Wait() {

}
