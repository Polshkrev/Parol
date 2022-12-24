from cryptography.fernet import Fernet

import pathlib

class Parol:

    def __init__(self, data_directory: str = "./data", key_filename: str = "KEY", password_filename: str = "passwords.db") -> None:
        self.key_file = pathlib.Path(f"{data_directory}/{key_filename}").absolute()
        self.password_file = pathlib.Path(f"{data_directory}/{password_filename}").absolute()
        if not self.password_file.parent.exists():
            self.password_file.parent.mkdir()
        elif not self.key_file.exists():
            ... # TODO: make a key generation method
        else:
            ... # TODO: make a load passwords method