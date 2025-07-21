package main

import (
	//"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	Tiles []*MapTile
}

func NewLevel() Level {
	l := Level{}
	l.Tiles = l.CreateTiles()
	return l
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (x * gd.ScreenHeight) + y
}

func GetTileFromIndex(x int, y int, level []*MapTile) *MapTile {
	t := level[GetIndexFromXY(x, y)]
	return t
}

func SetTile(x int, y int, level []*MapTile, tileType TileType) {
	tile := GetTileFromIndex(x, y, level)
	tile.SetTileType(tileType)
}

func GetAdjacentTiles(x int, y int, level []*MapTile) []*MapTile {
	gd := NewGameData()
	tiles := make([]*MapTile, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !(x+i < 0) && !(y+j < 0) && !(y+j >= gd.ScreenHeight) && !(y+j >= gd.ScreenWidth) && !(j == 0 && i == 0) {
				tiles = append(tiles, GetTileFromIndex(x+i, y+j, level))
			}
		}
	}
	return tiles
}

type VoronoiTile struct {
	tile *MapTile
	zone int
}

func PythagoreanDistance(presentTile *MapTile, targetTile *MapTile) int {
	a := math.Pow((float64(presentTile.PixelX - targetTile.PixelX)), 2)
	b := math.Pow((float64(presentTile.PixelY - targetTile.PixelY)), 2)
	return int(math.Sqrt(a + b))
}

func (level *Level) CreateTiles() []*MapTile {
	gd := NewGameData()
	tiles := make([]*MapTile, 0)
	//grass := NewTileType(false, "grass.png")
	sea := NewTileType("sea", true)
	island := NewTileType("island", true)
	grass := NewTileType("grass", true)
	sea.NewImage("sea.png")
	sea.NewImage("sea_dither_2.png")
	island.NewImage("island1.png")
	grass.NewImage("grass.png")
	//seed the entire map with sea
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := NewTile(x*gd.TileWidth, y*gd.TileHeight, sea)
			tiles = append(tiles, tile)
		}
	}
	zone_size := 10
	zone_height := gd.ScreenHeight / zone_size
	zone_width := gd.ScreenWidth / zone_size
	for i := 0; i < zone_width; i++ {
		for j := 0; j < zone_height; j++ {
			zone_type := GetRandomInt(3)
			//0 = open sea; 1 = archipelago; 2 = big island
			if zone_type == 1 {
				num_islands := GetDiceRoll(3)
				for a := 0; a < num_islands; a++ {
					random_x := GetRandomInt(zone_size-2) + 1 + i*zone_size
					random_y := GetRandomInt(zone_size-2) + 1 + j*zone_size
					SetTile(random_x, random_y, tiles, island)
					adj_tiles := GetAdjacentTiles(random_x, random_y, tiles)
					for _, adj_tile := range adj_tiles {
						if adj_tile.Type.Name == "sea" {
							adj_tile.TypeImageIndex = 1
						}
					}
				}
			}
			if zone_type == 2 {
				//create small voronoi map by placing 5 points
				voronoiPointsMap := make([]*VoronoiTile, 0)
				for a := 0; a < 5; a++ {
					random_x := GetRandomInt(zone_size-2) + 1 + i*zone_size
					random_y := GetRandomInt(zone_size-2) + 1 + j*zone_size
					voronoiTile := &VoronoiTile{tile: GetTileFromIndex(random_x, random_y, tiles), zone: a}
					voronoiPointsMap = append(voronoiPointsMap, voronoiTile)
				}
				//determining central voronoi zone
				central_voronoi_zone := 0
				current_max_distance := 99999
				for _, point := range voronoiPointsMap {
					distance_to_center := PythagoreanDistance(point.tile, GetTileFromIndex(zone_size/2+i*zone_size, zone_size/2+j*zone_size, tiles))
					if distance_to_center < current_max_distance {
						central_voronoi_zone = point.zone
						current_max_distance = distance_to_center
					}
				}
				//measure the distance and apply
				voronoiMap := make([]*VoronoiTile, 0)
				for x := 0; x < zone_size; x++ {
					for y := 0; y < zone_size; y++ {
						x_value := x + i*zone_size
						y_value := y + j*zone_size
						current_tile := GetTileFromIndex(x_value, y_value, tiles)
						nearest_point := 0
						//arbitrarily large, needs update
						nearest_closest_distance := 999999
						for _, voronoiPoint := range voronoiPointsMap {
							if PythagoreanDistance(voronoiPoint.tile, current_tile) < nearest_closest_distance {
								nearest_closest_distance = PythagoreanDistance(voronoiPoint.tile, current_tile)
								nearest_point = voronoiPoint.zone
							}
						}
						voronoiTile := &VoronoiTile{tile: current_tile, zone: nearest_point}
						voronoiMap = append(voronoiMap, voronoiTile)
					}
				}

				//selecting the central voronoi zone to make it all into grass
				for _, voronoiTile := range voronoiMap {
					if voronoiTile.zone == central_voronoi_zone {
						voronoiTile.tile.SetTileType(grass)
					}
				}
			}
		}
	}
	//making the adjacent tiles look like shallow water
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			current_tile := GetTileFromIndex(x, y, tiles)
			if current_tile.Type.Name == "grass" {
				adj_tiles := GetAdjacentTiles(x, y, tiles)
				for _, adj_tile := range adj_tiles {
					if adj_tile.Type.Name == "sea" {
						adj_tile.TypeImageIndex = 1
					}
				}
			}
		}
	}
	return tiles
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := GetTileFromIndex(x, y, level.Tiles)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image(), op)
		}
	}
}
