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
  QueryHelperUpdateAccount,
  QueryHelperUpdateAccountBalance,
  QueryHelperDeleteAccount,
  QueryHelperTransferAccount,
};
// free after done
const char *_Nonnull getQuery(enum QueryHelper type);
#endif