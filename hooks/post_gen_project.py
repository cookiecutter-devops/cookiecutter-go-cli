#!/usr/bin/env python
import os
import subprocess

PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)


def remove_file(filepath):
    os.remove(os.path.join(PROJECT_DIRECTORY, filepath))


def get_deps():
    subprocess.check_call(["go", "mod", "tidy"], cwd=PROJECT_DIRECTORY)

if __name__ == "__main__":
    if "{{ cookiecutter.git }}" == "y":
        subprocess.check_call(["git", "init"], cwd=PROJECT_DIRECTORY)
        subprocess.check_call(["git", "checkout", "-b", "main"], cwd=PROJECT_DIRECTORY)
    else:
        remove_file(".gitignore")
        remove_file(".pre-commit-config.yaml")

    get_deps()
