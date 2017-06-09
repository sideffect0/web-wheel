gobuild = go build -ldflags "-s -w" -o
mktar = tar zcvf

fmt:
	gofmt -w .

linux:
	GOOS=linux $(gobuild) dists/linux/web-wheel

release_linux: linux
	$(mktar) dists/web-wheel-linux.tar.gz dists/linux/web-wheel

darwin:
	GOOS=darwin $(gobuild) dists/osx/web-wheel

release_darwin: darwin
	$(mktar) dists/web-wheel-darwin.tar.gz dists/osx/web-wheel

windows:
	GOOS=windows $(gobuild) dists/windows/web-wheel.exe

release_windows: windows
	$(mktar) dists/web-wheel-windows.tar.gz dists/windows/web-wheel.exe

release: release_linux release_darwin release_windows
