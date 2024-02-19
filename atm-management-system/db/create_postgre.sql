CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL,
    active BOOLEAN DEFAULT TRUE NOT NULL
);

CREATE TABLE Accounts (
    account_id SERIAL PRIMARY KEY,
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
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    balance DOUBLE PRECISION NOT NULL DEFAULT 0,  -- Use DOUBLE PRECISION for floating-point values in PostgreSQL.
    country VARCHAR(19) NOT NULL,  -- Use CHAR(3) for fixed-length strings.
    phone VARCHAR(9) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    active BOOLEAN DEFAULT TRUE NOT NULL
);

CREATE TABLE Accounts_Transfers (
    transfer_id SERIAL PRIMARY KEY,
    receiver_user_id INTEGER NOT NULL,
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    transferred_account_id INTEGER NOT NULL,
    FOREIGN KEY (receiver_user_id) REFERENCES Users(user_id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    FOREIGN KEY (transferred_account_id) REFERENCES Accounts(account_id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE Transactions (
    transaction_id SERIAL PRIMARY KEY,
    sender_account_id INTEGER NOT NULL,
    receiver_account_id INTEGER NOT NULL,
    amount DOUBLE PRECISION NOT NULL,  -- Use DOUBLE PRECISION for floating-point values in PostgreSQL.
    date DATE NOT NULL,
    FOREIGN KEY (sender_account_id) REFERENCES Accounts(account_id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    FOREIGN KEY (receiver_account_id) REFERENCES Accounts(account_id) ON DELETE NO ACTION ON UPDATE NO ACTION
);


ALTER TABLE Users OWNER TO natheer;
ALTER TABLE Accounts OWNER TO natheer;
ALTER TABLE Accounts_Transfers OWNER TO natheer;
ALTER TABLE Transactions OWNER TO natheer;

-- add default users
INSERT INTO Users (name, password) VALUES ('Alice', 'q1w2e3r4t5y6');
INSERT INTO Users (name, password) VALUES ('Michel', 'q1w2e3r4t5y6');

-- insert default accounts
INSERT INTO Accounts (user_id, type, balance, country, phone) VALUES (0, 'savings',22432.52 , 'Africa', '291321234');
INSERT INTO Accounts (user_id, type, balance, country, phone) VALUES (1, 'fixed01', 10023.230000, 'Portugal', '123543455');
INSERT INTO Accounts (user_id, type, balance, country, phone) VALUES (1, 'saving', 12345.30, 'UK', '1234123');