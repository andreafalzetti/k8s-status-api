package pods

import (
	"errors"
	"sort"
	"strings"
	"time"
)

func sortByAge(items []Pod, reverse bool) []Pod {
	sort.Slice(items, func(i, j int) bool {
		timeI, _ := time.ParseDuration(items[i].Age)
		timeJ, _ := time.ParseDuration(items[j].Age)
		if reverse {
			return timeI > timeJ
		} else {
			return timeI < timeJ
		}
	})
	return items
}

func sortByRestarts(items []Pod, reverse bool) []Pod {
	sort.Slice(items, func(i, j int) bool {
		if reverse {
			return items[i].Restarts > items[j].Restarts
		} else {
			return items[i].Restarts < items[j].Restarts
		}
	})
	return items
}

func sortByName(items []Pod, reverse bool) []Pod {
	sort.Slice(items, func(i, j int) bool {
		if reverse {
			return items[i].Name > items[j].Name
		} else {
			return items[i].Name < items[j].Name
		}
	})
	return items
}

// TODO: extend to support multiple sort parameters
func SortPodsByParam(items []Pod, sortParam string) ([]Pod, error) {
	sortFunctions := map[string]func([]Pod, bool) []Pod{
		"age":      sortByAge,
		"restarts": sortByRestarts,
		"name":     sortByName,
	}

	reverse := false
	if strings.HasPrefix(sortParam, "-") {
		sortParam = sortParam[1:]
		reverse = true
	}

	if sortFunc, ok := sortFunctions[sortParam]; ok {
		return sortFunc(items, reverse), nil
	}

	return nil, errors.New("Invalid sort parameter")
}
