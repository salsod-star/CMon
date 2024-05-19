CREATE TYPE contribution_status AS ENUM ('pending', 'incomplete','complete');

CREATE TABLE IF NOT EXISTS contributions (
    id BIGSERIAL PRIMARY KEY,
    total_amount INT NOT NULL,
    current_paid_amount INT NOT NULL,
    outstanding INT NOT NULL,
    status contribution_status NOT NULL DEFAULT 'pending',
    time_paid TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
