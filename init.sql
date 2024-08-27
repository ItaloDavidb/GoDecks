-- Criação da tabela de usuários
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- Criação da tabela de cartas
CREATE TABLE cards (
    set_code VARCHAR(10),
    number VARCHAR(10),
    name VARCHAR(100),
    type VARCHAR(50),
    json_data JSONB,
    PRIMARY KEY (set_code, number)
);

-- Criação da tabela de cartas de usuário
CREATE TABLE user_cards (
    user_id INT,
    set_code VARCHAR(10),
    number VARCHAR(10),
    quantity INT DEFAULT 0,
    PRIMARY KEY (user_id, set_code, number),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (set_code, number) REFERENCES cards(set_code, number)
);

-- Criação da tabela de decks
CREATE TABLE decks (
    id SERIAL PRIMARY KEY,
    user_id INT,
    name VARCHAR(100),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Criação da tabela de cartas no deck
CREATE TABLE deck_cards (
    deck_id INT,
    set_code VARCHAR(10),
    number VARCHAR(10),
    quantity INT,
    PRIMARY KEY (deck_id, set_code, number),
    FOREIGN KEY (deck_id) REFERENCES decks(id) ON DELETE CASCADE,
    FOREIGN KEY (set_code, number) REFERENCES cards(set_code, number)
);
