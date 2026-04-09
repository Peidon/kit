CREATE TABLE user_profile (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    phone TEXT,
    location TEXT,
    city TEXT,
    country TEXT,
    linkedin TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_education (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    school TEXT NOT NULL,
    field TEXT,
    degree TEXT,
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_user_education
        FOREIGN KEY (user_id)
        REFERENCES user_profile(id)
        ON DELETE CASCADE
);

CREATE TABLE user_experience (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    company TEXT NOT NULL,
    title TEXT,
    start_date DATE,
    end_date DATE,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_user_experience
        FOREIGN KEY (user_id)
        REFERENCES user_profile(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_user_email ON user_profile(email);
CREATE INDEX idx_education_user_id ON user_education(user_id);
CREATE INDEX idx_experience_user_id ON user_experience(user_id);

ALTER TABLE user_profile
ADD CONSTRAINT email_format_check
CHECK (email LIKE '%@%');