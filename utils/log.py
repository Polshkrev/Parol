import logging
from typing import Optional
from enum import Enum

class LoggingLevel(Enum):

    DEBUG = logging.DEBUG
    INFO = logging.INFO
    WARNING = logging.WARNING
    ERROR = logging.ERROR
    CRITICAL = logging.CRITICAL

def _set_formatter(format: str = "%(asctime)s:%(name)s[%(levelname)s] - %(message)s") -> logging.Formatter:
    """Helper function to set the formatter of the logger. A logging class (or its attributes) is not needed to set the formatter."""
    return logging.Formatter(format)

class Logger:
    """Logger utility to help with logging data."""

    def __init__(self, name: Optional[str] = None, format: str = "%(asctime)s:%(name)s[%(levelname)s] - %(message)s", level: LoggingLevel = LoggingLevel.DEBUG) -> None:
        self.level = level
        self.format = format
        if name:
            self.logger = logging.getLogger(name)
        else:
            self.logger = logging.getLogger(__name__)

    def _set_level(self, level: LoggingLevel) -> None:
        """Helper method to set the level of the logger."""
        self.logger.setLevel(level.value)

    def add_console(self) -> None:
        """Adds the console to the logger."""
        console = logging.StreamHandler()
        console.setFormatter(_set_formatter(self.format))
        self.logger.addHandler(console)

    def add_file(self, file: str = "./log.log") -> None:
        """Adds a file to the logger."""
        file_handler = logging.FileHandler(file, encoding="utf-8")
        file_handler.setFormatter(_set_formatter(self.format))
        self.logger.addHandler(file_handler)

    def full_setup(self, filename: str = "./log.log") -> None:
        """Sets-up the logger with BOTH the console, and a file."""
        self._set_level(self.level)
        self.add_console()
        self.add_file(filename)

    def console_only(self) -> None:
        """Sets-up the logger that only logs to the console."""
        self._set_level(self.level)
        self.add_console()

    def file_only(self, filename: str = "./log.log") -> None:
        """Sets-up the logger that only logs to a file."""
        self._set_level(self.level)
        self.add_file(filename)

    def log(self, message: str, level: LoggingLevel = LoggingLevel.DEBUG) -> None:
        """Logs a given message with a given LoggingLevel. If no level is given, DEBUG will be used."""
        self.logger.log(level.value, message)