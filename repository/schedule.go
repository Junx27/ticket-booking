package repository

import (
	"context"
	"errors"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) entity.ScheduleRepository {
	return &ScheduleRepository{
		db: db,
	}
}
func (r *ScheduleRepository) GetUserID(id uint) (uint, error) {
	var schedule entity.Schedule
	if err := r.db.First(&schedule, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("schedule not found")
		}
		return 0, err
	}
	return schedule.UserID, nil
}
func (r *ScheduleRepository) GetManyByUser(ctx context.Context, userID uint) ([]interface{}, error) {

	schedules, err := r.GetMany(ctx, userID)
	if err != nil {

		return nil, err

	}
	result := make([]interface{}, len(schedules))
	for i, schedule := range schedules {

		result[i] = schedule

	}
	return result, nil

}

func (r *ScheduleRepository) GetMany(ctx context.Context, userId uint) ([]*entity.Schedule, error) {
	var schedules []*entity.Schedule
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *ScheduleRepository) GetManyCustomer(ctx context.Context) ([]*entity.Schedule, error) {
	var schedules []*entity.Schedule
	if err := r.db.WithContext(ctx).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *ScheduleRepository) GetOne(ctx context.Context, scheduleId uint) (*entity.Schedule, error) {
	schedule := &entity.Schedule{}
	if err := r.db.WithContext(ctx).Where("id = ?", scheduleId).First(&schedule).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *ScheduleRepository) CreateOne(ctx context.Context, schedule *entity.Schedule) (*entity.Schedule, error) {
	availableSeats := make([]int64, 0)
	for i := 1; i <= 5; i++ {
		availableSeats = append(availableSeats, int64(i))
	}
	schedule.AvailableSeats = pq.Int64Array(availableSeats)
	if err := r.db.WithContext(ctx).Create(schedule).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *ScheduleRepository) UpdateOne(ctx context.Context, scheduleId uint, updateData map[string]interface{}) (*entity.Schedule, error) {
	schedule := &entity.Schedule{}
	if err := r.db.WithContext(ctx).Where("id = ?", scheduleId).First(&schedule).Error; err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Model(&schedule).Updates(updateData).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *ScheduleRepository) UpdateSeatsStatus(ctx context.Context, scheduleId uint, seatsData map[int]interface{}) (*entity.Schedule, error) {
	schedule := &entity.Schedule{}
	if err := r.db.WithContext(ctx).First(&schedule, scheduleId).Error; err != nil {
		return nil, err
	}

	availableSeatsMap := make(map[int64]bool)
	for _, seat := range schedule.AvailableSeats {
		availableSeatsMap[seat] = true
	}

	bookedSeats := make(map[int]bool)
	for key, value := range seatsData {
		if status, ok := value.(string); ok {
			if status == "booked" {
				bookedSeats[key] = true
			} else if status == "cancel" {
				bookedSeats[key] = false
			}
		}
	}

	var updatedAvailableSeats []int64
	for seat, isBooked := range bookedSeats {
		if isBooked {
			delete(availableSeatsMap, int64(seat))
		} else {
			availableSeatsMap[int64(seat)] = true
		}
	}
	for seat := range availableSeatsMap {
		updatedAvailableSeats = append(updatedAvailableSeats, seat)
	}

	schedule.AvailableSeats = pq.Int64Array(updatedAvailableSeats)

	if err := r.db.WithContext(ctx).Save(&schedule).Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *ScheduleRepository) DeleteOne(ctx context.Context, scheduleId uint) error {
	if err := r.db.WithContext(ctx).Where("id = ?", scheduleId).Delete(&entity.Schedule{}).Error; err != nil {
		return err
	}
	return nil
}
