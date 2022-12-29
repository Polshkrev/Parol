import settings as config
import utils
import schif
import ui

import customtkinter as ctk

def _on_delete(root: ctk.CTk, logger: utils.Logger) -> None:
    root.quit()
    logger.log("Application ended.")

def run(settings: config.Settings, database: schif.Parol, logger: utils.Logger) -> None:

    ctk.set_appearance_mode(settings.configuration.appearance)
    ctk.set_default_color_theme(settings.configuration.theme)

    root = ctk.CTk()

    root_width, root_height = settings.configuration.geometry.split("x")

    gui = ui.GUI(root, settings, database, int(root_width), int(root_height))

    root.protocol("WM_DELETE_WINDOW", lambda: _on_delete(root, logger))

    logger.log("Application started.")

    gui.start()