# Substenv

Replace `$ENVVAR` or `${ENVVAR}` references in its input with values from the environment.

## Usage

```
substenv [<flags>] [<input>]

Flags:
  --help       Show help.
  -r, --regex  Use slower but less greedy regex parser
  --version    Show application version.

Args:
  [<input>]  Input file or stdin if not given
```

Install with

```
    $ go get github.com/ilkka/substenv
```
