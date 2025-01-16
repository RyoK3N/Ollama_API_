import toml
import subprocess
import os
import sys

def build_ollama_curl(command, cwd=None, shell=False):
    """
    Runs a shell command and handles errors.
    
    :param command: List or string of the command to run.
    :param cwd: Directory to run the command in.
    :param shell: Whether to run the command in the shell.
    """
    try:
        if isinstance(command, list):
            cmd_str = ' '.join(command)
        else:
            cmd_str = command
        print(f"Running command: {cmd_str}")
        subprocess.run(command, check=True, cwd=cwd, shell=shell)
    except subprocess.CalledProcessError as e:
        print(f"Error running command: {cmd_str}")
        print(f"Error: {e}")
        sys.exit(1)

def load_config(config_path="config.toml"):
    """
    Loads the configuration from a TOML file.
    
    :param config_path: Path to the TOML configuration file.
    :return: Dictionary containing the configuration.
    """
    if not os.path.exists(config_path):
        print(f"Configuration file '{config_path}' not found. Please run 'sys_info.py' first to generate it.")
        sys.exit(1)
    
    try:
        with open(config_path, "r") as f:
            config = toml.load(f)
        return config
    except Exception as e:
        print(f"Failed to read '{config_path}': {e}")
        sys.exit(1)

def verify_system_requirements(system_info):
    """
    Verifies if the system meets the requirements for installing Ollama.
    
    :param system_info: Dictionary containing system information.
    """
    supported_os = ["Linux", "Darwin", "Windows"]
    os_name = system_info.get("OS", "")
    
    if os_name not in supported_os:
        print(f"Unsupported OS: {os_name}. Supported OS are: {', '.join(supported_os)}.")
        sys.exit(1)
    
    # Add more system requirement checks if necessary
    print("System requirements verified.")

def install_ollama():
    """
    Installs Ollama using the provided curl command.
    """
    install_command = "curl -fsSL https://ollama.com/install.sh | sh"
    build_ollama_curl(install_command, shell=True)

def main():
    # Load configuration
    config = load_config()
    system_info = config.get("system_info", {})
    
    # Display system information
    print("System Information:")
    for key, value in system_info.items():
        print(f"  {key}: {value}")
    
    # Verify system requirements
    verify_system_requirements(system_info)
    
    # Proceed with installation
    print("\nStarting Ollama installation...")
    install_ollama()
    print("\nOllama installation completed successfully.")

if __name__ == "__main__":
    main()