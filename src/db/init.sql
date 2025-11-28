CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS mentor_recruitment_proposals;
DROP TABLE IF EXISTS contract_requests;
DROP TABLE IF EXISTS contracts;
DROP TABLE IF EXISTS mentor_recruitment_tags;
DROP TABLE IF EXISTS plan_tags;
DROP TABLE IF EXISTS skills;
DROP TABLE IF EXISTS mentor_recruitments;
DROP TABLE IF EXISTS plans;
DROP TABLE IF EXISTS careers;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS users;

-- カテゴリーテーブル
CREATE TABLE categories (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- ユーザーテーブル
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL CHECK (
        char_length(password) >= 12 AND
        password ~ '[a-zA-Z]' AND
        password ~ '[0-9]'
    ),
    self_introduction TEXT CHECK (char_length(self_introduction) <= 2000),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
CREATE INDEX idx_users_name ON users (name);

-- タグテーブル
CREATE TABLE tags (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- スキルテーブル
CREATE TABLE skills (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    evaluation INT NOT NULL,
    years_of_experience INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- 経歴テーブル
CREATE TABLE careers (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    detail TEXT NOT NULL,
    start_year INT,
    end_year INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
CREATE INDEX idx_careers_user_id ON careers (user_id);

-- メンタープランテーブル
CREATE TABLE plans (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id),
    description TEXT NOT NULL,
    status VARCHAR(255) NOT NULL,
    consultation_type VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    consultation_method VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- メンタープランとタグの中間テーブル
CREATE TABLE plan_tags (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    plan_id UUID NOT NULL REFERENCES plans(id),
    tag_id UUID NOT NULL REFERENCES tags(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- メンター募集テーブル
CREATE TABLE mentor_recruitments (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    consultation_type VARCHAR(255) NOT NULL,
    consultation_method VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    budget_from INT NOT NULL,
    budget_to INT NOT NULL,
    application_period DATE NOT NULL,
    status VARCHAR(255) NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- メンター募集とタグの中間テーブル
CREATE TABLE mentor_recruitment_tags (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    mentor_recruitment_id UUID NOT NULL REFERENCES mentor_recruitments(id),
    tag_id UUID REFERENCES tags(id),
    tag_text VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- メンター募集提案テーブル
CREATE TABLE mentor_recruitment_proposals (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    mentor_recruitment_id UUID NOT NULL REFERENCES mentor_recruitments(id),
    description TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- 契約リクエストテーブル
CREATE TABLE contract_requests (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    message TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    price_at_request INT NOT NULL,
    plan_id UUID NOT NULL REFERENCES plans(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
CREATE INDEX idx_contract_requests_user_id ON contract_requests (user_id);

-- 契約テーブル
CREATE TABLE contracts (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    plan_id UUID NOT NULL REFERENCES plans(id),
    message TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- 初期カテゴリーデータ挿入
INSERT INTO categories (id, name) VALUES
    (gen_random_uuid(), 'プログラミング'),
    (gen_random_uuid(), 'マーケティング'),
    (gen_random_uuid(), 'デザイン'),
    (gen_random_uuid(), 'ライティング'),
    (gen_random_uuid(), '動画・映像'),
    (gen_random_uuid(), 'ビジネス'),
    (gen_random_uuid(), '語学'),
    (gen_random_uuid(), 'ライフスタイル');

-- 初期タグデータ挿入
insert into tags (name) values ('PHP'), ('TS'), ('Go'), ('Java'), ('AWS');

