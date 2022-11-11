# HomeAssistant MQTT IOT
## In need of a better name

A simple Golang program that allows commandline based MQTT entities.
It uses HomeAssistant MQTT autodiscovery, thinly wrapping all of the native json fields for each object, handling topic creation, subscription, updating, and so on.
See `sample.config.json` and `sample.secrets.json` for examples.
Run `./ha-mqtt-iot -help` for help.

## Automatic Code Generation

Most of this tool is automatically generated from the Home Assistant documentation.
As such, should this documentation change, the code may be regenerated using `go generate`, or `go run ./helpers/`.
If this has been run before, clear the cache found in `./helpers/cache/`.

## Writing new configs

View the sample config for examples.
Most fields are inherited from the Home Assistant documentation, but to double check one can check the `./devices/internaldevice/*.go` files for the struct.
The program will fail on finding unknown fields.

## Install

### Dependencies 

* python3
* go

### Install

Use the famous linux 3-Step-Evergreen (TM)

This installes ha-mqtt-iot to your home directory according to the XDG_BASE_DIRECTORY standard and creates a systemd user service

```
./configure
make 
make install
```

INFO: A small help is displayed for the manual steps you have to do after the installation.

### Update

If you have ha-mqtt-iot installed before und you require and update use the following commands

This will create a new build, reinstalles ha-mqtt-iot <i>without</i> touching the secrets.json in your home

```
git pull .... 
./configure
make
make reinstall
```

### Uninstall 

This is sad. You like to remove ha-mqtt-iot:

```
make uninstall
```

### Helpers for development

#### Create a debug build

This produces a debug build in the ./build directory. Ajust the secrets.json to your needs and start the debug build inside the directory. 

```
./configure 
make
make debug

cd build
vim secrets.json 
[ ... ]

./ha-mqtt-iot
```

#### Only create the configuration 

This creates the configuration only 

```
./configure
make config
```

#### Clean up code direktory

Removes build directory, removes Makefile, removes ha-mqtt-iot binary

```
make clean
```

## Tested on:

- Arch Linux (amd64)
- Ubuntu Linux 22.10 (amd64)
- Raspbian Buster (arm7l)





