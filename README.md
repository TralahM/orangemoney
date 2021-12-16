## 4.1 Execution of operations

Each http (POST) request is sent in XML format.

## M2S REQUEST
The doM2S operation makes it possible to make a merchant payment between two merchants and a customer.

```xml
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doM2S>
         <subsmsisdn>0858606953</subsmsisdn>
         <PartnId>449</PartnId>
         <mermsisdn>0890033477</mermsisdn>
         <transid>566666662</transid>
         <currency>CDF</currency>
         <amount>100</amount>
      </ser:doM2S>
   </soapenv:Body>
</soapenv:Envelope>
```


## M2S REPONSE

```xml
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:doM2SResponse xmlns:ns2="http://services.ws1.com/">
         <return>
            <partnId/>
            <resultCode>200</resultCode>
            <resultDesc>Votre Transaction a reussie</resultDesc>
            <systemId>616062846331</systemId>
            <transId>566666662</transId>
         </return>
      </ns2:doM2SResponse>
   </S:Body>
</S:Envelope>
```



## S2M REQUEST
The doS2M operation allows payment to be made between a customer and an orange merchant number.

```xml
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doS2M>
         <subsmsisdn>0808599770</subsmsisdn>
         <PartnId>449</PartnId>
         <mermsisdn>0890033477</mermsisdn>
         <transid>0001646238236</transid>
         <currency>CDF</currency>
         <amount>100</amount>
         <callbackurl>http://callbackurlws/wscallbacktestService</callbackurl>
         <message_s2m> Vous avez sollicité un paiement pour un montant de 100 CDF  chez  Pâtisserie Lionel. Confirmez en tapant votre code PIN :  </message_s2m>
      </ser:doS2M>
   </soapenv:Body>
</soapenv:Envelope>
```


## S2M RESPONSE

```xml
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:doS2MResponse xmlns:ns2="http://services.ws1.com/">
         <return>
            <partnId>449</partnId>
            <resultCode>0</resultCode>
            <resultDesc>Paiement En Cours de traitement !</resultDesc>
            <systemId>730733438533</systemId>
            <transId>0001646238236</transId>
         </return>
      </ns2:doS2MResponse>
   </S:Body>
</S:Envelope>
```


## S2S REQUEST
The doS2S operation makes it possible to make a payment between two customers.

```xml
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doS2S>
         <subsmsisdn>0808599770</subsmsisdn>
         <PartnId>449</PartnId>
         <subsmsisdn2>0890033473</subsmsisdn2>
         <transid>12211471562111</transid>
         <currency>USD</currency>
         <amount>1</amount>
         <callbackurl>http://callbackurlws/wscallbacktestService</callbackurl>
         <message_s2s> Vous avez sollicité un paiement pour un montant de 100 CDF  chez  Pâtisserie Lionel. Confirmez en tapant votre code PIN :  </message_s2s>
    </ser:doS2S>
   </soapenv:Body>
</soapenv:Envelope>
```



## S2S RESPONSE

```xml
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:doS2SResponse xmlns:ns2="http://services.ws1.com/">
         <return>
            <partnId>449</partnId>
            <resultCode>0</resultCode>
            <resultDesc>Paiement en Cours de Traitement !</resultDesc>
            <systemId>071182538561</systemId>
            <transId>12211562111</transId>
         </return>
      </ns2:doS2SResponse>
   </S:Body>
</S:Envelope>
```


## M2M REQUEST
The doM2M operation allows you to make a merchant payment between two different merchants.

```xml
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
<soapenv:Header/>
   <soapenv:Body>
      <ser:doM2M>
         <mermsisdn1>0890033477</mermsisdn1>
         <PartnId>449</PartnId>
         <mermsisdn2>0841364721</mermsisdn2>
         <transid>1334455666</transid>
         <currency>CDF</currency>
         <amount>100</amount>
      </ser:doM2M>
   </soapenv:Body>
</soapenv:Envelope>
```

## M2M RESPONSE


```xml
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:doM2MResponse xmlns:ns2="http://services.ws1.com/">
         <return>
            <partnId/>
            <resultCode>200</resultCode>
            <resultDesc>Votre Transaction a reussie</resultDesc>
             <systemId>414895740867</systemId>
             <transId>1334455666</transId>
         </return>
      </ns2:doM2MResponse>
   </S:Body>
</S:Envelope>
```





## CHECK BALANCE  REQUEST
The TcheckBal operation allows the partner to check the balance of his merchant account.

```xml
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:TcheckBal>
         <subsmsisdn>0890033477</subsmsisdn>
         <currency>usd</currency>
         <PartnId>449</PartnId>
         <transid>4555566</transid>
      </ser:TcheckBal>
   </soapenv:Body>
</soapenv:Envelope>
```

## CHECK BALANCE  RESPONSE


```xml
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:TcheckBalResponse xmlns:ns2="http://services.ws1.com/">
         <return>
            <balance>13.86</balance>
            <partnId>449</partnId>
            <resultCode>200</resultCode>
            <resultDesc>Success</resultDesc>
            <transId>4555566</transId>
         </return>
      </ns2:TcheckBalResponse>
   </S:Body>
</S:Envelope>
```


## CHECK TRANSACTION STATUS  REQUEST
The doCheckTrans operation allows you to check the status of a completed transaction.

```xml
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doCheckTrans>
         <PartnId>449</PartnId>
         <mermsisdn>0890033477</mermsisdn>
         <transid>1334455666</transid>
      </ser:doCheckTrans>
   </soapenv:Body>
</soapenv:Envelope>
```

## CHECK TRANSACTION STATUS RESPONSE

```xml
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:doCheckTransResponse xmlns:ns2="http://services.ws1.com/">
         <return>
            <resultCode>00084</resultCode>
            <resultDesc>Aucun frais de service defini.</resultDesc>
            <transId>1334455666</transId>
         </return>
      </ns2:doCheckTransResponse>
   </S:Body>
</S:Envelope>
```




## 4.2 Webservice Callback url

### REQUEST

```xml
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://servicecb.ws.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doCallback>
         <subsmsisdn>0890033473</subsmsisdn>
         <PartnId>449</PartnId>
         <mermsisdn>0852478243</mermsisdn>
         <transid>100000</transid>
         <currency>CDF</currency>
         <amount>100</amount>
         <ResultCode>0</ResultCode>
         <ResultDesc>success</ResultDesc>
      </ser:doCallback>
   </soapenv:Body>
</soapenv:Envelope>
```

### REPONSE

```xml
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:doCallbackResponse xmlns:ns2="http://servicecb.ws.com/">
        <return>
            <partnId>449</partnId>
            <resultCode>0</resultCode>
            <resultDesc>success</resultDesc>
     </return>
      </ns2:doCallbackResponse>
   </S:Body>
</S:Envelope>
```




## 4.3 Webservice SENDsms

 This web services is developed according to the SOAP version 1.2 (Simple Object Access Protocol) protocol and is described by the following wsdl file:

http://SERVER_IP:PORT/wssendsms/SendsmsService?wsdl

### SMS SENDING REQUEST

```xml
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:sendSMS>
         <sender>Orange</sender>
         <msisdn>243890033473</msisdn>
         <message>Bonjour Test, vous allez bien?</message>
         <flash>optionel</flash>
         <PartnId>501</PartnId>
         <PartnName>Test</PartnName>
      </ser:sendSMS>
   </soapenv:Body>
</soapenv:Envelope>
```

### REPONSE

```xml
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
 <S:Body>
      <ns2:sendSMSResponse xmlns:ns2="http://services.ws.com/">
  <return>
            <message>message envoyé</message>
            <resultdesc>success</resultdesc>
            <resutcode>200</resutcode>
         </return>
      </ns2:sendSMSResponse>
  </S:Body>
</S:Envelope>
```




