# Caesar Cipher

> In cryptography, a [Caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher), also known as Caesar's cipher, the shift cipher, Caesar's code or Caesar shift, is one of the simplest and most widely known encryption techniques. It is a type of substitution cipher in which each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet. For example, with a left shift of 3, D would be replaced by A, E would become B, and so on. The method is named after Julius Caesar, who used it in his private correspondence.[1]

## Usage

The first two characters must be a decimal number that will be the password. The rest of the characters will be the message. Only lowercase characters are allowed
the quick brown fox jumps over the lazy dog

### Encode

```sh
$ echo -n "e07 hello world" | bf caesar.bf
 olssv dvysk
```

### Decode

```sh
$ echo -n "d07 olssv dvysk" | bf caesar.bf
 hello wolrd
```