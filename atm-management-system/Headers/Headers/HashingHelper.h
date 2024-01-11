#ifndef HASHINGHELPER_H
#define HASHINGHELPER_H

enum HashingHelperErorr {
    HashingHelperErorrNone,
    HashingHelperErorrFile,
    HashingHelperErorrMemory
};

enum HashingHelperErorr calc_sha256(const char * _Nonnull path, char output[_Nonnull 65]);

#endif