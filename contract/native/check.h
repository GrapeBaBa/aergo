/**
 * @file    check.h
 * @copyright defined in aergo/LICENSE.txt
 */

#ifndef _CHECK_H
#define _CHECK_H

#include "common.h"

#include "ast.h"
#include "ast_blk.h"
#include "ast_exp.h"
#include "ast_stmt.h"
#include "ast_id.h"
#include "ast_meta.h"

typedef struct check_s {
    ast_blk_t *root;

    ast_blk_t *blk;     // current block
} check_t;

void check(ast_t *ast, flag_t flag);

ast_id_t *check_search_id(check_t *ctx, int num, char *name);

#endif /* ! _CHECK_H */