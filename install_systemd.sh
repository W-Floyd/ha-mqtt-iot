#!/bin/bash

__project_name='ha-mqtt-iot'

__bin_file="${__project_name}"

__project_path="$(
    cd -- "$(dirname "$0")" >/dev/null 2>&1
    pwd -P
)"

__exe_path="${__project_path}/${__bin_file}"

(
    echo "[Unit]
Description=${__project_name}
After=network.target

[Service]
Type=simple
Restart=on-failure
ExecStart=${__exe_path}
WorkingDirectory=${__project_path}

[Install]
WantedBy=multi-user.target" | sudo tee "/etc/systemd/system/${__project_name}.service"
)

sudo systemctl daemon-reload
sudo systemctl enable "${__project_name}.service"
sudo systemctl start "${__project_name}.service"

exit
