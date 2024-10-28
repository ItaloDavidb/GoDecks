-- Criação da tabela users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Criação da tabela cards
CREATE TABLE cards (
    set_code VARCHAR NOT NULL,
    number VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    type VARCHAR,
    json_data TEXT,
    PRIMARY KEY (set_code, number)
);

-- Criação da tabela user_cards
CREATE TABLE user_cards (
    user_id INT NOT NULL,
    set_code VARCHAR NOT NULL,
    number VARCHAR NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, set_code, number),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (set_code, number) REFERENCES cards(set_code, number) ON DELETE CASCADE
);

-- Criação da tabela decks
CREATE TABLE decks (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Criação da tabela deck_cards
CREATE TABLE deck_cards (
    deck_id INT NOT NULL,
    set_code VARCHAR NOT NULL,
    number VARCHAR NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (deck_id, set_code, number),
    FOREIGN KEY (deck_id) REFERENCES decks(id) ON DELETE CASCADE,
    FOREIGN KEY (set_code, number) REFERENCES cards(set_code, number) ON DELETE CASCADE
);
