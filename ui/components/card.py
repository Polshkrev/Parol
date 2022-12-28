from __future__ import annotations

from dataclasses import dataclass, field

import customtkinter as ctk
import PIL.Image
import PIL.ImageTk
import pyperclip3

import typing

def _get_apearance() -> str:
    # TODO: make return type an enum
    return ctk.get_appearance_mode()

def _paint_frame(root: ctk.CTk, title: str, width: int = 150, height: int = 150) -> ctk.CTkFrame:
    frame = ctk.CTkFrame(root, width=width, height=height, border_width=0, corner_radius=10) #type: ignore
    title_label = ctk.CTkLabel(frame, text=title, width=0)
    title_label.place(relx=(frame.winfo_width() - .5), rely=0, anchor=ctk.N)
    title_label.configure(font=('Roboto', 12, "bold"))
    return frame

def _paint_title_label(frame: ctk.CTkFrame, title: str) -> ctk.CTkLabel:
    title_label = ctk.CTkLabel(frame, text=title, width=0)
    title_label.place(relx=(frame.winfo_width() - .5), rely=0, anchor=ctk.N)
    title_label.configure(font=('Roboto', 12, "bold"))
    return title_label

def _paint_copy_button(frame: ctk.CTkFrame, copy_function: typing.Callable[[], None]) -> ctk.CTkButton:
    copy = ctk.CTkButton(frame, text=None, command=copy_function, fg_color=frame.fg_color, hover_color=frame.master.fg_color, border_width=0, border=0, width=0, height=0, corner_radius=10) # type: ignore
    copy.place(anchor=ctk.NE, relx=.98, rely=.02)
    return copy

def _paint_word_label(frame: ctk.CTkFrame, text: str) -> ctk.CTkLabel:
    word = ctk.CTkLabel(frame, text="*"*len(text))
    word.place(anchor=ctk.CENTER, relx=.5, rely=.5)
    word.configure(font=('Roboto', 8))
    return word

def _paint_x_button(frame: ctk.CTkFrame, remove_function: typing.Callable[[Card], None], card: Card) -> ctk.CTkButton:
    # ! copied from template still testing if works
    return ctk.CTkButton(frame, text="\u00d7", command=lambda: remove_function(card), text_font=('Roboto', 12, "bold"), fg_color=frame.fg_color, hover_color=frame.fg_color, border_width=0, border=0, width=0, height=0, corner_radius=10) # type: ignore

def _load_image(black_image_directory: str, white_image_directory: str) -> PIL.ImageTk.PhotoImage:
    with PIL.Image.open(black_image_directory) as clip:
        black_image = PIL.ImageTk.PhotoImage(clip)

    with PIL.Image.open(white_image_directory) as clip:
        white_image = PIL.ImageTk.PhotoImage(clip)

    return black_image if _get_apearance() == "Light" else white_image

def _remove(card: Card) -> None:
    # ! see `_paint_x_button` above
    card.remove()

@dataclass
class Card:
    root: ctk.CTk
    black_image_directory: str
    white_image_directory: str
    title: str = "Lorem ipsum"
    text: typing.Optional[str] = None
    side: typing.Literal['left', 'right', 'top', 'bottom'] = ctk.LEFT
    removed: bool = field(default=False, repr=False)

    def setup(self, row: int, index: int) -> None:
        """Setup the card in the GUI."""
        self.frame = _paint_frame(self.root, self.title)
        self.title_label = _paint_title_label(self.frame, self.title)
        self.copy_image = _load_image(self.black_image_directory, self.white_image_directory)
        self.copy_button = _paint_copy_button(self.frame, self.copy_password)
        self.word_label = _paint_word_label(self.frame, self.text) # type: ignore
        self.frame.grid(row=row, column=index, sticky=ctk.NW, pady=5, padx=5)

    def update_title(self, title: str) -> None:
        """Update the title on the card."""
        self.title = title
        self.title_label.configure(require_redraw=True, text=self.title)
        self.frame.update()

    def update_text(self, text: str) -> None:
        """Update the text on the card."""
        self.text = text
        self.title_label.configure(require_redraw=True, text=self.text)
        self.frame.update()

    def copy_password(self) -> None:
        """Copy the password to the clipboard."""
        # ! may be insecure
        pyperclip3.copy(self.text)

    def remove(self) -> None:
        """Remove the card from the GUI."""
        self.removed = True
        for child in self.frame.winfo_children():
            child.destroy()
        # self.frame.lower()
        # self.frame.grid_forget()
        self.frame.destroy()
        self.root.update()

    # ! DEPRICATED
    def clear(self) -> None:
        """Clear the card from the GUI."""
        self.frame.grid_forget()