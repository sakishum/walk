// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type TabWidget struct {
	Widget        **walk.TabWidget
	Name          string
	StretchFactor int
	Row           int
	RowSpan       int
	Column        int
	ColumnSpan    int
	Font          Font
	Pages         []TabPage
}

func (tw TabWidget) Create(parent walk.Container) error {
	w, err := walk.NewTabWidget(parent)
	if err != nil {
		return err
	}

	return InitWidget(tw, w, func() error {
		var p *walk.TabPage

		for _, page := range tw.Pages {
			if page.Widget == nil {
				page.Widget = &p
			}

			if err := page.Create(nil); err != nil {
				return err
			}

			if err := w.Pages().Add(p); err != nil {
				return err
			}
		}

		if tw.Widget != nil {
			*tw.Widget = w
		}

		return nil
	})
}

func (tw TabWidget) CommonInfo() (name string, stretchFactor, row, rowSpan, column, columnSpan int) {
	return tw.Name, tw.StretchFactor, tw.Row, tw.RowSpan, tw.Column, tw.ColumnSpan
}

func (tw TabWidget) Font_() *Font {
	return &tw.Font
}
