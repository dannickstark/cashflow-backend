package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2"
)

// A function to save JSON data into a file
func SaveFile(data string, filename string) error {
	// Write the JSON data to the file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	l, err := f.WriteString(data)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// inspectRuntime tries to find the base executable directory and how it was run.
func inspectRuntime() (baseDir string, withGoRun bool) {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// probably ran with go run
		withGoRun = true
		baseDir, _ = os.Getwd()
	} else {
		// probably ran with go build
		withGoRun = false
		baseDir = filepath.Dir(os.Args[0])
	}
	return
}

func GetAbsolutePath(path string) string {
	baseDir, _ := inspectRuntime()
	return filepath.Join(baseDir, path)
}

// -------- Lists
func FilterList[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

// ----------[ Date ]
func CheckIfCorrecPace(date string, pace int, paceUnit PaceUnit) bool {
	// Set location
	utc, _ := time.LoadLocation(carbon.UTC)
	c := carbon.SetLocation(utc)
	t, _ := time.Parse(DefaultDateLayout, date)
	ct := carbon.CreateFromStdTime(t)

	switch paceUnit {
	case Day:
		return ct.AddDays(pace).StartOfDay().Eq(c.Now().StartOfDay())

	case Week:
		return ct.AddWeeks(pace).StartOfWeek().Eq(c.Now().StartOfWeek())

	case Month:
		return ct.AddMonths(pace).StartOfMonth().Eq(c.Now().StartOfMonth())

	case Year:
		return ct.AddYears(pace).StartOfYear().Eq(c.Now().StartOfYear())
	}

	return false
}

func ComputeNextDate(date string, pace int, paceUnit PaceUnit) string {
	t, _ := time.Parse(DefaultDateLayout, date)
	ct := carbon.CreateFromStdTime(t)

	switch paceUnit {
	case Day:
		return ct.AddDays(pace).ToDateTimeString()

	case Week:
		return ct.AddWeeks(pace).ToDateTimeString()

	case Month:
		return ct.AddMonths(pace).ToDateTimeString()

	case Year:
		return ct.AddYears(pace).ToDateTimeString()
	}

	return ""
}
