// #include <Headers/HashingHelper.h>
// #include <Headers/RandomHelpers.h>
// #include <stdio.h>

// #ifdef __APPLE__
// #include <CommonCrypto/CommonCrypto.h>

// enum HashingHelperErorr calc_sha256(const char * _Nonnull path, char output[_Nonnull 65]) {
//     FILE *fp = fopen(path, "rb");
//     if (fp == NULL) {
//         return HashingHelperErorrFile;
//     } 
//     long file_size = getFileSize(path);
//     if (file_size == -1) {
//         return HashingHelperErorrFile;
//     }
//     unsigned char *fdata = malloc(file_size);
//     // probably will never happen but couldddd
//     if (fdata == NULL) {
//         fclose(fp);
//         return HashingHelperErorrMemory;
//     }
//     fread(fdata,1, file_size, fp);
//     fclose(fp);
//     // actual hashing
//     unsigned char hash[CC_SHA256_DIGEST_LENGTH];
//     CC_SHA256(fdata, file_size, hash);
//     // to hex
//     for (int i = 0; i < CC_SHA256_DIGEST_LENGTH; i++) {
//         sprintf(output + i * 2, "%02x", hash[i]);
//     }
//     output[CC_SHA256_BLOCK_BYTES] = '\0';
//     return HashingHelperErorrNone;
// }

// #elif __linux__
// enum HashingHelperErorr calc_sha256(const char *path, char output[65]) {
    
// }

// #endif