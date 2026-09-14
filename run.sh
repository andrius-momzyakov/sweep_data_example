if [ -z "$1" ]
then
go run *.go
else
go run *.go -- "$1"
fi
