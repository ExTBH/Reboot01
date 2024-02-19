# Soruce Documentation

## HashingHelper.c
```c
enum HashingHelperErorr calc_sha256(const char * _Nonnull path, char output[_Nonnull 65]);
```
Calculates the sha256 of a given file at path, currently implemented only for OS X.

## RandomHelpers.c
```c
long getFileSize(const char * _Nonnull path);
```
Gets the size of a file at a given path, returns `-1` on error.

```c
bool isAlpha(const char *_Nonnull buff);
```
Checks if the given array is only alpha, forumla is `x > 32 && x < 126`.

```c
bool isValidName(const char *_Nonnull buff);
```
checks if a string is `[0-9a-zA-Z]`.

```c
_Nullable AccountType get_accountType(const char * _Nonnull type);
```
converts `type` to  `AccountType` or `NULL`.

```c
const char * _Nullable ullToChar(u_int64_t n);
```
convert a `u_int64_t` to a string, returns a `char[21]` or `NULL` and you should free it.

```c
u_int64_t charToUll(const char * _Nonnull s);
```
inverse of ullToChar, -1 in errors.

```c
double charToDouble(const char * _Nonnull s);
```
converts a string to a double, `-1` for errors

## databaseUtil.c
```c
enum ATMDBStatus atmdb_connect();
```
Attempts to connect to the database, must be called.

```c
void atmdb_close(void);
```
closes the database on a successfull connection, call it at the end of `main.c`

```c
static enum ATMDBStatus _verify(void);
```
unexported function called from `atmdb_connect()` to check the database strucure is valid.

```c
int count(const PGresult *res);
```
used to count results from queries, `-1` for erros.

```c
enum ATMDBStatus _fill_user(const PGresult * _Nonnull res, struct User * _Nonnull u);
```
just fills a user from a `res`, it doesnt validate it.

```c
enum ATMDBStatus get_user(struct User * _Nonnull u, bool withID);
```
a function to retrieve a user from the database taking a user structm if `withID` is `true` it will get the user with the id

```c
enum ATMDBStatus _get_user(__unused struct User * _Nonnull u, const char *buff,  const char *userQuery);
```
unexported function called from `get_user()` to not duplicate code.

```c
enum ATMDBStatus add_user(struct User * _Nonnull u);
```
Adds a new user to the database, i think it handles duplicated names and such.

```c
enum ATMDBStatus add_account(struct Account * _Nonnull acc);
```
Adds an account to the user at `acc->user->id`, think it handles errors


```c
enum ATMDBStatus get_accounts(struct User * _Nonnull u,  struct Account * _Nonnull * _Nonnull * _Nonnull accs,__unused int * _Nonnull count);
```
takes a pointer to an int `count` and an array of Accounts `accs` and gets all accounts for the user `u` and fills them in the array and update the `count`

```c
struct Account** _fill_accounts(PGresult *res,  int * _Nonnull count);
```
unexported function that takes a `res` (rows of accounts from the db), and returns an array of `Account` of them. used in `get_accounts()`

```c
enum ATMDBStatus delete_account(struct Account * _Nonnull acc);
```
deletes an account from the database

```c
enum ATMDBStatus transfer_account(struct Account * _Nonnull acc, struct User * _Nonnull receiver);
```
adds a new transfer row to a new user

```c
enum ATMDBStatus confirm_transfer_account(struct AccountTransfer * _Nonnull at);
```
used to confirm a transfer request