import os
from pyzbar.pyzbar import decode
from PIL import Image

# In-Memory Storage
user_records = []

# Lambda functions
create_user_record = lambda name, email: {"name": name, "email": email}
insert_user_record = lambda record: user_records.append(record)
delete_user_record = lambda name: [user_records.remove(user) for user in user_records if user['name'] == name]
fetch_all_user_records = lambda: user_records

# Function to scan and decode the SmartScan Code (QR Code)
def scan_qr_code(image_path):
    decoded_objects = decode(Image.open(image_path))
    for obj in decoded_objects:
        return obj.data.decode('utf-8')
    return None

# Function to register user from SmartScan (QR Code)
def RegisterUserFromSmartScan(image_path):
    user_data = scan_qr_code(image_path)
    if user_data:
        name, email = user_data.split(',')
        user_record = create_user_record(name, email)
        insert_user_record(user_record)
        print("User registered successfully.")
    else:
        print("No valid QR code found.")

# Menu-based interaction
def menu():
    while True:
        print("\nMenu:")
        print("1. Create new QR code")
        print("2. Register user from QR code")
        print("3. Delete user by name")
        print("4. Show all registered users")
        print("5. Exit")

        choice = input("Enter your choice: ")

        if choice == '1':
            name = input("Enter name: ")
            email = input("Enter email: ")
            from create_dummy_qr import create_qr_code
            create_qr_code(name, email)

        elif choice == '2':
            file_name = input("Enter the QR code file name (with extension): ")
            script_dir = os.path.dirname(os.path.abspath(__file__))
            img_path = os.path.join(script_dir, file_name)
            if os.path.exists(img_path):
                RegisterUserFromSmartScan(img_path)
            else:
                print("File not found.")

        elif choice == '3':
            name = input("Enter the name of the user to delete: ")
            delete_user_record(name)
            print("User deleted successfully.")

        elif choice == '4':
            all_users = fetch_all_user_records()
            print("Registered Users:", all_users)

        elif choice == '5':
            break

        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    menu()
