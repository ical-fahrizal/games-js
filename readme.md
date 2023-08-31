## service app
```
  ip : 103.175.220.6
  id : 
  name : games-js
  user-server : fahrizal
  pass-server : Ical@290983
  location app : /home/fahrizal/golang/games-js
```

## Build
```
	env GOOS=linux GOARCH=amd64 go build -o games-js.linux main.go

	env GOOS=darwin GOARCH=amd64 go build -o games-js.linux main.go

	go build -o games-js.linux main.go
```

## runnig
> nohup ./games-js.linux > games-js.out & 

