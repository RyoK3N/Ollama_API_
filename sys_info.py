import platform
import os
import sys
import toml

def get_system_info():
    system_info = {
        "OS": platform.system(),
        "OS_Version": platform.version(),
        "OS_Release": platform.release(),
        "Architecture": platform.architecture()[0],
        "Machine": platform.machine(),
        "Processor": platform.processor(),
        "Python_Version": sys.version.split()[0],
        "Python_Executable": sys.executable,
        "Environment_Variables": os.environ.get("PATH", "").split(os.pathsep),
    }
    return system_info

def save_to_toml(system_info, filename="config.toml"):
    config = {"system_info": system_info}
    with open(filename, "w") as toml_file:
        toml.dump(config, toml_file)
    print(f"System information saved to {filename}")

if __name__ == "__main__":
    info = get_system_info()
    save_to_toml(info)