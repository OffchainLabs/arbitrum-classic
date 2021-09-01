import os

path_prefix = "docs/sol_contract_docs/md_docs/"


black_list_dirs = ["arb-bridge-eth/libraries", "arb-bridge-eth/interfaces"]
black_list_files = [
    "ProxySetter.md",
    "TokenAddressHandler.md",
    "INode.md",
    "IRollup.md",
    "IEthERC20Bridge.md",
]


def get_all_files(dir):
    r = []
    subdirs = [x[0] for x in os.walk(dir)]
    for subdir in subdirs:
        files = next(os.walk(subdir))[2]
        if len(files) > 0:
            for file in files:
                r.append(os.path.join(subdir, file))
    return r


def generate_docs_by_dir(_dir):
    os.system(
        "yarn solidity-docgen -i packages/{_dir}/contracts -o {path_prefix}/{_dir} -t ./docs/sol_contract_docs".format(
            _dir=_dir, path_prefix=path_prefix
        )
    )


def generate_all_docs():
    os.system(
        "rm -rf {path_prefix} && mkdir {path_prefix}".format(path_prefix=path_prefix)
    )
    generate_docs_by_dir("arb-bridge-eth")
    generate_docs_by_dir("arb-bridge-peripherals")
    generate_docs_by_dir("arb-os")


# solgen's "exclude" breaks imports, so we delete unwanted docs post-generation
def remove_unwanted_docs():
    # remove blacklisted dirs
    for target in black_list_dirs:
        p = "rm -rf " + path_prefix + target
        os.system(p)
    all_files = get_all_files(path_prefix)

    for i, _file in enumerate(all_files):
        # remove blacklisted files
        for black_listed_file in black_list_files:
            if _file.endswith(black_listed_file):
                os.system("rm " + _file)
            else:
                # remove "empty" (header-only) files and test files
                if not os.path.exists(_file):
                    continue
                file_data = open(_file, "r")
                lines = len([line for line in file_data if len(line.strip()) > 0])
                is_test = _file.find("/test") > -1
                file_data.close()
                if lines <= 4 or is_test:
                    print("removing", _file)
                    os.system("rm " + _file)


def run():
    print("***                                  ***")
    print("***                                  ***")
    print("*** generating natspec solidity docs ***")
    print("***                                  ***")
    print("***                                  ***")
    generate_all_docs()
    remove_unwanted_docs()
    print(
        "***                                                                                             ***"
    )
    print(
        "***                                                                                             ***"
    )
    print(
        "*** success! run `pre-commit run --files docs/sol_contract_docs/md_docs/**/*.md` for final diff ***"
    )
    print(
        "***                                                                                             ***"
    )
    print(
        "***                                                                                             ***"
    )


if __name__ == "__main__":
    run()
