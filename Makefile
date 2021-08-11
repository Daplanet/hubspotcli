CC		:= go
CFLAGS		:=
TARGET		:= hscli
.PHONY: clean all

$(TARGET):
	@$(CC) build -o build/$@

deps:
	@$(CC) mod vendor

all: deps $(TARGET)

clean:
	@rm -rf build/
