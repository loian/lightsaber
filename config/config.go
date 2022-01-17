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

type Configuration struct {
	DisplayIndex    int             `json:"display_index"`
	SelectedMode    string          `json:"selected_mode"`
	Serial          Serial          `json:"serial"`
	ColorAdjustment ColorAdjustment `json:"color_adjustment"`
	LedGeometry     LedGeometry     `json:"led_geometry"`
	ScreenGrabber   ScreenGrabber   `json:"screen_grabber"`
}
