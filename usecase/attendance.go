package usecase

import "implementation-redis-golang/repository"

// UserService is responsible for user-related operations
type UserService struct {
	repo repository.AttendanceRepository
}

func NewUserService(repo repository.AttendanceRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) RecordAttendance(userID string) error {
	return u.repo.CheckIn(userID)
}

func (u *UserService) GetAttendance(userID string, limit int64) ([]string, error) {
	return u.repo.GetAttendance(userID, limit)
}
