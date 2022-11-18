// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func TestHealth(t *testing.T) {
	player := Player{name: "Mary", health: 100, maxHealth: 100}
	player.addHealth(100)
	if player.health < player.maxHealth {
		t.Errorf("Incorrect health value after addHealth:%v, want: %v ", player.health, player.maxHealth)
	}
	player.subtractHealth(300)
	if player.health < 0 {
		t.Errorf("Incorrect health value after subtractHealth: %v, want: 0", player.health)
	}
}

func TestEnergy(t *testing.T) {
	player := Player{name: "Alex", energy: 100, maxEnergy: 100}
	player.addEnergy(100)
	if player.energy < player.maxEnergy {
		t.Errorf("Incorrect health value after addHealth:%v, want: %v ", player.energy, player.maxEnergy)
	}
	player.subtractEnergy(300)
	if player.energy < 0 {
		t.Errorf("Incorrect health value after subtractHealth: %v, want: 0", player.energy)
	}
}
