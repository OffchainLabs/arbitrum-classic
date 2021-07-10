/*
 * Copyright 2021, Eigen
 * TODO: Lisence is add here
 */

#ifndef tee_task_hpp
#define tee_task_hpp

#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#include <arpa/inet.h>
#include <string.h>

#include <netinet/in.h>



int submit_task(const char* method, const char* args, const char* uid,
  const char* token, char** output, size_t* output_size);

int init(const char* pub, const char* pri, const char* conf, int32_t port1);

int release();

#endif