DROP TRIGGER IF EXISTS update_genres ON genres;
DROP FUNCTION IF EXISTS update_updated_at_column;

DROP TABLE IF EXISTS reactions;
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS user_film_lists;
DROP TABLE IF EXISTS film_images;
DROP TABLE IF EXISTS film_genres;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS films;
DROP TABLE IF EXISTS genres;

DROP TYPE IF EXISTS list_status;
DROP TYPE IF EXISTS film_status;

