# APT Proxy

> Lightweight **APT CHACHE PROXY** with only 2MB+ size!

<img src="screenshot/dockerhub.png" width="600"/>

APT Proxy is a Lightweight and Reliable APT packages (Ubuntu / Debian) cache tool, supports a large number of common system and Docker usage.

You can safely use it instead of [apt-cacher-ng](https://www.unix-ag.uni-kl.de/~bloch/acng/).

## Supported Systems and Architectures

- Linux: x86_64 / x86_32
- ARM: ARM64v8 / ARM32v6 / ARM32v7
- macOS: x86_64 / M1 ARM64v8

## Usage

Just run it:

```bash
./apt-proxy

2022/06/12 16:15:40 running apt-proxy
2022/06/12 16:15:41 Start benchmarking mirrors
2022/06/12 16:15:41 Finished benchmarking mirrors
2022/06/12 16:15:41 using fastest mirror https://mirrors.company.ltd/ubuntu/
2022/06/12 16:15:41 proxy listening on 0.0.0.0:3142
```

An APT proxy software with a cache function is started.

Then rewrite the command where you need to execute the `apt-get` command and execute it:

```bash
# `apt-get update` with apt-proxy service
http_proxy=http://your-domain-or-ip-address:3142 apt-get -o pkgProblemResolver=true -o Acquire::http=true update 
# `apt-get install vim -y` with apt-proxy service
http_proxy=http://your-domain-or-ip-address:3142 apt-get -o pkgProblemResolver=true -o Acquire::http=true install vim -y
```

When we need to execute the above commands repeatedly in batches, the speed of update and installation **will be greatly improved**.

## Docker

Just one command:

```bash
docker run -d --name=apt-proxy -p 3142:3142 soulteary/apt-proxy
```

## Options

View configuration items:

```bash
./apt-proxy -h

Usage of apt-proxy:
  -cachedir string
    	the dir to store cache data in (default "./.aptcache")
  -debug
    	whether to output debugging logging
  -host string
    	the host to bind to (default "0.0.0.0")
  -mirror string
    	the mirror for fetching packages
  -port string
    	the port to bind to (default "3142")
  -type string
    	select the type of system to cache: ubuntu/debian (default "ubuntu")
```

## [WIP] Development

Start the application in development mode:

```bash
go run apt-proxy.go
```

### Run Test And Get Coverage

```bash
# go test -cover ./...

?   	github.com/soulteary/apt-proxy	[no test files]
ok  	github.com/soulteary/apt-proxy/cli	2.492s	coverage: 68.4% of statements
ok  	github.com/soulteary/apt-proxy/linux	7.420s	coverage: 76.7% of statements
ok  	github.com/soulteary/apt-proxy/pkgs/httpcache	2.575s	coverage: 82.7% of statements
?   	github.com/soulteary/apt-proxy/pkgs/httplog	[no test files]
ok  	github.com/soulteary/apt-proxy/pkgs/stream.v1	1.238s	coverage: 100.0% of statements
ok  	github.com/soulteary/apt-proxy/pkgs/vfs	2.255s	coverage: 59.4% of statements
?   	github.com/soulteary/apt-proxy/proxy	[no test files]
```

View coverage report:

```
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# go test -coverprofile=coverage.out ./...
PASS
coverage: 86.7% of statements
ok  	github.com/soulteary/apt-proxy	0.485s

# go tool cover -html=coverage.out
```

### Ubuntu / Debian Debugging

```
http_proxy=http://192.168.33.1:3142 apt-get -o Debug::pkgProblemResolver=true -o Debug::Acquire::http=true update
http_proxy=http://192.168.33.1:3142 apt-get -o Debug::pkgProblemResolver=true -o Debug::Acquire::http=true install apache2
```

## Licenses, contains dependent software

This project is under the [Apache License 2.0](https://github.com/soulteary/apt-proxy/blob/master/LICENSE), and base on those software (or codebase).

- License NOT Found
    - [lox/apt-proxy](https://github.com/lox/apt-proxy#readme)
- MIT License
    - [lox/httpcache](https://github.com/lox/httpcache/blob/master/LICENSE)
    - [djherbis/stream](https://github.com/djherbis/stream/blob/master/LICENSE)
    - [stretchr/testify](https://github.com/stretchr/testify/blob/master/LICENSE)
- Mozilla Public License 2.0
    - [rainycape/vfs](https://github.com/rainycape/vfs/blob/master/LICENSE)
