CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;


CREATE TABLE IF NOT EXISTS master_tbl (
    id UUID PRIMARY KEY,
    created_at timestamptz NOT NULL,
    name character varying(100) NOT NULL,
    algorithm character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    special_key character varying(100) NOT NULL,
    plan character varying(100), -- Plan could be demo, pro , premium or basic
    count int8
);

CREATE TABLE IF NOT EXISTS user_tbl (
    id UUID PRIMARY KEY,
    created_at timestamptz NOT NULL,
    deleted_at timestamptz,
    deleted_by UUID,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL, -- This the login password use bcrypt to store it
    special_key  character varying(100) NOT NULL,
    public_key TEXT NOT NULL,
    private_key TEXT NOT NULL,
    isMaster BOOLEAN,
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
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    is_login BOOLEAN NOT NULL,
    is_master BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS creds_tbl (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    otp character varying(100) NOT NULL,
    is_used BOOLEAN 
);

COMMIT;