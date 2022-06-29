#include <immintrin.h>
#include <stdint.h>
#include <string.h>

#include "exports.h"

static char msg[] = "your message";

int subr()
{
    return 0xcafebabe;
}

int subr2(GoSlice *slice)
{
    if (!slice) return -1;
    char *data = slice->data;

    for (int i = 0; i < slice->len; i++) {
        data[i] = msg[i % sizeof msg];
    }

    return slice->len;
}

int subr3(GoSlice *slice)
{
    char *data;
    if (!slice) return 0;
    data = slice->data;
    if (!data) return 0;

    for (int i = 0; i < slice->len; i++) {
        data[i] = msg[i % sizeof msg];
    }
    return slice->len;
}

int subrSimd(GoSlice *slice, char c)
{
    unsigned char *buf, *rb;
    int nb;
    if (!slice) return 0;
    buf = slice->data;
    if (!buf) return 0;
    rb = buf;
    nb = slice->len;
    if (nb < 1) return 0;
    while (nb >= 16)
    {
        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;
        
        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;

        nb -= 16;
    }
    while (nb >= 8)
    {

        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;

        nb -= 8;
    }
    while (nb >= 4)
    {
        *rb++ = c; *rb++ = c;
        *rb++ = c; *rb++ = c;

        nb -= 4;
    }
    while (nb > 0)
    {
        *rb++ = c;
        nb -= 1;
    }
    return slice->len;
}