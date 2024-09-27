import os
import platform
import subprocess

def check_go_installation():
    try:
        subprocess.check_output(["go", "version"])
        return True
    except FileNotFoundError:
        return False

def main():
    if check_go_installation():
        print("Go is already installed.")
        subprocess.check_output(["go", "run", "main.go"])
    else:
        print("Go is not installed.")

if __name__ == "__main__":
    main()
