/*
 * Copyright 2022, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <avm_values/code.hpp>

void RunningCode::commitCodeToCore(
    const std::map<uint64_t, uint64_t>& segment_counts) const {
    const std::unique_lock<std::shared_mutex> lock(mutex);
    {
        // If our parent is an EphemeralBarrier, prune all the way down to the
        // CoreCode
        auto ephemeral = dynamic_cast<EphemeralBarrier*>(parent.get());
        if (ephemeral != nullptr) {
            parent = ephemeral->parent;
            while (true) {
                auto running_parent = dynamic_cast<RunningCode*>(parent.get());
                if (running_parent == nullptr) {
                    break;
                }
                parent = running_parent->getParent();
            }
        }
    }
    parent->commitCodeToCore(segment_counts);
    auto root_segments = parent->getRootSegments();
    auto it = segment_counts.lower_bound(impl->first_segment);
    auto end = segment_counts.lower_bound(impl->nextSegmentNum());
    for (; it != end; ++it) {
        auto segment = impl->getSegment(it->first);
        if (segment == nullptr) {
            continue;
        }
        auto inserted =
            root_segments.segments->insert(std::make_pair(it->first, segment));
        // Verify that the element didn't exist previously
        assert(inserted.second);
        if (!inserted.second) {
            throw std::runtime_error(
                "code segment id collision when filling in code");
        }
    }
    if (impl->nextSegmentNum() > *root_segments.next_segment_num) {
        *root_segments.next_segment_num = impl->nextSegmentNum();
    }
}
