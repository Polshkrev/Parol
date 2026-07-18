import typing

class RuntimeArguments(typing.Protocol):

    @property
    def ui(self) -> str:
        ...

    @property
    def verbose(self) -> bool:
        ...

    @property
    def language(self) -> str:
        ...