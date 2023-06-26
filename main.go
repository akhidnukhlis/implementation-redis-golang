package main

import (
	"fmt"
	"implementation-redis-golang/repository"
	"implementation-redis-golang/usecase"
)

func main() {
	// Initialize Redis client and attendance repository
	attendanceRepo := repository.NewRedisAttendanceRepository("localhost:6379", "password")

	// Create UserService instance
	userService := usecase.NewUserService(attendanceRepo)

	// Record attendance for a user
	userID := "User-1"
	err := userService.RecordAttendance(userID)
	if err != nil {
		fmt.Println("Error recording attendance:", err)
	}

	// Get attendance history for user
	attendanceLimit := int64(5)
	attendance, err := userService.GetAttendance(userID, attendanceLimit)
	if err != nil {
		fmt.Println("Error getting attendance:", err)
	} else {
		fmt.Printf("Attendance history for user %s:\n", userID)
		for i, att := range attendance {
			fmt.Printf("%d. %s\n", i+1, att)
		}
	}

	// Close Redis client connection
	err = attendanceRepo.Close()
	if err != nil {
		fmt.Println("Error closing Redis client:", err)
	}
}
