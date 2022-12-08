# port-scan
Scan for a port on local Network with golang

Using this tool you can scan a specific port on your internal network.
This tool works like this: it automatically detects the IP of your device in use, then it scans the range of /24 according to the IP of your device.

## Running 

As an example, The default port scanned by this tool is 25565:

```bash
./port-scan
```

You can also use this tool to expose custom port. For example, if you want scan port 22, all you have to do is:

```bash
./port-scan -p 22
```

## Last words

<img src="https://raw.githubusercontent.com/0x187/ClearText/main/68747470733a2f2f692e696d6775722e636f6d2f774d34553835682e6a7067.jpg">
