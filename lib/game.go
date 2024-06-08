package lib

import (
	"fmt"
)

func StartSession() {

	player, err := ReadPlayerData(filePath)
	if err != nil {
		fmt.Printf("Error reading player data: %v\n", err)
		return
	}

	fmt.Printf("Initial player data: %+v\n", player)

	Update(player)

	updatedPlayer, err := ReadPlayerData(filePath)
	if err != nil {
		fmt.Printf("Error reading updated player data: %v\n", err)
		return
	}

	fmt.Printf("Updated player data: %+v\n", updatedPlayer)
}
