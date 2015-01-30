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

func rectangleRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int, r, g, b, a uint8) bool {
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
	rect.x = x1
	rect.y = y1
	rect.w = x2 - x1
	rect.h = y2 - y1

	/*
	 * Draw
	 */
	var bm BlendMode
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	result := renderer.SetDrawBlendMode(renderer, bm)
	result |= renderer.SetDrawColor(renderer, r, g, b, a)
	result |= renderer.DrawRect(renderer, &rect)
	return result
}

func boxRGBA(renderer *sdl.Renderer, x1, y1, x2, y2 int, r, g, b, a uint8) bool {
	var tmp gl.int
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
	rect.x = x1
	rect.y = y1
	rect.w = x2 - x1 + 1
	rect.h = y2 - y1 + 1

	/*
	 * Draw
	 */
	var bm sdl.BlendMode
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	result := renderer.SetDrawBlendMode(renderer, bm)
	result |= renderer.SetDrawColor(renderer, r, g, b, a)
	result |= renderer.FillRect(renderer, &rect)
	return result
}

func pixelRGBA(renderer *sdl.Renderer, x, y int, r, g, b, a uint8) bool {
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	result := renderer.SetDrawBlendMode(renderer, bm)
	result |= renderer.SetDrawColor(renderer, r, g, b, a)
	result |= renderer.DrawPoint(renderer, x, y)
	return result
}

func vlineRGBA(renderer *sdl.Renderer, x, y1, y2 int, r, g, b, a uint8) bool {
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	result := renderer.SetDrawBlendMode(renderer, bm)
	result |= renderer.SetDrawColor(renderer, r, g, b, a)
	result |= renderer.DrawLine(renderer, x, y1, x, y2)
	return result
}

func hlineRGBA(renderer *sdl.Renderer, x1, x2, y int, r, g, b, a uint8) bool {
	if a == 255 {
		bm = sdl.BLENDMODE_NONE
	} else {
		bm = sdl.BLENDMODE_BLEND
	}
	result := renderer.SetDrawBlendMode(renderer, bm)
	result |= renderer.SetDrawColor(renderer, r, g, b, a)
	result |= renderer.DrawLine(renderer, x1, y, x2, y)
	return result
}
