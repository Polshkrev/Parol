import enum

class UI(enum.Enum):
    CLI = "cli"
    GUI = "gui"

LANGUAGE: dict[str, str] = {
    "english": "en",
    "russian": "ru"
}