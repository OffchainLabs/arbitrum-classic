//
//  checkpoint.cpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#include "avm/checkpoint.hpp"
#include <avm/tuple.hpp>
#include "avm/checkpointstorage.hpp"
#include "avm/machine.hpp"
#include "avm/machinestatedata.hpp"

MachineCheckPoints::MachineCheckPoints() {
    // storage = new CheckpointStorage   no work??
    storage.Intialize();
};

MachineCheckPoints::~MachineCheckPoints() {
    storage.Close();
}

// const machine?
Checkpoint MachineCheckPoints::SaveMachine(std::string machine_state_name,
                                           Machine& machine) {
    auto checkpoint_data = machine.getCheckPointData();
    auto pool = &machine.getPool();

    auto datastack_tuple =
        Tuple(checkpoint_data.stack.GetTupleRepresentation(pool),
              checkpoint_data.auxstack.GetTupleRepresentation(pool), pool);
    auto value_tuple =
        Tuple(checkpoint_data.staticVal, checkpoint_data.registerVal, pool);
    auto messagestack_tuple = Tuple(checkpoint_data.pendingInbox.messages,
                                    checkpoint_data.inbox.messages, pool);

    auto pc_codepoint = CodePoint();
    pc_codepoint.pc = checkpoint_data.pc;
    auto codepoint_tuple = Tuple(pc_codepoint, checkpoint_data.errpc, pool);

    auto list_of_token_pairs = checkpoint_data.balance.GetAllTokenPairs();
    std::vector<unsigned char> token_char_list;

    for (auto& pair : list_of_token_pairs) {
        auto tokentype = std::get<0>(pair);
        auto value = std::get<1>(pair);
        auto value_vector = ConvertToCharVector(value);

        token_char_list.insert(token_char_list.end(), std::begin(tokentype),
                               std::end(tokentype));
        token_char_list.insert(token_char_list.end(), value_vector.begin(),
                               value_vector.end());
    }

    auto status_value = (unsigned char)checkpoint_data.state;
    auto blockreasondata = SerializeBlockReason(checkpoint_data.blockReason);
}

// Machine MachineCheckPoints::RestoreMachine(std::string name, std::string
// contract_filename) {
//    auto machine = Machine(contract_filename);
//}
