import argparse
from setup.proto import RuntimeArguments
import setup.choices

def parse() -> RuntimeArguments:
    parser = argparse.ArgumentParser()
    parser.add_argument("ui", type=str, choices=setup.choices.UI, help="UI the programme will use.")
    args: RuntimeArguments = parser.parse_args() # type: ignore
    return args