gobuild = go build -ldflags "-s -w" -o
mktar = tar zcvf

linux:
	GOOS=linux $(gobuild) dists/linux/dev-server && $(mktar) dists/dev-server-linux.tar.gz dists/linux/dev-server

darwin:
	GOOS=darwin $(gobuild) dists/osx/dev-server && $(mktar) dists/dev-server-darwin.tar.gz dists/osx/dev-server

windows:
	GOOS=windows $(gobuild) dists/windows/dev-server.exe && $(mktar) dists/dev-server-windows.tar.gz dists/windows/dev-server.exe

all: linux darwin windows

