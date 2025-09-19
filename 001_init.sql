CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone_number VARCHAR(20),

);

CREATE TABLE cards (
    card_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    balance NUMERIC(12,2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'active',

);

CREATE TABLE topups (
    topup_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    card_id UUID REFERENCES cards(card_id) ON DELETE CASCADE,
    amount NUMERIC(12,2) NOT NULL,
    method VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE gates (
    gate_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    terminal_name VARCHAR(100) NOT NULL,
    gate_type VARCHAR(10) CHECK (gate_type IN ('IN','OUT')),
    location VARCHAR(255),
);

CREATE TABLE fare_matrix (
    fare_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    origin_gate_id UUID REFERENCES gates(gate_id) ON DELETE CASCADE,
    destination_gate_id UUID REFERENCES gates(gate_id) ON DELETE CASCADE,
    price NUMERIC(12,2) NOT NULL,
    UNIQUE(origin_gate_id, destination_gate_id)
);

CREATE TABLE trips (
    trip_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    card_id UUID REFERENCES cards(card_id) ON DELETE CASCADE,
    gate_in_id UUID REFERENCES gates(gate_id),
    gate_out_id UUID REFERENCES gates(gate_id),
    start_time TIMESTAMP DEFAULT NOW(),
    end_time TIMESTAMP,
    fare NUMERIC(12,2),
    status VARCHAR(20) DEFAULT 'ongoing'
);
