compile: format
	go build -o "./bin/waving-hands" ./src/main.go

format:
	gofmt -w ./src/*.*
	gofmt -w ./src/*/*.*

run: compile
	./bin/waving-hands

clean:
	rm ./bin/waving-hands
