import argparse
from setup.proto import RuntimeArguments
import setup.choices

# TODO: Add a factory to seperate creation from use of the UI element

def parse() -> RuntimeArguments:
    parser = argparse.ArgumentParser()
    parser.add_argument("ui", type=str, choices=setup.choices.UI, help="UI the programme will use.")
    parser.add_argument("-v", "--verbose", required=False, action="store_true", default=False, help="Logging will be displayed in the console. If not provided, logging will be sent to a file provided in the settings.")
    args: RuntimeArguments = parser.parse_args() # type: ignore
    return args