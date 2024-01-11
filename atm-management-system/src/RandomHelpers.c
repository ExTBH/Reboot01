#include <Headers/RandomHelpers.h>
#include <stdlib.h>
#include <string.h>

#define BUFF_SIZE 512

long getFileSize(const char * _Nonnull path) {
    FILE *fp = fopen(path, "r");
    if (fp == NULL) {
        return -1;
    } 
    // file length
    fseek(fp, 0, SEEK_END);
    long size = ftell(fp);
    fclose(fp);
    return size;
}

bool isAlpha(const char *_Nonnull buff) {
    unsigned long len = strlen(buff);
    for (unsigned long i = 0; i < len; i++) {
        if ((buff[i] < 0x20) && (buff[i] < 0x7E)) { return false; }
    }
    return true;
}

bool asciiReader(char * _Nonnull buff, size_t buff_size) {
    if (fgets(buff, buff_size, stdin) != NULL) {
        // Remove the newline character, if it exists
        size_t len = strlen(buff);
        if (len > 0 && buff[len - 1] == '\n') {
            buff[len - 1] = '\0';  // Replace '\n' with '\0'
        }
        return true;
    }
    return false;
}

bool isValidName(const char *_Nonnull buff) {
    unsigned long len = strlen(buff);
    for (unsigned long i = 0; i < len; i++) {
        // Check if the character is not a digit or a letter
        if (!(('0' <= buff[i] && buff[i] <= '9') ||
                ('a' <= buff[i] && buff[i] <= 'z') ||
                ('A' <= buff[i] && buff[i] <= 'Z'))) {
            return false;
        }
    }
    return true;
}

_Nullable AccountType get_accountType(const char * _Nonnull type) {
    if (strncmp(type, accountTypeFixed01, 7) == 0) {
        return accountTypeFixed01;
    } else if (strncmp(type, accountTypeFixed02, 7) == 0) {
        return accountTypeFixed02;
    } else if (strncmp(type, accountTypeFixed03, 7) == 0) {
        return accountTypeFixed03;
    } else if (strncmp(type, accountTypeSavings, 7) == 0) {
        return accountTypeSavings;
    } else if (strncmp(type, accountTypeCurrent, 7) == 0) {
        return accountTypeCurrent;
    } else {
        return NULL; // No match found
    }
}


const char * _Nullable ullToChar(u_int64_t n) {
    char *buffer = calloc(1, sizeof(char) * 21);  // Buffer to hold the string representation (max 20 digits + '\0')
    int chars_written = snprintf(buffer, 20, "%llu", n);
    if ((unsigned long)chars_written >= 20 || chars_written < 0) {
        free(buffer);
        return NULL;
    }
    return buffer;
}

u_int64_t charToUll(const char * _Nonnull s) {
    if (s == NULL || strlen(s) == 0) {
        return -1;
    }
    char *endptr;
    u_int64_t result = strtoull(s, &endptr, 10);
    if (*endptr != '\0') {
        return -1;
    }
    return result;
}

double charToDouble(const char * _Nonnull s) {
    if (s == NULL || strlen(s) == 0) {
        return -1;
    }
    char *endptr;
    double result = strtod(s, &endptr);
    if (*endptr != '\0') {
        return -1;
    }
    return result;
}