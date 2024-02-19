ALTER TABLE users_on_organizations
ADD COLUMN role VARCHAR(255) NOT NULL DEFAULT 'admin';