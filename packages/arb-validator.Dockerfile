### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM offchainlabs/cpp-base:0.2.5 as arb-avm-cpp

# Copy source code
COPY --chown=user arb-avm-cpp/ ./
# Copy build cache
COPY --chown=user --from=arb-validator /cpp-build build/
# Build arb-avm-cpp
RUN mkdir -p build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_TESTING=0 && \
    cmake --build . -j $(nproc) && \
    cd ../ && \
    ./scripts/install-cmachine-build

FROM offchainlabs/backend-base:0.2.8 as arb-validator-builder

# Build dependencies
COPY --chown=user arb-avm-cpp/go.* /home/user/arb-avm-cpp/
COPY --chown=user arb-util/go.* /home/user/arb-util/
COPY --chown=user arb-validator/go.* /home/user/arb-validator/
COPY --chown=user arb-validator-core/go.* /home/user/arb-validator-core/
COPY --chown=user arb-provider-go/go.* /home/user/arb-provider-go/
COPY --chown=user arb-checkpointer/go.* /home/user/arb-checkpointer/
COPY --chown=user arb-evm/go.* /home/user/arb-evm/
COPY --chown=user arb-tx-aggregator/go.* /home/user/arb-tx-aggregator/
RUN cd arb-validator && go mod download && cd ../arb-tx-aggregator && go mod download
# Copy source code
COPY --from=arb-avm-cpp /home/user/go.mod /home/user/go.sum /home/user/arb-avm-cpp/
COPY --from=arb-avm-cpp /home/user/cavm/*.h /home/user/arb-avm-cpp/cavm/
COPY --from=arb-avm-cpp /home/user/cmachine /home/user/arb-avm-cpp/cmachine/

COPY --chown=user arb-util/ /home/user/arb-util/
COPY --chown=user arb-avm-cpp/ /home/user/arb-avm-cpp/
COPY --chown=user arb-validator/ /home/user/arb-validator/
COPY --chown=user arb-validator-core/ /home/user/arb-validator-core/
COPY --chown=user arb-provider-go/ /home/user/arb-provider-go/
COPY --chown=user arb-checkpointer/ /home/user/arb-checkpointer/
COPY --chown=user arb-evm/ /home/user/arb-evm/
COPY --chown=user arb-tx-aggregator/ /home/user/arb-tx-aggregator/
# Copy build cache
COPY --from=arb-validator --chown=user /build /home/user/.cache/go-build
# Build arb-validator
RUN cd arb-validator && go install -v ./cmd/arb-validator && \
    cd ../arb-tx-aggregator && go install -v ./cmd/arb-tx-aggregator

FROM offchainlabs/cpp-base:0.2.5 as arb-validator
# Export binary

COPY --chown=user --from=arb-validator-builder /home/user/go/bin /home/user/go/bin

# Build cache
COPY --chown=user --from=arb-validator-builder /home/user/.cache/go-build /build
COPY --from=arb-avm-cpp /home/user/build /cpp-build

ENTRYPOINT ["/home/user/go/bin/arb-validator"]
EXPOSE 1235 1236
