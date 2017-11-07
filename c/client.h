#include <stdio.h>
#include <stdlib.h>
#include <tera_c.h>

tera_client_t* client_open(const char* conf_path, const char* log_prefix) {
	char *err = NULL;
	tera_client_t* cli = tera_client_open(conf_path, log_prefix, &err);
	if (err != NULL) {
		fprintf(stderr, "tera client open error: %s.\n", err);
		free(err);
	}
	return cli;
}

tera_table_t* table_open(tera_client_t* client, const char* table_name) {
	char *err = NULL;
	tera_table_t* table = tera_table_open(client, table_name, &err);
	if (err != NULL) {
		fprintf(stderr, "tera table open error: %s.\n", err);
		free(err);
	}
	return table;
}

