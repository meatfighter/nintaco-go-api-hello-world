# Nintaco Go API - Hello World Example

### Remote Control

To programmatically control Nintaco from an externally running Go program:

1. Start Nintaco and launch any game.
2. Open the Start Program Server window via Tools | Start Program Server...
3. Press Start Server.
4. From the command-line, execute the Go program.

In the case of this example program, if an internal socket connection was successfully established between it and Nintaco, then you should see "Hello, World!" and a bouncing ball displayed above the running game.  To disconnect, press Stop Server and/or break the example program (Ctrl+C). 