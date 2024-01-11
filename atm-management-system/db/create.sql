CREATE TABLE Users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(64) NOT NULL,
    active BOOLEAN DEFAULT 1 NOT NULL
);
CREATE TABLE Accounts (
    account_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    type VARCHAR(7) NOT NULL CHECK (
        type IN (
            'fixed01',
            'fixed02',
            'fixed03',
            'current',
            'savings'
        )
    ),
    date DATE NOT NULL,
    balance DOUBLE NOT NULL,
    country VARCHAR(3) NOT NULL CHECK (LENGTH(country) = 3),
    phone VARCHAR(10) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE
);
CREATE TABLE Accounts_Transfers (
    transfer_id INTEGER PRIMARY KEY AUTOINCREMENT,
    receiver_user_id INTEGER NOT NULL,
    date DATE NOT NULL,
    transfered_account_id INTEGER NOT NULL,
    FOREIGN KEY (receiver_user_id) REFERENCES Users(user_id) ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE,
    FOREIGN KEY (transfered_account_id) REFERENCES Accounts (account_id) ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE
);
CREATE TABLE Transactions (
    transaction_id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender_account_id INTEGER NOT NULL,
    receiver_account_id INTEGER NOT NULL,
    amount DOUBLE NOT NULL,
    date DATE NOT NULL,
    FOREIGN KEY (sender_account_id) REFERENCES Accounts(account_id) ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE,
    FOREIGN KEY (receiver_account_id) REFERENCES Accounts (account_id) ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE
);