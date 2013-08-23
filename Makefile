all:
	go get -d -v github.com/bruth/assert
	go get -d -v labix.org/v2/mgo/bson
	go build -v
