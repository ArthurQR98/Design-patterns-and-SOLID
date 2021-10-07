package repository

import "fmt"

type Repo struct{}

func (r *Repo) Add() {
	fmt.Println("Add in database")
}

func (r *Repo) Find() {
	fmt.Println("Finding from database")
}

func (r *Repo) Update() {
	fmt.Println("Update in database")
}

func (r *Repo) Remove() {
	fmt.Println("Deleting from database")
}
