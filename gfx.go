package main

/*

SDL2_gfxPrimitives.c: graphics primitives for SDL2 renderers

Copyright (C) 2012  Andreas Schiffler

This software is provided 'as-is', without any express or implied
warranty. In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

1. The origin of this software must not be misrepresented; you must not
claim that you wrote the original software. If you use this software
in a product, an acknowledgment in the product documentation would be
appreciated but is not required.

2. Altered source versions must be plainly marked as such, and must not be
misrepresented as being the original software.

3. This notice may not be removed or altered from any source
distribution.

Andreas Schiffler -- aschiffler at ferzkopp dot net

*/

import (
	"github.com/veandco/go-sdl2/sdl"
)

func rectangleRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int, r, g, b, a uint8) error {
	var tmp int
	var rect sdl.Rect

	/*
	 * Test for special cases of straight lines or single point
	 */
	if x1 == x2 {
		if y1 == y2 {
			return (pixelRGBA(renderer, x1, y1, r, g, b, a))
		} else {
			return (vlineRGBA(renderer, x1, y1, y2, r, g, b, a))
		}
	} else {
		if y1 == y2 {
			return (hlineRGBA(renderer, x1, x2, y1, r, g, b, a))
		}
	}

	/*
	 * Swap x1, x2 if required
	 */
	if x1 > x2 {
		tmp = x1
		x1 = x2
		x2 = tmp
	}

	/*
	 * Swap y1, y2 if required
	 */
	if y1 > y2 {
		tmp = y1
		y1 = y2
		y2 = tmp
	}

	/*
	 * Create destination rect
	 */
	rect.X = int32(x1)
	rect.Y = int32(y1)
	rect.W = int32(x2 - x1)
	rect.H = int32(y2 - y1)

	/*
	 * Draw
	 */
	var bm sdl.BlendMode
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	err := renderer.SetDrawBlendMode(bm)
	if err == nil {
		err = renderer.SetDrawColor(r, g, b, a)
	}
	if err == nil {
		err = renderer.DrawRect(&rect)
	}
	return err
}

func boxRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int, r, g, b, a uint8) error {
	var rect sdl.Rect

	/*
	 * Test for special cases of straight lines or single point
	 */
	if x1 == x2 {
		if y1 == y2 {
			return (pixelRGBA(renderer, x1, y1, r, g, b, a))
		} else {
			return (vlineRGBA(renderer, x1, y1, y2, r, g, b, a))
		}
	} else {
		if y1 == y2 {
			return (hlineRGBA(renderer, x1, x2, y1, r, g, b, a))
		}
	}

	/*
	 * Swap x1, x2 if required
	 */
	var tmp int
	if x1 > x2 {
		tmp = x1
		x1 = x2
		x2 = tmp
	}

	/*
	 * Swap y1, y2 if required
	 */
	if y1 > y2 {
		tmp = y1
		y1 = y2
		y2 = tmp
	}

	/*
	 * Create destination rect
	 */
	rect.X = int32(x1)
	rect.Y = int32(y1)
	rect.W = int32(x2 - x1 + 1)
	rect.H = int32(y2 - y1 + 1)

	/*
	 * Draw
	 */
	var bm sdl.BlendMode
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	err := renderer.SetDrawBlendMode(bm)
	if err == nil {
		err = renderer.SetDrawColor(r, g, b, a)
	}
	if err == nil {
		err = renderer.FillRect(&rect)
	}
	return err
}

func pixelRGBA(renderer *sdl.Renderer, x, y int, r, g, b, a uint8) error {
	var bm sdl.BlendMode
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	err := renderer.SetDrawBlendMode(bm)
	if err == nil {
		err = renderer.SetDrawColor(r, g, b, a)
	}
	if err == nil {
		err = renderer.DrawPoint(x, y)
	}
	return err
}

func vlineRGBA(renderer *sdl.Renderer, x, y1, y2 int, r, g, b, a uint8) error {
	var bm sdl.BlendMode
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	err := renderer.SetDrawBlendMode(bm)
	if err == nil {
		err = renderer.SetDrawColor(r, g, b, a)
	}
	if err == nil {
		err = renderer.DrawLine(x, y1, x, y2)
	}
	return err
}

func hlineRGBA(renderer *sdl.Renderer, x1, x2, y int, r, g, b, a uint8) error {
	var bm sdl.BlendMode
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	err := renderer.SetDrawBlendMode(bm)
	if err == nil {
		err = renderer.SetDrawColor(r, g, b, a)
	}
	if err == nil {
		err = renderer.DrawLine(x1, y, x2, y)
	}
	return err
}
