package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func SendEmail(email string) error {
	if rand.Intn(10) < 3 {
		return errors.New("stmp error")
	}

	time.Sleep(2 * time.Second)
	fmt.Println("email send to:", email)
	return nil
}
