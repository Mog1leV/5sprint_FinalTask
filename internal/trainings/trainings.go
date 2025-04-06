package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	str := strings.Split(datastring, ",")
	if len(str) != 3 {
		return errors.New("error during parsing")
	}
	steps, err := strconv.Atoi(str[0])
	if err != nil {
		return err
	}
	t.Steps = steps

	switch {
	case str[1] == "Бег":
		t.TrainingType = str[1]
	case str[1] == "Ходьба":
		t.TrainingType = str[1]
	default:
		return errors.New("unknown type of training")
	}
	time, err := time.ParseDuration(str[2])
	if err != nil {
		return err
	}
	t.Duration = time
	return nil

}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 {
		return "", errors.New("invalid number of steps")
	}
	distance := spentenergy.Distance(t.Steps)
	if t.Duration <= 0 {
		return "", errors.New("invalid training duration")
	}
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)
	switch t.TrainingType {
	case "Бег":
		calories := spentenergy.RunningSpentCalories(t.Steps, t.Height, t.Duration)
		message := fmt.Sprintf(
			"Тип тренировки: %s.\n"+
				"Длительность: %.2f ч.\n"+
				"Дистанция: %.2f км\n"+
				"Скорость: %.2f км/ч\n"+
				"Сожгли калорий: %.2f\n",
			t.TrainingType,
			t.Duration.Hours(),
			distance,
			averageSpeed,
			calories,
		)
		if calories == 0 {
			return "", errors.New("error in calorie calculations")
		}
		return message, nil
	case "Ходьба":
		calories := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		message := fmt.Sprintf(
			"Тип тренировки: %s.\n"+
				"Длительность: %.2f ч.\n"+
				"Дистанция: %.2f км\n"+
				"Скорость: %.2f км/ч\n"+
				"Сожгли калорий: %.2f\n",
			t.TrainingType,
			t.Duration.Hours(),
			distance,
			averageSpeed,
			calories,
		)
		if calories == 0 {
			return "", errors.New("error in calorie calculations")
		}
		return message, nil
	default:
		return "", errors.New("unknown type of training")
	}
}
