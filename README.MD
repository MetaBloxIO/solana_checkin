# Solana Checkin

This is a simple program that allows users to check in with a Solana account.

## Usage

1. Start the development server:

```bash
go run main.go
```

2. User prepare 3W data and call checkin API.

```
curl -X POST -H "Content-Type: application/json" \
    -d '{
        "userAddress": "3PcmXDanBD2wohL4zVoafAKwNm5whM5sSgL9QPnuw6oc",
        "did": "did:metablox:solana:3PcmXDanBD2wohL4zVoafAKwNm5whM5sSgL9QPnuw6oc",
        "location": "40.748817,-73.985428",
        "timestamp": "2006-01-02T15:04:05+07:00"
    }' \
    http://localhost:8080/v1/solana/3w/tx

```

Then you'll get a transaction like this:

```
{
    "code": 0,
    "message": "",
    "data": {
        "tx": "AnWTW7rBUQxrwaNgBeiucuxYpOi4KlN1Uyg3MBkw0MzYVcGrsN5RuBbouVUFBgSeyvXCCwNLntIC14/oP1YK2gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgACBKNzCV4JMip3azRgKyzqPTJxDqhfYdbpwSmrMiPHMDtmI4KYHigLvtwGpRVv7XQQv7mmJoa3CGZAhh0G76zk2hPC4ZSF8KZ12HPRcFIwYOa52RzM18cd8+F/ZDwLkNDotfKGW2Tj6IMCp38L2LW1pS/q3eOdXx/ArBn/Ueu7fxu8kbS9m37QbNariIOSMG0c3PP4Hc70tFzu6wbQF6OC5GMBAwMAAQJQ0f0E2frxzzJAAAAAAAAAADA0NTFjODQ2ZTNkMzUyM2MzNzZlMDRmMDNkZWU5NTZhODhlMTc2MDNlZGYyOGZiZWVkZmZmNDgyM2Q2YThmMGU="
    }
}
```

3. Parse transaction bytes and add your signature for sending tx to solana network.
   1. Decode the provided encodedData using SignedTx.decode.
   2. Sign the transaction's compiled message with the ed25519HDKeyPair.
   3. Add the new signature to the original signatures to create a new SignedTx object.
   4. Send the encoded new transaction using rpcClient.sendTransaction and obtain the transaction hash.
    
    Dart code example:
   
```dart
Future<String> signSubmitData(String encodedData, Ed25519HDKeyPair ed25519HDKeyPair) async {
  try {
    final signedTx = SignedTx.decode(encodedData);

    final signature = await ed25519HDKeyPair.sign(signedTx.compiledMessage.toByteArray());
    final signatures = [signedTx.signatures[0], signature];
    final newSignedTx = SignedTx(
      signatures: signatures,
      compiledMessage: signedTx.compiledMessage
    );

    print("signedTx: ${newSignedTx.toString()}");

    final transactionId = await rpcClient.sendTransaction(newSignedTx.encode());
    print("transactionId:$transactionId");
    return transactionId;
  } catch (e) {
    print("Error: ${e.toString()}");
  }
  return '';
}
```



