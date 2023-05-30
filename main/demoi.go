package main

import (
	"math"
)

// Tile
type Tile[T INumeric] struct {
	tileSize          T
	initialResolution T
	originShift       T
}

type INumeric interface {
	float64
}

func NewTile[T INumeric](tileSize T) *Tile[T] {
	t := new(Tile[T])
	if tileSize == 0 {
		tileSize = 256.0
	} else {
		tileSize = T(float64(tileSize / 1.0))
	}
	t.tileSize = tileSize
	t.initialResolution = T(2 * math.Pi * 6378137 / float64(t.tileSize))
	var (
		n T = 2.0
	)
	t.originShift = T(float64(2 * math.Pi * 6378137 / float64(n)))

	return t
}

func (t *Tile[T]) LatLonToMeters(lat, lon T) (mx, my T) {
	mx = lon * t.originShift / 180
	my = T(math.Log(math.Tan(float64(90+lat)*math.Pi/360)) / (math.Pi / 180))
	my = my * t.originShift / 180
	return
}

func (t *Tile[T]) MetersToLatLon(mx, my T) (lat, lon T) {
	lon = (mx / t.originShift) * 180
	lat = (my / t.originShift) * 180
	lat = T(180 / math.Pi * (2*math.Atan(math.Exp(float64(lat)*math.Pi/180)) - math.Pi/2))
	return
}

func (t *Tile[T]) Resolution(zoom T) T {
	return t.initialResolution / T(math.Pow(float64(2), float64(zoom)))
}

func (t *Tile[T]) PixelsToMeters(px, py, zoom T) (mx, my T) {
	res := t.Resolution(zoom)
	mx = px*res - t.originShift
	my = py*res - t.originShift
	return
}

func (t *Tile[T]) TileBounds(tx, ty, zoom T) (minx, miny, maxx, maxy T) {
	minx, miny = t.PixelsToMeters(tx*t.tileSize, ty*t.tileSize, zoom)
	maxx, maxy = t.PixelsToMeters((tx+1)*t.tileSize, (ty+1)*t.tileSize, zoom)
	return
}

func (t *Tile[T]) TileLatLonBounds(tx, ty, zoom T) (minLon, minLat, maxLon, maxLat T) {
	minx, miny, maxx, maxy := t.TileBounds(tx, ty, zoom)
	minLat, minLon = t.MetersToLatLon(minx, miny)
	maxLat, maxLon = t.MetersToLatLon(maxx, maxy)
	return
}

type Bounds struct {
	SRID int     `json:"srid"`
	Xmin float64 `json:"xmin"`
	Ymin float64 `json:"ymin"`
	Xmax float64 `json:"xmax"`
	Ymax float64 `json:"ymax"`
}

func main() {

}
