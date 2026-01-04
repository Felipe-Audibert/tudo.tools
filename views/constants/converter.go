package constants

import (
	"sort"

	"github.com/tkuchiki/go-timezone"
)

type SelectOption struct {
	Label string
	Value string
}

func GetDistanceConverterOptions() []SelectOption {
	return []SelectOption{
		{
			Label: "Milímetro (mm)",
			Value: "0.001",
		},
		{
			Label: "Centímetro (cm)",
			Value: "0.01",
		},
		{
			Label: "Decímetro (dm)",
			Value: "0.1",
		},
		{
			Label: "Metro (m)",
			Value: "1",
		},
		{
			Label: "Decâmetro (dam)",
			Value: "10",
		},
		{
			Label: "Hecâmetro (hm)",
			Value: "100",
		},
		{
			Label: "Kilômetro (km)",
			Value: "1000",
		},
	}
}

func GetWeightConverterOptions() []SelectOption {
	return []SelectOption{
		{
			Label: "Miligrama (g)",
			Value: "0.001",
		},
		{
			Label: "Grama (g)",
			Value: "1",
		},
		{
			Label: "Kilograma (Kg)",
			Value: "1000",
		},
	}
}

func GetTimeConverterOptions() []SelectOption {
	tz := timezone.New()
	timezones := tz.Timezones()

	seen := map[string]struct{}{}
	options := []SelectOption{}

	for _, values := range timezones {
		for _, value := range values {
			if _, ok := seen[value]; ok {
				continue
			}

			seen[value] = struct{}{}
			options = append(options, SelectOption{
				Label: value,
				Value: value,
			})
		}
	}

	sort.Slice(options, func(i, j int) bool {
		return options[i].Label < options[j].Label
	})

	return options
}
