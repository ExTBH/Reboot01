#include "Headers/ATM.h"
#include "Headers/header.h"
#include <Headers/databaseUtil.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void initMenu(struct User *u) {
  int option;
  system("clear");
  printf("\n\n\t\t======= ATM =======\n");
  printf("\n\t\t-->> Feel free to login / register :\n");
  printf("\n\t\t[1]- login\n");
  printf("\n\t\t[2]- register\n");
  printf("\n\t\t[3]- exit\n");
  printf("\n\n");
  while (scanf("%d", &option) != 1 || option < 1 || option > 3) {
    printf("Invalid input. Please enter a number between 1 and 3.\n");
    // Clear input buffer
    int c;
    while ((c = getchar()) != '\n' && c != EOF) {
      // Consume characters from input buffer until newline or EOF
    }
    // Prompt again for input
    printf("Choose an option (1-3): ");
  }

  switch (option) {
  case 1:
    loginMenu(u);
    char inputPass[65];
    enum ATMDBStatus status;
    strncpy(inputPass, u->password, 64);

    // userInfo = *u; // Copy the user information
    status = get_user(u, false);

    if (status != ATMDBStatusOK) {
      // Handle error
      // printf("Error retrieving user: %d\n", status);
      printf("\nWrong password!! or User Name\n");
      exit(EXIT_FAILURE);
    }
    // User retrieved successfully
    if (strncmp(u->password, inputPass, 64) == 0) {
      // strncmp(const char *, const char *, unsigned long)
      printf("\n\nPassword Match!\n");
    } else {
      printf("\nWrong password!! or User Name\n");
      exit(EXIT_FAILURE);
    }

    // printf("\t\t\t===== ID11 string %lu =====\n", u->id);
    // exit(EXIT_FAILURE);

    break;
  case 2:
    // student TODO : add your **Registration** function
    registerUser();
    break;

  case 3:
    system("clear");
    atmdb_close();
    exit(EXIT_SUCCESS);
    break;

  default:
    printf("Insert a valid operation!\n");
    break;
  }
};

void mainMenu(struct User u) {
  int option;
  system("clear");
  printf("\n\n\t\t======= ATM =======\n\n");
  printf("\n\t\t-->> Feel free to choose one of the options below <<--\n");
  printf("\n\t\t[1]- Create a new account\n");
  printf("\n\t\t[2]- Update account information\n");
  printf("\n\t\t[3]- Check accounts\n");
  printf("\n\t\t[4]- Check list of owned account\n");
  printf("\n\t\t[5]- Make Transaction\n");
  printf("\n\t\t[6]- Remove existing account\n");
  printf("\n\t\t[7]- Transfer ownership\n");
  printf("\n\t\t[8]- Exit\n");

  //   scanf("%d", &option);
  printf("\n\nChoose an option (1-8): ");
  while (scanf("%d", &option) != 1 || option < 1 || option > 8) {
    printf("Invalid input. Please enter a number between 1 and 8.\n");
    // Clear input buffer
    int c;
    while ((c = getchar()) != '\n' && c != EOF) {
      // Consume characters from input buffer until newline or EOF
    }
    // Prompt again for input
    printf("Choose an option (1-8): ");
  }

  //   flushInputBuffer();

  switch (option) {
  case 1:
    createNewAcc(&u);
    // add_account();
    break;
  case 2:
    // student TODO : add your **Update account information** function
    updateAccountInfo(&u);
    // here
    break;
  case 3:
    // student TODO : add your **Check the details of existing accounts**
    // function
    checkAccountDetails(&u);
    // here
    break;
  case 4:
    checkAllAccounts(&u);
    break;
  case 5:
    // student TODO : add your **Make transaction** function
    makeTransaction(&u);
    break;
  case 6:
    // student TODO : add your **Remove existing account** function
    removeAccount(&u);
    break;
  case 7:
    // student TODO : add your **Transfer owner** function
    transferOwnership(&u);
    break;
  case 8:
    exit(1);
    break;
  default:
    printf("Invalid operation!\n");
  }
};

int main() {
  enum ATMDBStatus rc = atmdb_connect();
  if (rc == ATMDBStatusCantConnect) {
    printf("Failed to connect to database\n");
    return EXIT_FAILURE;
  } else if (rc == ATMDBStatusCantVerify) {
    printf("Failed to verify database\n");
    return EXIT_FAILURE;
  }
  // printf("Choose an option (1-3): ");
  // int selectedOptions;

  // scanf("%d", &selectedOptions);
  // printf("You Selected Options: %d\n", selectedOptions);
  // if (selectedOptions < 1 || selectedOptions > 3) {
  //     puts("Invalid option sleected ");
  //     return  EXIT_FAILURE;
  // }
  // struct User newUser1 = {213, "aj", "aj123", false, NULL};
  // rc = add_user(&newUser1);
  // if (rc == ATMDBStatusUsersBadName) {
  //     puts("name is not valid");
  // } else if (rc == ATMDBStatusUsersConstraintName) {
  //     puts("username already exists");
  // }
  struct User u;
  initMenu(&u);

  mainMenu(u);

  atmdb_close();
  return EXIT_SUCCESS;
}

// This function is used to flush the input buffer
void flushInputBuffer() {
  int c;
  while ((c = getchar()) != '\n' && c != EOF) {
  } // Keep reading characters until a newline or end-of-file
}