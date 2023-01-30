import enum

class UI(enum.Enum):
    CLI = "cli"
    GUI = "gui"

    def __str__(self) -> str:
        return self.value

LANGUAGE: dict[str, str] = {
    "english": "en",
    "russian": "ru"
}