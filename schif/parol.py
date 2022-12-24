from cryptography.fernet import Fernet

import pathlib

from schif.context import SQLite

class Parol:

    def __init__(self, data_directory: str = "./data", key_filename: str = "KEY", password_filename: str = "passwords.db") -> None:
        self.key_file = pathlib.Path(f"{data_directory}/{key_filename}").absolute()
        self.password_file = pathlib.Path(f"{data_directory}/{password_filename}").absolute()
        if not self.password_file.parent.exists():
            self.password_file.parent.mkdir()
        elif not self.key_file.exists():
            self.create_key()
        else:
            ... # TODO: make a load passwords method

    def create_key(self) -> None:
        """Create an encryption key file."""
        with open(self.key_file, "wb") as f:
            f.write(Fernet.generate_key())

    def load_key(self) -> bytes:
        """Load an existing key file."""
        with open(self.key_file, 'rb') as f:
            return f.readline()

    def update_key_file(self, file: str) -> None:
        """Update the key file."""
        self.key_file = pathlib.Path(file).absolute()

    def update_password_file(self, file: str) -> None:
        """Update the password file."""
        self.password_file = pathlib.Path(file).absolute()