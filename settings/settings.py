from dataclasses import dataclass

import typing

import json
import yaml

Reader = typing.Callable[[typing.TextIO], dict[str, typing.Any]]
JSONReader = json.load
YAMLReader = yaml.safe_load

@dataclass
class Settings:
    key_filename: str
    password_filename: str
    log_directory: str
    data_directory: str
    verbose: bool
    black_image: str
    white_image: str

def read(filepath: str, filename: str, filetype: str = "json", reader: Reader = JSONReader) -> Settings:
    file = f"{filepath}/{filename}.{filetype}"
    with open(file, "r", encoding="utf-8") as f:
        data = reader(f)
        return Settings(**data)