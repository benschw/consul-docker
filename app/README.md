

	go build -o app
	docker build -t benschw/app .

	kubecfg -c /vagrant/app/app.yaml create replicationControllers