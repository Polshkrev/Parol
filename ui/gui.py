from __future__ import annotations

from dataclasses import dataclass, field

import customtkinter as ctk

from ui.components import Card
from ui.mvc import Controller
import settings

def _check_removed(card: Card) -> bool:
    return card.removed

def _get_removed(cards: list[Card]) -> list[Card]:
    return [card for card in cards if not _check_removed(card)]

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

def _paint_new_card(root: ctk.CTkToplevel, controller: Controller, site: str, password: str, verify_password: str) -> None:
    controller.add_card(site, password)
    controller.add_to_database(site, password)
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

def _paint_add_screen(screen: ctk.CTk, controller: Controller) -> None:
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

    sub_btn = ctk.CTkButton(add_screen, text="Submit", command=lambda: _paint_new_card(add_screen, controller, em.get(), ps.get(), vps.get()), state=ctk.DISABLED, width=200, hover=False)
    sub_btn.pack(ipadx=10, ipady=10, pady=5)

    _wait_for_input(add_screen, sub_btn, em, ps, vps)

@dataclass
class GUI:

    root: ctk.CTk
    settings: settings.Settings
    controller: Controller
    root_width: int = 856
    root_height: int = 482
    cards: list[Card] = field(default_factory=list)

    def setup(self) -> None:
        self.root.title("Parol")
        self.root.geometry(f"{self.root_width}x{self.root_height}")

        main_btn = ctk.CTkButton(self.root, text="+", hover=False, width=50, command=lambda: _paint_add_screen(self.root, self.controller))
        main_btn.place(anchor=ctk.NE, relx=1, x=-5, y=5)
        self._paint_cards()

    def start_main_loop(self) -> None:
        self.root.mainloop()

    def _validate(self, title: str) -> bool:
        return _validate_not_repeat(self.cards, title)

    def add_card(self, title: str, text: str) -> None:
        for card in self.cards:
            card.clear()
        card = Card(self.root, self.settings.black_clipboard_directory, self.settings.white_clipboard_directory, title, text)
        self.cards.append(card)
        self._paint_cards()

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