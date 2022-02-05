# .exe

A runtime collection embed in github with a web-extension.


## Requirements
- docker
- docker-compose


## How to start

Just after cloning the repo, you just have to run docker-compose

```bash
docker-compose up

# or if you have make installed
make up
```
This command will start piston and exe api at the same time

Then to install runtimes, you just have to see `./piston_cli/README.md`
Or execute those make commands :
```bash
# to install all runtimes
make install-runtimes IP="http://0.0.0.0:2000"

# to install only one runtime
make install-runtime IP="http://0.0.0.0:2000" RUNTIME="python=3.8.4"
```

## Author

- [sanix-darker](github.com/sanix-darker)
