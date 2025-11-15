
DROP TABLE IF EXISTS campaign_images;

CREATE TABLE campaign_images (
    id SERIAL PRIMARY KEY,
    campaign_id INTEGER,
    file_name VARCHAR(255),
    is_primary BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_campaign_images_campaign_id
    ON campaign_images (campaign_id);

ALTER TABLE campaign_images
    ADD CONSTRAINT fk_campaign_images_campaign
    FOREIGN KEY (campaign_id)
    REFERENCES campaigns(id)
    ON DELETE CASCADE;
