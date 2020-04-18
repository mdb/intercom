# intercom

A small logging package for use in Golang-based [Concourse resource type implementations](https://concourse-ci.org/implementing-resource-types.html)

Because Concourse resources communicate with Concourse over stdout, it's [recommended](https://concourse-ci.org/implementing-resource-types.html) that:

> Resources can emit logs to the user by writing to stderr. ANSI escape codes (coloring, cursor movement, etc.) will be interpreted properly by the web UI, so you should make your output pretty.

`intercom` does just this and writes to stderr.

## Usage Example

```golang
logger := intercom.NewLogger("debug")

// prints red text
logger.Errorf("foo")

// prints yellow text
logger.Warnf("bar")

// prints green text
logger.Infof("baz")

// prints blue text
logger.Debugf("bim")
```

All methods ultimately use `Fprintf`, so you can also do things like...

```golang
// prints 'foo bar' in red text
logger.Errorf("foo %s", "bar")
```
