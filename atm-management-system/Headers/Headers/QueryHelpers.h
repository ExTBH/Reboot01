#ifndef QUERYHELPER_H
#define QUERYHELPER_H

enum QueryHelper {
    QueryHelperValidate,
    QueryHelperAddUser,
    QueryHelperGetUserName,
    QueryHelperGetUserID,
    QueryHelperAddAccount,
    // QueryHelperGetAccount,
    QueryHelperGetAccounts,
    QueryHelperDeleteAccount,
    QueryHelperTransferAccount,
    QueryHelperConfirmTransferAccount,
};

// free after done
const char* _Nonnull getQuery(enum QueryHelper type);

#endif