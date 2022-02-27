# CLI Cobra

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

### Dependencies

* Describe any prerequisites, libraries, OS version, etc., needed before installing program.
* ex. Windows 10

### Installing


* Create a go module for CLI.
```
mkdir cli-cobra
go mod init satya.com/greetctl
```

* Get Cobra library
```
go get -u github.com/spf13/cobra/cobra
```

* Create a bare minimum skeleton

```
cobra init --pkg-name satya.com/greetctl
```

### Executing program

* How to run the program
* Step-by-step bullets

```
go run main.go
go build .
# use . to access the command
./greetctl --help
go install
# no need to use . to access the command
greetctl --help
```

## Help

Any advise for common problems or issues.
```
command to run if program contains helper info
```

## Authors

Contributors names and contact info

[@SatyaDillikar](https://twitter.com/SatyaDillikar)

## Version History

* 0.2
    * Various bug fixes and optimizations
    * See [commit change]() or See [release history]()
* 0.1
    * Initial Release

## License

N/A

## Acknowledgments

Inspiration, code snippets, etc.
* [kubectl](https://blog.knoldus.com/create-kubectl-like-cli-with-go-and-cobra/)
* [greetctl](https://github.com/knoldus/greetctl)