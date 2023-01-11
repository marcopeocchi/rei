# Valeera

A simple configurable... Dashboard?

Written in Go, extra-light on system-resources.

![home](https://i.ibb.co/Vxmc4L5/image.png)
![home](https://i.ibb.co/3SftxQz/image.png)

## Configurable

Easy configuration, just a YAML file

Create a file named `Valeerafile` under `/etc/valeera`

```sh
mkdir /etc/valeera
touch /etc/Valeerafile
# edit with your editor of choiche
```
Using this structure

```yaml
---
port: 8686
servername: Your Server Name
services:
- url: http://...
  name: Service1
- url: http://...
  name: Service2
```

Move it in `$PATH`

```sh
mv valeera /usr/bin/valeera
```

### Systemd configuration

```sh
nano /etc/systemd/system/valeera.service
```

```
[Unit]
Description=Valeera dashboard
After=network.target

[Service]
User=YOUR_USER
ExecStart=/usr/bin/valeera -c /etc/valeera/Valeerafile

[Install]
WantedBy=multi-user.target
```

```
systemctl enable valeera
systemctl start valeera
```

## Considerations

Only for **Linux** atm.

## Themable

6 Built-in themes, extendable `themes.js` file.

- Amber
- Icy
- Rosé
- Sakura
- Sky
- Spring