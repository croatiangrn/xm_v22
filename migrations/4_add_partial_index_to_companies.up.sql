ALTER TABLE companies DROP CONSTRAINT IF EXISTS uix_company_name;

CREATE UNIQUE INDEX uix_company_name ON companies (name) WHERE deleted_at IS NULL;