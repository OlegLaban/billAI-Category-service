CREATE TABLE IF NOT EXISTS categories (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id UUID DEFAULT NULL,
    category_name TEXT NOT NULL,
    is_default BOOLEAN DEFAULT false
);

CREATE INDEX idx_categories_user_id ON categories(user_id);

INSERT INTO categories (category_name, is_default)
VALUES 
    ('Продукты питания', true),
    ('Кафе и рестораны', true),
    ('Для дома и бытовая химия', true),
    ('Здоровье и принадлежности для ухода', true),
    ('Транспорт', true),
    ('Одежда и обувь', true),
    ('Развлечения и хобби', true),
    ('Образование', true),
    ('Домашние животные', true),
    ('Техника и электронника', true),
    ('Жилье и коммунальные услуги', true),
    ('Путешествия', true);