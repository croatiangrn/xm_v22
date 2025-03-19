CREATE TABLE companies
(
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name                VARCHAR(15)  NOT NULL,
    description         VARCHAR(3000),
    amount_of_employees INT          NOT NULL,
    registered          BOOLEAN      NOT NULL,
    type                company_type NOT NULL,
    CONSTRAINT uix_company_name UNIQUE (name)
);