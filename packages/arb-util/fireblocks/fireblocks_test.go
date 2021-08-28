/*
 * Copyright 2021, Offchain Labs, Inc.
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

package fireblocks

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks/accounttype"
	"testing"
)

func TestWalletMaps(t *testing.T) {
	fireblocksConfig := configuration.WalletFireblocks{
		APIKey:          "",
		AssetId:         "ETH",
		BaseURL:         "http://google.com",
		ExternalWallets: "0x77FEE9b17F2a8395C682d1b0548Ef35Bc10a234C:foobar,0xc87EAe2F7699fb1D94e89EE0D8AB1B338cf538E0:snafu",
		InternalWallets: "0xe83B89654C58E40EB04011063Fb6a6623ed23f45:fibbage",
		SourceAddress:   "0xADB898c59B2cf6eBD09d58A8cB4F176074556C65",
		SourceType:      "VAULT_ACCOUNT",
		SSLKey:          "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDDEELXabkjQD28\n0uV59YyoiRftYbKMBVtvsyJi1PcnDqbd2D1nKVApchq7/BBAgJHAw4iauVUvw7UV\neDC5absXwx7kMSOizlZx7QXKYXarE27v0tLdJNHZMJhJFDqVmSQqIbVczbUOF/bB\nOYPvI0GUI7JK6NXYw/Nkj0Gho4SSD68wh4dpqHFp0i2FRM6/YduLe/0X4myUqmgM\nKHcZxv5RFJ3MZ+y/jmVbL9+p+3xmTcRKAhWustVyfcihO1pcXK/oFyhc/P+A+My+\nQxHdP34LfkTLD7JlH9tnbQbJwcr2dcVEdgJdKfGxUIaFiE48H+6Ldt2X0Dy+Fui3\nAWwCK495AgMBAAECggEAZFAtf6ggDK/R0BRI7NJJVrbB9lLixj1PC3989IR+4bgY\nOkglM7M8RUQn2XrERNeZnR8xVhhvZZPowTTBIejszUiLz2Ax6lzgedAhxTUgGNRy\nnaS2ygFQZbgm6JYdlqddwMj7AtHLBYdvX2kbn1RrOiqtWvaoWsRSm+lMIDr+Rzoh\nrbUVQGnH8mLCJwnMQsK6bVkiSwx4WYelplcvbnI0mkwDm3FKfBgAxmdHDNaCUXE1\ndO0LCJNedBzhyY7BLN3vn1gw04fjtqid+fGtaPGqRmlNdLSTo8hNG6a2SLALfQ7Q\nzfne+WEYJZeXCX+Dl8u5HmKhVnayNtvVb8xIvwb1gQKBgQDnM+ZnUbNg65pmVwJs\n/LKTaoC8AJYjDncI6IdEK90sHgDKiCoBtGf9L1husFU4oh4l+8fqZRU8AYg9EY1L\nhzU0b7LqJ483Dnlq05kBhyvAW21Xyd0L64Ejz+PpX2YR8+Derm9kR3osUDW3q54+\nGnBaOUFoH8Ey7vy1+LzvmtbgEQKBgQDX/BeHRzuqEBq2axv9mTgxpwMUo6Oc15bc\nDTVPeh5kGFnSKWYy8yDmpx7SpfKQHox+F0aSwotx/id88ZPct/ph7mK1FaKYL+0B\n3x8S6A88KgrGNJkfQpq1GdO8dcmUMVDpfy4lcHaKYLgmQzmyQNGIeI3cRuQf7HO2\nSCfQvp+g6QKBgAd3AYVseeIzOTT8sjNapVllIurvCenv2aAMAfINiIKYJWZkpxaP\nAJyIHs7UGqxNb7PCQV4sDVd2pTNzkaBSqUEcKiatSE56xSjR1Crcdmkriog/GEf+\nPpktaSprzhveR/BoqWgPdr5fiJx5ig5vRBllp9r4y4Adf0NQ2KjT0wRBAoGBAM02\n9DAPqYv7QzZB7NV3CGJ8+jtZA/LZELjCP/3k3Q1j6av6s+UDBybAcVPaYu82Z2zC\ntsZo3E+SGXjonAIiOF9mPhkqllOdbcbSddbj8N3MYHJUtYxzH0WlZX3yOHZ6qRNC\nSNk/0xHFthJ820wXtD7DtJ4wKT5/zq3KPzJifI9JAoGASvy8VBbRgO8fsOEKwiAz\n3Go41ScTLyjPD/aO6W6cDUHeoJRzs1fy6OQMerxonBjVBDBsUcuA6pV1ron3Ihrl\ngw0kEzRPhw6vpnQX8P93yp2AZ700GGcBZ687gRkrWZRy2hwIenTnbGouN+o2vkDx\nNFV1lrE8e10w2HzGdtAOHdM=\n-----END PRIVATE KEY-----",
	}
	fb, err := New(fireblocksConfig)
	if err != nil {
		t.Fatal(err)
	}

	dest1 := fb.NewDestinationTransferUsingAddress("0x77FEE9b17F2a8395C682d1b0548Ef35Bc10a234C", "")
	if dest1.Type != accounttype.ExternalWallet {
		t.Errorf("dest1 type %v is not %v", dest1.Type, accounttype.ExternalWallet)
	}
	if dest1.Id != "foobar" {
		t.Errorf("dest1 id %s is not %s", dest1.Id, "foobar")
	}

	dest2 := fb.NewDestinationTransferUsingAddress("0xe83B89654C58E40EB04011063Fb6a6623ed23f45", "")
	if dest2.Type != accounttype.InternalWallet {
		t.Errorf("dest2 type %v is not %v", dest2.Type, accounttype.InternalWallet)
	}
	if dest2.Id != "fibbage" {
		t.Errorf("dest2 id %s is not %s", dest1.Id, "fibbage")
	}

	dest3 := fb.NewDestinationTransferUsingAddress("0x54d3173ef7DA5F1411661981A64cE74d46Fe0247", "tag3")
	if dest3.Type != accounttype.OneTimeAddress {
		t.Errorf("dest3 type %v is not %v", dest3.Type, accounttype.OneTimeAddress)
	}
	if dest3.OneTimeAddress.Address != "0x54d3173ef7DA5F1411661981A64cE74d46Fe0247" {
		t.Errorf("dest2 address %s is not %s", dest3.OneTimeAddress.Address, "f0x54d3173ef7DA5F1411661981A64cE74d46Fe0247ibbage")
	}
	if dest3.OneTimeAddress.Tag != "tag3" {
		t.Errorf("dest3 tag %s is not %s", dest3.OneTimeAddress.Tag, "tag3")
	}
}
func TestEmptyWalletMaps(t *testing.T) {
	fireblocksConfig := configuration.WalletFireblocks{
		APIKey:          "",
		AssetId:         "ETH",
		BaseURL:         "http://google.com",
		ExternalWallets: "",
		InternalWallets: "",
		SourceAddress:   "0xADB898c59B2cf6eBD09d58A8cB4F176074556C65",
		SourceType:      "VAULT_ACCOUNT",
		SSLKey:          "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDDEELXabkjQD28\n0uV59YyoiRftYbKMBVtvsyJi1PcnDqbd2D1nKVApchq7/BBAgJHAw4iauVUvw7UV\neDC5absXwx7kMSOizlZx7QXKYXarE27v0tLdJNHZMJhJFDqVmSQqIbVczbUOF/bB\nOYPvI0GUI7JK6NXYw/Nkj0Gho4SSD68wh4dpqHFp0i2FRM6/YduLe/0X4myUqmgM\nKHcZxv5RFJ3MZ+y/jmVbL9+p+3xmTcRKAhWustVyfcihO1pcXK/oFyhc/P+A+My+\nQxHdP34LfkTLD7JlH9tnbQbJwcr2dcVEdgJdKfGxUIaFiE48H+6Ldt2X0Dy+Fui3\nAWwCK495AgMBAAECggEAZFAtf6ggDK/R0BRI7NJJVrbB9lLixj1PC3989IR+4bgY\nOkglM7M8RUQn2XrERNeZnR8xVhhvZZPowTTBIejszUiLz2Ax6lzgedAhxTUgGNRy\nnaS2ygFQZbgm6JYdlqddwMj7AtHLBYdvX2kbn1RrOiqtWvaoWsRSm+lMIDr+Rzoh\nrbUVQGnH8mLCJwnMQsK6bVkiSwx4WYelplcvbnI0mkwDm3FKfBgAxmdHDNaCUXE1\ndO0LCJNedBzhyY7BLN3vn1gw04fjtqid+fGtaPGqRmlNdLSTo8hNG6a2SLALfQ7Q\nzfne+WEYJZeXCX+Dl8u5HmKhVnayNtvVb8xIvwb1gQKBgQDnM+ZnUbNg65pmVwJs\n/LKTaoC8AJYjDncI6IdEK90sHgDKiCoBtGf9L1husFU4oh4l+8fqZRU8AYg9EY1L\nhzU0b7LqJ483Dnlq05kBhyvAW21Xyd0L64Ejz+PpX2YR8+Derm9kR3osUDW3q54+\nGnBaOUFoH8Ey7vy1+LzvmtbgEQKBgQDX/BeHRzuqEBq2axv9mTgxpwMUo6Oc15bc\nDTVPeh5kGFnSKWYy8yDmpx7SpfKQHox+F0aSwotx/id88ZPct/ph7mK1FaKYL+0B\n3x8S6A88KgrGNJkfQpq1GdO8dcmUMVDpfy4lcHaKYLgmQzmyQNGIeI3cRuQf7HO2\nSCfQvp+g6QKBgAd3AYVseeIzOTT8sjNapVllIurvCenv2aAMAfINiIKYJWZkpxaP\nAJyIHs7UGqxNb7PCQV4sDVd2pTNzkaBSqUEcKiatSE56xSjR1Crcdmkriog/GEf+\nPpktaSprzhveR/BoqWgPdr5fiJx5ig5vRBllp9r4y4Adf0NQ2KjT0wRBAoGBAM02\n9DAPqYv7QzZB7NV3CGJ8+jtZA/LZELjCP/3k3Q1j6av6s+UDBybAcVPaYu82Z2zC\ntsZo3E+SGXjonAIiOF9mPhkqllOdbcbSddbj8N3MYHJUtYxzH0WlZX3yOHZ6qRNC\nSNk/0xHFthJ820wXtD7DtJ4wKT5/zq3KPzJifI9JAoGASvy8VBbRgO8fsOEKwiAz\n3Go41ScTLyjPD/aO6W6cDUHeoJRzs1fy6OQMerxonBjVBDBsUcuA6pV1ron3Ihrl\ngw0kEzRPhw6vpnQX8P93yp2AZ700GGcBZ687gRkrWZRy2hwIenTnbGouN+o2vkDx\nNFV1lrE8e10w2HzGdtAOHdM=\n-----END PRIVATE KEY-----",
	}
	fb, err := New(fireblocksConfig)
	if err != nil {
		t.Fatal(err)
	}

	dest3 := fb.NewDestinationTransferUsingAddress("0x54d3173ef7DA5F1411661981A64cE74d46Fe0247", "tag3")
	if dest3.Type != accounttype.OneTimeAddress {
		t.Errorf("dest3 type %v is not %v", dest3.Type, accounttype.OneTimeAddress)
	}
	if dest3.OneTimeAddress.Address != "0x54d3173ef7DA5F1411661981A64cE74d46Fe0247" {
		t.Errorf("dest2 address %s is not %s", dest3.OneTimeAddress.Address, "f0x54d3173ef7DA5F1411661981A64cE74d46Fe0247ibbage")
	}
	if dest3.OneTimeAddress.Tag != "tag3" {
		t.Errorf("dest3 tag %s is not %s", dest3.OneTimeAddress.Tag, "tag3")
	}
}
