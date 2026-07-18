import sqlite3

class SQLite:
    
    def __init__(self, file: str) -> None:
        self.file = file
    
    def __enter__(self) -> sqlite3.Cursor:
        self.connection = sqlite3.connect(self.file)
        return self.connection.cursor()
    
    def __exit__(self, exc_type, exc_value, traceback) -> None:
        self.connection.commit()
        self.connection.close()