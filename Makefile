# 定义变量
BUILD_DIR=build
EXE_FILE=$(BUILD_DIR)/riscv64-ld
GO_BUILD=go build -o $(EXE_FILE) ./src
GCC=riscv64-linux-gnu-gcc
CFLAGS=-c
SRC_DIR=tests/csrc
OUT_DIR=$(BUILD_DIR)/out
OBJ_FILES=$(patsubst $(SRC_DIR)/%.c,$(OUT_DIR)/%.o,$(wildcard $(SRC_DIR)/*.c))

GREEN=\033[0;32m
NC=\033[0m  # No Color

# 默认目标
.PHONY: all
all: build test run

# build伪命令
.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	$(GO_BUILD)

# test伪命令
.PHONY: test
test: $(OBJ_FILES)
	@echo "get all obj file"

# 依赖目标，编译C文件为object文件
$(OUT_DIR)/%.o: $(SRC_DIR)/%.c
	@mkdir -p $(OUT_DIR)
	$(GCC) $(CFLAGS) $< -o $@

# run伪命令
.PHONY: run
run: test build
	@for obj in $(OBJ_FILES); do \
		echo "$(GREEN)Testing$(NC) $(EXE_FILE) $$obj"; \
		$(EXE_FILE) $$obj; \
	done
	@echo "$(GREEN)test finished$(NC)"

# 清理生成文件
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
