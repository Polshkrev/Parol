import schif
import utils
import os

MENU = """(1) Create a new key
(2) Load an existing password file
(3) Add a new password
(4) Get a password
(5) Update password directory
(6) Update key directory
(7) Remove a password
(8) Update password
(c) Clear the screen
(q) Quit"""

def run(manager: schif.Parol, logger: utils.Logger) -> None:
    utils.post_event("application_start", "cli")
    while True:
        print(MENU)
        choice = input(">> ").strip()
        match choice:
            case "1":
                manager.create_key()
                logger.log("Key file created.")
            case "2":
                key = manager.load_key()
                logger.log("Key file loaded.")
                manager.load_passwords(key)
                logger.log("Passwords loaded.")
            case "3":
                site = input("Enter the site: ")
                password = input("Enter the password: ")
                manager.add_password(site, password)
                utils.post_event("add_password", site)
            case "4":
                site = input("What site do you want: ")
                try:
                    password = manager.get_password(site)
                except KeyError:
                    logger.log(f"Site: {site} not found in database.", utils.LoggingLevel.ERROR)
                    continue
                else:
                    print(f"Password for {site} is: {password}.")
            case "5":
                path = input("Enter a path: ").strip()
                manager.update_password_file(path)
                logger.log("Password file updated.")
            case "6":
                path = input("Enter a path: ").strip()
                manager.update_key_file(path)
                logger.log(f"Key file updated to: {path}.")
            case "7":
                password = input("Enter a site: ")
                manager.remove_password(password)
                utils.post_event("remove_password", password)
            case "8":
                site = input("Enter site: ").strip()
                password = input("New password: ").strip()
                manager.update_password(site, password)
                utils.post_event("change_password", site)
            case "c":
                os.system('cls' if os.name == 'nt' else 'clear')
                continue
            case "q":
                utils.post_event("application_end", None)
                print("Bye")
                break
            case other:
                logger.log(f"Invalid choice {other} used.", utils.LoggingLevel.ERROR)