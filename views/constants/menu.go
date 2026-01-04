package constants

type Menu struct {
	Url      string
	Title    string
	Icon     string
	Children *[]Menu
}

func GetMenus() []Menu {
	return []Menu{
		{
			Icon:  "fa-solid fa-arrow-right-arrow-left",
			Title: "Converter",
			Children: &[]Menu{
				{
					Url:   "/",
					Title: "Tamanhos/Distâncias",
					Children: &[]Menu{
						{
							Url:   "/Metros",
							Title: "Metros",
						},
						{
							Url:   "/test3",
							Title: "Test3",
						},
					},
				},
				{
					Url:   "/convert/distance",
					Title: "Distância",
					Icon:  "fa-solid fa-ruler-combined",
				},
				{
					Url:   "/convert/time",
					Title: "Fuso horário",
					Icon:  "fa-regular fa-clock",
				},
				{
					Url:   "/convert/weight",
					Title: "Peso",
					Icon:  "fa-solid fa-balance-scale"},
			},
		},
	}
}
