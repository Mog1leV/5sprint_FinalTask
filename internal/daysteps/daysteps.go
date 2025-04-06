package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	str := strings.Split(datastring, ",")
	if len(str) != 2 {
		return errors.New("data parsing error")
	}
	steps, err := strconv.Atoi(str[0])
	if err != nil {
		return err
	}
	ds.Steps = steps
	time, err := time.ParseDuration(str[1])
	if err != nil {
		return err
	}
	ds.Duration = time
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	time := ds.Duration
	steps := ds.Steps
	if time <= 0 {
		return "", errors.New("duration value error")
	}
	distance := spentenergy.Distance(steps)
	calories := spentenergy.WalkingSpentCalories(steps, ds.Weight, ds.Height, time)
	if calories <= 0 {
		return "", errors.New("error in calorie calculations")
	}
	message := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps,
		distance,
		calories,
	)
	return message, nil

}
