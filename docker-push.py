import subprocess
from pathlib import Path

# --- SETTINGS ---
username = "sumanchapai"
image_name = "gw"
local_docker_image_name = image_name
remote_docker_image_name = f"{username}/{image_name}"

version_file = Path(".dockerversion")


# --- HELPERS ---
def yes(msg: str = "Do you want to proceed?") -> bool:
    got = input(f"{msg} (Y/n): ").strip().lower()
    return got in ("", "y", "yes")


def ask_and_run_cmd(cmd: str):
    print("\nAttempting to run:")
    print(" ", cmd)
    if yes():
        subprocess.run(cmd, shell=True, check=True)
    else:
        print("Aborted.")
        exit(1)


# --- MAIN ---
if not version_file.exists():
    print(
        version_file, "doesn't exist. create one and specify your latest version there"
    )
    exit(1)  # <-- FIXED: os.exit() does not exist; use exit() or sys.exit()

with open(version_file) as fd:
    current_version = fd.read().strip()

print("Current version:", current_version)

new_version = input("Enter new version: ").strip()
if not new_version:
    print("No version entered. Exiting.")
    exit(1)

# Update the version file
with open(version_file, "w") as fd:
    fd.write(new_version)

ask_and_run_cmd(
    f"docker buildx build "
    f"--platform linux/amd64,linux/arm64 "
    f"-t {remote_docker_image_name}:{new_version} "
    f"-t {remote_docker_image_name}:latest "
    f"--push ."
)
