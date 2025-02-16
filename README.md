# inkube
A simple Kubernetes monitoring tool using a Pi Zero W, an eInk screen &amp; Golang

## Usage

Create a config file:

```toml
[cluster]
server="<api_endpoint>"
certificate_authority_data="<base64_encoded_ca_data>"

[auth]
client_certificate_data="<base64_encoded_certificate>"
client_key_data="<base64_encoded_key>"

[namespace]
default="default"

[display]
refresh="60"
```

Start the application (with the config file in the current working directory):
```shell
./inkube
```
