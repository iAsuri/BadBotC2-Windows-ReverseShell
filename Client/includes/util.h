#pragma once

#include <strings.h>
#include <winsock.h>

// XorBytes
/*
 * Params: string/Characters that you want Xored, char length, XorKey
*/
void xorbytes(char *bytes, int size, int key);

// write_xor
/*
 * Params: Text To Xor, winSocked, xorkey
 * Description: This function uses the xorbytes function to xor the text then is sent over socket
*/
int write_xor(char *text, SOCKET winsock, int key);