build-and-copy:
	GOOS=windows GOARCH=amd64 go build -v .
	-rm go-issue-68268.upx.exe
	upx -o go-issue-68268.upx.exe go-issue-68268.exe
	cp -v go-issue-68268.* ~/utm-shared/
