linux:
	GOOS=linux go build -ldflags "-s -w" -o dists/linux/dev-server && tar zcvf dists/dev-server-linux.tar.gz dists/linux/dev-server

darwin:
	GOOS=darwin go build -ldflags "-s -w" -o dists/osx/dev-server && tar zcvf dists/dev-server-darwin.tar.gz dists/osx/dev-server

windows:
	GOOS=windows go build -ldflags "-s -w" -o  dists/windows/dev-server.exe && tar zcvf dists/dev-server-windows.tar.gz dists/windows/dev-server.exe

all: linux darwin windows

