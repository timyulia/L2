package handler

import (
	"calendar"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const dateLayout = "2006-01-02"

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	event, err := unmarshalEvent(r)
	if err != nil {
		throwError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal event: %s", err))
		return
	}

	err = h.repos.CreateEvent(event)
	if err != nil {
		throwError(w, http.StatusServiceUnavailable, fmt.Errorf("could not insert this event: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "event created", []calendar.Event{event})
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	event, err := unmarshalEvent(r)
	if err != nil {
		throwError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal event: %s", err))
		return
	}
	err = h.repos.UpdateEvent(event)
	if err != nil {
		throwError(w, http.StatusServiceUnavailable, fmt.Errorf("could not update this event: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "event updated", []calendar.Event{event})
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	event, err := unmarshalEvent(r)
	if err != nil {
		throwError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal event: %s", err))
		return
	}
	err = h.repos.DeleteEvent(event.ID)
	if err != nil {
		throwError(w, http.StatusServiceUnavailable, fmt.Errorf("could not delete this event: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "event deleted", []calendar.Event{event})
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	userID, date := dateHandler(w, r)
	if userID == 0 {
		return
	}
	events, err := h.repos.ForMonth(userID, date)
	if err != nil {
		throwError(w, http.StatusServiceUnavailable, fmt.Errorf("could not get events for a day: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "events:", events)
}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	userID, date := dateHandler(w, r)
	if userID == 0 {
		return
	}
	events, err := h.repos.ForWeek(userID, date)
	if err != nil {
		throwError(w, http.StatusServiceUnavailable, fmt.Errorf("could not get events for a week: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "events:", events)
}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	userID, date := dateHandler(w, r)
	if userID == 0 {
		return
	}
	events, err := h.repos.ForMonth(userID, date)
	if err != nil {
		throwError(w, http.StatusServiceUnavailable, fmt.Errorf("could not get events for a month: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "events:", events)
}

func dateHandler(w http.ResponseWriter, r *http.Request) (int, time.Time) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return 0, time.Time{}
	}

	required := []string{"user_id", "date"}

	v := r.URL.Query()
	values, err := paramValues(v, required...)
	if err != nil {
		throwError(w, http.StatusBadRequest, err)
		return 0, time.Time{}
	}

	date, err := time.Parse(dateLayout, values["date"])
	if err != nil {
		throwError(
			w, http.StatusBadRequest, fmt.Errorf("invalid date format: %s, expected %s", values["date"], dateLayout),
		)
		return 0, time.Time{}
	}

	userID, err := strconv.Atoi(values["user_id"])
	if err != nil || userID < 1 {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %s", values["user_id"]))
		return 0, time.Time{}
	}
	return userID, date
}
