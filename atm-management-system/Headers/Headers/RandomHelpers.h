#ifndef RANDOMHELPERS_H
#define RANDOMHELPERS_H

#include <Headers/ATM.h>
#include <stdio.h>
#include <stdbool.h>

// return the file size, or -1 on error
long getFileSize(const char * _Nonnull path);

// checks if a string is valid alpha
bool isAlpha(const char *_Nonnull buff);

// checks if a string is [0-9a-zA-Z]
bool isValidName(const char *_Nonnull buff);

// Reads a string from stdin, with support for spaces
bool asciiReader(char * _Nonnull buff, size_t buff_size);

// converts `type` to  `AccountType` or `NULL`
_Nullable AccountType get_accountType(const char * _Nonnull type);

// convert a `u_int64_t` to a string, returns a `char[21]` or `NULL` and you should free it
const char * _Nullable ullToChar(u_int64_t n);

// inverse of ullToChar, -1 in errors
u_int64_t charToUll(const char * _Nonnull s);

double charToDouble(const char * _Nonnull s);

#endif