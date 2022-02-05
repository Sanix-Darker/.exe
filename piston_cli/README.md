Using docker you can build the piston-cli:

```bash
docker build -t piston_cli:latest -f Dockerfile .
```

And then run it just like `cli/index.js $arg1 $arg1` instead of having to install node yourself... etc

```bash
# to run for having the list of runtimes
docker run -it piston_cli:latest -u http://192.168.1.26:2000 ppman list

# to install runtimes
docker run -it piston_cli:latest -u http://192.168.1.26:2000 ppman install $(cat runtimes.txt | tr , " ")
```
