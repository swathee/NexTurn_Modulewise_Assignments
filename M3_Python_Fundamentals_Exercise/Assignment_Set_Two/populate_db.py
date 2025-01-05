import sqlite3

def populate_data():
    conn = sqlite3.connect("bookbuddy.db")
    cursor = conn.cursor()
    sample_books = [
        ("To Kill a Mockingbird", "Harper Lee", 1960, "Fiction"),
        ("1984", "George Orwell", 1949, "Dystopian"),
        ("The Catcher in the Rye", "J.D. Salinger", 1951, "Fiction")
    ]
    cursor.executemany("INSERT INTO books (title, author, published_year, genre) VALUES (?, ?, ?, ?)", sample_books)
    conn.commit()
    conn.close()

populate_data()
