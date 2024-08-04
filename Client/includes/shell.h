#pragma once

// Define types for terminal() function return
#define FailedPipe -1
#define SocketClose 0

// Opens a Pipe And Writes To The Socket
int terminal(char *cmd, SOCKET winsock);