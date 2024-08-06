# mainLibraryManagement.py

from LibraryBooks import Library, populate_library

def main():
    # Create library instance and populate with sample data
    library = Library()
    populate_library(library)

    # Demonstrate adding a book
    new_book = {
        'title': 'New Book Title',
        'author': 'New Author',
        'publisher': 'New Publisher',
        'volume': 1,
        'year': 2024,
        'isbn': '9789999999999'
    }
    library.add_book(**new_book)
    print("Added new book:", library.get_book_details('9789999999999'))

    # Demonstrate removing a book
    removed = library.remove_book('9789999999999')
    print("Removed book with ISBN '9789999999999':", removed)

    # Demonstrate getting book details
    book_details = library.get_book_details('9781119456339')
    print("Details of book with ISBN '9781119456339':", book_details)

if __name__ == "__main__":
    main()
