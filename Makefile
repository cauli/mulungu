app: .
	docker build -t mulungu .
kill:
	-docker rm -f $$(docker ps -a | grep mulungu | awk '{print $$1}')