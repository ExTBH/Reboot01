# List of source files in src/ directory
sources := $(wildcard src/*.c)

# Generate object file names based on source files
objects := $(patsubst src/%.c,build/%.o,$(sources))
CFLAGS = -Wall -Wextra -Werror -IHeaders

POSTGRE_PATH = ""

ifeq ($(shell uname -s), Darwin)
	CFLAGS += -I/Library/PostgreSQL/15/include
	POSTGRE_PATH = /Library/PostgreSQL/15/lib
else
	CFLAGS += -I/usr/include/postgresql
	POSTGRE_PATH = /usr/lib/postgresql
endif


.PHONY: build install commands

install: atm
	@echo "Installing to dist/db"
	@mkdir -p dist
	@cp build/atm.out dist/

atm : $(objects)
	@echo "Compiling and Linking $(objects)"
	clang -o build/atm.out $(objects) -L$(POSTGRE_PATH) -lpq


build/%.o: src/%.c | build
	clang $(CFLAGS) -c $< -o $@

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
