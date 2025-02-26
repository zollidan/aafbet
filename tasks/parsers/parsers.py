import subprocess

def run_parser_1():
    process = subprocess.run(["tasks/parsers/parser1/index.exe"])
    
    return process


def run_parser_2():
    process = subprocess.run(["tasks/parsers/parser2/index.exe"])
    
    return process