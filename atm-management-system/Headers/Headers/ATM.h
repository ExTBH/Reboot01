#ifndef ATM_H
#define ATM_H

#include <stdbool.h>
#include <stdint.h>

typedef const char *AccountType;

static AccountType accountTypeFixed01 = "fixed01";
static AccountType accountTypeFixed02 = "fixed02";
static AccountType accountTypeFixed03 = "fixed03";
static AccountType accountTypeSavings = "savings";
static AccountType accountTypeCurrent = "current";

struct User {
    uint64_t id;
    char name[51];
    char password[65];
    bool active;
    struct Account **accounts;
};

struct Account {
    uint64_t id;
    struct User *user;
    AccountType type;
    char date[11]; // 31-10-2023 + '\0'
    double balance;
    char country[20];
    char phone[10];
};

struct AccountTransfer {
    uint64_t id;
    uint64_t receiver;
    uint64_t acc_to_transfer;
    uint64_t date;
};

struct Transaction {
    uint64_t id;
    struct Account *sender;
    struct Account *receiver;
    double amount;
    uint64_t date;
};

#endif