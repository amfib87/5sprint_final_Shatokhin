package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

const (
	run  = "Бег"
	walk = "Ходьба"
)

func (t *Training) Parse(datastring string) (err error) {
	sl := strings.Split(datastring, ",")

	if len(sl) != 3 {
		return fmt.Errorf("len of datastring != 3")
	}

	steps, err := strconv.Atoi(sl[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("wrong steps <= 0")
	}

	duration, err := time.ParseDuration(sl[2])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return fmt.Errorf("wrong duration <= 0")
	}

	t.Steps = steps
	t.TrainingType = sl[1]
	t.Duration = duration

	return
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	switch t.TrainingType {
	case run:
		cal, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		calories = cal

	case walk:
		cal, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		calories = cal

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	text := fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)

	return text, nil
}
