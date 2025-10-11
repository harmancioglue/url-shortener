-- Migration: 001_create_urls_table.sql
-- Description: Create the main URLs table for storing URL mappings
-- Version: 1.0.0
-- Date: 2025-01-11

-- Create database if it doesn't exist
CREATE DATABASE IF NOT EXISTS url_shortener 
CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;

-- Use the database
USE url_shortener;

-- Drop table if it exists (for clean migration)
DROP TABLE IF EXISTS urls;

-- Create the urls table
CREATE TABLE urls (
    id           bigint auto_increment primary key,
    short_code   varchar(255) not null,
    original_url text not null,
    created_at   timestamp default CURRENT_TIMESTAMP not null,
    updated_at   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    expires_at   timestamp null,
    click_count  bigint default 0 null,
    is_active    tinyint(1) default 1 null,
    constraint short_code unique (short_code)
) collate = utf8mb4_unicode_ci;

-- Create indexes for performance
CREATE INDEX idx_short_code ON urls (short_code);
CREATE INDEX idx_created_at ON urls (created_at);
CREATE INDEX idx_is_active ON urls (is_active);
CREATE INDEX idx_expires_at ON urls (expires_at);

-- Add comments for documentation
ALTER TABLE urls 
    COMMENT = 'Stores URL mappings for the URL shortener service';

ALTER TABLE urls 
    MODIFY COLUMN id bigint auto_increment comment 'Primary key - auto-incrementing identifier';

ALTER TABLE urls 
    MODIFY COLUMN short_code varchar(255) not null comment 'Unique short code for the URL';

ALTER TABLE urls 
    MODIFY COLUMN original_url text not null comment 'Original long URL to redirect to';

ALTER TABLE urls 
    MODIFY COLUMN created_at timestamp default CURRENT_TIMESTAMP not null comment 'Timestamp when URL was created';

ALTER TABLE urls 
    MODIFY COLUMN updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment 'Timestamp when URL was last updated';

ALTER TABLE urls 
    MODIFY COLUMN expires_at timestamp null comment 'Optional expiration timestamp for the URL';

ALTER TABLE urls 
    MODIFY COLUMN click_count bigint default 0 null comment 'Number of times the URL has been accessed';

ALTER TABLE urls 
    MODIFY COLUMN is_active tinyint(1) default 1 null comment 'Flag indicating if URL is active (1=active, 0=inactive)';

-- Insert sample data for testing (optional)
INSERT INTO urls (short_code, original_url) VALUES 
    ('example1', 'https://www.example.com'),
    ('test123', 'https://github.com/harmancioglue/url-shortener'),
    ('demo456', 'https://www.google.com');

-- Migration completed
SELECT 'Migration 001_create_urls_table.sql completed successfully' as status;