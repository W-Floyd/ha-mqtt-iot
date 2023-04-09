# Config Layout

```json
{
    "default": {
        "node_id": "name",
        "device_name": "Name"
    },
    "mqtt": [
        {
            "name": "Broker 1",
            "broker": "tcp://example.com:1883",
            "username": "foo",
            "password": "bar",

        }
    ]

}
```

See [schema](config.schema.json), and test UI in [UI Schema](https://ui-schema.bemit.codes/examples/Main-Demo)