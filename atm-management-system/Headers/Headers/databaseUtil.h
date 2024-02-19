#ifndef DATABASEUTIL_H
#define DATABASEUTIL_H
#include <stdio.h>
#include <stdbool.h>
#include "ATM.h"
enum ATMDBStatus {
    // MARK: - Generics
    ATMDBStatusOK,
    ATMDBStatusNotOK,
    ATMDBStatusCantConnect,
    ATMDBStatusCantVerify,
    // MARK: - Users
    ATMDBStatusUsersNull,
    ATMDBStatusUsersConstraintID,
    ATMDBStatusUsersConstraintName,
    ATMDBStatusUsersBadName,
    // MARK: - Accounts
    ATMDBStatusAccountsNull,
    ATMDBStatusAccountsUserNotFound,
    ATMDBStatusAccountsBadType
};
enum ATMDBStatus atmdb_connect();
void atmdb_close(void);
// MARK: - User Related
// adds a new User, ignoring the ID and active fields
enum ATMDBStatus add_user(struct User * _Nonnull u);
enum ATMDBStatus get_user(struct User * _Nonnull u, bool withID);
// MARK: - Accounts Related
enum ATMDBStatus add_account(struct Account * _Nonnull acc);
enum ATMDBStatus get_accounts(struct User * _Nonnull u,  struct Account * _Nonnull * _Nonnull * _Nonnull accs, int * _Nonnull count);
enum ATMDBStatus delete_account(struct Account * _Nonnull acc);
enum ATMDBStatus update_account(const unsigned long acc_id, const char country[_Nonnull 20], const char phone  [_Nonnull 10]);
enum ATMDBStatus update_account_balance(const unsigned long acc_id, const double balance);

// MARK: - Transfers Related
enum ATMDBStatus transfer_account(struct Account * _Nonnull acc, struct User * _Nonnull receiver);
#endif