# Language Randomizer

On AoC 2022 day 3, I have decided to build a random generator from language yaml file of github linguist. I have also decided to make it cool looking CLI.
This satisfies the need for randomly picking languages. Since linguist's language yaml file is a bit odd, I had to clean up some data formats and add --select flag to my CLI to
pick multiple languages at once.

## Screenshot

![language-randomizer-ss](https://user-images.githubusercontent.com/22800416/205465903-51ec0847-30cc-47f3-b361-feded432e67c.png)

## Installation

Grab the binary from github releases or clone the repository and run it from source.
```sh
    git clone https://github.com/mrwormhole/language-randomizer.git
    go run main.go
```

## Usage

To specify only once
```sh
    ./Language\ Randomizer
```

To specify more than 1

```sh
  ./Language\ Randomizer --select=2
```
