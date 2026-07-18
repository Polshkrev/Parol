import settings as config
import schif
import ui
import setup
import events

def main(settings: config.Settings) -> None:
    args = setup.parse()

    manager = schif.Parol(
        data_directory=settings.data_directory,
        key_filename=settings.key_filename,
        password_filename=settings.password_filename
    )

    logger = setup.logger(settings.log_directory, verbose=args.verbose)

    language = settings.language or args.language.lower()

    setup.language(language=language, language_folder=settings.configuration.language_folder)

    events.setup_log_event_handlers(logger)

    if args.ui is setup.UI.GUI:
        events.setup_gui_event_handlers(manager)
        ui.run_gui(settings, manager)
        
    else:
        ui.run_cli(manager, logger)

if __name__ == "__main__":
    settings = config.read(filepath=config.Folder, filename="settings", filetype="yaml", reader=config.YAMLReader)
    main(settings)