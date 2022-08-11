compile: format
	go build -o "./bin/waving-hands" ./main.go

format:
	gofmt -w ./main.go
	gofmt -w ./gamestate/*.*

run: compile
	./bin/waving-hands

clean:
	rm ./bin/waving-hands
