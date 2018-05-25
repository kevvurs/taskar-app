# Taskar App
The contextual awareness project, which supports the autonomous
wheelchair project, uses this service as a receptor for sensor data.
the microcontroller connects to this app periodically and sends the
readings from all the peripheral sensor devices.

This app is written in Go, and it starts a server to receive the data.
The app will also host a UI webapp for visualizing data. Data is persisted
in a MySQL 5.7 instance.

## Installing
Install the following for the best experience:
 *  Go  https://golang.org/
 *  Glide  https://glide.sh/

Setup `$GOPATH` and check that `$GOROOT` is correct.
Then use the following command to start.
```
go get github.com/kevvurs/taskar-app
cd "$GOPATH"/src/github.com/kevvurs/taskar-app
glide install
go build && ./taskar-app
```

## Description
_TODO_
