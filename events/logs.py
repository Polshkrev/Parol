import utils

def handle_application_start(ui: str, logger: utils.Logger) -> None:
    logger.log(f"Application started with {ui.upper()}.")

def handle_password_add(site: str, logger: utils.Logger) -> None:
    logger.log(f"A new password was added to the database for {site}.", utils.LoggingLevel.INFO)

def handle_password_update(site: str, logger: utils.Logger) -> None:
    logger.log(f"Password for {site} has been updated.", utils.LoggingLevel.INFO)

def handle_password_removed(site: str, logger: utils.Logger) -> None:
    logger.log(f"Password for {site} was removed from the database.", utils.LoggingLevel.INFO)

def handle_application_end(logger: utils.Logger) -> None:
    logger.log("Application ended.")

def setup_log_event_handlers(logger: utils.Logger) -> None:
    utils.subscribe("application_start", lambda ui: handle_application_start(ui, logger))
    utils.subscribe("add_password", lambda site: handle_password_add(site, logger))
    utils.subscribe("change_password", lambda site: handle_password_update(site, logger))
    utils.subscribe("remove_password", lambda site: handle_password_removed(site, logger))
    utils.subscribe("application_end", lambda _: handle_application_end(logger))