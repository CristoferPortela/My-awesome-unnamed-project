package model

type ModelMenu struct {
	DisplayName string
	Model       interface{}
}

func RegisterModel() ([]ModelMenu, error) {

	models := []ModelMenu{
		{DisplayName: "Usuarios", Model: User{}},
		{DisplayName: "Configurações", Model: Setting{}},
	}

	return models, nil
}
