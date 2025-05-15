package utils

import (
    "regexp"
    "time"
)

func IsValidEmail(email string) bool {
    re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    return re.MatchString(email)
}

func IsValidDeadline(deadline time.Time) bool {
    return deadline.After(time.Now())
}
