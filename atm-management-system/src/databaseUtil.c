#include <Headers/databaseUtil.h>
#include <Headers/RandomHelpers.h>
#include <Headers/QueryHelpers.h>
#include <stdio.h>
#include <strings.h>
#include <stdlib.h>
#include <libpq-fe.h>

#define DB_PATH "db/atm.db"
#define VALIDATE_TABLES_QUERY "db/validate.sql"
#define TABLES_COUNT 4


static enum ATMDBStatus _verify(void);
int count(const PGresult *res);
enum ATMDBStatus _get_user(struct User * _Nonnull u, const char *buff,  const char *queryPath);
enum ATMDBStatus _fill_user(const PGresult * _Nonnull res, struct User * _Nonnull u);

PGconn *conn = NULL;

// attempts to connect and verify the tables 
enum ATMDBStatus atmdb_connect() {
    conn = PQconnectdb("dbname=atmdb user=natheer password=1234");
    if (PQstatus(conn) != CONNECTION_OK) {
        fprintf(stderr, "Connection to database failed: %s", PQerrorMessage(conn));
        atmdb_close();
        return ATMDBStatusCantConnect;
    }
    if ( _verify() == ATMDBStatusCantVerify) {
        atmdb_close();
        return ATMDBStatusCantVerify;
    }
    return ATMDBStatusOK;
}

void atmdb_close(void) {
    PQfinish(conn);
}

// verifies the tables existence in the database
static enum ATMDBStatus _verify(void) {
    const char *query = getQuery(QueryHelperValidate);

    PGresult *res = PQexec(conn, query);
    free((char*)query);
    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        fprintf(stderr, "Query execution failed: %s", PQresultErrorMessage(res));
        PQclear(res);
        atmdb_close();
        return ATMDBStatusCantVerify;
    }
    int matches = count(res);
    if ( matches == -1) {
        fprintf(stderr, "An Error happened while verifying tables\n");
        return ATMDBStatusCantVerify;
    }
    if ( matches != 4) {
        fprintf(stderr, "Found %d instead of 4 while veryfying\n", matches);
        return  ATMDBStatusCantVerify;
    }
    return ATMDBStatusOK;
}

// used to count results, `-1` is error
int count(const PGresult *res) {
    int rowCount = PQntuples(res);
    if (rowCount > 0) {
        return  atoi(PQgetvalue(res, 0, 0));
    }
    return -1;
}

// just fills a user from a `res`, it doesnt validate it
enum ATMDBStatus _fill_user(const PGresult * _Nonnull res, struct User * _Nonnull u) {
// get and convert ID to int
    char *id_str = PQgetvalue(res, 0, 0);
    uint64_t id = charToUll(id_str);
    
    if (id < 0) {
        puts("ID is not valid");
        // handle
        return ATMDBStatusNotOK;
    }
    u->id = id;
    // set name and pass
    strncpy(u->name, PQgetvalue(res, 0, 1), 50);
    strncpy(u->password, PQgetvalue(res, 0, 2), 64);
    u->name[50] = '\0';
    u->password[64] = '\0';
    // set active, t = true, f = false
    if (strncmp("t", PQgetvalue(res, 0, 3) , 1) == 0) {
        u->active = true;
    } else {
        u->active = false;
    }
    return ATMDBStatusOK;
}

enum ATMDBStatus _get_user(__unused struct User * _Nonnull u, const char *buff,  const char *userQuery) {
    PGresult *res = PQexecParams(conn, userQuery, 1, NULL, &buff, NULL, NULL, 0);
    free((char*)userQuery);
    if (PQresultStatus(res) != PGRES_TUPLES_OK || PQntuples(res) != 1) {
        PQclear(res);
        return ATMDBStatusNotOK;
    }
    enum ATMDBStatus st = _fill_user(res, u);
    PQclear(res);
    return st;
}

enum ATMDBStatus add_user(struct User * _Nonnull u) {
    if (u == NULL) {return ATMDBStatusUsersNull;}
    // validate name
    if (strlen(u->name) <= 0 || !isValidName(u->name)) {
        return ATMDBStatusUsersBadName;
    }

    const char *userQuery = getQuery(QueryHelperAddUser);
    // prepare query
    const char *params[2] = {u->name, u->password};
    PGresult *res = PQexecParams(conn, userQuery, 2, NULL, params, NULL, NULL, 0);
    free((char*)userQuery);
    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        // catch Name Insert vilation
        const char *sqlState = PQresultErrorField(res, PG_DIAG_SQLSTATE);
        if (sqlState && strcmp(sqlState, "23505") == 0) {
            PQclear(res);
            return ATMDBStatusUsersConstraintName;
        }
        fprintf(stderr, "Query execution failed: %s", PQresultErrorMessage(res));
        PQclear(res);
        return ATMDBStatusNotOK;
    }
    _fill_user(res, u);
    PQclear(res);
    return ATMDBStatusOK;
}

enum ATMDBStatus get_user(struct User * _Nonnull u, bool withID) {
    if (u == NULL) {return ATMDBStatusUsersNull;}
    enum ATMDBStatus st;
    const char *query;
    if (withID) {
        // uint64 to string
        const char* buffer = ullToChar(u->id);
        if (buffer == NULL) { return ATMDBStatusNotOK;}

        query = getQuery(QueryHelperGetUserID);
        st = _get_user(u, buffer, query);

        free((char*)buffer);
    } else {
        query = getQuery(QueryHelperGetUserName);
        st= _get_user(u, u->name, query);
    }
    free((char*)query);
    return st;
}

enum ATMDBStatus add_account(struct Account * _Nonnull acc) {
    if ( acc == NULL) { return ATMDBStatusAccountsNull;}
    
    const char *userQuery = getQuery(QueryHelperAddAccount);
    // uint64 to string
    const char* buffer = ullToChar(acc->user->id);
    if (buffer == NULL) { return ATMDBStatusNotOK;}
    
    const char *params[4] = {buffer, acc->type, acc->country, acc->phone};
    PGresult *res = PQexecParams(conn, userQuery, 4, NULL, params, NULL, NULL, 0);

    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        const char *sqlState = PQresultErrorField(res, PG_DIAG_SQLSTATE);
        // catch Non existent User ID
        if (sqlState && strcmp(sqlState, "23503") == 0) {
            PQclear(res);
            free((char*)buffer);
            return ATMDBStatusAccountsUserNotFound;
        }
        // catch Bad Account type
        else if (sqlState && strcmp(sqlState, "23514") == 0) {
            PQclear(res);
            free((char*)buffer);
            return ATMDBStatusAccountsBadType;
        }
        fprintf(stderr, "Query execution failed: %s", PQresultErrorMessage(res));
        PQclear(res);
        free((char*)buffer);
        return ATMDBStatusNotOK;
    }
    PQclear(res);
    free((char*)buffer);
    return  ATMDBStatusOK;
}

struct Account** _fill_accounts(PGresult *res,  int * _Nonnull count) {
    *count = PQntuples(res);
    struct Account **accounts = calloc(*count, sizeof(struct Account*));
    if (accounts == NULL) {
        // Handle memory allocation failure
        return NULL;
    }

    for (int row = 0; row < *count; row++) {
        accounts[row] = calloc(1, sizeof(struct Account));
        if (accounts[row] == NULL) {
            // Handle memory allocation failure
            return NULL;
        }

        u_int64_t id = charToUll(PQgetvalue(res, row, 0));
        if (id < 0) {
            puts("ID is not valid");
            // Handle the error
            return NULL;
        }
        accounts[row]->id = id;

        AccountType type = get_accountType(PQgetvalue(res, row, 1));
        if (type == NULL) {
            puts("Can't get account type");
            // Handle the error
            return NULL;
        }
        accounts[row]->type = type;

        strncpy(accounts[row]->date, PQgetvalue(res, row, 2), 10);

        double balance = charToDouble(PQgetvalue(res, row, 3));
        if (balance < 0) {
            puts("Can't convert balance to double");
            // Handle the error
            return NULL;
        }
        accounts[row]->balance = balance;
    }

    return accounts;
}

enum ATMDBStatus get_accounts(struct User * _Nonnull u,  struct Account * _Nonnull * _Nonnull * _Nonnull accs,__unused int * _Nonnull count) {
    if (u == NULL) {return ATMDBStatusUsersNull;}
    const char *userQuery = getQuery(QueryHelperGetAccounts);
    // uint64 to string
    const char* buffer = ullToChar(u->id);
    if (buffer == NULL) { return ATMDBStatusNotOK;}
    
    const char *params[1] = {buffer};
    PGresult *res = PQexecParams(conn, userQuery, 1, NULL, params, NULL, NULL, 0);

    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        fprintf(stderr, "Query execution failed: %s", PQresultErrorMessage(res));
        PQclear(res);
        free((char*)buffer);
        return  ATMDBStatusNotOK;
    }
    *accs = _fill_accounts(res, count);
    PQclear(res);
    return ATMDBStatusOK;
}

enum ATMDBStatus delete_account(struct Account * _Nonnull acc) {
if (acc == NULL) {return ATMDBStatusAccountsNull;}
    const char *userQuery = getQuery(QueryHelperDeleteAccount);
    // uint64 to string
    const char* buffer = ullToChar(acc->id);
    if (buffer == NULL) { return ATMDBStatusNotOK;}
    
    const char *params[1] = {buffer};
    PGresult *res = PQexecParams(conn, userQuery, 1, NULL, params, NULL, NULL, 0);

    if (PQresultStatus(res) != PGRES_COMMAND_OK) {
        fprintf(stderr, "Query execution failed: %s", PQresultErrorMessage(res));
        PQclear(res);
        free((char*)buffer);
        return  ATMDBStatusNotOK;
    }

    return  ATMDBStatusOK;
}

enum ATMDBStatus transfer_account(struct Account * _Nonnull acc, struct User * _Nonnull receiver) {
    const char *userQuery = getQuery(QueryHelperTransferAccount);
    const char* recID = ullToChar(receiver->id);
    const char* accID = ullToChar(acc->id);

    const char *params[2] = {recID, accID};

    PGresult *res = PQexecParams(conn, userQuery, 2, NULL, params, NULL, NULL, 0);
    
    PQclear(res);
    free((char*)userQuery);
    free((char*)recID);
    free((char*)accID);
    return ATMDBStatusOK;
}

enum ATMDBStatus confirm_transfer_account(struct AccountTransfer * _Nonnull at) {
    const char *userQuery = getQuery(QueryHelperConfirmTransferAccount);
    // uint64 to string
    const char* buffer = ullToChar(at->id);
    if (buffer == NULL) { return ATMDBStatusNotOK;}
    
    const char *params[1] = {buffer};
    PGresult *res = PQexecParams(conn, userQuery, 1, NULL, params, NULL, NULL, 0);

    if (PQresultStatus(res) != PGRES_COMMAND_OK) {
        fprintf(stderr, "Query execution failed: %s", PQresultErrorMessage(res));
        PQclear(res);
        free((char*)buffer);
        return  ATMDBStatusNotOK;
    }
    return  ATMDBStatusOK;
}