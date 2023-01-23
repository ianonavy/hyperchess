# hyperchess

Chess engine go brrrrrrrrrr

## Run locally

Start one engine on TCP port 5556 (white):

```bash
socat TCP4-LISTEN:5556,reuseaddr,fork EXEC:"$(which stockfish)"
```

Start one engine on TCP port 5557 (black):

```bash
socat TCP4-LISTEN:5557,reuseaddr,fork EXEC:"$(which stockfish)"
```

Compile and run:

```bash
go build && ./hyperchess
```

Start a game:

```bash
curl http://localhost:8080/chess/
```
