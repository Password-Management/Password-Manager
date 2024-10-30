CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;


CREATE TABLE IF NOT EXISTS master_tbl (
    id UUID PRIMARY KEY,
    name character varying(100) NOT NULL,
    algorithm character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    special_key character varying(100) NOT NULL,
    count int8
);

CREATE TABLE IF NOT EXISTS user_tbl (
    id UUID PRIMARY KEY,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL, -- This the login password use bcrypt to store it
    special_key  character varying(100) NOT NULL,
    public_key TEXT NOT NULL,
    private_key TEXT NOT NULL,
    master_id UUID,
    CONSTRAINT fk_user FOREIGN KEY (master_id) REFERENCES master_tbl(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS psswrd_tbl (
    id UUID PRIMARY KEY,
    website_name character varying(100) NOT NULL,
    password character varying(500) NOT NULL,  --This is the password of the webiste mentioned (Encyption format)
    user_id UUID, 
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES user_tbl(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS login_tbl (
    user_id UUID,
    master_id UUID,
    is_login BOOLEAN NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES user_tbl(id) ON DELETE CASCADE,
    CONSTRAINT fk_master FOREIGN KEY (master_id) REFERENCES master_tbl(id) ON DELETE CASCADE
);

COMMIT;