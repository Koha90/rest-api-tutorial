CREATE TABLE public.author (
                               id uuid primary key default gen_random_uuid(),
                               name varchar(100) not null
);

CREATE TABLE public.book (
                             id uuid primary key default gen_random_uuid(),
                             name varchar(100) not null,
                             author_id uuid not null,
                             CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);

INSERT INTO author (name) VALUES ('Народ'); -- aa6fd453-ac18-4ae4-95b4-9dbf08e26fc1
INSERT INTO author (name) VALUES ('Джоан Роулинг'); -- 431973c4-eebc-4a34-9aa1-4c7faef538b5
INSERT INTO author (name) VALUES ('Джак Лондон'); -- 70a530a3-097a-4b86-b52f-b001285b9d39

INSERT INTO book (name, author_id) VALUES ('колобок', 'aa6fd453-ac18-4ae4-95b4-9dbf08e26fc1');
INSERT INTO book (name, author_id) VALUES ('Гарри Поттер', '431973c4-eebc-4a34-9aa1-4c7faef538b5');
INSERT INTO book (name, author_id) VALUES ('Белый Клык', '70a530a3-097a-4b86-b52f-b001285b9d39');