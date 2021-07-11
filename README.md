## About

Brainfuck-Go is a [Brainfuck](http://en.wikipedia.org/wiki/Brainfuck) interpreter written in Go.

## Usage

Compile using go build and run:

```sh
./bf filename
```

## Examples

```sh
./bf hw.bf
```

Increase the steps for long programs

```sh
curl -s "http://www.99-bottles-of-beer.net/download/1718" > bottles.bf
./bf --max-steps 1000000  bottles.bf 
```

see [caesar_cipher](./caesar_cipher)

## License

[The MIT License (MIT)](http://opensource.org/licenses/mit-license.php)
