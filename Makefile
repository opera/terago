TAGS = mock # mock | prod
PROGRAMS = 
EXAMPLES = kvstore_example plugin_example
PLUGINS = terago.so

all: clean $(PROGRAMS) $(EXAMPLES) $(PLUGINS)

test: check

kvstore_example:
	go build -tags $(TAGS) -o kvstore_example examples/kvstore.go

plugin_example:
	go build -tags $(TAGS) -o plugin_example examples/plugin.go

terago.so:
	go build -buildmode plugin -tags $(TAGS) plugin/terago.go

check:
	go test -tags $(TAGS)

clean:
	@rm -rf $(PROGRAMS) $(EXAMPLES) $(PLUGINS)
