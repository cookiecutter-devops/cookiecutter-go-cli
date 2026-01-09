#!/usr/bin/env python
import os
import subprocess
from pathlib import Path

PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)
env = os.environ.copy()
env["GOSUMDB"] = "off"

def get_deps():
    mod = Path("go.mod").read_text()

    deps= set()

    in_require = False
    for line in mod.splitlines():
        line = line.strip()
        if line.startswith("require ("):
            in_require = True
            continue
        if not in_require:
            continue
        if line.startswith(")"):
            in_require = False
            continue
        if "// indirect" in line:
            continue

        deps.add(line.strip().split()[0])

    return list(deps)

if __name__ == "__main__":
    if "{{ cookiecutter.git }}" == "y":
        subprocess.check_call(["git", "init"], cwd=PROJECT_DIRECTORY)
        subprocess.check_call(["git", "checkout", "-b", "main"], cwd=PROJECT_DIRECTORY)
    else:
        Path(".gitignore").unlink()
        Path(".pre-commit-config.yaml").unlink()

    _ = subprocess.check_call(["go", "mod", "tidy"])

    for dep in get_deps():
        _ = subprocess.check_call(["go", "get", dep])
