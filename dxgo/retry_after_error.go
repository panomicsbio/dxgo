package dxgo

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type RetryAfterError struct {
	response http.Response
}

func (err RetryAfterError) Error() string {
	return fmt.Sprintf(
		"Request to %s fail %s (%d)",
		err.response.Request.RequestURI,
		err.response.Status,
		err.response.StatusCode,
	)
}

func ParseRetryAfter(retryAfter string) (time.Time, error) {
	if dur, err := ParseSeconds(retryAfter); err == nil {
		now := time.Now()
		return now.Add(dur), nil
	}
	if dt, err := ParseHTTPDate(retryAfter); err == nil {
		return dt, nil
	}
	return time.Time{}, errors.New("Retry-After value must be seconds integer or HTTP date string")
}

func ParseSeconds(retryAfter string) (time.Duration, error) {
	seconds, err := strconv.ParseInt(retryAfter, 10, 64)
	if err != nil {
		return time.Duration(0), err
	}
	if seconds < 0 {
		return time.Duration(0), errors.New("negative seconds not allowed")
	}
	return time.Second * time.Duration(seconds), nil
}

func ParseHTTPDate(retryAfter string) (time.Time, error) {
	parsed, err := time.Parse(time.RFC1123, retryAfter)
	if err != nil {
		return time.Time{}, err
	}
	return parsed, nil
}
