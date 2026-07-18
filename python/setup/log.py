import polutils

from datetime import datetime

import pathlib

def _make_log_directory(log_folder: str) -> None:
    path = pathlib.Path(log_folder).absolute()
    if not path.exists():
        path.mkdir()

def _make_logger(log_folder: str, date: str, verbose: bool) -> polutils.Logger:
    logger = polutils.Logger(__name__)
    log_file = f"{log_folder}/{date}.log"
    if not verbose:
        logger.file_only(log_file)
    if verbose:
        logger.full_setup(log_file)
    return logger

def logger(log_folder: str, verbose: bool = False) -> polutils.Logger:
    now = str(datetime.now().date())
    _make_log_directory(log_folder)
    logger = _make_logger(log_folder, now, verbose)
    return logger