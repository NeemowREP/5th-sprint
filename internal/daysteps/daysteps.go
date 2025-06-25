package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	sliceData := strings.Split(datastring, ",")
	if len(sliceData) != 2 {
		return errors.New("incorrect number of elements")
	}
// Проверка на пустую строку
if strings.TrimSpace(sliceData[0]) == "" && sliceData[0] != "" {
	return errors.New("steps value cannot be empty")
}

	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return err
	} else if steps <= 0 {
		return fmt.Errorf("incorrect value of steps, must be positive - have: %d", steps)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(sliceData[1])
	if err != nil {
		return err
	} else if duration <= 0 {
		return fmt.Errorf("invalid duration value: %v (must be positive)", duration)
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	if ds.Steps <= 0 {
		return "", fmt.Errorf("invalid steps: %d (must be positive)", ds.Steps)
	} else if ds.Duration <= 0 {
		return "", fmt.Errorf("invalid duration: %v (must be positive)", ds.Duration)
	} else if ds.Personal.Weight <= 0 || ds.Personal.Height <= 0 {
		return "", fmt.Errorf("invalid weight or height (must be positive)")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("miscalculation of calories: %v", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
