Simple webdav server using the golang.org/x/net/webdav package

## Usage 
```
docker run -d -v $PWD:$PWD -w $PWD --name webdav -p 8080:8080 lalyos/webdav
```