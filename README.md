# Nintaco Go API - Hello World Example

### About

[The Nintaco NES/Famicom emulator](https://nintaco.com/) provides [a Go API](https://github.com/meatfighter/nintaco-go-api) that enables externally running programs to control the emulator at a very granular level. This example displays "Hello, World!" and a bouncing ball above a running game.

### External Control

To control Nintaco from a Go program:

1. Start Nintaco and launch a game.
2. Open the Start Program Server window via Tools | Start Program Server...
3. Press Start Server.
4. From the command-line, execute the Go program that uses the Nintaco API.

To disconnect, press Stop Server and/or break the example program (Ctrl+C). 