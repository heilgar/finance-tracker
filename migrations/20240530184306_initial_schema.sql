-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
    ID SERIAL PRIMARY KEY,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE Accounts (
    ID SERIAL PRIMARY KEY,
    user_id INT,
    account_name VARCHAR(100) NOT NULL,
    account_type VARCHAR(10) NOT NULL CHECK (account_type IN ('Bank', 'Debit', 'Credit', 'Savings', 'Cash', 'Crypto')),
    balance DECIMAL(15, 2) DEFAULT 0.00,
    currency VARCHAR(10) DEFAULT 'USD',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(ID)
);

CREATE TABLE Categories (
    ID SERIAL PRIMARY KEY,
    user_id INT,
    category_name VARCHAR(50) NOT NULL,
    category_type VARCHAR(10) NOT NULL CHECK (category_type IN ('Income', 'Expense', 'Transfer')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(ID)
);

CREATE TABLE Transactions (
    ID SERIAL PRIMARY KEY,
    user_id INT,
    account_id INT,
    category_id INT,
    transaction_type VARCHAR(10) NOT NULL CHECK (transaction_type IN ('Income', 'Expense', 'Transfer')),
    amount DECIMAL(15, 2) NOT NULL,
    transaction_date DATE NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(ID),
    FOREIGN KEY (account_id) REFERENCES Accounts(ID),
    FOREIGN KEY (category_id) REFERENCES Categories(ID)
);

CREATE TABLE Assets (
    ID SERIAL PRIMARY KEY,
    user_id INT,
    account_id INT,
    asset_name VARCHAR(50) NOT NULL,
    quantity DECIMAL(15, 8) NOT NULL,
    purchase_price DECIMAL(15, 2) NOT NULL,
    current_price DECIMAL(15, 2),
    purchase_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(ID),
    FOREIGN KEY (account_id) REFERENCES Accounts(ID)
);

CREATE TABLE CryptoTrades (
    ID SERIAL PRIMARY KEY,
    user_id INT,
    account_id INT,
    asset_id INT,
    trade_type VARCHAR(10) NOT NULL CHECK (trade_type IN ('Buy', 'Sell')),
    quantity DECIMAL(15, 8) NOT NULL,
    price_per_unit DECIMAL(15, 2) NOT NULL,
    trade_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(ID),
    FOREIGN KEY (account_id) REFERENCES Accounts(ID),
    FOREIGN KEY (asset_id) REFERENCES Assets(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE CryptoTrades;
DROP TABLE Assets; 
DROP TABLE Transactions;
DROP TABLE Categories;
DROP TABLE Accounts;
DROP TABLE Users;
-- +goose StatementEnd
