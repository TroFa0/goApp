package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type pair struct {
	movie string
	mTime time.Time
}

var movieDurability int = 3
var movies []pair
var dayCount int

func p(temp string) int {
	x, _ := strconv.Atoi(temp)
	return x
}
func init() {
	file, _ := os.Open("movieSchedule.txt")
	tempTime := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		arr := strings.Split(temp, " ")
		lenA := len(arr)

		t := time.Date(p(arr[lenA-5]), time.Month(p(arr[lenA-4])), p(arr[lenA-3]), p(arr[lenA-2]), p(arr[lenA-1]), 0, 0, time.Now().UTC().Location())
		if t.After(tempTime) {
			tempTime = t
		}

		lenB := lenA - 5
		var name string = ""
		for i := 0; i < lenB; i++ {
			name += arr[i]
			name += " "
		}
		movies = append(movies, pair{name, t})
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
