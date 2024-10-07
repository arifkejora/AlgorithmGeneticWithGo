package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Name string `json:"name"`
}

type ScheduleInput struct {
	Employees   []Employee `json:"employees"`
	Population  int        `json:"population"`
	Generations int        `json:"generations"`
}

type Schedule struct {
	Name string            `json:"name"`
	Days map[string]string `json:"days"`
}

var weekdays = []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}

func main() {
	r := gin.Default()

	r.POST("/generate-schedule", GenerateSchedule)

	r.Run(":8080")
}

func GenerateSchedule(c *gin.Context) {
	var input ScheduleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schedules := GeneticAlgorithm(input.Employees, input.Population, input.Generations)

	c.JSON(http.StatusOK, schedules)
}

func GeneticAlgorithm(employees []Employee, populationSize int, generations int) []Schedule {
	rand.Seed(time.Now().UnixNano())
	var bestSchedules []Schedule

	wfhAssigned := make(map[string]bool)

	population := make([][]Schedule, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = generateRandomSchedule(employees, &wfhAssigned)
	}

	for gen := 0; gen < generations; gen++ {
		population = evolvePopulation(population)
	}

	bestSchedules = population[0]
	return bestSchedules
}

func generateRandomSchedule(employees []Employee, wfhAssigned *map[string]bool) []Schedule {
	var schedules []Schedule
	for _, emp := range employees {
		wfhDay := getUniqueWFHDay(wfhAssigned)

		days := make(map[string]string)
		for _, day := range weekdays {
			if day == wfhDay {
				days[day] = "WFH"
			} else {
				days[day] = "WFO"
			}
		}

		schedules = append(schedules, Schedule{Name: emp.Name, Days: days})
	}
	return schedules
}

func getUniqueWFHDay(wfhAssigned *map[string]bool) string {
	var availableDays []string
	for _, day := range weekdays {
		if !(*wfhAssigned)[day] {
			availableDays = append(availableDays, day)
		}
	}
	rand.Shuffle(len(availableDays), func(i, j int) {
		availableDays[i], availableDays[j] = availableDays[j], availableDays[i]
	})
	selectedWFHDay := availableDays[0]
	(*wfhAssigned)[selectedWFHDay] = true
	return selectedWFHDay
}

func evolvePopulation(population [][]Schedule) [][]Schedule {
	return population
}
