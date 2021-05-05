
#include <avm/machinestate/machinestate.hpp>
#include <avm/machinestate/runwasm.hpp>

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <wasm.h>
#include <wasmtime.h>

wasm_trap_t* cb_get_length(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    printf("Calling back closure...\n");

    results->data[0].kind = WASM_I32;
    results->data[0].of.i32 = (int32_t)dta->buffer_len;
    return NULL;
}

wasm_trap_t* cb_set_length(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    printf("Calling back closure...\n");

    if (args->data[0].kind == WASM_I32) {
        dta->buffer_len = args->data[0].of.i32;
    } else if (args->data[0].kind == WASM_I64) {
        dta->buffer_len = args->data[0].of.i64;
    }
    return NULL;
}

wasm_trap_t* cb_usegas(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    printf("Calling back closure...\n");

    if (args->data[0].kind == WASM_I32) {
        dta->gas_left -= args->data[0].of.i32;
    } else if (args->data[0].kind == WASM_I64) {
        dta->gas_left -= args->data[0].of.i64;
    }
    return NULL;
}

wasm_trap_t* cb_get_buffer(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t offset;
    printf("Calling back closure...\n");

    if (args->data[0].kind == WASM_I32) {
        offset = args->data[0].of.i32;
    } else if (args->data[0].kind == WASM_I64) {
        offset = args->data[0].of.i64;
    }
    results->data[0].kind = WASM_I32;
    results->data[0].of.i32 = (int32_t)dta->buffer.get(offset);
    return NULL;
}

wasm_trap_t* cb_set_buffer(void* env,
                           const wasm_val_vec_t* args,
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t offset;
    uint8_t v;
    printf("Calling back closure...\n");

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
                           wasm_val_vec_t* results) {
    WasmEnvData* dta = (WasmEnvData*)env;
    uint64_t offset;
    uint8_t v;
    printf("Calling back closure...\n");

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

void exit_with_error(wasmtime_error_t* error, wasm_trap_t* trap) {
    wasm_byte_vec_t error_message;
    if (error != NULL) {
        wasmtime_error_message(error, &error_message);
        wasmtime_error_delete(error);
    } else {
        wasm_trap_message(trap, &error_message);
        wasm_trap_delete(trap);
    }
    fprintf(stderr, "error %.*s\n", (int)error_message.size,
            error_message.data);
    wasm_byte_vec_delete(&error_message);
    exit(1);
}

RunWasm::RunWasm() {
    printf("Initializing... ????\n");
    wasm_engine_t* engine = wasm_engine_new();
    assert(engine != NULL);
    printf("Initialized...%x \n", engine);

    // With an engine we can create a *store* which is a long-lived group of
    // wasm modules.
    wasm_store_t* store = wasm_store_new(engine);
    assert(store != NULL);
    printf("Store...%x \n", store);

    // Read our input file, which in this case is a wasm text file.
    FILE* file = fopen("/home/sami/arb-os/wasm-tests/test-buffer.wasm", "r");
    assert(file != NULL);
    fseek(file, 0L, SEEK_END);
    size_t file_size = ftell(file);
    fseek(file, 0L, SEEK_SET);
    wasm_byte_vec_t wasm;
    wasm_byte_vec_new_uninitialized(&wasm, file_size);
    fread(wasm.data, file_size, 1, file);
    fclose(file);

    // Now that we've got our binary webassembly we can compile our module.
    printf("Compiling module...\n");
    wasm_module_t* module = NULL;
    wasmtime_error_t* error = wasmtime_module_new(engine, &wasm, &module);
    wasm_byte_vec_delete(&wasm);
    if (error != NULL) {
        std::cerr << "failed to compile module\n";
        exit_with_error(error, NULL);
    }

    WasmEnvData* env = &this->data;

    // Create external functions
    printf("Creating get len callback...\n");
    wasm_functype_t* callback_type_getlen =
        wasm_functype_new_0_1(wasm_valtype_new_i32());
    wasm_func_t* callback_func1 = wasm_func_new_with_env(
        store, callback_type_getlen, cb_get_length, (void*)env, NULL);
    wasm_functype_delete(callback_type_getlen);

    printf("Creating set len callback...\n");
    wasm_functype_t* callback_type_setlen =
        wasm_functype_new_1_0(wasm_valtype_new_i32());
    wasm_func_t* callback_func2 = wasm_func_new_with_env(
        store, callback_type_setlen, cb_set_length, (void*)env, NULL);
    wasm_func_t* callback_func6 = wasm_func_new_with_env(
        store, callback_type_setlen, cb_usegas, (void*)env, NULL);
    wasm_functype_delete(callback_type_setlen);

    printf("Creating get buf callback...\n");
    wasm_functype_t* callback_type_getbuf =
        wasm_functype_new_1_1(wasm_valtype_new_i32(), wasm_valtype_new_i32());
    wasm_func_t* callback_func3 = wasm_func_new_with_env(
        store, callback_type_getbuf, cb_get_buffer, (void*)env, NULL);
    wasm_functype_delete(callback_type_getbuf);

    printf("Creating set buf callback...\n");
    wasm_functype_t* callback_type_setbuf =
        wasm_functype_new_2_0(wasm_valtype_new_i32(), wasm_valtype_new_i32());
    wasm_func_t* callback_func4 = wasm_func_new_with_env(
        store, callback_type_setbuf, cb_set_buffer, (void*)env, NULL);

    printf("Creating write extra callback...\n");
    wasm_func_t* callback_func5 = wasm_func_new_with_env(
        store, callback_type_setbuf, cb_write_extra, (void*)env, NULL);
    wasm_functype_delete(callback_type_setbuf);

    printf("Instantiating module...\n");

    wasm_importtype_vec_t import_vec;
    wasm_module_imports(module, &import_vec);

    wasm_instance_t* instance = NULL;
    wasm_extern_t* imports[import_vec.size];
    for (int i = 0; i < import_vec.size; i++) {
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
        }
    }
    /*
    wasm_extern_t* imports[] = {
        wasm_func_as_extern(callback_func1),
        wasm_func_as_extern(callback_func2),
        wasm_func_as_extern(callback_func3),
        wasm_func_as_extern(callback_func4),
        wasm_func_as_extern(callback_func5)
    };*/
    wasm_extern_vec_t imports_vec;
    printf("Extracting export...\n");
    wasm_extern_vec_new(&imports_vec, import_vec.size, imports);
    error = wasmtime_instance_new(store, module, &imports_vec, &instance, &this->trap);
    if (instance == NULL)
        exit_with_error(error, trap);

    // Lookup our `run` export function
    printf("Extracting export...\n");
    wasm_extern_vec_t externs;
    wasm_instance_exports(instance, &externs);
    run = wasm_extern_as_func(externs.data[0]);

}

std::pair<Buffer, uint64_t> RunWasm::run_wasm(Buffer buf, uint64_t len) {
    data.buffer = buf;
    data.buffer_len = len;
    wasm_val_t arg_params[1];
    arg_params[0].kind = WASM_I64;
    arg_params[0].of.i64 = 123;
    wasm_val_vec_t args_vec;
    wasm_val_vec_new(&args_vec, 1, arg_params);

    wasm_val_t res_params[1];
    wasm_val_vec_t results_vec;
    res_params[0].kind = WASM_I32;
    res_params[0].of.i64 = 123;
    wasm_val_vec_new(&results_vec, 1, res_params);

    wasmtime_error_t* error = wasmtime_func_call(run, &args_vec, &results_vec, &trap);
    if (error != NULL || trap != NULL)
        exit_with_error(error, trap);
    return {data.buffer, data.buffer_len};

}

RunWasm runner;

std::pair<Buffer, uint64_t> run_wasm(Buffer buf, uint64_t len) {
    auto a = runner.run_wasm(Buffer(), 123);
    return a;
}
