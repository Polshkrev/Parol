import argparse

UI_CHOICES = [
    "cli",
    "gui"
]

def parse() -> None:

    parser = argparse.ArgumentParser()

    parser.add_argument("ui", type=str, choices=UI_CHOICES, help="UI the programme will use.")
    # parser.add_argument("-l", "--language", type=str, required=False, choices=SUPPORTED_LANGUAGES, default="en", help="The language in which the programme will be run. If not provided, the programme will be run in English.")
    parser.add_argument("-v", "--verbose", required=False, action="store_true", default=False, help="Logging will be displayed in the console. If not provided, logging will be sent to a file provided in the settings.")
    parser.parse_args()