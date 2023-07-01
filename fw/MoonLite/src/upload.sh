
export GOPATH=${PWD}/..
#tinygo flash -target=amfoc01 main.go -v
tinygo build -target=./targets/amfoc01.json -o amfoc.uf2 main.go 
tinygo flash -target=./targets/amfoc01.json main.go
