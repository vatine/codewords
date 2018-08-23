# Codewords

This is the base code for a workshop on how to instrument source code for Prometheus monitoring. In order to successfully instrument code, you will need to have at least a rough idea of what's important. To help with that, this document contains the design constraints, as well as a rough guide to the code.

## Design brief

The codewords service generates code phrases, constucted by taking an adjective and a noun (so things like "blue orange", "excited wombat", "sad aardvark", ...). It is important that a given code word is never repeated. It is expected that code words will be requested at an average rate (over one day) of one per minute, with a peak rate no higher than one code word every 10 seconds. It is expected that code words are available within 750 ms of the initial request being made (99.9% of the time).

## Code tour

The codewords service is split into a front-end (can be instantiated multiple times and load-balanced for availability) and a backend (not designed to work in parallel). The bulk of the code for the frontend lives in frontend/ and the bulk of the code for the bakcend lives in backend/. However, the actual source code for the main binaries are in cmd/frontend.go and cmd/backend.go respectively.

The frontend(s) and backend communicate using grpc, the user(s) and frontend(s) use HTTP.

## Running

Tested on the standard Centos 7.5 image from Digital Ocean:

```bash
yum -y install git golang
```

And Ubuntu 18.04:

```bash
apt-get update
apt-get install golang-go
useradd -m -s /bin/bash bunty
su - bunty
```

Get required dependencies:
```bash
git clone https://github.com/vatine/codewords.git
cd codewords/cmd
go get
```

Launch both the front-end and back-end commands that are in the cmd directory:
```bash
go run backend.go &   # run in background
go run frontend.go
```
