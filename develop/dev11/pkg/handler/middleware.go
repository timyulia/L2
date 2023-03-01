package handler

import (
	"calendar"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func EventLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("Method: %s URI: %s Time: %s", r.Method, r.RequestURI, now)

		},
	)
}

func paramValues(values url.Values, dict ...string) (map[string]string, error) {
	found := make(map[string]string)
	missedValues := make([]string, 0)

	for _, v := range dict {
		value := values.Get(v)
		if value == "" {
			missedValues = append(missedValues, v)
		}

		found[v] = value
	}

	if len(missedValues) != 0 {
		return map[string]string{}, fmt.Errorf("not enough parameters: %s", strings.Join(missedValues, ", "))
	}

	return found, nil
}

func unmarshalEvent(r *http.Request) (calendar.Event, error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return calendar.Event{}, err
	}

	var event calendar.Event

	err = json.Unmarshal(b, &event)
	if err != nil {
		return calendar.Event{}, err
	}

	return event, validateEvent(event)
}

func validateEvent(event calendar.Event) error {
	if event.ID < 1 {
		return errors.New("invalid event ID")
	}
	if event.UserID < 1 {
		return errors.New("invalid user ID")
	}
	return nil
}
