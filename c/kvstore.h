#include <stdio.h>
#include <stdlib.h>
#include <tera_c.h>
bool table_put_kv_sync(tera_table_t* table, const char* key, int keylen,
                       const char* value, int vallen, int ttl) {
    char* err = NULL;
    bool ret = tera_table_put_kv(table, key, keylen, value, vallen, ttl, &err);
    if (!ret) {
        fprintf(stderr, "tera put kv error: %s.\n", err);
        free(err);
    }
    return ret;
}

void put_callback(void* mu) {
    int64_t err = tera_row_mutation_get_status_code((tera_row_mutation_t*)mu);
    if (err != 0) {
        fprintf(stderr, "tera put kv error: %d.\n", (int)err);
    }
    tera_row_mutation_destroy(mu);
}

void table_put_kv_async(tera_table_t* table, const char* key, int keylen,
                        const char* value, int vallen, int ttl) {
    char* err = NULL;
    tera_row_mutation_t* mu = tera_row_mutation(table, key, keylen);
    tera_row_mutation_put_kv(mu, value, vallen, ttl);
    tera_row_mutation_set_callback(mu, put_callback);
    tera_table_apply_mutation(table, mu);
}

char* table_get_kv_sync(tera_table_t* table, const char* key, int keylen, int* vallen) {
    uint64_t vlen = 0;
    char* err = NULL;
    char* value;
    bool ret = tera_table_get(table, key, keylen, "", "", 0, &value, &vlen, &err, 0);
    if (ret) {
        *vallen = (int)vlen;
    } else {
        *vallen = -1;
        fprintf(stderr, "tera get kv error: %s.\n", err);
        free(err);
    }
    return value;
}

bool table_delete_kv_sync(tera_table_t* table, const char* key, int keylen) {
    bool ret = tera_table_delete(table, key, keylen, "", "", 0);
    if (!ret) {
        fprintf(stderr, "tera delete error.\n");
    }
    return ret;
}

char* scanner_key(tera_result_stream_t* stream, int* keylen) {
    uint64_t len = 0;
    char* key;
    tera_result_stream_row_name(stream, &key, &len);
    *keylen = (int)len;
    return key;
}

char* scanner_value(tera_result_stream_t* stream, int* vallen) {
    uint64_t len = 0;
    char* value;
    tera_result_stream_value(stream, &value, &len);
    *vallen = (int)len;
    return value;
}
