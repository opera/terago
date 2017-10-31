TAGS = mock # mock | prod
PROGRAMS = 
EXAMPLES = plugin_example
PLUGINS = terago.so

all: clean $(PROGRAMS) $(EXAMPLES) $(PLUGINS)

test: check

plugin_example:
	go build -o plugin_example examples/plugin.go

terago.so:
	go build -buildmode plugin plugin/terago.go

check:
	go test -v plugin_test.go interface.go

clean:
	@rm -rf $(PROGRAMS) $(EXAMPLES) $(PLUGINS)
