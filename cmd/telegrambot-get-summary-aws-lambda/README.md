windows:


$Env:GOOS = "linux"

$Env:GOARCH="arm64"

go build -o bootstrap main.go

zip bootstrap.zip bootstrap 
