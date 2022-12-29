import settings as config
import schif
import ui
import setup

def main(settings: config.Settings) -> None:
    logger = setup.logger(settings.log_directory, verbose=settings.verbose)
    manager = schif.Parol(
        data_directory=settings.data_directory,
        key_filename=settings.key_filename,
        password_filename=settings.password_filename
    )
    # ui.run_cli(manager, logger)
    # ui.run_gui(settings, manager, logger)

if __name__ == "__main__":
    settings = config.read(filepath="./settings", filename="settings", filetype="yaml", reader=config.YAMLReader)
    main(settings)