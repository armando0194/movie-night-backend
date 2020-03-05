package cron

import "github.com/robfig/cron"

func New() {
	// Add new movie night every wednesday at night
	cron.New().AddFunc("* 23 * * 3", AddMovieNight)
}

func AddMovieNight() {

}
