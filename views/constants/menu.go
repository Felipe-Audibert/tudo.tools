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
			Title: "Convertedores",
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
			},
		},
		{
			Title: "Test2",
			Children: &[]Menu{
				{
					Url:   "/test3",
					Title: "Test3",
				},
			},
		},
		{
			Url:   "/test3",
			Title: "Teste",
		},
	}
}
