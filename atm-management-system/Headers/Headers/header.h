#include "Headers/ATM.h"
#include <stdbool.h>


// authentication functions
void loginMenu(struct User *u);
void readPassword(char pass[50]);
// void registration(char a[50], char pass[50]);
// void registerMenu(char a[50], char pass[50]);
// const char *getPassword(struct User u);

// system function
void registerUser();
void mainMenu(struct User u);
void checkAccountDetails(struct User* u);
void checkAllAccounts(struct User* u);
void makeTransaction(struct User* u);
void updateAccountInfo(struct User* u);
void removeAccount(struct User* u);
void transferOwnership(struct User* u);

void createNewAcc(struct User *u);
void success(struct User u);
void mainOrExit(struct User u);

// file
char *readString(char *print);
double readBalance(char *prompt);
char *readPhoneNum(char *prompt);

char* readDate();
void flushInputBuffer();
void printInterestAmount(struct Account acc);
