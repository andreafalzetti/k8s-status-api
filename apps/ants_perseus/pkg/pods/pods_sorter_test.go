package pods

import (
	"errors"
	"reflect"
	"testing"
)

func TestSortPodsByParam(t *testing.T) {
	testCases := []struct {
		name           string
		pods           []Pod
		sortParam      string
		expectedResult []Pod
		expectedError  error
	}{
		{
			name: "Invalid sort parameter",
			pods: []Pod{
				{Name: "pod3", Age: "3d", Restarts: 1},
				{Name: "pod1", Age: "1d", Restarts: 2},
				{Name: "pod2", Age: "2d", Restarts: 3},
			},
			sortParam:      "some-invalid-param",
			expectedResult: nil,
			expectedError:  errors.New("Invalid sort parameter"),
		},
		{
			name: "Sort by name ascending",
			pods: []Pod{
				{Name: "pod-c", Age: "3d", Restarts: 1},
				{Name: "pod-a", Age: "1d", Restarts: 2},
				{Name: "pod-b", Age: "2d", Restarts: 3},
			},
			sortParam: "name",
			expectedResult: []Pod{
				{Name: "pod-a", Age: "1d", Restarts: 2},
				{Name: "pod-b", Age: "2d", Restarts: 3},
				{Name: "pod-c", Age: "3d", Restarts: 1},
			},
			expectedError: nil,
		},
		{
			name: "Sort by name descending",
			pods: []Pod{
				{Name: "pod-c", Age: "3d", Restarts: 1},
				{Name: "pod-a", Age: "1d", Restarts: 2},
				{Name: "pod-b", Age: "2d", Restarts: 3},
				{Name: "pod-d", Age: "16h", Restarts: 10},
			},
			sortParam: "-name",
			expectedResult: []Pod{
				{Name: "pod-d", Age: "16h", Restarts: 10},
				{Name: "pod-c", Age: "3d", Restarts: 1},
				{Name: "pod-b", Age: "2d", Restarts: 3},
				{Name: "pod-a", Age: "1d", Restarts: 2},
			},
			expectedError: nil,
		},
		{
			name: "Sort by age ascending",
			pods: []Pod{
				{Name: "pod-c", Age: "3h30m15s", Restarts: 0},
				{Name: "pod-a", Age: "1h10m5s", Restarts: 0},
				{Name: "pod-b", Age: "2h20m10s", Restarts: 0},
			},
			sortParam: "age",
			expectedResult: []Pod{
				{Name: "pod-a", Age: "1h10m5s", Restarts: 0},
				{Name: "pod-b", Age: "2h20m10s", Restarts: 0},
				{Name: "pod-c", Age: "3h30m15s", Restarts: 0},
			},
			expectedError: nil,
		},
		{
			name: "Sort by age descending",
			pods: []Pod{
				{Name: "pod3", Age: "3h30m15s", Restarts: 1},
				{Name: "pod1", Age: "1h10m5s", Restarts: 2},
				{Name: "pod2", Age: "2h20m10s", Restarts: 3},
			},
			sortParam: "-age",
			expectedResult: []Pod{
				{Name: "pod3", Age: "3h30m15s", Restarts: 1},
				{Name: "pod2", Age: "2h20m10s", Restarts: 3},
				{Name: "pod1", Age: "1h10m5s", Restarts: 2},
			},
			expectedError: nil,
		},
		{
			name: "Sort by restarts ascending",
			pods: []Pod{
				{Name: "pod3", Age: "3d", Restarts: 1},
				{Name: "pod1", Age: "1d", Restarts: 2},
				{Name: "pod2", Age: "2d", Restarts: 3},
			},
			sortParam: "restarts",
			expectedResult: []Pod{
				{Name: "pod3", Age: "3d", Restarts: 1},
				{Name: "pod1", Age: "1d", Restarts: 2},
				{Name: "pod2", Age: "2d", Restarts: 3},
			},
			expectedError: nil,
		},
		{
			name: "Sort by restarts descending",
			pods: []Pod{
				{Name: "pod3", Age: "3d", Restarts: 1},
				{Name: "pod1", Age: "1d", Restarts: 2},
				{Name: "pod2", Age: "2d", Restarts: 3},
			},
			sortParam: "-restarts",
			expectedResult: []Pod{
				{Name: "pod2", Age: "2d", Restarts: 3},
				{Name: "pod1", Age: "1d", Restarts: 2},
				{Name: "pod3", Age: "3d", Restarts: 1},
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := SortPodsByParam(tc.pods, tc.sortParam)
			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("Expected result %v, but got %v", tc.expectedResult, result)
			}

			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error %v, but got %v", tc.expectedError, err)
			}
		})
	}
}
