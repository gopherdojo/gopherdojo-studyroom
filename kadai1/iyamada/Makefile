NAME	:= convert
SRCS	:= main.go

all: $(NAME)

$(NAME): $(SRCS)
	go build -o $(NAME) $(SRCS)

clean:
	rm -rf $(NAME)

re: clean all

test:
	go test ./imgconv/... -covermode=count -coverprofile=imgconv/c.out
	go tool cover -html=imgconv/c.out -o imgconv/coverage.html

.PHONY: all clean re test
