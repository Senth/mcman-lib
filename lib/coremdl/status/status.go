// Package status various server status code containing more information why
// the request failed.
package status

import "time"

// Status used for indicating various server errors to the frontend
type Status struct {
	Code  Code        `json:"code"`
	Trace string      `json:"trace"`
	Data  interface{} `json:"data"`
}

// Statuses is a list of Status
type Statuses []Status

// Code used for indicating various server errors to the frontend
type Code string

// Error Status codes
const (
	CodeServerError       Code = "server_error"
	CodeModrinthError     Code = "modrinth_error"
	CodeModrinthDown      Code = "modrinth_down"
	CodeModrinthRateLimit Code = "modrinth_rate_limit"
	CodeModrinthNotFound  Code = "modrinth_not_found"
	CodeCurseError        Code = "curse_error"
	CodeCurseDown         Code = "curse_down"
	CodeCurseRateLimit    Code = "curse_rate_limit"
	CodeCurseNotFound     Code = "curse_not_found"
)

func NewServerErrors() Statuses {
	return Statuses{{Code: CodeServerError}}
}

// RateLimit additional data for
type RateLimit struct {
	Reset time.Time `json:"reset"`
}

// Append add statuses to the existing statuses
// If both existing and appending statuses are nil, this function returns nil
func (s Statuses) Append(statuses Statuses) Statuses {
	// Return nil if there are no statuses
	if statuses == nil {
		if s == nil {
			return nil
		}

		return s
	}

	newStatuses := append(s, statuses...)

	// Remove duplicates
	m := make(map[Code]Status)
	for _, status := range newStatuses {
		existing, ok := m[status.Code]
		if ok {
			// Always replace existing data with newer data
			if status.Data != nil {
				existing.Data = status.Data
				m[status.Code] = existing
			}
		} else {
			m[status.Code] = status
		}
	}

	// Convert map to slice
	newStatuses = make(Statuses, 0, len(m))
	for _, status := range m {
		newStatuses = append(newStatuses, status)
	}

	return newStatuses
}
