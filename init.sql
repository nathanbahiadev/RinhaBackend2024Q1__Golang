BEGIN;

CREATE UNLOGGED TABLE IF NOT EXISTS CLIENTS (
    ID SERIAL PRIMARY KEY,
    ACCOUNT_LIMIT INTEGER NOT NULL DEFAULT 0,
    BALANCE INTEGER NOT NULL DEFAULT 0
);

CREATE UNLOGGED TABLE IF NOT EXISTS TRANSACTIONS (
    ID SERIAL PRIMARY KEY,
    CLIENT_ID INTEGER NOT NULL,
    VALUE INTEGER NOT NULL,
    TYPE VARCHAR(1) NOT NULL,
    DESCRIPTION VARCHAR(10) NOT NULL,
    CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (CLIENT_ID) REFERENCES CLIENTS(ID)
);

CREATE INDEX IDX_TRANSACTIONS ON TRANSACTIONS (ID DESC);
CREATE UNIQUE INDEX CLIENTS_PK ON CLIENTS(ID);
CREATE UNIQUE INDEX TRANSACTIONS_PK ON TRANSACTIONS(ID);

INSERT INTO CLIENTS (ACCOUNT_LIMIT, BALANCE) VALUES (100000, 0);
INSERT INTO CLIENTS (ACCOUNT_LIMIT, BALANCE) VALUES (80000, 0);
INSERT INTO CLIENTS (ACCOUNT_LIMIT, BALANCE) VALUES (1000000, 0);
INSERT INTO CLIENTS (ACCOUNT_LIMIT, BALANCE) VALUES (10000000, 0);
INSERT INTO CLIENTS (ACCOUNT_LIMIT, BALANCE) VALUES (500000, 0);

COMMIT;
