import argparse
from setup.proto import RuntimeArguments
import setup.choices

# TODO: Add a factory to seperate creation from use of the UI element

def parse() -> RuntimeArguments:
    parser = argparse.ArgumentParser()
    parser.add_argument("ui", type=setup.choices.UI, metavar=f'\173{",".join([choice.value for choice in setup.choices.UI])}\175', choices=[choice for choice in setup.choices.UI], help="UI the programme will use.")
    parser.add_argument("-l", "--language", type=str, required=False, choices=setup.choices.LANGUAGE.keys(), default="english", help="The language in which the programme will be run. If not provided, the programme will be run in English.")
    parser.add_argument("-v", "--verbose", required=False, action="store_true", default=False, help="Logging will be displayed in the console. If not provided, logging will be sent to a file provided in the settings.")
    args: RuntimeArguments = parser.parse_args() # type: ignore
    return args