#include <iostream>
#include <string>
#include "engine/executor.h"
#include "engine/query.h"
#include "engine/storage.h"

struct Args {
    std::string file;
    std::string query;
};

bool parseArgs(int argc, char* argv[], Args& args) {
    for (int i = 1; i < argc; ++i) {
        std::string arg = argv[i];
        if (arg == "--file" && i + 1 < argc) {
            args.file = argv[++i];
        } else if (arg == "--query" && i + 1 < argc) {
            args.query = argv[++i];
        } else {
            return false;
        }
    }

    return !args.file.empty() && !args.query.empty();
}

int main(int argc, char* argv[]) {
    Args args;
    if (!parseArgs(argc, argv, args)) {
        std::cerr << "Usage: " << argv[0] << " --file <path> --query '<command>'" << std::endl;
        return 1;
    }

    Storage storage;
    try {
        storage.load(args.file);
    } catch (const std::exception& e) {
        std::cerr << "error: cannot load " << args.file << ": " << e.what() << std::endl;
        return 1;
    }

    try {
        Array tokens = tokenize(args.query);
        bool changed = false;
        std::string result = execute(storage, tokens, changed);
        if (changed) {
            storage.save(args.file);
        }
        std::cout << "-> " << result << std::endl;
    } catch (const std::exception& e) {
        std::cerr << "error: " << e.what() << std::endl;
        return 1;
    }

    return 0;
}
