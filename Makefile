compile:
	go build -o "./bin/waving-hands" ./src/main

run: compile
	./bin/waving-hands

clean:
	rm ./bin/waving-hands
