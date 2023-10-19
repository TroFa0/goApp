package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type pair struct {
	movie string
	mTime time.Time
}

type Movie struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
	Hour  int    `json:"hour"`
	Min   int    `json:"min"`
}
type Movies struct {
	Movies []Movie `json:"movies"`
}

var movieDurability int = 3
var movies []pair
var dayCount int

func p(temp string) int {
	x, _ := strconv.Atoi(temp)
	return x
}
func init() {
	jsonFile, _ := os.Open("movies.json")
	var moviesJson Movies
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &moviesJson)
	tempTime := time.Now()
	for i := 0; i < len(moviesJson.Movies); i++ {
		t := time.Date(moviesJson.Movies[i].Year, time.Month(moviesJson.Movies[i].Month), moviesJson.Movies[i].Day, moviesJson.Movies[i].Hour,
			moviesJson.Movies[i].Min, 0, 0, time.Now().UTC().Location())
		if t.After(tempTime) {
			tempTime = t
		}
		movies = append(movies, pair{moviesJson.Movies[i].Name, t})
	}
	dayCount = tempTime.YearDay() - time.Now().YearDay() + 1
}

func main() {
	var name string
	fmt.Println("Enter username: ")
	fmt.Scan(&name)
	getMovies(name)
}
func toTime(str string) time.Time {
	time, _ := time.Parse(time.RFC3339, str)
	return time
}

func MovieByOneTime(time time.Time) {
	for i := 0; i < len(movies); i++ {
		if movies[i].mTime.Month() == time.Month() &&
			movies[i].mTime.Day() == time.Day() {
			fmt.Println(movies[i].movie, movies[i].mTime)
		}
	}
}

func MovieByTwoTimes(start time.Time, end time.Time) {
	for i := 0; i < len(movies); i++ {
		temp := start.Add(time.Hour * time.Duration(movieDurability))
		if temp.Before(end) && movies[i].mTime.Month() == start.Month() &&
			movies[i].mTime.Day() == start.Day() && movies[i].mTime.After(start) {
			temp2 := movies[i].mTime.Add(time.Hour * time.Duration(movieDurability))
			if temp2.Before(end) {
				fmt.Println(movies[i].movie, movies[i].mTime)
			}
		}
	}
}

func getMovies(user string) {
	aut(user)
	fmt.Println("Фільми на які ви можете піти:")
	temp := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Now().UTC().Location())
	for i := 0; i <= dayCount; i++ {
		events := getEventsDay(user, temp)
		if len(events) == 0 {
			MovieByOneTime(temp)
			temp = temp.AddDate(0, 0, 1)
			continue
		} else {
			for j := 0; j < len(events); j++ {
				start := events[j].Start.DateTime
				end := events[j].End.DateTime
				var tStart time.Time
				var tEnd time.Time
				if start == "" {
					continue
				}
				if j == 0 {
					tStart = temp
					tEnd = toTime(start)
					MovieByTwoTimes(tStart, tEnd)
				}
				if j == len(events)-1 {
					event, ex := getFirstEvent(user, temp)
					var temp2 time.Time

					if ex {
						if event.Start.DateTime != "" {
							temp2 = toTime(event.Start.DateTime)
						} else {
							temp2 = temp
							temp2 = temp2.Add(time.Hour * 28)
						}
					} else {
						temp2 = temp
						temp2 = temp2.Add(time.Hour * 28)
					}

					tStart = toTime(end)
					MovieByTwoTimes(tStart, temp2)
					continue
				}
				startNext := events[j+1].Start.DateTime
				MovieByTwoTimes(toTime(end), toTime(startNext))
			}
		}
		temp = temp.AddDate(0, 0, 1)
	}
}
