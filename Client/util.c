#include "./includes/util.h"

// XorBytes
/*
 * Params: string/Characters that you want Xored, char length, XorKey
 */
void xorbytes(char *bytes, int size, int key)
{
    for (int i = 0; i < size; i++)
        bytes[i] ^= key;
}

// write_xor
/*
 * Params: Text To Xor, winSocked, xorkey
 * Description: This function uses the xorbytes function to xor the text then is sent over socket
*/
int write_xor(char *text, SOCKET winsock, int key)
{
    // Length of text
    int textlen = strlen(text);

    // Xor Bytes
    xorbytes(text, textlen, key);
    return send(winsock, text, textlen, 0);
}