package daysteps

import (
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
	sl := strings.Split(datastring, ",")

	if len(sl) != 2 {
		return fmt.Errorf("len of datastring != 2")
	}

	steps, err := strconv.Atoi(sl[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return fmt.Errorf("wrong steps <= 0")
	}

	duration, err := time.ParseDuration(sl[1])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return fmt.Errorf("wrong duration <= 0")
	}

	ds.Steps = steps
	ds.Duration = duration
	return

}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	text := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, ds.Steps, distance, calories)

	return text, nil
}
