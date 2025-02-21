package board

type Stone int

const (
    Empty Stone = iota  // 0
    Black              // 1
    White              // 2
)

type Board struct {
    Size  int       // Size of the board (typically 19)
    Grid  [][]Stone // 2D slice to store the stones
}

func NewBoard() Board {
	grid := make([][]Stone, 19)
	for i:= range grid {
		grid[i] = make([]Stone, 19)
	}

	return Board{
		Size: 19,
		Grid: grid,
	}
}

func (b Board) GetStone(x, y int) *Stone {
	return &b.Grid[x][y]
}


func (b *Board) PlaceStone(stone Stone, x, y int) {
	spotToPlace := b.GetStone(x, y)

	if(stone == Empty || *spotToPlace != Empty){
		return 
	}

	*spotToPlace = stone
}