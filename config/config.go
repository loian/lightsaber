package config

var DiscoveredPort = "";

type Serial struct {
	//Port string `json:"port" binding:"required"`
	Baud int    `json:"baud" binding:"required,oneof=9600 19200 38400 57600 115200"`
}

type LedGeometry struct {
	Offset *int `json:"offset" binding:"required,numeric"`
	Right  *int `json:"right" binding:"required,numeric"`
	Top    *int `json:"top" binding:"required,numeric"`
	Left   *int `json:"left" binding:"required,numeric"`
	Bottom *int `json:"bottom" binding:"required,numeric"`
}

type ColorAdjustment struct {
	Darken     *float64 `json:"darken" binding:"required,numeric"`
	Brighten   *float64 `json:"brighten" binding:"required,numeric"`
	Saturate   *float64 `json:"saturate" binding:"required,numeric"`
	Desaturate *float64 `json:"desaturate" binding:"required,numeric"`
	Hue        *float64 `json:"hue" binding:"required,numeric"`
}

type Margins struct {
	Right  *int `json:"right" binding:"required,numeric"`
	Top    *int `json:"top" binding:"required,numeric"`
	Left   *int `json:"left" binding:"required,numeric"`
	Bottom *int `json:"bottom" binding:"required,numeric"`
}

type Size struct {
	Width  *int `json:"width" binding:"required,numeric"`
	Height *int `json:"height" binding:"required,numeric"`
}

type ScreenGrabber struct {
	Size   Size    `json:"size" binding:"required,dive"`
	Margin Margins `json:"margin" binding:"required,dive"`
}

type Swirl struct {
	PulseDepth         *float64 `json:"pulse_depth" binding:"required,numeric"`          //0 - 0.5
	PulseSpeed         *float64 `json:"pulse_speed" binding:"required,numeric"`          //0.03 0.06 0.09 0.12 0.15 0.18 0.21 0.24 0.27 0.3
	ColorRotationSpeed *int     `json:"color_rotation_speed" binding:"required,numeric"` //1.2.3.5.6.7.8.9.10
}

type Vader struct {
	Speed *int `json:"speed" binding:"required,numeric"` //1...100
}

type RGB struct {
	R byte `json:"r"`
	G byte `json:"g"`
	B byte `json:"b"`
}

type Backlight RGB

type Custom struct {
	Leds []RGB `json:"leds" binding:"required"`
}

type Configuration struct {
	DisplayIndex    *int            `json:"display_index" binding:"required,numeric"`
	SelectedMode    *string         `json:"selected_mode" binding:"required"`
	Serial          Serial          `json:"serial" binding:"required,dive"`
	ColorAdjustment ColorAdjustment `json:"color_adjustment" binding:"required"`
	LedGeometry     LedGeometry     `json:"led_geometry" binding:"required"`
	ScreenGrabber   ScreenGrabber   `json:"screen_grabber" binding:"required,dive"`
	Swirl           Swirl           `json:"swirl" binding:"required"`
	Vader           Vader           `json:"vader" binding:"required"`
	Backlight       Backlight       `json:"backlight" binding:"required"`
	Custom          Custom          `json:"custom_scene" binding:"required"`
}
