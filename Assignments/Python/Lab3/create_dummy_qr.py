import qrcode
import os

def create_qr_code(name, email):
    # User data to encode in the QR code
    user_data = f"{name},{email}"

    # Create QR code
    qr = qrcode.QRCode(
        version=1,
        error_correction=qrcode.constants.ERROR_CORRECT_L,
        box_size=10,
        border=4,
    )
    qr.add_data(user_data)
    qr.make(fit=True)

    # Create an image from the QR Code instance
    img = qr.make_image(fill='black', back_color='white')

    # Save the QR code image in the same folder as the script
    script_dir = os.path.dirname(os.path.abspath(__file__))
    img_name = f"{name.replace(' ', '_')}_qr.png"
    img_path = os.path.join(script_dir, img_name)
    img.save(img_path)

    print(f"QR code saved to: {img_path}")
    return img_path

if __name__ == "__main__":
    name = input("Enter name: ")
    email = input("Enter email: ")
    create_qr_code(name, email)
