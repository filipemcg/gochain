# GoChain

Really basic block chain using Go.

Concept by Anders Brownworth:
- Part1: https://youtu.be/_160oMzblY8?si=4UpdqjvXa9_FXvdV
- Part2: https://youtu.be/xIDL_akeras?si=1WNEODoB130b188v

## How to use

Run using Go:
```sh
go run ./src
```

Creating your first block:
```sh
curl --location 'http://localhost:9000/blocks' \
    --header 'Content-Type: application/json' \
    --data '{"data": "My First Block"}'
```

Getting block info:
```sh
curl --location --request GET 'http://localhost:9000/blocks/000076403cc7726b72b23eae2e43f60324b232da4674f01674773e7f5301ffa0' \
    --header 'Content-Type: application/json' \
    --data '{"data": "Hello"}'
```

