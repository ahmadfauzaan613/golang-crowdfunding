-- Drop tabel jika ada
DROP TABLE IF EXISTS campaigns;

-- Buat tabel
CREATE TABLE campaigns (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    name VARCHAR(255),
    short_description VARCHAR(255),
    description TEXT,
    perks TEXT,
    backer_count INTEGER DEFAULT 0,
    goal_amount INTEGER DEFAULT 0,
    current_amount INTEGER DEFAULT 0,
    slug VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_campaigns_user_id
    ON campaigns (user_id);

ALTER TABLE campaigns
    ADD CONSTRAINT fk_campaigns_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE SET NULL;
