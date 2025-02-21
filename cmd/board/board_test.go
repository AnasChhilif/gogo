package board

import (
    "testing"
)

func TestNewBoard(t *testing.T) {
    board := NewBoard()
    
    // Test board size
    if board.Size != 19 {
        t.Errorf("Expected board size to be 19, got %d", board.Size)
    }
    
    // Test if all positions are empty
    for i := 0; i < board.Size; i++ {
        for j := 0; j < board.Size; j++ {
            if board.Grid[i][j] != Empty {
                t.Errorf("Expected position (%d,%d) to be Empty, got %d", i, j, board.Grid[i][j])
            }
        }
    }
}

func TestGetStone(t *testing.T) {
    board := NewBoard()
    
    // Test getting stone at various positions
    testCases := []struct {
        x, y     int
        expected Stone
    }{
        {0, 0, Empty},
        {18, 18, Empty},
        {5, 5, Empty},
    }

    for _, tc := range testCases {
        stone := board.GetStone(tc.x, tc.y)
        if *stone != tc.expected {
            t.Errorf("Expected stone at (%d,%d) to be %d, got %d", tc.x, tc.y, tc.expected, *stone)
        }
    }
}

func TestPlaceStone(t *testing.T) {
    board := NewBoard()
    
    testCases := []struct {
        name     string
        x, y     int
        stone    Stone
        expected Stone
    }{
        {"Place black stone", 0, 0, Black, Black},
        {"Place white stone", 5, 5, White, White},
        {"Place empty stone", 1, 1, Empty, Empty},
        {"Place stone on occupied spot", 0, 0, White, Black}, // Should remain Black from first test
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            board.PlaceStone(tc.stone, tc.x, tc.y)
            result := *board.GetStone(tc.x, tc.y)
            
            if tc.name == "Place stone on occupied spot" {
                // Verify that we can't place a stone on an occupied spot
                if result != tc.expected {
                    t.Errorf("Expected stone at (%d,%d) to remain %d, got %d", 
                        tc.x, tc.y, tc.expected, result)
                }
            } else if tc.stone != Empty && result != tc.stone {
                t.Errorf("Expected stone at (%d,%d) to be %d, got %d", 
                    tc.x, tc.y, tc.stone, result)
            }
        })
    }
}

func TestBoardBoundaries(t *testing.T) {
    board := NewBoard()
    
    // Test that we can place stones at the board boundaries
    testPositions := []struct {
        x, y  int
        stone Stone
    }{
        {0, 0, Black},       // Top-left corner
        {18, 0, White},      // Top-right corner
        {0, 18, White},      // Bottom-left corner
        {18, 18, Black},     // Bottom-right corner
    }

    for _, pos := range testPositions {
        board.PlaceStone(pos.stone, pos.x, pos.y)
        result := *board.GetStone(pos.x, pos.y)
        if result != pos.stone {
            t.Errorf("Failed to place stone at boundary (%d,%d): expected %d, got %d",
                pos.x, pos.y, pos.stone, result)
        }
    }
}