CREATE TABLE IF NOT EXISTS user_challenges (
    user_id UUID NOT NULL,
    challenge_id UUID NOT NULL REFERENCES sustainability_challenges(id),
    progress DECIMAL(10, 2) DEFAULT 0,
    completed_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (user_id, challenge_id)
);