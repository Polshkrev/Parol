import schif
import pkutils
import i18n
import os

def _print_menu() -> None:
    print(i18n.t("ui.cli.menu.new_key"))
    print(i18n.t("ui.cli.menu.load_file"))
    print(i18n.t("ui.cli.menu.add_password"))
    print(i18n.t("ui.cli.menu.retrieve_password"))
    print(i18n.t("ui.cli.menu.update_password_file"))
    print(i18n.t("ui.cli.menu.update_key_file"))
    print(i18n.t("ui.cli.menu.remove_password"))
    print(i18n.t("ui.cli.menu.update_password"))
    print(i18n.t("ui.cli.menu.clear_screen"))
    print(i18n.t("ui.cli.menu.quit"))

def run(manager: schif.Parol, logger: pkutils.globals.Logger) -> None:
    pkutils.patterns.post("application_start", "cli")
    while True:
        _print_menu()
        choice = input(">> ").strip()
        if choice == i18n.t("ui.cli.menu.quit"):
            pkutils.patterns.post("application_end", None)
            print(i18n.t("general.logs.bye"))
            break
        elif choice == i18n.t("ui.cli.menu.clear_screen"):
            os.system('cls' if os.name == 'nt' else 'clear')
            continue
        match choice:
            case "1":
                manager.create_key()
                i18n.t("ui.cli.options.key_file")
            case "2":
                key = manager.load_key()
                logger.log(i18n.t("ui.cli.options.load_key"))
                manager.load_passwords(key)
                logger.log(i18n.t("ui.cli.options.password_loaded"))
            case "3":
                site = input(i18n.t("ui.cli.prompts.site"))
                password = input(i18n.t("ui.cli.prompts.password"))
                manager.add_password(site, password)
                pkutils.patterns.post("add_password", site)
            case "4":
                site = input(i18n.t("ui.cli.prompts.site"))
                try:
                    password = manager.get_password(site)
                except KeyError:
                    logger.log(i18n.t("ui.cli.errors.password_not_found", site=site), pkutils.globals.LoggingLevel.ERROR)
                    continue
                else:
                    print(i18n.t("ui.cli.options.password_found", site=site, password=password))
            case "5":
                path = input(i18n.t("ui.cli.prompts.path")).strip()
                manager.update_password_file(path)
                logger.log(i18n.t("ui.cli.options.password_file_update"))
            case "6":
                path = input(i18n.t("ui.cli.prompts.path")).strip()
                manager.update_key_file(path)
                logger.log(i18n.t("ui.cli.options.key_file_update"))
            case "7":
                site = input(i18n.t("ui.cli.prompts.site"))
                manager.remove_password(site)
                pkutils.patterns.post("remove_password", site)
            case "8":
                site = input(i18n.t("ui.cli.prompts.site")).strip()
                password = input(i18n.t("ui.cli.prompts.new_password")).strip()
                manager.update_password(site, password)
                pkutils.patterns.post("change_password", site)
            case "c" if i18n.get("locale") == "en":
                os.system('cls' if os.name == 'nt' else 'clear')
                continue
            case "q" if i18n.get("locale") == "en":
                pkutils.patterns.post("application_end", None)
                print("Bye")
                break
            case other:
                logger.log(f"Invalid choice {other} used.", pkutils.globals.LoggingLevel.ERROR)