from cryptography.fernet import Fernet

import pathlib

from schif.context import SQLite, sqlite3

import typing

CreationFunction = typing.Callable[[], None]

def _check_parent(root: pathlib.Path) -> bool:
    """Check if the parent of a given file exists."""
    return root.parent.exists()

def _make_parent(root: pathlib.Path) -> None:
    """Make the parent directory of a given path if it doesn't exist already."""
    if _check_parent(root):
        return
    root.parent.mkdir()

def _check_root(root: pathlib.Path) -> bool:
    """Check if a given file exists."""
    return root.exists()

def _make_root(file: pathlib.Path, create_function: CreationFunction) -> None:
    """Make a given file if it doesn't exist already and it's parent directory exists."""
    if _check_root(root=file):
        return
    create_function()

def _check_file(path: pathlib.Path, create_function: CreationFunction) -> None:
    """Interface to simplify the control flow of file creation."""
    _make_parent(root=path)
    _make_root(file=path, create_function=create_function)

class Parol:

    def __init__(self, data_directory: str = "./data", key_filename: str = "KEY.key", password_filename: str = "passwords.db") -> None:
        self.key_file = pathlib.Path(f"{data_directory}/{key_filename}").absolute()
        self.password_file = pathlib.Path(f"{data_directory}/{password_filename}").absolute()
        _check_file(self.password_file, self._create_table)
        _check_file(self.key_file, self.create_key)
        self.load_passwords(self.load_key())

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
        path = pathlib.Path(file).absolute()
        _check_file(path, self.create_key)
        self.key_file = path

    def update_password_file(self, file: str) -> None:
        """Update the password file."""
        path = pathlib.Path(file).absolute()
        _check_file(path, self._create_table)
        self.key_file = path

    def dump_passwords_to_file(self, key: bytes, data: dict[str, str]) -> None:
        """Dump passwords to a file."""
        encrypt = {site: Fernet(key).encrypt(password.encode()).decode() for site, password in data.items()}
        with SQLite(file=str(self.password_file)) as cursor:
            try:
                cursor.executemany("""INSERT OR IGNORE INTO passwords VALUES (?,?)""", encrypt.items())
            except sqlite3.OperationalError:
                self._create_table()

    def load_passwords(self, key: bytes) -> dict[str, str]:
        """Load passwords from a file."""
        with SQLite(file=str(self.password_file)) as cursor:
            try:
                cursor.execute("""SELECT * FROM passwords""")
            except sqlite3.OperationalError:
                self._create_table()
                cursor.execute("""SELECT * FROM passwords""")
            raw: list[tuple[str, str]] = cursor.fetchall()
            return {site: Fernet(key).decrypt(password.encode()).decode() for site, password in raw}

    def add_password(self, site: str, password: str) -> None:
        """Add a given password for a given site into the database."""
        key = self.load_key()
        self.dump_passwords_to_file(key, {site: password})

    def update_password(self, site: str, password: str) -> None:
        """Update a given site's password."""
        key = self.load_key()
        encrypted = (site, Fernet(key).encrypt(password.encode()).decode())
        site, password = encrypted
        with SQLite(file=str(self.password_file)) as cursor:
            cursor.execute("""UPDATE passwords SET password = ? WHERE site = ?""", (password, site))

    def remove_password(self, site: str) -> None:
        """Remove a password from a file."""
        with SQLite(file=str(self.password_file)) as cursor:
            cursor.execute("""DELETE FROM passwords WHERE site=?""", (site,))

    def get_password(self, site: str) -> str:
        """Retrieve the password for a given site."""
        key = self.load_key()
        passwords = self.load_passwords(key)
        return passwords[site]

    # ! May be insecure (if one knows the encrypted password, one can reverse engineer and unencrypt it)
    def print_passwords(self) -> None:
        """Print the encrypted passwords in the database."""
        passwords = self._load_encrypted_passwords()
        for index, (site, password) in enumerate(passwords.items(), start=1):
            print(f"{index}. {site} - {password}")

    def _load_encrypted_passwords(self) -> dict[str, str]:
        with SQLite(file=str(self.password_file)) as cursor:
            cursor.execute("""SELECT * FROM passwords""")
            raw = cursor.fetchall()
            return {site: password for site, password in raw}

    def _create_table(self) -> None:
        with SQLite(file=str(self.password_file)) as cursor:
            cursor.execute("""CREATE TABLE IF NOT EXISTS passwords(site TEXT NOT NULL, password TEXT)""")