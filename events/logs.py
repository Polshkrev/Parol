import pkutils
import i18n

def handle_application_start(ui: str, logger: pkutils.globals.Logger) -> None:
    logger.log(i18n.t("general.logs.application_start", ui_name=ui.upper()))

def handle_password_add(site: str, logger: pkutils.globals.Logger) -> None:
    logger.log(i18n.t("general.logs.new_password", site=site), pkutils.globals.LoggingLevel.INFO)

def handle_password_update(site: str, logger: pkutils.globals.Logger) -> None:
    logger.log(i18n.t("general.logs.password_update", site=site), pkutils.globals.LoggingLevel.INFO)

def handle_password_removed(site: str, logger: pkutils.globals.Logger) -> None:
    logger.log(i18n.t("general.logs.remove_password", site=site), pkutils.globals.LoggingLevel.INFO)

def handle_application_end(logger: pkutils.globals.Logger) -> None:
    logger.log(i18n.t("general.logs.application_end"))

def setup_log_event_handlers(logger: pkutils.globals.Logger) -> None:
    pkutils.patterns.subscribe("application_start", lambda ui: handle_application_start(ui, logger))
    pkutils.patterns.subscribe("add_password", lambda site: handle_password_add(site, logger))
    pkutils.patterns.subscribe("change_password", lambda site: handle_password_update(site, logger))
    pkutils.patterns.subscribe("remove_password", lambda site: handle_password_removed(site, logger))
    pkutils.patterns.subscribe("application_end", lambda _: handle_application_end(logger))