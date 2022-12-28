from ui.components.card import Card
from ui.mvc.controller import Controller
import typing

class View(typing.Protocol):

    def setup(self, controller: Controller) -> None:
        """Sets up the view based on a Controller."""

    def start_main_loop(self) -> None:
        """Starts the main loop for the application."""

    def add_card(self, title: str, text: str) -> None:
        """Adds a card to the UI."""

    def remove_card(self, card: Card) -> None:
        ...

    def get_cards(self) -> list[Card]:
        ...

    def find_card(self, title: str) -> Card | None:
        ...

    def get_removed(self) -> list[Card]:
        ...