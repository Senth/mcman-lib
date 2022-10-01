package status

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatuses_Append(t *testing.T) {
	testCases := []struct {
		name      string
		existing  Statuses
		appending Statuses
		expected  Statuses
	}{
		{
			name:      "Return early when existing and appending are nil",
			existing:  nil,
			appending: nil,
			expected:  nil,
		},
		{
			name:      "Return existing when appending is nil",
			existing:  Statuses{Status{Code: CodeServerError}},
			appending: nil,
			expected:  Statuses{Status{Code: CodeServerError}},
		},
		{
			name:      "Return appending when existing is nil",
			existing:  nil,
			appending: Statuses{Status{Code: CodeServerError}},
			expected:  Statuses{Status{Code: CodeServerError}},
		},
		{
			name:      "Combine existing and appending",
			existing:  Statuses{Status{Code: CodeModrinthError}},
			appending: Statuses{Status{Code: CodeCurseError}},
			expected:  Statuses{Status{Code: CodeModrinthError}, Status{Code: CodeCurseError}},
		},
		{
			name:      "Combine multiple of the same into one",
			existing:  Statuses{Status{Code: CodeModrinthError}},
			appending: Statuses{Status{Code: CodeModrinthError}},
			expected:  Statuses{Status{Code: CodeModrinthError}},
		},
		{
			name:      "Combine multiple of the same into one use the data from existing where it is not nil",
			existing:  Statuses{Status{Code: CodeModrinthError, Data: "existing"}},
			appending: Statuses{Status{Code: CodeModrinthError}},
			expected:  Statuses{Status{Code: CodeModrinthError, Data: "existing"}},
		},
		{
			name:      "Combine multiple of the same into one use the data from appending where it is not nil",
			existing:  Statuses{Status{Code: CodeModrinthError}},
			appending: Statuses{Status{Code: CodeModrinthError, Data: "appending"}},
			expected:  Statuses{Status{Code: CodeModrinthError, Data: "appending"}},
		},
		{
			name:      "Combine multiple of the same into one use the data from appending if both are not nil",
			existing:  Statuses{Status{Code: CodeModrinthError, Data: "existing"}},
			appending: Statuses{Status{Code: CodeModrinthError, Data: "appending"}},
			expected:  Statuses{Status{Code: CodeModrinthError, Data: "appending"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.existing.Append(tc.appending)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}
