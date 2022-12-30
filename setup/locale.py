import i18n

from setup.choices import LANGUAGE

def gen(language: str, language_folder: str) -> None:
    """Initialize the iternationalization library."""
    i18n.set("locale", LANGUAGE[language.lower()])
    i18n.load_path.append(language_folder)