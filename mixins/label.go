// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"gxui"
	"gxui/math"
	"gxui/mixins/base"
	"strings"
)

type LabelOuter interface {
	base.ControlOuter
}

type Label struct {
	base.Control

	outer               LabelOuter
	font                gxui.Font
	color               gxui.Color
	horizontalAlignment gxui.HorizontalAlignment
	verticalAlignment   gxui.VerticalAlignment
	multiline           bool
	text                string
}

func (l *Label) Init(outer LabelOuter, theme gxui.Theme, font gxui.Font, color gxui.Color) {
	l.Control.Init(outer, theme)
	l.outer = outer
	l.font = font
	l.color = color
	l.horizontalAlignment = gxui.AlignLeft
	l.verticalAlignment = gxui.AlignMiddle
	// Interface compliance test
	_ = gxui.Label(l)
}

func (l *Label) Text() string {
	return l.text
}

func (l *Label) SetText(text string) {
	if l.text != text {
		l.text = text
		l.outer.Relayout()
	}
}

func (l *Label) Color() gxui.Color {
	return l.color
}

func (l *Label) SetColor(color gxui.Color) {
	if l.color != color {
		l.color = color
		l.outer.Redraw()
	}
}

func (l *Label) Multiline() bool {
	return l.multiline
}

func (l *Label) SetMultiline(multiline bool) {
	if l.multiline != multiline {
		l.multiline = multiline
		l.outer.Relayout()
	}
}

func (l *Label) DesiredSize(min, max math.Size) math.Size {
	t := l.text
	if !l.multiline {
		t = strings.Replace(t, "\n", " ", -1)
	}
	s := l.font.Measure(t)
	return s.Clamp(min, max)
}

func (l *Label) SetHorizontalAlignment(horizontalAlignment gxui.HorizontalAlignment) {
	if l.horizontalAlignment != horizontalAlignment {
		l.horizontalAlignment = horizontalAlignment
		l.Redraw()
	}
}

func (l *Label) HorizontalAlignment() gxui.HorizontalAlignment {
	return l.horizontalAlignment
}

func (l *Label) SetVerticalAlignment(verticalAlignment gxui.VerticalAlignment) {
	if l.verticalAlignment != verticalAlignment {
		l.verticalAlignment = verticalAlignment
		l.Redraw()
	}
}

func (l *Label) VerticalAlignment() gxui.VerticalAlignment {
	return l.verticalAlignment
}

// parts.DrawPaint overrides
func (l *Label) Paint(c gxui.Canvas) {
	r := l.outer.Bounds().Size().Rect()
	t := l.text
	if !l.multiline {
		t = strings.Replace(t, "\n", " ", -1)
	}
	c.DrawText(l.font, t, l.color, r, l.horizontalAlignment, l.verticalAlignment)
}