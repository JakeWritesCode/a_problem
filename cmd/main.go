package main

import "a_problem/matching/domain"

func main() {
	d := domain.NewDomain()
	d.CreateLastActive()
	//d.SeedCoreData()
	//d.SeedParticipants()
}
