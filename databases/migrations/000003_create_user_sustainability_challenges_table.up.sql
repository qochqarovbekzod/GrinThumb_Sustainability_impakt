CREATE TABLE IF NOT EXISTS sustainability_challenges (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    goal_amount DECIMAL(10, 2),
    goal_unit VARCHAR(20),
    start_date DATE,
    end_date DATE
);