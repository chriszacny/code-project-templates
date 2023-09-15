# Makefile

# Directories
SRCDIR = src
BINDIR = bin

# Source files
SOURCES = $(wildcard $(SRCDIR)/*.py)

# Binary name
TARGET = $(BINDIR)/program

# Default target
all:
	docker build -t pythonapp .
	docker run -ti --rm pythonapp

# Building the binary
#$(TARGET): $(SOURCES)
#	@mkdir -p $(BINDIR)
#	cp $(SOURCES) $(TARGET)
#	chmod +x $(TARGET)

# Clean the binary
clean:
	@rm -rf $(BINDIR)

.PHONY: all clean
