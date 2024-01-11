# List of source files in src/ directory
sources := $(wildcard src/*.c)

# Generate object file names based on source files
objects := $(patsubst src/%.c,build/%.o,$(sources))
CFLAGS = -Wall -Wextra -Werror -I/Library/PostgreSQL/15/include -IHeaders

.PHONY: build install commands

install: atm
	@echo "Installing to dist/db"
	@mkdir -p dist/db
	@cp build/atm.out dist/

atm : $(objects)
	@echo "Compiling and Linking $(objects)"
	@cc -o build/atm.out $(objects) -lncurses -lmenu -L/Library/PostgreSQL/15/lib -lpq


build/%.o: src/%.c | build
	@cc $(CFLAGS) -c $< -o $@

build:
	@mkdir -p build


clean :
	@echo "Deleting build/ atm.out *.plist compile_commands.json .cache"
	@rm -rf build/ atm.out *.plist compile_commands.json .cache/

analyze:
	@echo "Analyzing Source files"
	@clang --analyze $(CFLAGS) src/*.c

commands:
	@bear -- make
