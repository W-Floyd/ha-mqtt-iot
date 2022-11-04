# Script documentation 

Summery about the used scripts ha-mqtt-iot can adress


## bedtime.sh

Switches the bedtime gnome extension on or off. 

Call: 
```
bedtime.sh command ON
bedtime.sh command OFF
bedtime.sh command-state 
```

Dependency: 
```
gnome extension "bedtime-mode"
```

## camera-sensor.sh
Checks if any webcam (/dev/video*) is currently in use on the endpoint by using fuser;

Call: 
```
camera-sensor.sh get
```

Dependency: 
```
fuser
```

## camera.sh

Creates a shot with the webcam. The screenshot is returned as base64 encoded picture.

Call: 
```
camera.sh
```

Dependency: 
```
mplayer
```
## ip.sh
Returns the IP address of given interface

Call:
```
ip.sh wlp64s0$
```

Dependency: 
```
jq
```

## laptop-backlight.sh


Call: 
```
laptop-backlight.sh
```

Dependency: 
```

```

## micro.sh;
Gets and sets the DEFAULT microphone volume and toggle on and off by using pactl (pulseaudio)

Call: 
```
micro.sh set-state <1 or 0> #
micro.sh get-state #
micro.sh set-sensitivity <volume 0-100> #
micro.sh get-sensitivity $
```
Dependency:
```
pulsaudio / pactl;
```

## monitor.sh


Call: 
```
monitor.sh
```

Dependency: 
```

```
## run-in-user-session.sh


Call: 
```
run-in-user-session.sh
```

Dependency: 
```

```


## screenlock.sh

Locks the screen of the system by using loginctl. State is reported by dbus request. At the moment this is limited to GNOME and KDE

Call: 
```
screenlock.sh command lock
screenlock.sh command unlock
screenlock.sh state
```

Dependency: 
```
loginctl
gnome desktop
and/or
kde desktop
```

## screen.sh

Creates a screenshot of the current screen. This is returned as base64 encoded string.

Call: 
```
screen.sh
```

Dependency: 
```
gnome-screenshot
```

## speaker.sh

Gets and sets the DEFAULT microphone volume and toggle on and off by using pactl (pulseaudio)

Call: 
```
speaker.sh set-state <1 or 0> #
speaker.sh get-state #
speaker.sh set-volume <volume 0-100> #
speaker.sh get-volume $
```
Dependency:
```
pulsaudio / pactl;
```

## system.sh

Get different systemstats 

Call: 
```
system.sh get-cpu             # returns cpu usage in %
system.sh get-ram             # returns ram usage in %
system.sh get-battery-status  # returns battery status in %
system.sh get-charging-status # returns if battery is charging (ON) or not (OFF)
system.sh get-board <pi>      # returns the mainboard identifier (only pi at the moment)
```

Dependency: 
```
awk
mpstat
```

## systemd-service.sh

Checks the status of a systemd-service (freely definable)

Call: 
```
systemd-service.sh <your-systemd-service>

e.g.: systemd-service.sh sshd  # returns if sshd is running
      systemd-service.sh nginx # returns if nginx is running
```

Dependency: 
```
systemd
```

## temperature.sh

Returns cpu temprature (thermal_zone0) in Celsius

Call: 
```
temperature.sh get-cpu-temp
```

Dependency: 
```
python3
```

## theme-nightthemeswitcher.sh


Call: 
```
theme-nightthemeswitcher.sh
```

Dependency: 
```

```
## theme.sh


Call: 
```
theme.sh
```

Dependency: 
```

```
## wifi.sh

Returns signal quality and signal strength (in dbm) of an active wifi interface. If the interface is not active it returns down

Call: 
```
wifi.sh get-signal
wifi.sh get-quality
```

Dependency: 
```
iwconfig
iw
```
## window.sh

Returns the active window 

Call: 
```
window.sh
```

Dependency: 
```
xdotool 
```

**WARNING**: xdotool depends on X11 - encounter problems with wayland