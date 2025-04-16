package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

const (
	numRounds  = 11
	numCourts  = 3
	numAttemps = 1000000
)

var (
	players = []Player{
		{Name: "A", Rating: 1.0},
		{Name: "B", Rating: 1.2},
		{Name: "C", Rating: 1.4},
		{Name: "D", Rating: 1.6},
		{Name: "E", Rating: 1.8},
		{Name: "F", Rating: 2.0},
		{Name: "G", Rating: 2.2},
		{Name: "H", Rating: 2.4},
		{Name: "I", Rating: 2.6},
		{Name: "J", Rating: 2.8},
		{Name: "K", Rating: 3.0},
		{Name: "L", Rating: 3.2},
	}

	// https://americano-padel.com/r/168ACC25-3D27-46FF-9C0E-9FA3B655AF4433
	lineup = [][][]int{
		{
			{11, 4, 3, 2},
			{8, 5, 7, 0},
			{10, 1, 6, 9},
		},
		{
			{8, 6, 4, 9},
			{3, 10, 11, 1},
			{2, 0, 7, 5},
		},
		{
			{3, 5, 2, 10},
			{7, 9, 4, 6},
			{8, 11, 0, 1},
		},
		{
			{4, 1, 11, 0},
			{8, 2, 9, 10},
			{6, 5, 7, 3},
		},
		{
			{7, 8, 3, 1},
			{11, 6, 4, 5},
			{10, 0, 9, 2},
		},
		{
			{11, 2, 6, 0},
			{3, 9, 7, 1},
			{8, 4, 5, 10},
		},
		{
			{5, 1, 4, 10},
			{8, 0, 6, 3},
			{7, 2, 11, 9},
		},
		{
			{8, 9, 11, 5},
			{7, 4, 2, 1},
			{6, 10, 3, 0},
		},
		{
			{4, 0, 7, 10},
			{11, 3, 9, 5},
			{8, 1, 6, 2},
		},
		{
			{5, 2, 6, 1},
			{8, 10, 7, 11},
			{9, 0, 4, 3},
		},
		{
			{8, 3, 4, 2},
			{9, 1, 5, 0},
			{7, 6, 11, 10},
		},
	}
)

type Player struct {
	Name   string
	Rating float64
}

func main() {
	var (
		bestAvg  float64
		bestMax  float64
		bestList []Player
	)

	for a := range numAttemps {
		list := shufflePlayers()
		maxDiff, avgDiff := listStat(list)
		fmt.Printf("[%d] max: %.3f, avg: %.3f, list: %v\n", a, maxDiff, avgDiff, list)

		if bestMax == 0 || maxDiff < bestMax {
			bestAvg = avgDiff
			bestMax = maxDiff
			bestList = list
		}
	}

	fmt.Printf("Best - max: %.3f, avg: %.3f, list: %v\n", bestMax, bestAvg, bestList)
	printList(bestList)
}

func getPlayer(list []Player, round, court, position int) Player {
	index := lineup[round][court][position]
	return list[index]
}

func shufflePlayers() []Player {
	list := make([]Player, len(players))
	copy(list, players)
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
}

func listStat(list []Player) (maxDiff float64, avgDiff float64) {
	var sumDiff float64
	var count int
	for r := range numRounds {
		for c := range numCourts {
			player1 := getPlayer(list, r, c, 0)
			player2 := getPlayer(list, r, c, 1)
			player3 := getPlayer(list, r, c, 2)
			player4 := getPlayer(list, r, c, 3)

			diff := math.Abs(player1.Rating + player2.Rating - player3.Rating - player4.Rating)
			if diff > maxDiff {
				maxDiff = diff
			}
			sumDiff += diff
			count++
		}
	}
	avgDiff = sumDiff / float64(count)
	return
}

func printList(list []Player) {
	for r := range numRounds {
		fmt.Printf("Round %d\n", r+1)
		for c := range numCourts {
			player1 := getPlayer(list, r, c, 0)
			player2 := getPlayer(list, r, c, 1)
			player3 := getPlayer(list, r, c, 2)
			player4 := getPlayer(list, r, c, 3)

			fmt.Printf("  Court %d: %s (%.1f) / %s (%.1f) – %s (%.1f) / %s (%.1f)\n", c+1,
				player1.Name, player1.Rating, player2.Name, player2.Rating, player3.Name, player3.Rating, player4.Name, player4.Rating)
		}
		fmt.Println()
	}
}
