from flask import Flask, request, jsonify
import sqlite3

app = Flask(__name__)

def connect_db():
    return sqlite3.connect("bookbuddy.db")

# Initialize the database
def init_db():
    with connect_db() as conn:
        cursor = conn.cursor()
        cursor.execute('''CREATE TABLE IF NOT EXISTS books (
                            id INTEGER PRIMARY KEY AUTOINCREMENT,
                            title TEXT NOT NULL,
                            author TEXT NOT NULL,
                            published_year INTEGER NOT NULL,
                            genre TEXT NOT NULL
                          )''')
        conn.commit()

@app.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    required_fields = ['title', 'author', 'published_year', 'genre']

    for field in required_fields:
        if field not in data:
            return jsonify({"error": "Invalid data", "message": f"'{field}' is required"}), 400

    try:
        title = data['title']
        author = data['author']
        published_year = int(data['published_year'])
        genre = data['genre']

        with connect_db() as conn:
            cursor = conn.cursor()
            cursor.execute('INSERT INTO books (title, author, published_year, genre) VALUES (?, ?, ?, ?)',
                           (title, author, published_year, genre))
            conn.commit()
            return jsonify({"message": "Book added successfully", "book_id": cursor.lastrowid}), 201
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route('/books', methods=['GET'])
def get_books():
    with connect_db() as conn:
        cursor = conn.cursor()
        cursor.execute('SELECT * FROM books')
        books = cursor.fetchall()
        return jsonify([{
            "id": book[0],
            "title": book[1],
            "author": book[2],
            "published_year": book[3],
            "genre": book[4]
        } for book in books])

@app.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    with connect_db() as conn:
        cursor = conn.cursor()
        cursor.execute('SELECT * FROM books WHERE id = ?', (book_id,))
        book = cursor.fetchone()

        if book:
            return jsonify({
                "id": book[0],
                "title": book[1],
                "author": book[2],
                "published_year": book[3],
                "genre": book[4]
            })
        else:
            return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

@app.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    data = request.get_json()

    with connect_db() as conn:
        cursor = conn.cursor()
        cursor.execute('SELECT * FROM books WHERE id = ?', (book_id,))
        if not cursor.fetchone():
            return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

        updates = []
        params = []
        for key in ['title', 'author', 'published_year', 'genre']:
            if key in data:
                updates.append(f"{key} = ?")
                params.append(data[key])

        if updates:
            params.append(book_id)
            cursor.execute(f'UPDATE books SET {", ".join(updates)} WHERE id = ?', params)
            conn.commit()
            return jsonify({"message": "Book updated successfully"})
        else:
            return jsonify({"error": "Invalid data", "message": "No valid fields to update"}), 400

@app.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    with connect_db() as conn:
        cursor = conn.cursor()
        cursor.execute('SELECT * FROM books WHERE id = ?', (book_id,))
        if not cursor.fetchone():
            return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

        cursor.execute('DELETE FROM books WHERE id = ?', (book_id,))
        conn.commit()
        return jsonify({"message": "Book deleted successfully"})

if __name__ == '__main__':
    init_db()
    app.run(debug=True)
