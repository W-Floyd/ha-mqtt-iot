# HomeAssistant MQTT IOT
## In need of a better name

A simple Golang program that allows commandline based MQTT entities.
It uses HomeAssistant MQTT autodiscovery, thinly wrapping all of the native json fields for each object, handling topic creation, subscription, updating, and so on.
See `sample.config.json` and `sample.secrets.json` for examples.
Run `./ha-mqtt-iot -help` for help.

## Automatic Code Generation

Most of this tool is automatically generated from the Home Assistant documentation.
As such, should this documentation change, the code may be regenerated using `go generate`, or `go run ./helpers/generate`.
If this has been run before, clear the cache found in `./helpers/generate/cache/`.

## Writing new configs

View the sample config for examples.
Most fields are inherited from the Home Assistant documentation, but to double check one can check the `./devices/internaldevice/*.go` files for the struct.
The program will fail on finding unknown fields.