CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    slack_id TEXT NOT NULL,

    role user_role NOT NULL DEFAULT 'user'
);

