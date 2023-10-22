-- посты
DROP TABLE IF EXISTS news;
CREATE TABLE news (
     id SERIAL PRIMARY KEY,
     title TEXT NOT NULL,
     content TEXT NOT NULL,
     pubtime BIGINT NOT NULL DEFAULT 0,
     link TEXT NOT NULL 
);

-- комментарии
DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    news_id INT,
    content TEXT NOT NULL,
    pubtime INT NOT NULL DEFAULT 0
);
INSERT INTO comments(news_id,content)  VALUES (1,'комментарий1');
INSERT INTO comments(news_id,content)  VALUES (2,'комментарий2');
INSERT INTO comments(news_id,content)  VALUES (3,'комментарий3');

-- цензурирование
DROP TABLE IF EXISTS blacklist;
CREATE TABLE IF NOT EXISTS blacklist (
     id SERIAL PRIMARY KEY,
     word TEXT
);
INSERT INTO blacklist (word) VALUES ('qwerty');
INSERT INTO blacklist (word) VALUES ('йцукен');
INSERT INTO blacklist (word) VALUES ('zxvbnm');
