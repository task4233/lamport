build:
	(cd client; go build . ; cd ..)
	(cd server; go build . ; cd ..)

run:
	./server/server &
	./client/client &

clean:
	pkill server
	pkill client
