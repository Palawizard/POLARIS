package utils

var currentTurn int // tracks the active battle turn

// SendTurn sets the active battle turn.
func SendTurn(turn int) { currentTurn = turn }

// GetTurn returns the last turn set via SendTurn.
func GetTurn() int { return currentTurn }
