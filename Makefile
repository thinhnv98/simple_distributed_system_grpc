sync-proto:
	cp -r proto A
	cp -r proto B

gen:
	make -C A gen
	make -C B gen

run-server:
	make -C A run-server

run-client:
	make -C B run-client