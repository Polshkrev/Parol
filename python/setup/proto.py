import typing
import setup.choices

class RuntimeArguments(typing.Protocol):

    @property
    def ui(self) -> setup.choices.UI:
        ...

    @property
    def verbose(self) -> bool:
        ...

    @property
    def language(self) -> str:
        ...