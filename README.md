# dash-hooker

Amazon Dash Button's arp request hooker.

## Usage

```
dash-hooker
```

## Options

```
# dash-hooker --help
Usage of dash-hooker:
  -v    verbose output
```

## Requirements

[google/gopacket](https://github.com/google/gopacket) uses pcap.

```
yum install -y libpcap-devel
```

This program requires **Root Permission**.

## Installation

You can use `go get`.

```
go get -u github.com/zuiurs/dash-hooker
```

## Preparing

- var.go
  - `device`: define your network device name
  - `mac`: define your Dash Button MAC
- routine.go
  - write your code

## Search your Dash Button MAC Address

Use `tcpdump`.

```
sudo tcpdump -ne -i <network_interface> arp
```

and search the OUI.

|Amazon OUI|
|:-:|
|00:FC:8B|
|0C:47:C9|
|18:74:2E|
|34:D2:70|
|40:B4:CD|
|44:65:0D|
|50:F5:DA|
|68:37:E9 (My Dush Button MAC)|
|68:54:FD|
|74:75:48|
|74:C2:46|
|78:E1:03|
|84:D6:D0|
|88:71:E5|
|A0:02:DC|
|AC:63:BE|
|B4:7C:9C|
|F0:27:2D|
|F0:D2:F1|
|FC:65:DE|
|FC:A6:67|

(reference of https://mac.uic.jp)

## Systemd Setting

Edit dash-hooker.service and replace `ExecStart` to your dash-hooker executalbe path.

```
vim dash-hooker.service
```

Install and Start.

```
sudo cp -i dash-hooker.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable dash-hooker.service
sudo systemctl start dash-hooker.service
```

## TODO

- make tool of searching dash button MAC

## License

This software is released under the MIT License, see LICENSE.txt.
