import schif
import ui.components.card
import ui.gui
import utils
import customtkinter as ctk

def handle_new_card(root: ctk.CTkToplevel, gui: ui.gui.GUI, site: str, password: str) -> None:
    gui.add_card(site, password)
    gui.database.add_password(site, password)
    root.destroy()

def handle_x_button_click(card: ui.components.card.Card, database: schif.Parol) -> None:
    card.remove()
    database.remove_password(card.title)
    utils.post_event("remove_password", card.title)

def setup_gui_event_handlers(database: schif.Parol) -> None:
    utils.subscribe("new_card", lambda args: handle_new_card(*args))
    utils.subscribe("remove_card", lambda card: handle_x_button_click(card, database))