

	go build -o app
	docker build -t benschw/app .



	vagrant up
	ssh node1

	cd /etc/systemd/system
	fleetctl submit myapp\@.service
	fleetctl start myapp@1
	fleetctl start myapp@2
	fleetctl list-units










	docker port myapp-1 8080 | awk -F: '{print $2}'

	docker inspect --format '{{ .NetworkSettings.Gateway }}' myapp-1