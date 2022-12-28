from dataclasses import dataclass

import schif

from ui.components.card import Card

from ui.mvc.view import View

@dataclass
class Controller:

    database: schif.Parol
    view: View

    def add_card(self, site: str, password: str) -> None:
        self.view.add_card(site, password)

    def add_to_database(self, site: str, password: str) -> None:
        self.database.add_password(site, password)

    def remove_cards(self) -> None:
        cards = self.view.get_removed()
        if cards:
            for card in cards:
                self.remove_card(card)

    def remove_card(self, card: Card) -> None:
        self.database.remove_password(card.title)
        self.view.remove_card(card)

    def start(self) -> None:
        for title, text in self.database.load_passwords(self.database.load_key()).items():
            self.view.add_card(title, text)
        self.view.setup(self)
        self.view.start_main_loop()