package constants

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
