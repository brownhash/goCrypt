# goCrypt
A layer above your regular terminal editor to write/edit/read encrypted files

### Usage
```shell script
git clone https://github.com/sharma1612harshit/goCrypt
cd goCrypt
go run main.go [options]
```

### Building Executables
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
   if -k is not provided then the default key is used
```
