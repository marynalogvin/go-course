//--Summary:
//  Write a program to display server status.
//
//--Requirements:

//* Perform the following status changes and display server info:

package main

import "fmt"

const (
	Online int = iota
	Offline
	Maintenance
	Retired
)

// * Create a function to print server status, including:
//   - Number of servers
//   - Number of servers for each status (Online, Offline, Maintenance, Retired)
func printServerStatus(serverStatuses map[string]int) {
	fmt.Println("Number of services: ", len(serverStatuses))
	statuses := make(map[int]int)
	for _, status := range serverStatuses {
		switch status {
		case Online:
			statuses[Online] += 1
		case Offline:
			statuses[Offline] += 1
		case Maintenance:
			statuses[Maintenance] += 1
		case Retired:
			statuses[Retired] += 1
		default:
			panic("unknown server status")
		}
	}
	fmt.Println("Online: ", statuses[Online])
	fmt.Println("Offline: ", statuses[Offline])
	fmt.Println("Maintenance: ", statuses[Maintenance])
	fmt.Println("Retired: ", statuses[Retired])
}
func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	//* Store the existing slice of servers in a map
	//* Default all of the servers to `Online`
	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}
	//  - display server info
	printServerStatus(serverStatus)
	//  - change `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired
	//  - change `aiur` to `Offline`
	serverStatus["aiur"] = Offline
	// - display server info
	printServerStatus(serverStatus)
	// - change all servers to `Maintenance`
	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}
	// - display server info
	printServerStatus(serverStatus)

}
