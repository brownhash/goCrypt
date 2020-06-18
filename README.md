# goCrypt
A layer above your regular terminal editor to write/edit/read encrypted files

### Usage
Set your preferred terminal editor as an environment variable, else it will use `vim`
```shell script
export EDITOR='<name of preferd editor>'
```
```shell script
git clone https://github.com/sharma1612harshit/goCrypt
cd goCrypt
go run main.go [options]
```

### [Docker Image](https://hub.docker.com/r/sharma1612harshit/gocrypt)
```shell script
docker pull sharma1612harshit/gocrypt
```
or
```shell script
git clone https://github.com/sharma1612harshit/goCrypt
cd goCrypt
docker build --tag gocrypt .
```
then,
```shell script
docker run -it -v <localpath>:/tmp/ gocrypt [option] /tmp/<filename> -k <encryption key>
```


### Building Executables
```shell script
git clone https://github.com/sharma1612harshit/goCrypt -b <latest tag>-binaries

Here the directories are of the form (os-architecture)
choose your os and architecture and use the precompiled binary
```

OR create your own binary,

```shell script
env GOOS=target-OS GOARCH=target-architecture go build package-import-path
```
Available options:

| target-OS     | target-architecture |
| ------------- | ------------------- |
| darwin	    | 386                 |
| darwin	    | amd64               |
| darwin	    | arm                 |
| darwin	    | arm64               |
| dragonfly	    | amd64               |
| freebsd	    | 386                 |
| freebsd	    | amd64               |
| freebsd	    | arm                 |
| linux	        | 386                 |
| linux	        | amd64               |
| linux	        | arm                 |
| linux	        | arm64               |
| linux	        | ppc64               |
| linux	        | ppc64le             |
| linux	        | mips                |
| linux	        | mipsle              |
| linux	        | mips64              |
| linux	        | mips64le            |
| netbsd	    | 386                 |
| netbsd	    | amd64               |
| netbsd	    | arm                 |
| openbsd	    | 386                 |
| openbsd	    | amd64               |
| openbsd	    | arm                 |
| plan9	        | 386                 |
| plan9	        | amd64               |
| solaris	    | amd64               |

### Help
```shell script
Usage: gocrypt [option] [supporting options]

Options:
1) -w (write a file): gocrypt -w filename [optional arguments]
2) -r (read a encrypted file): gocrypt -r filename [optional arguments]
3) -h (help): gocrypt -h

Optional Arguments:
1) -k (encryption key): 
          gocrypt -w filename -k encryption_key
          gocrypt -r filename -k encryption_key
   if key is not provided then the default key is used

2) -f (file name of file containing encryption key)
          gocrypt -w filename -f /home/username/encryptionkey.json
          gocrypt -r filename -f /home/username/encryptionkey.json
```

Writing keys in file, while using `-f` option: <br />
The file should be a valid json file with a key value pair where both key and value are strings.

```shell script
{
    "key": "MyENCRYptionkey123"
}
```
