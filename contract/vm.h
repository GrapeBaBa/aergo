#ifndef _VM_H
#define _VM_H

#include <lualib.h>
#include <lauxlib.h>
#include <luajit.h>
#include "sqlite3-binding.h"

typedef struct blockchain_ctx {
    char *stateKey;
    char *sender;
    char *contractId;
    char *txHash;
    unsigned long long blockHeight;
    long long timestamp;
    char *node;
    int confirmed;
    int isQuery;
    int transferFailed;
    sqlite3 *db;
    int service;
} bc_ctx_t;

lua_State *vm_newstate();
int vm_isnil(lua_State *L, int idx);
void vm_getfield(lua_State *L, const char *name);
void vm_remove_construct(lua_State *L, const char *constructName);
const char *vm_loadbuff(lua_State *L, const char *code, size_t sz, bc_ctx_t *bc_ctx);
const char *vm_pcall(lua_State *L, int argc, int* nresult);
const char *vm_get_json_ret(lua_State *L, int nresult);
const char *vm_tostring(lua_State *L, int idx);
const char *vm_copy_result(lua_State *L, lua_State *target, int cnt);
void bc_ctx_delete(bc_ctx_t *bcctx);

#endif /* _VM_H */
