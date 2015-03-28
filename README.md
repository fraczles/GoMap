go build emerging.go cmap.go
./emerging -readers=2 -askers=2 -askdelay=10 -reducedelay=100 -infiles="pg1041.txt,pg1103.txt"

