CREATE TYPE contribution_status AS ENUM ('pending', 'incomplete','complete');

CREATE TABLE IF NOT EXISTS contributions (
    id BIGSERIAL PRIMARY KEY,
    total_amount DECIMAL(10, 2) NOT NULL,
    current_paid_amount DECIMAL(10, 2) NOT NULL,
    outstanding DECIMAL(10, 2) NOT NULL,
    status contribution_status NOT NULL DEFAULT 'pending',
    interval TEXT,
    time_paid TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
