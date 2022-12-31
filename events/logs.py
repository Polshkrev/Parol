import utils
import i18n

def handle_application_start(ui: str, logger: utils.Logger) -> None:
    logger.log(i18n.t("general.logs.application_start", ui_name=ui.upper()))

def handle_password_add(site: str, logger: utils.Logger) -> None:
    logger.log(i18n.t("general.logs.new_password", site=site), utils.LoggingLevel.INFO)

def handle_password_update(site: str, logger: utils.Logger) -> None:
    logger.log(i18n.t("general.logs.password_update", site=site), utils.LoggingLevel.INFO)

def handle_password_removed(site: str, logger: utils.Logger) -> None:
    logger.log(i18n.t("general.logs.remove_password", site=site), utils.LoggingLevel.INFO)

def handle_application_end(logger: utils.Logger) -> None:
    logger.log(i18n.t("general.logs.application_end"))

def setup_log_event_handlers(logger: utils.Logger) -> None:
    utils.subscribe("application_start", lambda ui: handle_application_start(ui, logger))
    utils.subscribe("add_password", lambda site: handle_password_add(site, logger))
    utils.subscribe("change_password", lambda site: handle_password_update(site, logger))
    utils.subscribe("remove_password", lambda site: handle_password_removed(site, logger))
    utils.subscribe("application_end", lambda _: handle_application_end(logger))