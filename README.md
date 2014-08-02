# Ferryman

Its the Ferryman's job to move expired jobs from delayed queue to urgent queue. Moving happens in bulk fashion, so if there are several expired jobs, they are popped and pushed all at once. See [Rivers](https://github.com/marconi/rivers) for more information.

## Running Tests

Tests can be run with [Docker](http://www.docker.com) using [Fig](http://www.fig.sh):

```sh
$ fig run --rm test
```

## License

[http://marconi.mit-license.org](http://marconi.mit-license.org)
