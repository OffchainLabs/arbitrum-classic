pragma solidity ^0.4.16;

import "./Seriality.sol";


contract SerialitySample is Seriality {

    function testSample1() public returns(int n1, int8 n2, uint24 n3, string n4,string n5) {
        
        bytes memory buffer = new  bytes(200);
        string memory out4  = new string(32);        
        string memory out5  = new string(32);
        n4 = new string(32);
        n5 = new string(32);
        int     out1 = 34444445;
        int8    out2 = 87;
        uint24  out3 = 76545;
        out4 = "Copy kon lashi";
        out5 = "Bia inja dahan service";

        // Serializing
        uint offset = 200;

        intToBytes(offset, out2, buffer);
        offset -= sizeOfInt(8);

        uintToBytes(offset, out3, buffer);
        offset -= sizeOfUint(24);

        stringToBytes(offset, bytes(out5), buffer);
        offset -= sizeOfString(out5);

        stringToBytes(offset, bytes(out4), buffer);
        offset -= sizeOfString(out4);       

        intToBytes(offset, out1, buffer);
        offset -= sizeOfInt(256);
        // Deserializing
        offset = 200; 
            
        n2 = bytesToInt8(offset, buffer);
        offset -= sizeOfInt(8);

        n3 = bytesToUint24(offset, buffer);
        offset -= sizeOfUint(24);

        bytesToString(offset, buffer, bytes(n5));
        offset -= sizeOfString(out5);

        bytesToString(offset, buffer, bytes(n4));
        offset -= sizeOfString(out4);

        n1 = bytesToInt256(offset, buffer);
    }
    
    function testSample2() public returns(int8 n1, int24 n2, uint32 n3, int128 n4, address n5, address n6) {
        
        bytes memory buffer = new bytes(64);
        int8    out1 = -12;
        int24   out2 = 838860;
        uint32  out3 = 333333333;
        int128  out4 = -44444444444;
        address out5 = 0x15B7926835A7C2FD6D297E3ADECC5B45F7309F59;
        address out6 = 0x1CB5CF010E407AFC6249627BFD769D82D8DBBF71;
        
        // Serializing
        uint offset = 64;
        
        intToBytes(offset, out1, buffer);
        offset -= sizeOfInt(8);
        
        intToBytes(offset, out2, buffer);
        offset -= sizeOfUint(24);
        
        uintToBytes(offset, out3, buffer);
        offset -= sizeOfInt(32);
        
        intToBytes(offset, out4, buffer);
        offset -= sizeOfUint(128);
        
        addressToBytes(offset, out5, buffer);
        offset -= sizeOfAddress();
        
        addressToBytes(offset, out6, buffer);

        // Deserializing
        offset = 64; 
       
        n1 = bytesToInt8(offset, buffer);
        offset -= sizeOfInt(8);
        
        n2 = bytesToInt24(offset, buffer);
        offset -= sizeOfUint(24);
        
        n3 = bytesToUint8(offset, buffer);
        offset -= sizeOfInt(32);
        
        n4 = bytesToInt128(offset, buffer);
        offset -= sizeOfUint(128);
        
        n5 = bytesToAddress(offset, buffer);
        offset -= sizeOfAddress();
        
        n6 = bytesToAddress(offset, buffer);
    }
}
