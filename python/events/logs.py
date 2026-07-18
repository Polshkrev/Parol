import polutils
import i18n

def handle_application_start(ui: str, logger: polutils.Logger) -> None:
    logger.log(i18n.t("general.logs.application_start", ui_name=ui.upper()))

def handle_password_add(site: str, logger: polutils.Logger) -> None:
    logger.log(i18n.t("general.logs.new_password", site=site), polutils.LoggingLevel.INFO)

def handle_password_update(site: str, logger: polutils.Logger) -> None:
    logger.log(i18n.t("general.logs.password_update", site=site), polutils.LoggingLevel.INFO)

def handle_password_removed(site: str, logger: polutils.Logger) -> None:
    logger.log(i18n.t("general.logs.remove_password", site=site), polutils.LoggingLevel.INFO)

def handle_application_end(logger: polutils.Logger) -> None:
    logger.log(i18n.t("general.logs.application_end"))

def setup_log_event_handlers(logger: polutils.Logger) -> None:
    polutils.patterns.event.subscribe("application_start", lambda ui: handle_application_start(ui, logger))
    polutils.patterns.event.subscribe("add_password", lambda site: handle_password_add(site, logger))
    polutils.patterns.event.subscribe("change_password", lambda site: handle_password_update(site, logger))
    polutils.patterns.event.subscribe("remove_password", lambda site: handle_password_removed(site, logger))
    polutils.patterns.event.subscribe("application_end", lambda _: handle_application_end(logger))