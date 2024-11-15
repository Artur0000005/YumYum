package ftracker // драсте

import (
    "fmt"
    "math"
)

const (
    lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   //
)

func distance ( action int ) float64 {
    return float64 ( action ) * lenStep / mInKm
}

func meanSpeed ( action int, duration float64) float64 {

    if duration == 0 {
		return 0
	}
	distance := distance(action)
	return distance / duration
}

func ShowTrainingInfo ( action int, trainingType string, duration, weight, height float64, lenghPool, coiuntPool int ) string {

    switch {
    case trainingType == "Бег":

        distance := distance(action)                               
		speed := meanSpeed(action, duration)                       
		calories := RunningSpentCalories(action, weight, duration) 

        return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)

    case trainingType == "Ходьба":

        distance := distance(action)                                       
		speed := meanSpeed(action, duration)                               
		calories := WalkingSpentCalories(action, duration, weight, height) 

        return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)

    case trainingType == "Плавание":

        distance := distance(action)                                               
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)               
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)

    return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)

        default:

            return "неизвестный тип тренировки"

    const (
        runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79
    )

    func RunningSpentCalories(action int, weight, duration float64) float64 {        if duration == 0 {
            return 0
        }
        return (runningCaloriesMeanSpeedMultiplier * distance(action) / duration * runningCaloriesMeanSpeedShift) * weight / mInKm * duration * minInH
    }
    const (
        walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029
    )
    func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
        if duration == 0 {
            return 0
        }
        avgSpeedInMinutes := (distance(action) / duration) * kmhInMsec
        return (walkingCaloriesWeightMultiplier*weight + (math.Pow(avgSpeedInMinutes, 2)/height*cmInM)*walkingSpeedHeightMultiplier*weight) * duration * minInH
    }
    const (
        swimmingCaloriesMeanSpeedShift   = 1.1 // среднее количество сжигаемых колорий при плавании относительно скорости.
	swimmingCaloriesWeightMultiplier = 2
    )
    func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
        if duration == 0 {
            return 0
        }
        return float64(lengthPool) * float64(countPool) / mInKm / duration
    }
    func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {

        if duration == 0 {
            return 0
        }
        return (swimmingMeanSpeed(lengthPool, countPool, duration) + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
    }
}
