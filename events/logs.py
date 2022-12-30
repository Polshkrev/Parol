import utils

def handle_application_start(ui: str, logger: utils.Logger) -> None:
    logger.log(f"Application started with {ui.upper()}.")

def handle_application_end(logger: utils.Logger) -> None:
    logger.log("Application ended.")

def handle_key_file_created(key_file: str, logger: utils.Logger) -> None:
    logger.log(f"Key file created at {key_file}")

def setup_log_event_handlers(logger: utils.Logger) -> None:
    utils.subscribe("application_start", lambda ui: handle_application_start(ui, logger))
    utils.subscribe("key_created", lambda file: handle_key_file_created(file, logger))
    utils.subscribe("application_end", lambda _: handle_application_end(logger))