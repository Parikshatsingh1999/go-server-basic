package logs

import (
	"fmt"
	"os"
	"time"
)

func AddLog(message string) {
	pathName := "logs/loggs.txt"

	// Open the file for writing (create if it doesn't exist)
	file, err := os.OpenFile(pathName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error logging message:", message, "in opening file:", err)
		return
	}
	_, err = file.WriteString(message + " at time " + time.Now().Format("2006-01-02 15:04:05") + "\n")
	if err != nil {
		fmt.Println("Error writing message", message, "to file:", err)
		return
	}
	defer file.Close() // Defer closing the file until the function exits

	// same as defer, written at the end itself
	//file.Close()
}
