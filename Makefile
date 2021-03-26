build :
	rm -rf ./dist
	GOOS=linux GOARCH=amd64 go build -o ./dist/postman-to-markdown_linux   -ldflags "-s -w"
	GOOS=darwin GOARCH=amd64 go build -o ./dist/postman-to-markdown_darwin   -ldflags "-s -w"
	cd dist && upx -9 ./postman-to-markdown_linux && upx -9 ./postman-to-markdown_darwin  
	cd dist && tar czvf linux.tgz ./postman-to-markdown_linux 
	cd dist && tar czvf mac.tgz ./postman-to-markdown_darwin
