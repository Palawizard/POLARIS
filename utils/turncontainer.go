package utils

var currentturn int

func SendTurn(turn int) {
	currentturn = turn
}

func GetTurn() int {
	return currentturn
}
