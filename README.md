# Algorithm Genetic With Go

This program is a web application built using the Go programming language and the Gin framework to generate employee work schedules using a genetic algorithm. The application allows users to input employee names and produces optimal work schedules where each employee has 1 work-from-home (WFH) day and 4 in-office (WFO) days per week.

## Program Description

The goal of this program is to manage and schedule employee working hours efficiently. Each employee will receive one randomly assigned WFH day, while the remaining days of the week will be designated as WFO. The genetic algorithm is used to optimize this scheduling process, ensuring that no two employees share the same WFH day.

## What is a Genetic Algorithm?

A genetic algorithm is a search and optimization technique inspired by the biological process of evolution. It utilizes concepts such as natural selection, reproduction, and mutation to find optimal solutions for a problem. In the context of scheduling, this algorithm will:
1. **Generate an initial population** of random work schedules.
2. **Evaluate** the quality of each schedule based on specific criteria (ensuring all employees have 1 WFH day).
3. **Reproduce** and **modify** schedules to create new generations until the desired results are achieved.

## Process Overview

1. **Input**: Users provide a JSON object containing employee names, the population size, and the number of generations for the genetic algorithm.
2. **Schedule Generation**: The algorithm creates a random schedule for each employee, ensuring each has 1 WFH day and 4 WFO days.
3. **Output**: The program returns a JSON response containing the generated schedules for each employee.

## API Endpoints

### Generate Schedule

- **Endpoint**: `/generate-schedule`
- **Method**: `POST`
- **Request Parameters**:
    ```json
    {
        "employees": [
            {
                "name": "Arif"
            },
            {
                "name": "Intan"
            },
            {
                "name": "Savira"
            },
            {
                "name": "Handa"
            },
            {
                "name": "Ari"
            }
        ],
        "population": 1,
        "generations": 100000000
    }
    ```

- **Response Example**:
    ```json
    [
        {
            "name": "Arif",
            "days": {
                "Monday": "WFO",
                "Tuesday": "WFO",
                "Wednesday": "WFH",
                "Thursday": "WFO",
                "Friday": "WFO"
            }
        },
        {
            "name": "Intan",
            "days": {
                "Monday": "WFO",
                "Tuesday": "WFH",
                "Wednesday": "WFO",
                "Thursday": "WFO",
                "Friday": "WFO"
            }
        },
        {
            "name": "Savira",
            "days": {
                "Monday": "WFO",
                "Tuesday": "WFO",
                "Wednesday": "WFO",
                "Thursday": "WFH",
                "Friday": "WFO"
            }
        },
        {
            "name": "Handa",
            "days": {
                "Monday": "WFO",
                "Tuesday": "WFO",
                "Wednesday": "WFO",
                "Thursday": "WFO",
                "Friday": "WFH"
            }
        },
        {
            "name": "Ari",
            "days": {
                "Monday": "WFH",
                "Tuesday": "WFO",
                "Wednesday": "WFO",
                "Thursday": "WFO",
                "Friday": "WFO"
            }
        }
    ]
    ```

## Conclusion

This project aims to facilitate the scheduling process for employees, ensuring fairness in WFH assignments while optimizing their working hours. The genetic algorithm provides a robust approach to solving the scheduling problem, allowing for effective management of employee workloads.

## Requirements

- Go (version 1.16 or later)
- Gin framework

## How to Run

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go run main.go`.
4. Use a tool like Postman or cURL to test the API endpoint.

