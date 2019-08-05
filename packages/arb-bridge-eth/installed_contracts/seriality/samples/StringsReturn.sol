pragma solidity ^0.4.16;

import "./Seriality.sol";

contract StringsReturn is Seriality {
  
    function stringCaller() public returns(  string memory out1,
                                            string memory out2,
                                            string memory out3,
                                            string memory out4,
                                            string memory out5)
                                            {
        
        bytes memory buffer = new bytes(320);
        uint offset = stringCallee(buffer);
        
        //deserializing
        out1 = new string (getStringSize(offset, buffer));
        bytesToString(offset, buffer, bytes(out1));
        offset -= sizeOfString(out1);
        
        out2 = new string (getStringSize(offset, buffer));
        bytesToString(offset, buffer, bytes(out2));
        offset -= sizeOfString(out2);
        
        out3 = new string (getStringSize(offset, buffer));
        bytesToString(offset, buffer, bytes(out3));
        offset -= sizeOfString(out3);
        
        out4 = new string (getStringSize(offset, buffer));
        bytesToString(offset, buffer, bytes(out4));
        offset -= sizeOfString(out4);
        
        out5 = new string (getStringSize(offset, buffer));
        bytesToString(offset, buffer, bytes(out5));
      
    }
    
    function stringCallee(bytes memory buffer) public returns (uint buffer_size) {
    
        string memory out1  = new string(32); 
        string memory out2  = new string(32);        
        string memory out3  = new string(32);
        string memory out4  = new string(32);        
        string memory out5  = new string(32);
        
        out1 = "Come on baby lets dance!";
        out2 = "May I buy you a drink?";
        out3 = "I am an itinerant programmer";
        out4 = "Inam javab lashi!";
        out5 = "Bia inja dahan service";

        // Serializing
        buffer_size = sizeOfString(out5) +
                       sizeOfString(out4) + 
                       sizeOfString(out3) + 
                       sizeOfString(out2) +
                       sizeOfString(out1);
                           
        uint offset = buffer_size;

        stringToBytes(offset, bytes(out1), buffer);
        offset -= sizeOfString(out1); 
        
        stringToBytes(offset, bytes(out2), buffer);
        offset -= sizeOfString(out2);

        stringToBytes(offset, bytes(out3), buffer);
        offset -= sizeOfString(out3); 
        
        stringToBytes(offset, bytes(out4), buffer);
        offset -= sizeOfString(out4); 
        
        stringToBytes(offset, bytes(out5), buffer);
        
        return buffer_size;
    }    
}
