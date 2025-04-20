CREATE TYPE film_status AS ENUM ('not_yet_aired', 'airing', 'finished_airing');
CREATE TYPE list_status AS ENUM ('plan_to_watch', 'watching', 'completed', 'on_hold', 'dropped');

CREATE TABLE genres (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(30) UNIQUE NOT NULL,
    
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE films (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(100) NOT NULL,
    synopsis        TEXT NOT NULL,
    status          film_status NOT NULL,
    total_episodes  INT NOT NULL,
    
    release_date    DATE NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    username        VARCHAR(30) UNIQUE NOT NULL,
    email           VARCHAR(100) UNIQUE NOT NULL,
    display_name    VARCHAR(30) NOT NULL,
    bio             TEXT,
    password        VARCHAR(100) NOT NULL,
    role            VARCHAR(100) NOT NULL,
    
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE film_genres (
    film_id         INT NOT NULL,
    genre_id        INT NOT NULL,

    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (film_id, genre_id),
    FOREIGN KEY (film_id) REFERENCES films(id) ON DELETE CASCADE,
    FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
);  

CREATE TABLE film_images (
    id              SERIAL PRIMARY KEY,
    film_id         INT NOT NULL,
    path            VARCHAR(255) NOT NULL,
    status          BOOLEAN DEFAULT FALSE,
    
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (film_id) REFERENCES films(id) ON DELETE CASCADE
);

CREATE TABLE user_film_lists (
    user_id         UUID NOT NULL,
    film_id         INT NOT NULL,
    status          list_status NOT NULL,
    visibility      BOOLEAN DEFAULT TRUE,

    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, film_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (film_id) REFERENCES films(id) ON DELETE CASCADE
);

CREATE TABLE reviews (
    id              SERIAL PRIMARY KEY,
    user_id         UUID NOT NULL,
    film_id         INT NOT NULL,
    rating          INT NOT NULL,
    comment         TEXT,

    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (film_id) REFERENCES films(id) ON DELETE CASCADE,
    UNIQUE (user_id, film_id)
);

CREATE TABLE reactions (
    review_id       INT NOT NULL,
    user_id         UUID NOT NULL,
    status          BOOLEAN NOT NULL,
    
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (review_id, user_id),
    FOREIGN KEY (review_id) REFERENCES reviews(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE OR REPLACE FUNCTION update_updated_at_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_genres
BEFORE UPDATE ON genres
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();