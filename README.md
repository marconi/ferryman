# Ferryman

Its the Ferryman's job to move expired jobs from delayed queue to urgent queue. Moving happens in bulk fashion, so if there are several expired jobs, they are popped and pushed all at once. See [Rivers](https://github.com/marconi/rivers) for more information.

## Building and Running

```sh
$ cd cmd/ferryman
$ go build && ./ferryman
```

See `--help` for more info.

## Running Tests

Tests can be run with [Docker](http://www.docker.com) using [Fig](http://www.fig.sh):

```sh
$ fig run --rm test
```

## Status
- Not used in any real project yet
- Coverage is at 70.0%

## License

[http://marconi.mit-license.org](http://marconi.mit-license.org)
