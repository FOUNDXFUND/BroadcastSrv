# BroadcastSrv
a golang broadcast server for game,web,etc...

# Dependence
* gin https://github.com/gin
* nano https://github.com/lonnng/nano

# Run Test
* go run main.go
* open browser: http://localhost:3250/web/
* send msg to user:curl -i -d '{"uid":3}' http://127.0.0.1:3000/broadcast/single // checkout your uid
