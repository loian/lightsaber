package config

type Serial struct {
	Port string `json:"port"`
	Baud int    `json:"baud"`
}

type LedGeometry struct {
	Offset int `json:"offset"`
	Right  int `json:"right"`
	Top    int `json:"top"`
	Left   int `json:"left"`
	Bottom int `json:"bottom"`
}

type ColorAdjustment struct {
	DarkenPercentage float64 `json:"darken_percentage"`
}

type Margins struct {
	Right  int `json:"right"`
	Top    int `json:"top"`
	Left   int `json:"left"`
	Bottom int `json:"bottom"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type ScreenGrabber struct {
	Size   Size    `json:"size"`
	Margin Margins `json:"margin"`
}

type Swirl struct {
	PulseDepth         float64 //0 - 0.5
	PulseSpeed         float64 //0.03 0.06 0.09 0.12 0.15 0.18 0.21 0.24 0.27 0.3
	ColorRotationSpeed int     //1.2.3.5.6.7.8.9.10
}

type Vader struct {
	Speed int //1...100
}

type RGB struct {
	R byte `json:"r"`
	G byte `json:"g"`
	B byte `json:"b"`
}

type Backlight RGB

type Custom struct {
	Leds []RGB
}

type Configuration struct {
	DisplayIndex    int             `json:"display_index"`
	SelectedMode    string          `json:"selected_mode"`
	Serial          Serial          `json:"serial"`
	ColorAdjustment ColorAdjustment `json:"color_adjustment"`
	LedGeometry     LedGeometry     `json:"led_geometry"`
	ScreenGrabber   ScreenGrabber   `json:"screen_grabber"`
	Swirl           Swirl           `json:"swirl"`
	Vader           Vader           `json:"vader"`
	Backlight       Backlight       `json:"backlight"`
	Custom          Custom          `json:"custom"`
}
