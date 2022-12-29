from dataclasses import dataclass, field

import customtkinter as ctk

import settings as config
from ui.components import Card

@dataclass
class GUI:
    root: ctk.CTk
    settings: config.Settings
    root_width: int = 856
    root_height: int = 452
    cards: list[Card] = field(default_factory=list, repr=False)

    def setup(self) -> None:
        self.root.title("Parol")
        self.root.geometry(f"{self.root_width}x{self.root_height}")

        main_button = ctk.CTkButton(self.root, text="+", hover=False, width=50) # TODO: add command
        main_button.place(anchor=ctk.NE, relx=1, x=-5, y=5)
        self._paint_cards()
        self.start_mainloop()
    
    def start_mainloop(self) -> None:
        self.root.mainloop()

    def add_card(self, title: str, text: str) -> None:
        for card in self.cards:
            card.clear()
        card = Card(self.root, "./ui/assets/black.png", "./ui/assets/white.png", title, text)
        self.cards.append(card)
        self._paint_cards()

    def _paint_cards(self) -> None:
        row = 0
        index = 0 # enumerate doesn't work, i've tried.
        for card in self.cards:
            if index % 5 == 0 and index != 0:
                row += 1
                index -= 5
            index += 1
            if card.removed:
                continue
            card.setup(index, row)

    def find_card(self, title: str) -> Card | None:
        for card in self.cards:
            if title.lower() not in card.title.lower():
                continue
            return card

    def get_removed(self) -> list[Card]:
        return [card for card in self.cards if card.removed]

    def remove_card(self, card: Card) -> None:
        self.cards.remove(card)

    def get_cards(self) -> list[Card]:
        return self.cards