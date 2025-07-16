CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(32) NOT NULL UNIQUE,
    description VARCHAR(255)
);
