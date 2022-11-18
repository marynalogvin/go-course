//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

// * Implement a player having the following statistics:
//   - Health, Max Health
//   - Energy, Max Energy
//   - Name
type Player struct {
	name              string
	health, maxHealth int
	energy, maxEnergy int
}

func (player *Player) addHealth(amount int) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Printf("%v added %v health = %v\n", player.name, amount, player.health)
}

func (player *Player) subtractHealth(amount int) {
	if player.health-amount < player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Printf("%v applied %v damage = %v\n", player.name, amount, player.health)
}
func (player *Player) addEnergy(amount int) {
	player.energy += amount
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Printf("%v added %v energy = %v\n", player.name, amount, player.energy)
}

func (player *Player) subtractEnergy(amount int) {
	if player.energy-amount < player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Printf("%v lost %v energy = %v\n", player.name, amount, player.energy)
}

func main() {
	player1 := Player{"Jacob", 5, 100, 4, 100}

	player1.addHealth(50)
	player1.addEnergy(105)
	player1.subtractHealth(70)
	player1.subtractEnergy(99)
}
