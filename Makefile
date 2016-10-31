gobuild = go build -ldflags "-s -w" -o
mktar = tar zcvf

fmt:
	gofmt -w .

linux:
	GOOS=linux $(gobuild) dists/linux/dev-server

release_linux: linux
	$(mktar) dists/dev-server-linux.tar.gz dists/linux/dev-server

darwin:
	GOOS=darwin $(gobuild) dists/osx/dev-server

release_darwin: darwin
	$(mktar) dists/dev-server-darwin.tar.gz dists/osx/dev-server

windows:
	GOOS=windows $(gobuild) dists/windows/dev-server.exe

release_windows: windows
	$(mktar) dists/dev-server-windows.tar.gz dists/windows/dev-server.exe

release: release_linux release_darwin release_windows
