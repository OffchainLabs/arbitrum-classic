// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

#ifndef EIGENTEE_H
#define EIGENTEE_H

#ifdef __cplusplus
extern "C" {
#endif

#include <sys/socket.h>

#include "tee/visibility.h"

typedef struct eigen eigen_t;
typedef struct eigen_enclave_info eigen_enclave_info_t;
typedef struct eigen_task eigen_task_t;
typedef struct eigen_taskinfo eigen_taskinfo_t;
typedef struct eigen_auditor_set eigen_auditor_set_t;

typedef struct sockaddr sockaddr_t;

typedef enum eigen_task_status_t {
  TASK_CREATED,
  TASK_READY,
  TASK_RUNNING,
  TASK_FINISHED,
  TASK_FAILED,
} eigen_task_status_t;

// Initialize logger
EIGENTEE_API int eigen_init();

// MesaTEE Context APIs
EIGENTEE_API eigen_t *eigen_context_new(const eigen_enclave_info_t *enclave_info_ptr,
                                           const char* user_id, const char* user_token,
                                           sockaddr_t * tms_addr);

EIGENTEE_API eigen_t* eigen_context_new2(const eigen_enclave_info_t* enclave_info_ptr,
                                           const char* user_id, const char* user_token,
                                           const char* tms_addr_ptr/*ip:port*/);

EIGENTEE_API int eigen_context_free(eigen_t *ctx_ptr);

// MesaTEE EnclaveInfo APIs
EIGENTEE_API eigen_enclave_info_t *
eigen_enclave_info_load(eigen_auditor_set_t *auditors_ptr, const char *enclave_info_file_path_ptr);
EIGENTEE_API int eigen_enclave_info_free(eigen_enclave_info_t *enclave_info_ptr);

// Auditor APIs
EIGENTEE_API eigen_auditor_set_t *eigen_auditor_set_new();
EIGENTEE_API int eigen_auditor_set_add_auditor(eigen_auditor_set_t *ptr,
                                                const char *pub_key_path, const char *sig_path);
EIGENTEE_API int eigen_auditor_set_free(eigen_auditor_set_t *ptr);

// MesaTEE Task APIs
EIGENTEE_API eigen_task_t *eigen_create_task(eigen_t *ctx_ptr, const char *func_name_ptr);
EIGENTEE_API int eigen_task_free(eigen_task_t *eigen_task_ptr);
EIGENTEE_API int eigen_task_invoke_with_payload(eigen_task_t *eigen_task_ptr, const char *payload_buf_ptr,
                                                 int payload_buf_len, char *result_buf_ptr, int result_buf_len);

#ifdef __cplusplus
} /* extern C */
#endif

#endif // EIGENTEE_H
