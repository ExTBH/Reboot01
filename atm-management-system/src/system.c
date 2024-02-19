#include "Headers/ATM.h"
#include "Headers/RandomHelpers.h"
#include "Headers/databaseUtil.h"
#include "Headers/header.h"
#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

const char *readAccountType();
double TransectionAmount();

char *readString(char *print) {
  char str[50];

  while (1) {
     flushInputBuffer(); // Clear the input buffer

    printf("%s", print);
    if (scanf("%49s", str) == 1) {
      // Input was successful
      // Clear the input buffer to remove any remaining characters
    // flushInputBuffer(); // Clear the input buffer

      return strdup(str); // Return a dynamically allocated copy of the string
    } else {
      // Input was not successful
      printf("Invalid input. Please try again: ");
      
    }
    flushInputBuffer(); // Clear the input buffer
  } // Loop until a valid input is provided
}
void success(struct User u) {
  printf("\n✔ Success!\n\n");
  mainOrExit(u);
}

bool isValidUserName(char *username) {
  size_t length = strlen(username); // Use size_t for the length

  // Cast the size_t length to an int for comparison
  for (int i = 0; i < (int)length; i++) {
    if (!isalpha(username[i]) && username[i] != ' ') {
      return false;
    }
  }
  return true;
}

void registerUser() {
  struct User newUser;
  enum ATMDBStatus status;
  system("clear");
  while (1) {
    printf("\t\t\t===== Register =====\n");
    // Prompt and read the username
    char *username = readString("\n\n\n\n\n\t\t\t\tEnter username: ");
    strcpy(newUser.name, username);
    free(username); // Free the dynamically allocated memory
    if (!isValidUserName(newUser.name)) {
      printf("Account creation failed. Invalid username. Please enter a valid "
             "username.\n");
      continue;
    }

    // Prompt and read the password
    char password[50];
    readPassword(password);
    strcpy(newUser.password, password);
    status = add_user(&newUser);
    if (status == ATMDBStatusUsersBadName) {
      puts("name is not valid");
    } else if (status == ATMDBStatusUsersConstraintName) {
      puts("username already exists");
    } else {
      break;
    }
  }
  success(newUser);
}

void mainOrExit(struct User u) {
  int option;
  while (1) {
    printf("Enter 1 to go to the main menu and 0 to exit!\n");
    scanf("%d", &option);
    flushInputBuffer();

    if (option == 1) {
      mainMenu(u);
      break; // Exit the loop after mainMenu() is executed
    } else if (option == 0) {
      exit(1);
    } else {
      printf("Insert a valid operation!\n");
    }
  }
}

void createNewAcc(struct User *u) {
  struct Account newAcc;
  bool creationResult;
  enum ATMDBStatus status;
  newAcc.user = u;
  system("clear");

  do {
    printf("\t\t\t===== New Account =====\n");

    // Prompt and read the account type
    newAcc.type = readAccountType();

    newAcc.balance = readBalance("\nEnter amount to deposit: $");

    // Prompt and read the Date
    // char *date = readDate();
    // strcpy(newAcc.date, date);
    // free(date); // Free the dynamically allocated memory

    // Prompt and read the username
    char *coun = readString("\nEnter the country: ");
    strcpy(newAcc.country, coun);
    free(coun); // Free the dynamically allocated memory

    // newAcc.phone = readPhoneNum("\nEnter the phone number:");
    char *phoneNum = readPhoneNum("\nEnter the phone number:");
    strcpy(newAcc.phone, phoneNum);
    // free(phoneNum); // Free the dynamically allocated memory

    // creationResult = sql_create_account(*u, newAcc);

    status = add_account(&newAcc);
    printf("\t\t\t===== HERE =====\n");

    if (status != ATMDBStatusOK) {
      printf("Account creation failed. Please try again.\n %d \n", status);
    } else {
      creationResult = true;
    }
  } while (!creationResult);
  success(*u);
}

const char *readAccountType() {
  int choice;

  printf("Select Account Type:\n");
  printf("1. Fixed01\n");
  printf("2. Fixed02\n");
  printf("3. Fixed03\n");
  printf("4. Savings\n");
  printf("5. Current\n");

  printf("Enter your choice: ");

  while (1) {
     flushInputBuffer();
    scanf("%d", &choice);

    switch (choice) {
    case 1:
      return "fixed01";
    case 2:
      return "fixed02";
    case 3:
      return "fixed03";
    case 4:
      return "savings";
    case 5:
      return "current";
    default:
      printf("Invalid account type. Please enter a valid choice.\n");
    }
  }
}

double readBalance(char *prompt) {
  double value;
  int result;

  while (1) {
     flushInputBuffer();
    printf("%s", prompt);
    result = scanf("%lf", &value);

    // Check if the input was successfully converted to a double
    if (result == 1) {
      // Check if the value is a positive number with at most two decimal places
      if (value <= 0) {
        printf("Invalid input. Please enter a valid positive number.\n");
      } else if (value > 9999999) {
        printf("Amount is too large.\n");
        // } else if (hasMoreThanTwoDecimalPlaces(value)) {
        //     printf("Invalid balance. Please enter a valid number with at most
        //     two decimal places.\n");
      } else {
        break; // Valid input, break out of the loop
      }
    } else {
      printf("Invalid input. Please enter a valid number.\n");
    }
  }

  return value;
}

// char *readDate() {
//   struct Date {
//     int day;
//     int month;
//     int year;
//   };

//   struct Date date;

//   while (1) {
//     printf("Enter date (dd/mm/yyyy): ");
//     if (scanf("%d/%d/%d", &date.day, &date.month, &date.year) == 3) {
//       // Check if the input values are within valid ranges
//       if (date.day >= 1 && date.day <= 31 && date.month >= 1 &&
//           date.month <= 12 && date.year >= 1900 && date.year <= 2100) {
//         break; // Valid input, break out of the loop
//       } else {
//         printf("Invalid date. Please try again.\n");
//       }
//     } else {
//       printf("Invalid input format. Please use dd/mm/yyyy format.\n");
//       flushInputBuffer();
//     }
//   }

//   // Convert the date to the desired format "%Y/%m/%d"
//   char formattedDate[11];
//   sprintf(formattedDate, "%04d/%02d/%02d", date.year, date.month, date.day);

//   return strdup(formattedDate);
// }

char *readPhoneNum(char *prompt) {
  char *phone =
      malloc(11 * sizeof(char)); // Allocate memory for the phone number
  if (phone == NULL) {
    // Handle memory allocation failure
    fprintf(stderr, "Failed to allocate memory for phone number\n");
    exit(EXIT_FAILURE);
  }

  printf("%s ", prompt);
  if (scanf("%10s", phone) != 1) {
    fprintf(stderr, "Error reading phone number\n");
    exit(EXIT_FAILURE);
  }

  // Check if the input is a valid phone number
  size_t len = strlen(phone);
  if (len < 7 || len > 10) {
    fprintf(stderr,
            "Invalid input. Phone number must be between 7 and 10 digits.\n");
    exit(EXIT_FAILURE);
  }

  for (size_t i = 0; i < len; i++) {
    if (!isdigit(phone[i])) {
      fprintf(stderr,
              "Invalid input. Phone number must contain only digits.\n");
      exit(EXIT_FAILURE);
    }
  }

  return phone;
}
void checkAccountDetails(struct User *u) {
  char input[50];
  system("clear");
  while (1) {
    printf("\t\t\t===== Check Account Details =====\n");
    printf("Enter the account ID you want to check or \\back to return: ");
    scanf("%s", input);
    if (strcmp(input, "\\back") == 0) {
      mainOrExit(*u);
      return;
    }

    struct Account **accs;
    // struct Account *account;
    int index = -1;
    int accs_count;

    // Call get_accounts to retrieve the accounts
    enum ATMDBStatus status = get_accounts(u, &accs, &accs_count);
    if (status != ATMDBStatusOK) {
      // HANDLE ERROR
      exit(1);
    }
    if (accs_count == 0) {
      printf("You dont have any accounts");
      continue;
    }
    for (int i = 0; i < accs_count; i++) {
      if (charToUll(input) == accs[i]->id) {
        index = i;
        break;
      }
    }
    if (index == -1) {
      printf("Account not found.\n");
      continue;
    }

    // if (accs[index]->country == ) {
    //   printf("Account country not found.\n");
    //   continue;
    // }

    // if (accs[index]->user->id != u->id) {
    //   printf("The account does not belong to the user.\n");
    //   continue;
    // }
    // Print account details
    // strcpy(accs[index]->country, "test");
    printf("Account ID: %ld\n", accs[index]->id);
    printf("Country: %s\n", accs[index]->country);
    printf("Phone number: %s\n", accs[index]->phone);
    printf("Amount deposited: $%.2lf\n", accs[index]->balance);
    printf("Account type: %s\n", accs[index]->type);
    printf("Creation Date: %s\n", accs[index]->date);
    printInterestAmount(*accs[index]);
  }
}

void printInterestAmount(struct Account acc) {
  double interestAmount = 0.0;

  if (acc.type == accountTypeSavings) {
    // Calculate monthly interest
    interestAmount = (acc.balance * 0.07 / 12);
    printf("You will gain $%.2lf of interest on day 10 of every month.\n",
           interestAmount);
  } else if (acc.type == accountTypeFixed01) {
    // Calculate interest for one year from account creation date
    // Extracting year, month, and day
    int year = atoi(acc.date);
    int month = atoi(acc.date + 5); // Skip the first 5 characters (year-) to get the month
    int day = atoi(acc.date + 8);   // Skip the first 8 characters (year-month-) to get the day

    // Adding one year to the date
    year++;
    interestAmount = (acc.balance * 0.04);
    printf("You will gain $%.2lf interest on %02d/%02d/%04d (one year from account creation).\n",interestAmount, day, month, year);
  } else if (acc.type == accountTypeFixed02) {
    // Calculate interest for two years from account creation date
 // Extracting year, month, and day
    int year = atoi(acc.date);
    int month = atoi(acc.date + 5); // Skip the first 5 characters (year-) to get the month
    int day = atoi(acc.date + 8);   // Skip the first 8 characters (year-month-) to get the day

    // Adding one year to the date
    year+=2;

    interestAmount = (acc.balance * 0.05 * 2);
    printf("You will gain $%.2lf interest on %02d/%02d/%04d (two years from "
           "account creation).\n",
           interestAmount, day, month, year);
  } else if (acc.type == accountTypeFixed03) {
    // Calculate interest for three years from account creation date
    int year = atoi(acc.date);
    int month = atoi(acc.date + 5); // Skip the first 5 characters (year-) to get the month
    int day = atoi(acc.date + 8);   // Skip the first 8 characters (year-month-) to get the day

    // Adding one year to the date
    year+=3;

    interestAmount = (acc.balance * 0.08 * 3);
    printf("You will gain $%.2lf interest on %02d/%02d/%04d (three years from "
           "account creation).\n",
           interestAmount, day, month, year);
  } else {
    printf(
        "You will not get interests because the account is of type current.\n");
  }
}

void checkAllAccounts(struct User *u) {
  system("clear");
  struct Account **accs;
  int accs_count;

  // Call get_accounts to retrieve the accounts
  enum ATMDBStatus status = get_accounts(u, &accs, &accs_count);
  if (status != ATMDBStatusOK) {
    // HANDLE ERROR
    exit(1);
  }
  if (accs_count == 0) {
    printf("You dont have any accounts");
  } else {
    printf("Owned Account IDs:\n");
    for (int i = 0; i < accs_count; i++) {
      printf("Account ID: %lu\n", accs[i]->id);
    }
  }

  success(*u);
}

void makeTransaction(struct User *u) {
  char input[50];
  system("clear");
  while (1) {
    printf("\t\t\t===== Make Transaction =====\n");
    printf("Enter the account ID for the transaction or \\back to return: ");
    scanf("%s", input);
    if (strcmp(input, "\\back") == 0) {
      mainOrExit(*u);
      return;
    }

    struct Account **accs;
    // struct Account *account;
    int index = -1;
    int accs_count;

    // Call get_accounts to retrieve the accounts
    enum ATMDBStatus status = get_accounts(u, &accs, &accs_count);
    if (status != ATMDBStatusOK) {
      // HANDLE ERROR
      exit(1);
    }
    if (accs_count == 0) {
      printf("You dont have any accounts");
      continue;
    }
    for (int i = 0; i < accs_count; i++) {
      if (charToUll(input) == accs[i]->id) {
        index = i;
        break;
      }
    }

    if (index == -1) {
      printf("Account not found.\n");
      continue;
    }

    // if (accs[index]->user == NULL) {
    //   printf("Account not found.\n");
    //   continue;
    // }

    // if (accs[index]->user->id != u->id) {
    //   printf("The account does not belong to the user.\n");
    //   continue;
    // }

    // Check if the account type disallows transactions
    if (accs[index]->type == accountTypeFixed01 ||
        accs[index]->type == accountTypeFixed02 ||
        accs[index]->type == accountTypeFixed03) {
      printf("Transactions are only allowed for accounts of type saving and "
             "current.\n");
      continue;
    }

    double amount = TransectionAmount();
    if (accs[index]->balance + amount < 0) {
      printf("Transaction failed. Insufficient funds.\n");
      continue;
    }
    double newValue = accs[index]->balance + amount;
    char newValue_str[20];
    sprintf(newValue_str, "%.2lf", newValue);
    status =
        update_account_balance(accs[index]->id, charToDouble(newValue_str));
    if (status != ATMDBStatusOK) {
      // HANDLE ERROR
      printf("Transaction failed. Please try again.\n");
      // exit(1);
    } else {
      break;
    }
  }
  success(*u);
}

double TransectionAmount() {
  int choice;

  printf("Select an option:\n");
  printf("1. Deposit\n");
  printf("2. Withdraw\n");

  printf("Enter your choice: ");

  while (1) {
    flushInputBuffer();
    scanf("%d", &choice);

    if (choice == 1 || choice == 2) {
      double amount =
          readBalance(choice == 1 ? "Enter the amount to deposit: "
                                  : "Enter the amount to withdraw: ");
      return choice == 1 ? amount : -amount;
    } else {
      printf("Invalid option. Please enter a valid choice.\n");
    }
  }
}

void updateAccountInfo(struct User *u) {
  char input[50];
  system("clear");
  while (1) {
    printf("\t\t\t===== Update Account Info =====\n");
    printf("Enter the account ID you want to update or \\back to return: ");
    scanf("%s", input);
    if (strcmp(input, "\\back") == 0) {
      mainOrExit(*u);
      return;
    }

    struct Account **accs;
    // struct Account *account;
    int index = -1;
    int accs_count;

    // Call get_accounts to retrieve the accounts
    enum ATMDBStatus status = get_accounts(u, &accs, &accs_count);
    if (status != ATMDBStatusOK) {
      // HANDLE ERROR
      exit(1);
    }
    if (accs_count == 0) {
      printf("You dont have any accounts");
      continue;
    }

    for (int i = 0; i < accs_count; i++) {
      if (charToUll(input) == accs[i]->id) {
        // printf("input %s || acc ID %lu.\n", input, accs[i]->id);

        index = i;
        break;
      }
    }
    if (index == -1) {
      printf("Account not found.\n");
      continue;
    }
    // if (accs[index]->user == NULL) {
    //   printf("Account not found.\n");
    //   continue;
    // }

    // if (accs[index]->user->id != u->id) {
    //   printf("The account does not belong to the user.\n");
    //   continue;
    // }

    char *newCountry = readString("\nEnter updated country: ");

    char *newPhoneNum = readPhoneNum("\nEnter updated phone number: ");

    status = update_account(accs[index]->id, newCountry, newPhoneNum);
    if (status != ATMDBStatusOK) {
      // HANDLE ERROR
      printf("update account failed. Please try again.\n");
      // exit(1);
    } else {
      break;
    }
  }
  success(*u);
}

void removeAccount(struct User *u) {
  char input[50];
  system("clear");
  while (1) {
    printf("\t\t\t===== Remove Account =====\n");
    printf("Enter the account ID you want to remove or \\back to return: ");
    scanf("%s", input);
    if (strcmp(input, "\\back") == 0) {
      mainOrExit(*u);
      return;
    }

    struct Account **accs;
    // struct Account *account;
    int index = -1;
    int accs_count;

    // Call get_accounts to retrieve the accounts
    enum ATMDBStatus status = get_accounts(u, &accs, &accs_count);
    if (status != ATMDBStatusOK) {
      // HANDLE ERROR
      exit(1);
    }
    if (accs_count == 0) {
      printf("You dont have any accounts");
      continue;
    }

    for (int i = 0; i < accs_count; i++) {
      if (charToUll(input) == accs[i]->id) {
        // printf("input %s || acc ID %lu.\n", input, accs[i]->id);

        index = i;
        break;
      }
    }
    if (index == -1) {
      printf("Account not found.\n");
      continue;
    }

    // if (accs[index]->user == NULL) {
    //   printf("Account not found.\n");
    //   continue;
    // }

    // if (accs[index]->user->id != u->id) {
    //   printf("The account does not belong to the user.\n");
    //   continue;
    // }

    status = delete_account(accs[index]);
    if (status != ATMDBStatusOK) {
      // HANDLE ERROR
      printf("Remove account failed. Please try again.\n");
      // exit(1);
    } else {
      break;
    }
  }
  success(*u);
}

void transferOwnership(struct User *u) {
  char input[50];
  system("clear");
  while (1) {
    printf("\t\t\t===== Transfer Ownership =====\n");
    printf("Enter the account ID you want to transfer or \\back to return: ");
    scanf("%s", input);
    if (strcmp(input, "\\back") == 0) {
      mainOrExit(*u);
      return;
    }

    struct Account **accs;
    // struct Account *account;
    int index = -1;
    int accs_count;

    // Call get_accounts to retrieve the accounts
    enum ATMDBStatus status = get_accounts(u, &accs, &accs_count);
    if (status != ATMDBStatusOK) {
      // HANDLE ERROR
      exit(1);
    }
    if (accs_count == 0) {
      printf("You dont have any accounts");
      continue;
    }

    for (int i = 0; i < accs_count; i++) {
      if (charToUll(input) == accs[i]->id) {
        // printf("input %s || acc ID %lu.\n", input, accs[i]->id);

        index = i;
        break;
      }
    }
    if (index == -1) {
      printf("Account not found.\n");
      continue;
    }

    // if (accs[index]->user == NULL) {
    //   printf("Account not found.\n");
    //   continue;
    // }

    // if (accs[index]->user->id != u->id) {
    //   printf("The account does not belong to the user.\n");
    //   continue;
    // }

    struct User new_owner;
    char *newID[50];

    *newID = readString("Enter the id of the new owner: ");
    new_owner.id = charToUll(*newID);
    
    status = get_user(&new_owner, true);
    if (status != ATMDBStatusOK) {
      // Handle error
      printf("Account not found.\n");
      continue;
    }

    status = transfer_account(accs[index], &new_owner);
    if (status != ATMDBStatusOK) {
      // HANDLE ERROR
      printf("Remove account failed. Please try again.\n");
      // exit(1);
    } else {
      break;
    }
  }
  success(*u);
}
