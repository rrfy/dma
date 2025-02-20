#!/usr/bin/env python3
import argparse
import os
import subprocess
import sys
import tempfile
import zipfile
import glob

GREEN = "\033[92m"
RED = "\033[91m"
YELLOW = "\033[93m"
BLUE = "\033[94m"
RESET = "\033[0m"

LANGUAGE_CONFIG = {
    ".py": {
        "compile": None,
        "run": lambda src, exe: ["python3", src],
    },
    ".cpp": {
        "compile": None,
        "run": lambda src, exe: [exe],
    },
    ".c": {
        "compile": None,
        "run": lambda src, exe: [exe],
    },
    ".go": {
        "compile": lambda src, exe: ["go", "build", "-o", exe, src],
        "run": lambda src, exe: [exe],
    },
    ".java": {
        "compile": lambda src, exe: ["javac", src],
        "run": lambda src, exe: ["java", exe],
    },
    ".kt": {
        "compile": lambda src, exe: ["kotlinc", src, "-include-runtime", "-d", exe],
        "run": lambda src, exe: ["java", "-jar", exe],
    },
    ".js": {
        "compile": None,
        "run": lambda src, exe: ["node", src],
    },
    ".rb": {
        "compile": None,
        "run": lambda src, exe: ["ruby", src],
    },
    ".cs": {
        "compile": lambda src, exe: ["csc", "/nologo", "/out:" + exe, src],
        "run": lambda src, exe: [exe] if os.name == "nt" else ["mono", exe],
    },
}

def compile_code(src_file, lang_conf, work_dir):
    ext = os.path.splitext(src_file)[1]
    exe = None
    if ext == ".java":
        exe = os.path.splitext(os.path.basename(src_file))[0]
    elif ext == ".kt":
        exe = os.path.join(work_dir, "solution.jar")
    elif ext == ".cs":
        exe = os.path.join(work_dir, "solution.exe")
    elif lang_conf["compile"] is not None:
        exe = os.path.join(work_dir, "solution_exec")

    if lang_conf["compile"]:
        compile_cmd = lang_conf["compile"](src_file, exe)
        print(f"{BLUE}Compiling with command: {' '.join(compile_cmd)}{RESET}")
        result = subprocess.run(compile_cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
        if result.returncode != 0:
            print(f"{RED}Compilation failed:{RESET}")
            print(result.stderr)
            sys.exit(1)
    return exe

def run_test(run_cmd, input_data, timeout=2):
    try:
        result = subprocess.run(run_cmd, input=input_data, capture_output=True, text=True, timeout=timeout)
        return result.stdout, result.stderr, result.returncode
    except subprocess.TimeoutExpired:
        return "", f"Time Limit Exceeded after {timeout} seconds", -1

def main():
    parser = argparse.ArgumentParser(description="Run solution code against tests from an archive.")
    parser.add_argument("source", help="Path to the user solution code file.")
    parser.add_argument("tests_archive", help="Path to the ZIP archive containing tests (input and output files).")
    parser.add_argument(
        "--cpp-compiler", default="g++",
        help="C++ compiler command to use (default: g++)"
    )
    parser.add_argument(
        "--c-compiler", default="gcc",
        help="C compiler command to use (default: gcc)"
    )
    args = parser.parse_args()

    src_file = os.path.abspath(args.source)
    tests_archive = os.path.abspath(args.tests_archive)
    ext = os.path.splitext(src_file)[1]

    if ext not in LANGUAGE_CONFIG:
        print(f"{RED}Unsupported file extension: {ext}{RESET}")
        sys.exit(1)

    lang_conf = LANGUAGE_CONFIG[ext]

    if ext == ".cpp":
        lang_conf["compile"] = lambda src, exe: [args.cpp_compiler, src, "-O2", "-o", exe]
    elif ext == ".c":
        lang_conf["compile"] = lambda src, exe: [args.c_compiler, src, "-O2", "-o", exe]

    with tempfile.TemporaryDirectory() as work_dir:
        exe = compile_code(src_file, lang_conf, work_dir)

        tests_dir = os.path.join(work_dir, "tests")
        os.makedirs(tests_dir, exist_ok=True)
        with zipfile.ZipFile(tests_archive, "r") as z:
            z.extractall(tests_dir)
        print(f"{GREEN}Extracted tests to {tests_dir}{RESET}")

        input_files = sorted(glob.glob(os.path.join(tests_dir, "*.in")))
        if not input_files:
            print(f"{RED}No test input files found in the archive.{RESET}")
            sys.exit(1)

        total_tests = len(input_files)
        passed = 0

        for input_path in input_files:
            test_name = os.path.splitext(os.path.basename(input_path))[0]
            expected_path = os.path.join(tests_dir, test_name + ".out")
            if not os.path.exists(expected_path):
                print(f"{YELLOW}Warning: Expected output file for {test_name} not found. Skipping.{RESET}")
                continue

            with open(input_path, "r") as f:
                input_data = f.read()
            with open(expected_path, "r") as f:
                expected_output = f.read()

            run_cmd = lang_conf["run"](src_file, exe)
            stdout, stderr, retcode = run_test(run_cmd, input_data)

            stdout_norm = stdout.strip()
            expected_norm = expected_output.strip()

            if retcode != 0:
                print(f"{RED}Test {test_name} FAILED: Non-zero exit code {retcode}{RESET}")
                print(f"{YELLOW}Stderr:{RESET}\n{stderr}")
                sys.exit(1)

            if stdout_norm != expected_norm:
                print(f"{RED}Test {test_name} FAILED{RESET}")
                print(f"{BLUE}Input:{RESET}\n{input_data}")
                print(f"{BLUE}Expected Output:{RESET}\n{expected_output}")
                print(f"{BLUE}Your Output:{RESET}\n{stdout}")
                sys.exit(1)
            else:
                print(f"{GREEN}Test {test_name} PASSED{RESET}")
                passed += 1

        print(f"{GREEN}All {passed}/{total_tests} tests passed!{RESET}")

if __name__ == "__main__":
    main()
