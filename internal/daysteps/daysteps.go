package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	stepLength = 0.65
	mInKm      = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("неверный формат строки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("неверное количество шагов: %w", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше 0")
	}

	dur, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("неверный формат продолжительности: %w", err)
	}
	if dur <= 0 {
		return 0, 0, errors.New("продолжительность должна быть больше 0")
	}

	return steps, dur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps,
		distanceKm,
		calories,
	)
}
