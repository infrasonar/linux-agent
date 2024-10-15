[![CI](https://github.com/infrasonar/linux-agent/workflows/CI/badge.svg)](https://github.com/infrasonar/linux-agent/actions)
[![Release Version](https://img.shields.io/github/release/infrasonar/linux-agent)](https://github.com/infrasonar/linux-agent/releases)

# InfraSonar Linux Agent

Documentation: https://docs.infrasonar.com/collectors/agents/linux/

## Environment variables

Environment                 | Default                       | Description
----------------------------|-------------------------------|-------------------
`STORAGE_PATH`              | `HOME/.infrasonar/`           | Path where files are stored _(not used when `ASSET_ID` is set)_.
`TOKEN`                     | _required_                    | Token used for authentication _(This MUST be a container token)_.
`ASSET_NAME`                | _none_                        | Initial Asset Name. This will only be used at the announce. Once the asset is created, `ASSET_NAME` will be ignored.
`ASSET_ID`                  | _none_                        | Asset Id _(If not given, the asset Id will be stored and loaded from file)_.
`API_URI`                   | https://api.infrasonar.com    | InfraSonar API.
`SKIP_VERIFY`				| _none_						| Set to `1` or something else to skip certificate validation.
`CHECK_SYSTEM_INTERVAL`     | `300`                         | Interval in seconds for the `system` check.


## Build
```
CGO_ENABLED=0 go build -o linux-agent
```

## Installation

Download the latest release:
```bash
$ wget https://github.com/infrasonar/linux-agent/releases/download/v0.1.0/linux-agent
```

Ensure the binary is executable:
```
chmod +x linux-agent
```

Copy the binary to `/usr/sbin/infrasonar-linux-agent`

```
$ sudo cp linux-agent /usr/sbin/infrasonar-linux-agent
```

### Using Systemd

```bash
$ sudo touch /etc/systemd/system/infrasonar-linux-agent.service
$ sudo chmod 664 /etc/systemd/system/infrasonar-linux-agent.service
```

Using you favorite editor, add the content below to the file created

```
[Unit]
Description=InfraSonar Linux Agent
Wants=network.target

[Service]
Environment="TOKEN=<YOUR TOKEN HERE>"
# Environment="ASSET_ID=<YOUR ASSET ID>"
# Environment="STORAGE_PATH=<PATH_TO_STORE_ASSET_FILE>"
ExecStart=/usr/sbin/infrasonar-linux-agent

[Install]
WantedBy=multi-user.target
```

Reload systemd

```bash
$ sudo systemctl daemon-reload
```

Install the service
```bash
$ sudo systemctl enable infrasonar-linux-agent
```

You may want to start/stop or view the status
```bash
$ sudo systemctl start infrasonar-linux-agent
$ sudo systemctl stop infrasonar-linux-agent
$ sudo systemctl status infrasonar-linux-agent
```

View logging:
```bash
$ journalctl -u infrasonar-linux-agent
```

