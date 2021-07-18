build:
	go build .

run:
	./lamport 0 &
	echo "wait 1sec"; sleep 1;
	./lamport 1 &

clean:
	pkill lamport

