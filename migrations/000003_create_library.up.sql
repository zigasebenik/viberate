CREATE TABLE library (
    id SERIAL PRIMARY KEY,          
    user_id INT NOT NULL,            
    book_id INT NOT NULL,            
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,           
    CONSTRAINT fk_book
        FOREIGN KEY (book_id)
        REFERENCES books(id)
        ON DELETE CASCADE            
);