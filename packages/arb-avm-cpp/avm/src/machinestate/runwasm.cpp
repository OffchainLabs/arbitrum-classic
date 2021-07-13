
#include <avm/machinestate/machinestate.hpp>
#include <avm/machinestate/runwasm.hpp>

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
// #include <wasmer/wasmer.hh>
#include <wasmer/wasmer.h>
// #include <wasmtime.h>

wasm_trap_t* cb_get_length(void* env,
                           const wasm_val_vec_t*,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    // printf("get len... %i\n", (int32_t)dta->buffer_len);

    results->data[0].kind = WASM_I32;
    results->data[0].of.i32 = (int32_t)dta->buffer_len;
    return NULL;
}

wasm_trap_t* cb_set_length(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;
    // printf("set len...\n");

    if (args->data[0].kind == WASM_I32) {
        dta->buffer_len = args->data[0].of.i32;
    } else if (args->data[0].kind == WASM_I64) {
        dta->buffer_len = args->data[0].of.i64;
    }
    return NULL;
}

wasm_trap_t* cb_usegas(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;
    // printf("use gas... %li %li\n", dta->gas_left, args->size);

    if (args->data[0].kind == WASM_I32) {
        dta->gas_left -= args->data[0].of.i32;
    } else if (args->data[0].kind == WASM_I64) {
        dta->gas_left -= args->data[0].of.i64;
    }
    // printf("used gas... %li %li\n", dta->gas_left, args->size);
    return NULL;
}

value simple_table_aux(int level) {
    if (level == 0) {
        return Tuple(0,0,0,0,0,0,0,0);
    }
    return Tuple(
        simple_table_aux(level - 1),
        simple_table_aux(level - 1),
        simple_table_aux(level - 1),
        simple_table_aux(level - 1),
        simple_table_aux(level - 1),
        simple_table_aux(level - 1),
        simple_table_aux(level - 1),
        simple_table_aux(level - 1)
    );
}

const int LEVEL = 5;

value simple_table() {
    return simple_table_aux(LEVEL - 1);
}

wasm_trap_t* cb_uint_immed(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;

    if (args->data[0].kind == WASM_I32) {
        // read uint from memory
        auto mem = (const char*)wasm_memory_data(dta->memory);
        mem += + args->data[0].of.i32;
        uint256_t num = deserializeUint256t(mem);
        // std::cerr << "load num " << num << "\n";
        dta->immed = std::make_shared<value>(num);
    }
    return NULL;
}

wasm_trap_t* cb_special_immed(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;

    Tuple t(Buffer(), 0);

    // std::cerr << "special immed\n";
    dta->immed = std::make_shared<value>(t);
    return NULL;
}

wasm_trap_t* cb_global_immed(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;

    Tuple t(0, 0, 0, 0, 0, 0, 0, 0);

    // std::cerr << "global immed\n";
    dta->immed = std::make_shared<value>(t);
    return NULL;
}

wasm_trap_t* cb_push_insn(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;

    if (args->data[0].kind == WASM_I32) {
        dta->insn->push_back(Operation(OpCode(args->data[0].of.i32)));
    }
    return NULL;
}

wasm_trap_t* cb_push_immed(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;

    if (args->data[0].kind == WASM_I32) {
        dta->insn->push_back(Operation(OpCode(args->data[0].of.i32), *dta->immed));
    }
    return NULL;
}

wasm_trap_t* cb_cptable(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;

    uint64_t len = dta->insn->size();

    if (args->data[0].kind == WASM_I32) {
        dta->table.push_back(std::pair<uint64_t, uint64_t>(args->data[0].of.i32, len));
    }
    return NULL;
}

wasm_trap_t* cb_get_buffer(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t offset;
    // printf("read buf...\n");

    if (args->data[0].kind == WASM_I32) {
        offset = args->data[0].of.i32;
    } else if (args->data[0].kind == WASM_I64) {
        offset = args->data[0].of.i64;
    }
    results->data[0].kind = WASM_I32;
    results->data[0].of.i32 = ((int32_t)dta->buffer.get(offset)) & 0xff;
    return NULL;
}

wasm_trap_t* cb_rvec(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t offset;
    uint64_t ptr;
    uint64_t len;
    // printf("read buf...\n");

    if (args->data[0].kind == WASM_I32) {
        auto mem = (uint8_t*)wasm_memory_data(dta->memory);
        ptr = args->data[0].of.i32;
        offset = args->data[1].of.i32;
        len = args->data[2].of.i32;
        for (int i = 0; i < len; i++) {
            mem[ptr+i] = dta->buffer.get(offset+i);
        }
    }
    return NULL;
}

wasm_trap_t* cb_tuplebytes(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t idx;
    uint64_t ptr;
    // std::cerr << "tuplebytes...\n";

    if (args->data[0].kind == WASM_I32) {
        auto mem = (uint8_t*)wasm_memory_data(dta->memory);
        ptr = args->data[0].of.i32;
        idx = args->data[1].of.i32;

        auto immed = *dta->immed;
        auto t = std::get_if<Tuple>(&immed);
        auto buf = t->get_element(idx);
        auto num = std::get_if<uint256_t>(&buf);

        std::vector<uint8_t> buffer;
        marshal_uint256_t(*num, buffer);

        for (int i = 0; i < 32; i++) {
            mem[ptr+i] = buffer[i];
        }
    }
    return NULL;
}

wasm_trap_t* cb_tuple2bytes(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t idx2;
    uint64_t idx;
    uint64_t ptr;
    // std::cerr << "tuple2bytes...\n";

    if (args->data[0].kind == WASM_I32) {
        auto mem = (uint8_t*)wasm_memory_data(dta->memory);
        ptr = args->data[0].of.i32;
        idx = args->data[1].of.i32;
        idx2 = args->data[2].of.i32;

        auto immed = *dta->immed;
        auto t = std::get_if<Tuple>(&immed);
        auto t1 = t->get_element(idx);
        auto t2 = std::get_if<Tuple>(&t1);
        auto buf = t2->get_element(idx2);
        auto num = std::get_if<uint256_t>(&buf);

        std::vector<uint8_t> buffer;
        marshal_uint256_t(*num, buffer);

        for (int i = 0; i < 32; i++) {
            mem[ptr+i] = buffer[i];
        }
    }
    return NULL;
}

wasm_trap_t* cb_tuple2buffer(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t idx2;
    uint64_t idx;
    uint64_t ptr;
    uint64_t len;
    // std::cerr << "tuple2buffer...\n";

    if (args->data[0].kind == WASM_I32) {
        auto mem = (uint8_t*)wasm_memory_data(dta->memory);
        ptr = args->data[0].of.i32;
        idx = args->data[1].of.i32;
        idx2 = args->data[2].of.i32;
        len = args->data[3].of.i32;

        auto immed = *dta->immed;
        auto t = std::get_if<Tuple>(&immed);
        auto t1 = t->get_element(idx);
        auto t2 = std::get_if<Tuple>(&t1);
        auto buf = t2->get_element(idx2);
        auto buffer = std::get_if<Buffer>(&buf);

        for (int i = 0; i < len; i++) {
            mem[ptr+i] = buffer->get(i);
        }
    }
    return NULL;
}

wasm_trap_t* cb_wvec(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t offset;
    uint64_t ptr;
    uint64_t len;
    // std::cerr << "write vec...\n";

    if (args->data[0].kind == WASM_I32) {
        auto mem = (uint8_t*)wasm_memory_data(dta->memory);
        ptr = args->data[0].of.i32;
        offset = args->data[1].of.i32;
        len = args->data[2].of.i32;
        Buffer b = dta->buffer;
        std::vector<uint8_t> bytes;
        for (int i = 0; i < len; i++) {
            bytes.push_back(mem[ptr+i]);
            if ((offset+i) % 32 == 31) {
                auto loc = offset + i + 1 - bytes.size();
                // std::cerr << "offset " << offset << " i " << i << " size " << bytes.size() << " loc " << loc << "\n";
                b = b.set_many(loc, bytes);
                bytes.clear();
            }
        }
        auto loc = offset + len - bytes.size();
        // std::cerr << "offset " << offset << " size " << bytes.size() << " loc " << loc << "\n";
        if (bytes.size() > 0) b = b.set_many(loc, bytes);
        dta->buffer = b;
    }
    return NULL;
}

wasm_trap_t* cb_set_buffer(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t offset;
    uint8_t v;
    // printf("write buf...\n");

    if (args->data[0].kind == WASM_I32) {
        offset = args->data[0].of.i32;
    } else if (args->data[0].kind == WASM_I64) {
        offset = args->data[0].of.i64;
    }
    if (args->data[1].kind == WASM_I32) {
        v = args->data[1].of.i32;
    } else if (args->data[1].kind == WASM_I64) {
        v = args->data[1].of.i64;
    }
    dta->buffer = dta->buffer.set(offset, v);
    return NULL;
}

wasm_trap_t* cb_write_extra(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t*) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t offset;
    uint8_t v;
    // printf("wextra...\n");

    if (args->data[0].kind == WASM_I32) {
        offset = args->data[0].of.i32;
    } else if (args->data[0].kind == WASM_I64) {
        offset = args->data[0].of.i64;
    }
    if (args->data[1].kind == WASM_I32) {
        v = args->data[1].of.i32;
    } else if (args->data[1].kind == WASM_I64) {
        v = args->data[1].of.i64;
    }
    if (dta->extra.size() <= offset) {
        dta->extra.resize(offset+1, 0);
    }
    dta->extra[offset] = v;
    return NULL;
}

RunWasm::RunWasm(std::string fname) {
    data = new WasmEnvData();
    // Read our input file, which in this case is a wasm text file.
    FILE* file = fopen(fname.c_str(), "r");
    assert(file != NULL);
    fseek(file, 0L, SEEK_END);
    size_t file_size = ftell(file);
    fseek(file, 0L, SEEK_SET);
    wasm_byte_vec_t wasm;
    wasm_byte_vec_new_uninitialized(&wasm, file_size);
    fread(wasm.data, file_size, 1, file);
    fclose(file);
    // std::cerr << "File ????\n";
    init(wasm);
}

RunWasm::RunWasm(std::vector<uint8_t> &buf) {
    data = new WasmEnvData();
    wasm_byte_vec_t wasm;
    wasm_byte_vec_new_uninitialized(&wasm, buf.size());
    for (int i = 0; i < buf.size(); i++) {
        wasm.data[i] = buf[i];
    }
    // std::cerr << "Got size " << buf.size() << "\n";
    init(wasm);
}

void RunWasm::init(wasm_byte_vec_t wasm) {
    // printf("Initializing... ????\n");
    wasm_config_t* config = wasm_config_new();
    wasm_config_set_compiler(config, CRANELIFT);
    // wasm_config_set_compiler(config, SINGLEPASS);
    wasm_engine_t* engine = wasm_engine_new_with_config(config);
    assert(engine != NULL);
    // printf("Initialized...%x \n", engine);

    // With an engine we can create a *store* which is a long-lived group of
    // wasm modules.
    wasm_store_t* store = wasm_store_new(engine);
    assert(store != NULL);
    // printf("Store...%x \n", store);

    // Now that we've got our binary webassembly we can compile our module.
    printf("Compiling module...\n");
    wasm_module_t* module = wasm_module_new(store, &wasm);
    wasm_byte_vec_delete(&wasm);
    if (module == NULL) {
        std::cerr << "failed to compile module\n";
        return;
    }

    WasmEnvData* env = this->data;

    // Create external functions
    // printf("Creating get len callback...\n");
    wasm_functype_t* callback_type_getlen =
        wasm_functype_new_0_1(wasm_valtype_new_i32());
    wasm_func_t* callback_func1 = wasm_func_new_with_env(
        store, callback_type_getlen, cb_get_length, (void*)env, NULL);
    wasm_functype_delete(callback_type_getlen);

    // printf("Creating set len callback...\n");
    wasm_functype_t* callback_type_setlen =
        wasm_functype_new_1_0(wasm_valtype_new_i32());
    wasm_func_t* callback_func2 = wasm_func_new_with_env(
        store, callback_type_setlen, cb_set_length, (void*)env, NULL);
    wasm_func_t* callback_func6 = wasm_func_new_with_env(
        store, callback_type_setlen, cb_usegas, (void*)env, NULL);

    wasm_func_t* callback_func_uint = wasm_func_new_with_env(
        store, callback_type_setlen, cb_uint_immed, (void*)env, NULL);
    wasm_func_t* callback_func_special = wasm_func_new_with_env(
        store, callback_type_setlen, cb_special_immed, (void*)env, NULL);
    wasm_func_t* callback_func_global = wasm_func_new_with_env(
        store, callback_type_setlen, cb_global_immed, (void*)env, NULL);
    wasm_func_t* callback_func_push = wasm_func_new_with_env(
        store, callback_type_setlen, cb_push_insn, (void*)env, NULL);
    wasm_func_t* callback_func_push_immed = wasm_func_new_with_env(
        store, callback_type_setlen, cb_push_immed, (void*)env, NULL);
    wasm_func_t* callback_func_cptable = wasm_func_new_with_env(
        store, callback_type_setlen, cb_cptable, (void*)env, NULL);

    wasm_functype_delete(callback_type_setlen);

    // printf("Creating get buf callback...\n");
    wasm_functype_t* callback_type_getbuf =
        wasm_functype_new_1_1(wasm_valtype_new_i32(), wasm_valtype_new_i32());
    wasm_func_t* callback_func3 = wasm_func_new_with_env(
        store, callback_type_getbuf, cb_get_buffer, (void*)env, NULL);
    wasm_functype_delete(callback_type_getbuf);

    // printf("Creating set buf callback...\n");
    wasm_functype_t* callback_type_setbuf =
        wasm_functype_new_2_0(wasm_valtype_new_i32(), wasm_valtype_new_i32());
    wasm_func_t* callback_func4 = wasm_func_new_with_env(
        store, callback_type_setbuf, cb_set_buffer, (void*)env, NULL);

    // printf("Creating write extra callback...\n");
    wasm_func_t* callback_func5 = wasm_func_new_with_env(
        store, callback_type_setbuf, cb_write_extra, (void*)env, NULL);
    wasm_func_t* callback_tuplebytes = wasm_func_new_with_env(
        store, callback_type_setbuf, cb_tuplebytes, (void*)env, NULL);
    wasm_functype_delete(callback_type_setbuf);

    wasm_functype_t* callback_type_rvec =
        wasm_functype_new_3_0(wasm_valtype_new_i32(), wasm_valtype_new_i32(), wasm_valtype_new_i32());
    wasm_func_t* callback_func_rvec = wasm_func_new_with_env(
        store, callback_type_setbuf, cb_rvec, (void*)env, NULL);
    wasm_func_t* callback_func_wvec = wasm_func_new_with_env(
        store, callback_type_setbuf, cb_wvec, (void*)env, NULL);
    wasm_func_t* callback_tuple2bytes = wasm_func_new_with_env(
        store, callback_type_setbuf, cb_tuple2bytes, (void*)env, NULL);
    wasm_functype_delete(callback_type_rvec);

    wasm_valtype_t* ps[4] = {wasm_valtype_new_i32(), wasm_valtype_new_i32(), wasm_valtype_new_i32(), wasm_valtype_new_i32()};
    wasm_valtype_vec_t params, results;
    wasm_valtype_vec_new(&params, 4, ps);
    wasm_valtype_vec_new_empty(&results);
    wasm_functype_t* callback_type_tuple2buffer = wasm_functype_new(&params, &results);
    wasm_func_t* callback_func_tuple2buffer = wasm_func_new_with_env(
        store, callback_type_tuple2buffer, cb_tuple2buffer, (void*)env, NULL);
    wasm_functype_delete(callback_type_tuple2buffer);

    // printf("Instantiating module...\n");

    wasm_importtype_vec_t import_vec;
    wasm_module_imports(module, &import_vec);

    std::cerr << "Making imports " << import_vec.size << "\n";

    wasm_extern_t* imports[import_vec.size];
    for (uint64_t i = 0; i < import_vec.size; i++) {
        auto imp = import_vec.data[i];
        auto name = wasm_importtype_name(imp);
        std::string str = name->data;
        if (str.find("read") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func3);
        } else if (str.find("write") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func4);
        } else if (str.find("usegas") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func6);
        } else if (str.find("wextra") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func5);
        } else if (str.find("getlen") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func1);
        } else if (str.find("setlen") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func2);
        } else if (str.find("uintimmed") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func_uint);
        } else if (str.find("globalimmed") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func_global);
        } else if (str.find("specialimmed") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func_special);
        } else if (str.find("pushimmed") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func_push_immed);
        } else if (str.find("pushinst") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func_push);
        } else if (str.find("cptable") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func_cptable);
        } else if (str.find("rvec") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func_rvec);
        } else if (str.find("wvec") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func_wvec);
        } else if (str.find("tuplebytes") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_tuplebytes);
        } else if (str.find("tuple2bytes") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_tuple2bytes);
        } else if (str.find("tuple2buffer") != std::string::npos) {
            imports[i] = wasm_func_as_extern(callback_func_tuple2buffer);
        } else {
            imports[i] = wasm_func_as_extern(callback_func2);
        }
    }
    wasm_extern_vec_t imports_vec;
    // printf("Extracting export...\n");
    wasm_extern_vec_new(&imports_vec, import_vec.size, imports);
    wasm_instance_t* instance = wasm_instance_new(store, module, &imports_vec, &this->trap);
    if (instance == NULL) {
        std::cerr << "Cannot instantiate\n";
        exit(1);
    }

    // Get memory from instance


    // Lookup our `run` export function
    // printf("Extracting export...\n");
    wasm_extern_vec_t externs;
    wasm_instance_exports(instance, &externs);
    for (uint64_t i = 0; i < externs.size; i++) {
        auto kind = wasm_extern_kind(externs.data[i]);
        if (kind == WASM_EXTERN_FUNC) {
            run = wasm_extern_as_func(externs.data[i]);
        } else if (kind == WASM_EXTERN_MEMORY) {
            std::cerr << "found memory\n";
            data->memory = wasm_extern_as_memory(externs.data[i]);
        }
    }

    // printf("externs %i\n", externs.size);

}

WasmResult RunWasm::run_wasm(Buffer buf, uint64_t len) {
    data->buffer = buf;
    data->buffer_len = len;
    wasm_val_t arg_params[1];
    arg_params[0].kind = WASM_I64;
    arg_params[0].of.i64 = 123;
    wasm_val_vec_t args_vec;
    wasm_val_vec_new_empty(&args_vec);

    wasm_val_t res_params[1];
    wasm_val_vec_t results_vec;
    res_params[0].kind = WASM_I32;
    res_params[0].of.i64 = 123;
    wasm_val_vec_new(&results_vec, 1, res_params);

    data->gas_left = 1000000;
    data->extra.resize(0);

    data->table = std::vector<std::pair<uint64_t, uint64_t>>();
    data->immed = std::make_shared<value>(0);
    data->insn = std::make_shared<std::vector<Operation>>();

    std::cerr << "Running wasm\n";
    if (wasm_func_call(run, &args_vec, &results_vec)) {
        std::cerr << "Error running wasm\n";
    }
    std::cerr << "Ran wasm\n";

    return {data->buffer_len, data->buffer, data->extra, data->gas_left, data->immed, data->insn, data->table};

}

WasmResult RunWasm::run_wasm(Buffer buf, uint64_t len, value v) {
    data->buffer = buf;
    data->buffer_len = len;
    wasm_val_t arg_params[1];
    arg_params[0].kind = WASM_I64;
    arg_params[0].of.i64 = 123;
    wasm_val_vec_t args_vec;
    wasm_val_vec_new_empty(&args_vec);

    wasm_val_t res_params[1];
    wasm_val_vec_t results_vec;
    res_params[0].kind = WASM_I32;
    res_params[0].of.i64 = 123;
    wasm_val_vec_new(&results_vec, 1, res_params);

    data->gas_left = 1000000;
    data->extra.resize(0);

    data->table = std::vector<std::pair<uint64_t, uint64_t>>();
    data->immed = std::make_shared<value>(v);
    data->insn = std::make_shared<std::vector<Operation>>();

    // std::cerr << "Running wasm\n";
    if (wasm_func_call(run, &args_vec, &results_vec)) {
        std::cerr << "Error running wasm\n";
    }
    // std::cerr << "Ran wasm\n";

    return {data->buffer_len, data->buffer, data->extra, data->gas_left, data->immed, data->insn, data->table};

}

// RunWasm runner("/home/sami/arb-os/wasm-tests/test-buffer.wasm");

std::pair<Buffer, uint64_t> run_wasm(Buffer buf, uint64_t len) {
    return {buf, len};
    // auto a = runner.run_wasm(Buffer(), 123);
    // return a;
}
