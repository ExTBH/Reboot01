#include <Headers/QueryHelpers.h>
#include <string.h>
const char *VALIDATE_TABLES_SQL = "SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE' AND TABLE_NAME IN ('users', 'accounts', 'accounts_transfers', 'transactions');";
const char *ADD_USER_QUERY = "INSERT INTO users (name, password) VALUES ($1, $2) RETURNING *;";
const char *GET_USER_WITHNAME_QUERY = "SELECT user_id, name, password, active FROM Users WHERE name = $1;";
const char *GET_USER_WITHID_QUERY = "SELECT user_id, name, password, active FROM Users WHERE user_id = $1;";
const char *ADD_ACCOUNT_QUERY = "INSERT INTO accounts (user_id, type, country, phone, balance) VALUES ($1, $2, $3, $4, $5) RETURNING *;";
const char *UPDATE_ACCOUNT_QUERY = "UPDATE Accounts SET phone = $1, country = $2 WHERE account_id = $3;";
const char *UPDATE_ACCOUNT_BALANCE_QUERY = "UPDATE Accounts SET balance = $1 WHERE account_id = $2;";
// const char *GET_ACCOUNT_QUERY = "SELECT account_id, type, date, balance, country, phone FROM accounts WHERE $1 = account_id;";
const char *GET_ACCOUNTS_QUERY = "SELECT account_id, type, date, balance, country, phone FROM accounts WHERE active = TRUE AND $1 = user_id;";
const char *DELETE_ACCOUNT_QUERY = "UPDATE accounts SET active = FALSE WHERE account_id = $1;";
const char *TRANSFER_ACCOUNT_QUERY = "UPDATE Accounts SET user_id = $1 WHERE account_id = $2;";

const char* _Nonnull getQuery(enum QueryHelper type) {
    switch (type) {
        case QueryHelperValidate:
            return strdup(VALIDATE_TABLES_SQL);
        case QueryHelperAddUser:
            return strdup(ADD_USER_QUERY);
        case QueryHelperGetUserName:
            return strdup(GET_USER_WITHNAME_QUERY);
        case QueryHelperGetUserID:
            return strdup(GET_USER_WITHID_QUERY);
        case QueryHelperAddAccount:
            return strdup(ADD_ACCOUNT_QUERY);
        // case QueryHelperGetAccount:
        //     return strdup(GET_ACCOUNT_QUERY);
        case QueryHelperGetAccounts:
            return strdup(GET_ACCOUNTS_QUERY);
        case QueryHelperDeleteAccount:
            return strdup(DELETE_ACCOUNT_QUERY);
        case QueryHelperTransferAccount:
            return strdup(TRANSFER_ACCOUNT_QUERY);
        case QueryHelperUpdateAccount:
            return strdup(UPDATE_ACCOUNT_QUERY);
        case QueryHelperUpdateAccountBalance:
            return strdup(UPDATE_ACCOUNT_BALANCE_QUERY);
        default:
            __builtin_unreachable();
    }
}