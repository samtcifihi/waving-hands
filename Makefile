compile:
	go build -o "./bin/waving-hands" ./src/main

test: compile
	./bin/waving-hands

clean:
	rm ./bin/waving-hands
