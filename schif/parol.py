from cryptography.fernet import Fernet

class Parol:

    def __init__(self, data_directory: str = "./data", key_file: str = "./KEY", password_file: str = "./passwords.db") -> None:
        self.data_directory = data_directory
        self.key_file = key_file
        self.password_file = password_file