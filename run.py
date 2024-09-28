import platform
import subprocess


def check_go_installation():
    try:
        subprocess.check_output(["go", "version"])
        return True
    except FileNotFoundError:
        return False


def install_go():
    sys = platform.system().lower()
    if sys == "windows":
        url = "https://dl.google.com/go/go1.23.0.windows-amd64.msi"
        subprocess.check_output(["powershell", "-Command", f"Invoke-WebRequest -Uri {url} -OutFile go.msi"])
        subprocess.check_output(["msiexec", "/i", "go.msi"])
    elif sys == "darwin":  # macOS
        subprocess.check_output(["brew", "install", "go"])


def main():
    if not check_go_installation():
        print("Golang is not installed")
        print("Running installation...")
        install_go()
        print("Success!")
    print("Running project...")
    subprocess.check_output(["go", "run", "main.go"])


if __name__ == "__main__":
    main()
