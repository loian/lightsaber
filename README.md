# lightsaber

A simple server to control adalights led strips.
Not yet mature, absolutely untested, but it seems working.

to run it `lightsaber 8877`


## installation (MacOS)
install the go toolchain (brew install go)
Build the server
`go build`

copy the binary to the user Application folder
`cp lightsaber ~/Applications/`

copy the plist to the launch agent directory
`cp loian.lightsaber.plist ~/Library/LaunchAgents/`

register the daemon
`launchctl load ~/Library/LaunchAgents/loian.lightsaber.plist`

start it
`launchctl start loian.lightsaber`


## Endpoints:
- POST /stop
- POST /start

Example of a start request body:

```json
{
  "display_index": 0,
  "selected_mode": "backlight",
  "serial": {
    "port": "/dev/tty.usbserial-110",
    "baud": 115200
  },
  "color_adjustment": {
    "darken": 0.0,
    "brighten": 0.0,
    "saturate": 0.0,
    "desaturate": 0.0,
    "hue": 0.0
  },
  "led_geometry": {
    "offset": 0,
    "right": 22,
    "top": 48,
    "left": 22,
    "bottom": 48
  },
  "screen_grabber": {
    "margin": {
      "right": 150,
      "top": 100,
      "left": 150,
      "bottom": 100
    },
    "size": {
      "width": 450,
      "height": 400
    }
  },
  "swirl": {
    "pulse_depth": 0.2,
    "pulse_speed": 0.09,
    "color_rotation_speed": 5
  },
  "vader": {
    "speed": 40
  },
  "backlight": {
    "r": 255,
    "g": 255,
    "b": 0
  },
  "custom_scene": {
    "leds": [
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 0,
        "g": 255,
        "b": 255
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      },
      {
        "r": 255,
        "g": 0,
        "b": 0
      }
    ]
  }
}

```
