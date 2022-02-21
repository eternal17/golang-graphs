package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Graph structure
type Graph struct {
	rooms []*Room
}

//Room structure
type Room struct {
	key      string
	adjacent []*Room
}

//Add Room to a graph
func (g *Graph) AddRoom(k string) {
	if contains(g.rooms, k) {
		err := fmt.Errorf("Room %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.rooms = append(g.rooms, &Room{key: k})
	}
}

//getRoom returns a pointer to the Room key integer
func (g *Graph) getRoom(k string) *Room {
	for i, v := range g.rooms {
		if v.key == k {
			return g.rooms[i]
		}
	}
	return nil
}

//contains checks if the Room key exists

func contains(s []*Room, k string) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// handling all the errors with specific messages
func allErrors() {
	// check if file contains start and end
	for i := range readAntsFile("ants.txt") {
		if readAntsFile("ants.txt")[i] != "##start" {
			fmt.Println("ERROR: invalid data format, no start room found")
		}
	}

}

func readAntsFile(filename string) []string {
	file, _ := os.Open(filename)
	scanned := bufio.NewScanner(file)
	scanned.Split(bufio.ScanLines)

	var lines []string

	for scanned.Scan() {
		lines = append(lines, scanned.Text())
	}
	return lines
}

func main() {
	test := &Graph{}

	allErrors()

	// reading all rooms
	// adding all rooms
	for i, line := range readAntsFile("ants.txt") {
		if strings.Contains(string(line), " ") {
			test.AddRoom(strings.Split(readAntsFile("ants.txt")[i], " ")[0])
		}
		// adding all edges from and to rooms
		if strings.Contains(string(line), "-") {
			test.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[0], strings.Split(readAntsFile("ants.txt")[i], "-")[1])
		}

	}

	// add all edges

	test.Print()
}

//Add edge to the graph. deals with a directional graph only

func (g *Graph) AddEdge(from, to string) {
	//get Room
	fromRoom := g.getRoom(from)
	toRoom := g.getRoom(to)

	//check error
	if fromRoom == nil || toRoom == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromRoom.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom)
	}
	//add edge
}

//Print will print the adjacent list for each Room of the graph

func (g *Graph) Print() {
	// fmt.Println(readAntsFile("ants.txt"))

	for i, v := range g.rooms {
		if i == 0 {
			fmt.Printf("\n Start Room %v : ", v.key)

		} else if i == len(g.rooms)-1 {
			fmt.Printf("\n End Room %v : ", v.key)

		} else {
			fmt.Printf("\n Room %v : ", v.key)

		}
		for _, v := range v.adjacent {
			fmt.Printf(" %v", v.key)
		}
	}
	fmt.Println()
}
