package util

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

func GenerateUUID() string {
	return uuid.New().String()
}


func GenerateRandomScore() int {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Intn(801) + 150 
}