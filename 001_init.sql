CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(100) UNIQUE
);

CREATE TABLE IF NOT EXISTS cards (
    card_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    balance NUMERIC(12,2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'active'
);

CREATE TABLE IF NOT EXISTS topups (
    topup_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    card_id UUID REFERENCES cards(card_id) ON DELETE CASCADE,
    amount NUMERIC(12,2) NOT NULL,
    channel VARCHAR(50),
    status VARCHAR(20) DEFAULT 'completed',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS terminals (
    terminal_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    location VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS gates (
    gate_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    terminal_id UUID REFERENCES terminals(terminal_id) ON DELETE CASCADE,
    type VARCHAR(10) CHECK (type IN ('IN','OUT')),
    status VARCHAR(20) DEFAULT 'active'
);

CREATE TABLE IF NOT EXISTS transactions (
    event_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    card_id UUID REFERENCES cards(card_id),
    gate_id UUID REFERENCES gates(gate_id),
    event_type VARCHAR(50) NOT NULL,
    fare NUMERIC(12,2),
    timestamp TIMESTAMP DEFAULT NOW(),
    synced BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS trips (
    trip_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    card_id UUID REFERENCES cards(card_id),
    start_event_id UUID REFERENCES transactions(event_id),
    end_event_id UUID REFERENCES transactions(event_id),
    fare NUMERIC(12,2),
    status VARCHAR(20) DEFAULT 'ongoing'
);

CREATE TABLE IF NOT EXISTS fare_matrix (
    fare_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    origin_terminal_id UUID REFERENCES terminals(terminal_id),
    destination_terminal_id UUID REFERENCES terminals(terminal_id),
    price NUMERIC(12,2) NOT NULL,
    effective_date DATE DEFAULT CURRENT_DATE,
    UNIQUE(origin_terminal_id, destination_terminal_id, effective_date)
);

CREATE TABLE IF NOT EXISTS sync_log (
    sync_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    gate_id UUID REFERENCES gates(gate_id),
    last_synced_at TIMESTAMP DEFAULT NOW(),
    records_uploaded INTEGER DEFAULT 0
);
