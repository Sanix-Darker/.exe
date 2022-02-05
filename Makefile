up:
	docker-compose up

down:
	docker-compose down

# to run the installation
# make install-runtimes IP="http://0.0.0.0:2000"
install-runtimes:
	cd piston_cli && docker run -it piston_cli:latest -u $(IP) ppman install $(cat runtimes.txt | tr , " ")


# to run the installation
# make install-runtime IP="http://0.0.0.0:2000" RUNTIME="python=3.8.4"
install-runtime:
	cd piston_cli && docker run -it piston_cli:latest -u $(IP) ppman install $(RUNTIME)

