# aesecryptor

Encryptor/Decryptor using "AES/ECB/PKCS5Padding" algorithm by golang

# How to use

```golang
import "github.com/spotlight21c/aesencryptor"

plain := "This-is-plain-text"
key := "1234567890abcdef"
encValue, _ := aesencryptor.Encrypt(plain, key)
encodedString := base64.StdEncoding.EncodeToString(encValue)
```

This package is compatible with this JAVA CODE

```java
String ALGORITHM = "AES/ECB/PKCS5Padding";

String plainString = "This-is-plain-text";
String encryptionKey = "1234567890abcdef";

Key key = new SecretKeySpec(encryptionKey.getBytes(), ALGORITHM);

try {
    Cipher c = Cipher.getInstance(ALGORITHM);
    c.init(Cipher.ENCRYPT_MODE, key);
    byte[] encValue = c.doFinal(plainString.getBytes());
    String encodedString = new String(Base64.encodeToString(encValue, Base64.DEFAULT));
    System.out.print(encodedString);
} catch (Exception e) {
    
}
```

# Reference

https://www.bbsmax.com/A/kjdwvyer5N/