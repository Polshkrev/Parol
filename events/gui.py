import schif
import ui.components.card
import utils

def handle_x_button_click(card: ui.components.card.Card, database: schif.Parol) -> None:
    card.remove()
    database.remove_password(card.title)

def setup_gui_event_handlers(database: schif.Parol) -> None:
    utils.subscribe("remove_card", lambda card: handle_x_button_click(card, database))