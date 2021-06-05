# USB-Test
USB device testing 

This is a Golang tool for testing read and write functions in USB devices.

## Table of Contents
   * [Running From Binary File](#Running-From-BinaryFile)
      * [CLI Usage](#cli-usage)
   * [Installation](#installation)
      * [Compiling on Linux](#compiling-on-linux)
      * [Compiling on MacOS](#compiling-on-macos)

---
## Running From Binary File
### CLI usage

Download `usb-test` binary from release pages

```bash
./usb-test --help # to get help
./usb-test --sim-r=4 # sets simultaneous number of read (default runtime.GOMAXPROCS(0)) - alias(r)
./usb-test --sim-w=4 # sets simultaneous number of write (default runtime.GOMAXPROCS(0)) - alias(w)
./usb-test --root-dir="/mnt/" # sets the mount directory for usb drives (default "/mnt/") - alias(d)

./usb-test --sim-r=4 --sim-w=4 --root-dir="/mnt/
```

## Installation
usb-test is developed and tested on `go version go1.16.5 linux/amd64`

## Compiling on Linux
### Install Go
* Go to [Go Downloads](https://golang.org/dl/) and download from the featured downloads for linux, something like `go1.16.2.linux-amd64.tar.gz`
* Extract the archive and install, you may require root or sudo \
For example: \
   ```rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz```
* Add /usr/local/go/bin to the PATH environment variable. \
Or just use this for a quick check \
```export PATH=$PATH:/usr/local/go/bin```
* Verify that you've installed Go by opening a command prompt and typing the following command: `go version`\
\
Following [this](https://golang.org/doc/install) for more.

#### Install `git` using apt

```bash
sudo apt update
sudo apt install git
```
You can use any other code editor, for installing sublime run these commands:
#### Install `sublime` using apt
```bash
sudo apt update
sudo apt install sublime-text
```

### Usage
Run the program.
```bash
git clone https://github.com/watercompany/usb-test.git
cd usb-test
go mod download
go run main.go
```

Open source code in editor
```!bash
cd usb-test
subl ./
```

## Compiling on MacOS
### Install Go
* Go to [Go Downloads][golang] and download from the featured downloads for Apple macOS, something like `go1.16.2.darwin-amd64.pkg`
* Open the package file you downloaded and follow the prompts to install Go.
* Verify that you've installed Go by opening a command prompt and typing the following command: `go version`\
Following [this](https://golang.org/doc/install) for more.

#### Install `git` with brew

```!bash
brew insall git
```

#### Install `sublime` with brew
```!bash
brew install --cask sublime-text
```

### Usage
Run the program.
```bash
git clone clone https://github.com/watercompany/usb-test.git
cd usb-test
go mod download
go run main.go
```

Open source code in editor
```!bash
cd usb-test
subl ./
```

---

