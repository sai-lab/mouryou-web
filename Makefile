GO_BUILDOPT := -ldflags '-s -w'

gom:
	go get github.com/mattn/gom
	gom install

run:
	gom run *.go ${ARGS}

fmt:
	gom exec goimports -w *.go lib/*/*.go

bindata:
	gom exec go-bindata-assetfs ./app/... ./assets/... ./lib/views/...

debugdata:
	gom exec go-bindata-assetfs -debug=true ./app/... ./assets/... ./lib/views/...

build: fmt bindata
	gom build $(GO_BUILDOPT) -o bin/mouryou-web *.go

clean:
	rm -f bin/mouryou-web

link:
	mkdir -p $(GOPATH)/src/github.com/sai-lab
	ln -s $(CURDIR) $(GOPATH)/src/github.com/sai-lab/mouryou-web

unlink:
	rm $(GOPATH)/src/github.com/sai-lab/mouryou-web
	rmdir $(GOPATH)/src/github.com/sai-lab
