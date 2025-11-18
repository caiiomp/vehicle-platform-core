CREATE TABLE IF NOT EXISTS vehicles (
    id TEXT PRIMARY KEY,
    brand TEXT NOT NULL,
    model TEXT NOT NULL,
    year INT NOT NULL,
    color TEXT NOT NULL,
    price DECIMAL(12,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sales (
    id TEXT PRIMARY KEY,
    vehicle_id TEXT UNIQUE,
    buyer_document_number TEXT NOT NULL,
    sold_at TIMESTAMP NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (vehicle_id) REFERENCES vehicles (id)
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON vehicles
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON sales
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
