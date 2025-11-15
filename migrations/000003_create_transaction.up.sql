-- Drop tabel jika ada
DROP TABLE IF EXISTS transactions;

-- Buat tabel
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    campaign_id INTEGER,
    user_id INTEGER,
    amount INTEGER DEFAULT 0,
    status VARCHAR(255),
    code VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_transactions_campaign_id
    ON transactions (campaign_id);

CREATE INDEX idx_transactions_user_id
    ON transactions (user_id);

ALTER TABLE transactions
    ADD CONSTRAINT fk_transactions_campaign
    FOREIGN KEY (campaign_id)
    REFERENCES campaigns(id)
    ON DELETE SET NULL;

ALTER TABLE transactions
    ADD CONSTRAINT fk_transactions_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE SET NULL;
