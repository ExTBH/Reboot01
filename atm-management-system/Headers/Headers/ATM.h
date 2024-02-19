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
  unsigned long id;
  char name[51];
  char password[65];
  bool active;
  struct Account **accounts;
};

struct Account {
  unsigned long id;
  struct User *user;
  AccountType type;
  char date[11]; // 31-10-2023 + '\0'
  double balance;
  char country[20];
  char phone[10];
};

struct AccountTransfer {
  unsigned long id;
  unsigned long receiver;
  unsigned long acc_to_transfer;
  unsigned long date;
};

struct Transaction {
  unsigned long id;
  struct Account *sender;
  struct Account *receiver;
  double amount;
  unsigned long date;
};

#endif