from dataclasses import dataclass, field

import typing

import json
import yaml

Reader = typing.Callable[[typing.TextIO], dict[str, typing.Any]]
JSONReader = json.load
YAMLReader = yaml.safe_load
Folder = "./settings/public"

@dataclass
class Configuration:
    appearance: str
    theme: str
    geometry: str
    black_image: str
    white_image: str

@dataclass
class Settings:
    key_filename: str
    password_filename: str
    log_directory: str
    data_directory: str
    verbose: bool
    private_file: str = field(default_factory=str, repr=False)

    def __post_init__(self) -> None:
        self.configuration = self._get_config(reader=JSONReader)

    def _get_config(self, reader: Reader = JSONReader) -> Configuration:
        with open(self.private_file, "r", encoding="utf-8") as f:
            data = reader(f)
            return Configuration(**data)

def read(filepath: str, filename: str, filetype: str = "json", reader: Reader = JSONReader) -> Settings:
    file = f"{filepath}/{filename}.{filetype}"
    with open(file, "r", encoding="utf-8") as f:
        data = reader(f)
        return Settings(**data)