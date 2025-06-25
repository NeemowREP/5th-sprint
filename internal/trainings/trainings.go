package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)
type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {

// Парсим полученуую строку в слайс строк
	sliceData := strings.Split(datastring, ",")

// Проверка на количество данных
	if len(sliceData) != 3 {
		return errors.New("incorrect number of elements")
	}
// Переводим тип данных (из стринг в инт), проверяем на ошибку и положительное значение
// Присваиваем значение в структуру
	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return err
	} else if steps <= 0 {
		return fmt.Errorf("incorrect value of steps, must be positive - have: %d", steps)
	}
	t.Steps = steps

// Проверяем на пустое значение типа тренировок, если все корректно, присваиваем значение в структуру
	trainingType := sliceData[1]
	if trainingType == "" {
		return errors.New("incorrect type of training")
	}
	t.TrainingType = trainingType

// Проверяем корректность строки времени и присваиваем значение в структуру
	duration, err := time.ParseDuration(sliceData[2])
	if err != nil {
		return err
	} else if duration <= 0 {
		return fmt.Errorf("invalid duration value: %v (must be positive)", duration)
	}
	t.Duration = duration
	return nil
}



func (t Training) ActionInfo() (string, error) {

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки %s", t.TrainingType)
	}

	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	text := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
	t.TrainingType,
	t.Duration.Hours(),
	distance,
	speed,
	calories)

	return text, nil
}
