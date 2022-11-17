package models

type Light struct {
	On bool
}

func NewLight(on bool) Light {
	return Light{on}
}

func (light *Light) IsOn() bool {
	return light.On
}

func (light *Light) TurnOn() *Light {
	light.On = true
	return light
}

func (light *Light) TurnOff() *Light {
	light.On = false
	return light
}

func (light *Light) Toggle() *Light {
	light.On = !light.On
	return light
}
