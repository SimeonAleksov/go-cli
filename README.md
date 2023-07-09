# go-cli

CLI to consume kafka events.

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/go-cli
```

However for that you'll need to adjust the .env with the correct host names, we can also run docker:
```shell
$ docker compose up
```

### Testing

``make test``