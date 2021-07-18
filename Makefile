build:
	go build .

run:
	./lamport 0 &
	./lamport 1 &

clean:
	pkill lamport

