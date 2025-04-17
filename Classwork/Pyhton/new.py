import re
def validate_email(email):
    pattern = re.compile(r'^[a-zA-Z0-9_+.-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9]')

    if re.findall(pattern, email):
        return True
    else:
        return False
    
email= "user@example.com"
print(validate_email(email))