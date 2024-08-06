# LibraryBooks.py

class Library:
    def __init__(self):
        self.books = {}

    def add_book(self, title, author, publisher, volume, year, isbn):
        """Adds a book to the library."""
        self.books[isbn] = {
            'title': title,
            'author': author,
            'publisher': publisher,
            'volume': volume,
            'year': year,
            'isbn': isbn
        }

    def remove_book(self, isbn):
        """Removes a book from the library by ISBN."""
        if isbn in self.books:
            del self.books[isbn]
            return True
        return False

    def get_book_details(self, isbn):
        """Gets the details of a book by ISBN."""
        return self.books.get(isbn, None)

# Pre-populate with 25 books (sample data)
def populate_library(library):
    books = [
        {'title': 'Operating System Concepts', 'author': 'Abraham Silberschatz', 'publisher': 'Wiley', 'volume': 10, 'year': 2020, 'isbn': '9781119456339'},
        {'title': 'Modern Operating Systems', 'author': 'Andrew S. Tanenbaum', 'publisher': 'Pearson', 'volume': 4, 'year': 2020, 'isbn': '9780133591620'},
        {'title': 'Data Structures and Algorithm Analysis', 'author': 'Mark Allen Weiss', 'publisher': 'Pearson', 'volume': 4, 'year': 2020, 'isbn': '9780132847377'},
        {'title': 'Introduction to Algorithms', 'author': 'Thomas H. Cormen', 'publisher': 'MIT Press', 'volume': 3, 'year': 2022, 'isbn': '9780262046305'},
        {'title': 'Algorithms Unlocked', 'author': 'Thomas H. Cormen', 'publisher': 'MIT Press', 'volume': 1, 'year': 2021, 'isbn': '9780262518802'},
        {'title': 'Hands-On Machine Learning with Scikit-Learn, Keras, and TensorFlow', 'author': 'Aurélien Géron', 'publisher': 'O\'Reilly Media', 'volume': 2, 'year': 2022, 'isbn': '9781492032649'},
        {'title': 'Deep Learning', 'author': 'Ian Goodfellow', 'publisher': 'MIT Press', 'volume': 1, 'year': 2020, 'isbn': '9780262035613'},
        {'title': 'Machine Learning Yearning', 'author': 'Andrew Ng', 'publisher': 'Deeplearning.ai', 'volume': 1, 'year': 2020, 'isbn': '9781365840733'},
        {'title': 'Pattern Recognition and Machine Learning', 'author': 'Christopher Bishop', 'publisher': 'Springer', 'volume': 1, 'year': 2020, 'isbn': '9780387310732'},
        {'title': 'Probabilistic Machine Learning: An Introduction', 'author': 'Kevin P. Murphy', 'publisher': 'MIT Press', 'volume': 1, 'year': 2022, 'isbn': '9780262046824'},
        {'title': 'Artificial Intelligence: A Modern Approach', 'author': 'Stuart Russell', 'publisher': 'Pearson', 'volume': 4, 'year': 2020, 'isbn': '9780134610992'},
        {'title': 'Deep Learning with Python', 'author': 'François Chollet', 'publisher': 'Manning Publications', 'volume': 2, 'year': 2021, 'isbn': '9781617296864'},
        {'title': 'Machine Learning with PyTorch and Scikit-Learn', 'author': 'Sebastian Raschka', 'publisher': 'Packt Publishing', 'volume': 1, 'year': 2021, 'isbn': '9781801819312'},
        {'title': 'Data Structures and Algorithms Made Easy', 'author': 'Narasimha Karumanchi', 'publisher': 'CareerMonk Publications', 'volume': 5, 'year': 2021, 'isbn': '9788193245279'},
        {'title': 'Python Data Structures and Algorithms', 'author': 'Benjamin Baka', 'publisher': 'Packt Publishing', 'volume': 1, 'year': 2020, 'isbn': '9781786467355'},
        {'title': 'Fundamentals of Data Structures in C', 'author': 'Ellis Horowitz', 'publisher': 'Universities Press', 'volume': 2, 'year': 2020, 'isbn': '9788173716058'},
        {'title': 'Data Structures and Algorithms Using Python', 'author': 'Rance D. Necaise', 'publisher': 'Wiley', 'volume': 1, 'year': 2021, 'isbn': '9781118902741'},
        {'title': 'Machine Learning: A Probabilistic Perspective', 'author': 'Kevin P. Murphy', 'publisher': 'MIT Press', 'volume': 1, 'year': 2021, 'isbn': '9780262018029'},
        {'title': 'Data Structures and Algorithms in Python', 'author': 'Michael T. Goodrich', 'publisher': 'Wiley', 'volume': 1, 'year': 2021, 'isbn': '9781118290274'},
        {'title': 'Machine Learning in Action', 'author': 'Peter Harrington', 'publisher': 'Manning Publications', 'volume': 1, 'year': 2020, 'isbn': '9781617290183'},
        {'title': 'Introduction to the Theory of Computation', 'author': 'Michael Sipser', 'publisher': 'Cengage Learning', 'volume': 3, 'year': 2021, 'isbn': '9781281504945'},
        {'title': 'Machine Learning: The Art and Science of Algorithms that Make Sense of Data', 'author': 'Peter Flach', 'publisher': 'Cambridge University Press', 'volume': 1, 'year': 2020, 'isbn': '9781107096394'},
        {'title': 'Operating Systems: Principles and Practice', 'author': 'Thomas Anderson', 'publisher': 'Recursive Books', 'volume': 2, 'year': 2020, 'isbn': '9780985673529'},
        {'title': 'Operating Systems: Three Easy Pieces', 'author': 'Remzi H. Arpaci-Dusseau', 'publisher': 'Arpaci-Dusseau Books', 'volume': 1, 'year': 2021, 'isbn': '9781985086593'},
        {'title': 'Machine Learning for Absolute Beginners', 'author': 'Oliver Theobald', 'publisher': 'Scatterplot Press', 'volume': 3, 'year': 2021, 'isbn': '9781549617218'}
    ]
    for book in books:
        library.add_book(**book)

if __name__ == "__main__":
    lib = Library()
    populate_library(lib)
