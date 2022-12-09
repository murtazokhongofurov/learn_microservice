CREATE TABLE IF NOT EXISTS addresses(
    id UUID PRIMARY KEY NOT NULL,
    store_id UUID REFERENCES stores(id) NOT NULL,
    country TEXT NOT NULL,
    street TEXT NOT NULL
);