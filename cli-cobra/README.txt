https://blog.knoldus.com/create-kubectl-like-cli-with-go-and-cobra/

https://github.com/knoldus/greetctl


mkdir cli-cobra
# Create a go module for CLI.
go mod init satya.com/greetctl

# Get Cobra library
go get -u github.com/spf13/cobra/cobra

# Create a bare minimum skeleton
cobra init --pkg-name satya.com/greetctl

go run main.go
go build .
# use . to access the command
./greetctl --help
go install
# no need to use . to access the command
greetctl --help
