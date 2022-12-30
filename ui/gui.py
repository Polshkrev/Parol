from __future__ import annotations

from dataclasses import dataclass, field

import customtkinter as ctk

import settings as config
import schif
from ui.components import Card


def _validate_not_repeat(cards: list[Card], title: str) -> bool:
    title_search = [card for card in cards if title.lower() == card.title.lower()]
    if not title_search:
        return True
    else:
        return False

def _validate(str1: str, str2: str) -> bool:
    if str1 != str2:
        return False
    elif "script" in str1.lower():
        return False
    elif "script" in str2.lower():
        return False
    else:
        return True
    # return str1 == str2 or "script" not in str1.lower() or "script" not in str2.lower()

def _paint_new_card(root: ctk.CTkToplevel, gui: GUI, site: str, password: str, verify_password: str) -> None:
    gui.add_card(site, password)
    gui.database.add_password(site, password)
    root.destroy()

def _check_length(*entries: ctk.CTkEntry) -> bool:
    len_entries = [entry for entry in entries if len(entry.get()) > 1]
    if not len_entries:
        return False
    if len(len_entries) < len(entries):
        return False
    else:
        return True

def _wait_for_input(root: ctk.CTkToplevel, target_button: ctk.CTkButton, *entries: ctk.CTkEntry, target_trigger: str = "<Return>", timestep: int = 50) -> None:
    if _check_length(*entries):
        target_button.configure(state=ctk.NORMAL)
        root.bind(target_trigger, target_button.clicked)
    else:
        target_button.configure(state=ctk.DISABLED)
        root.unbind(target_trigger)
    root.after(timestep, lambda: _wait_for_input(root, target_button, *entries, target_trigger=target_trigger, timestep=timestep))

def _paint_add_screen(screen: ctk.CTk, gui: GUI) -> None:
    add_screen = ctk.CTkToplevel(screen)
    add_screen.title("Add a Password")
    add_screen.geometry("250x250")
    add_screen.wm_resizable(False, False)

    em = ctk.CTkEntry(add_screen, placeholder_text="Site", width=200)
    em.pack(ipadx=10, ipady=10, pady=(15, 5))

    ps = ctk.CTkEntry(add_screen, placeholder_text="Password", show="*", width=200)
    ps.pack(ipadx=10, ipady=10, pady=5)

    vps = ctk.CTkEntry(add_screen, placeholder_text="Verify Password", show="*", width=200)
    vps.pack(ipadx=10, ipady=10, pady=5)

    sub_btn = ctk.CTkButton(add_screen, text="Submit", command=lambda: _paint_new_card(add_screen, gui, em.get(), ps.get(), vps.get()), state=ctk.DISABLED, width=200, hover=False)
    sub_btn.pack(ipadx=10, ipady=10, pady=5)

    _wait_for_input(add_screen, sub_btn, em, ps, vps)


@dataclass
class GUI:
    root: ctk.CTk
    settings: config.Settings
    database: schif.Parol
    root_width: int = 856
    root_height: int = 452
    cards: list[Card] = field(default_factory=list, repr=False)

    def setup(self) -> None:
        self.root.title("Parol")
        self.root.geometry(f"{self.root_width}x{self.root_height}")

        main_button = ctk.CTkButton(self.root, text="+", hover=False, width=50, command=lambda: _paint_add_screen(self.root, self))
        main_button.place(anchor=ctk.NE, relx=1, x=-5, y=5)
    
    def start_mainloop(self) -> None:
        self.root.mainloop()

    def start(self) -> None:
        for title, text in self.database.load_passwords(self.database.load_key()).items():
            self.add_card(title, text)
        self.setup()
        self.start_mainloop()

    def add_card(self, title: str, text: str) -> None:
        for card in self.cards:
            card.clear()
        card = Card(self.root, self.settings.configuration.black_image, self.settings.configuration.white_image, title, text)
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