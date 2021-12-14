package orangesdk

// 3.2 Connection to the web service
// 4.1 Execution of operations

// m2sreq request
var m2sreq = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doM2S>
         <subsmsisdn>{{.Subsmsisdn}}</subsmsisdn>
         <PartnId>{{.PartnID}}</PartnId>
         <mermsisdn>{{.Mermsisdn}}</mermsisdn>
         <transid>{{.Transid}}</transid>
         <currency>{{.Currency}}</currency>
         <amount>{{.Amount}}</amount>
      </ser:doM2S>
   </soapenv:Body>
</soapenv:Envelope>
`

// m2sres REPONSE
var m2sres = `
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:doM2SResponse xmlns:ns2="http://services.ws1.com/">
         <return>
            <partnId>449</partnId>
            <resultCode>200</resultCode>
            <resultDesc>Votre Transaction a reussie</resultDesc>
            <systemId>616062846331</systemId>
            <transId>566666662</transId>
         </return>
      </ns2:doM2SResponse>
   </S:Body>
</S:Envelope>
`

// s2mreq REQUEST
// You asked for a payment for an amount of 100 CDFs in Lionel pastry. Confirm by typing your PIN code:
var s2mreq = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doS2M>
         <subsmsisdn>{{.Subsmsisdn}}</subsmsisdn>
         <PartnId>{{.PartnID}}</PartnId>
         <mermsisdn>{{.Mermsisdn}}</mermsisdn>
         <transid>{{.Transid}}</transid>
         <currency>{{.Currency}}</currency>
         <amount>{{.Amount}}</amount>
         <callbackurl>{{.Callbackurl}}</callbackurl>
         <message_s2m>{{.MessageS2m}}  </message_s2m>
      </ser:doS2M>
   </soapenv:Body>
</soapenv:Envelope>
`

// s2mres RESPONSE
var s2mres = `
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
`

// s2sreq request
var s2sreq = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doS2S>
         <subsmsisdn>{{.Subsmsisdn}}</subsmsisdn>
         <PartnId>{{.PartnID}}</PartnId>
         <subsmsisdn2>{{.Subsmsisdn2}}</subsmsisdn2>
         <transid>{{.Transid}}</transid>
         <currency>{{.Currency}}</currency>
         <amount>{{.Amount}}</amount>
         <callbackurl>{{.Callbackurl}}</callbackurl>
         <message_s2s>{{.MessageS2s}}</message_s2s>
    </ser:doS2S>
   </soapenv:Body>
</soapenv:Envelope>
`

// s2sres RESPONSE
var s2sres = `
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
`

// m2mreq REQUEST
var m2mreq = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
<soapenv:Header/>
   <soapenv:Body>
      <ser:doM2M>
         <mermsisdn1>{{.Mermsisdn1}}</mermsisdn1>
         <PartnId>{{.PartnID}}</PartnId>
         <mermsisdn2>{{.Mermsisdn2}}</mermsisdn2>
         <transid>{{.Transid}}</transid>
         <currency>{{.Currency}}</currency>
         <amount>{{.Amount}}</amount>
      </ser:doM2M>
   </soapenv:Body>
</soapenv:Envelope>
`

// m2mres RESPONSE
var m2mres = `
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
`

// tcheckbalreq CHECK BALANCE  REQUEST
var tcheckbalreq = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:TcheckBal>
         <subsmsisdn>{{.Subsmsisdn}}</subsmsisdn>
         <currency>{{.Currency}}</currency>
         <PartnId>{{.PartnID}}</PartnId>
         <transid>{{.Transid}}</transid>
      </ser:TcheckBal>
   </soapenv:Body>
</soapenv:Envelope>
`

// checkbalres CHECK BALANCE RESPONSE
var checkbalres = `
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
`

// checktxstatusreq CHECK TRANSACTION STATUS  REQUEST
var checktransreq = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws1.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doCheckTrans>
         <PartnId>{{.PartnID}}</PartnId>
         <mermsisdn>{{.Mermsisdn}}</mermsisdn>
         <transid>{{.Transid}}</transid>
      </ser:doCheckTrans>
   </soapenv:Body>
</soapenv:Envelope>
`

// checktxstatusres CHECK TRANSACTION STATUS RESPONSE
var checktxstatusres = `
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
`

// callbackreq REQUEST
var callbackreq = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://servicecb.ws.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:doCallback>
         <subsmsisdn>{{.Subsmsisdn}}</subsmsisdn>
         <PartnId>{{.PartnID}}</PartnId>
         <mermsisdn>{{.Mermsisdn}}</mermsisdn>
         <transid>{{.Transid}}</transid>
         <currency>{{.Currency}}</currency>
         <amount>{{.Amount}}</amount>
         <ResultCode>{{.ResultCode}}</ResultCode>
         <ResultDesc>{{.ResultDesc}}</ResultDesc>
      </ser:doCallback>
   </soapenv:Body>
</soapenv:Envelope>
`

// callbackres REPONSE
var callbackres = `
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
`

// sendsmsreq SMS SENDING REQUEST
var sendsmsreq = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://services.ws.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:sendSMS>
         <sender>{{.Sender}}</sender>
         <msisdn>{{.Msisdn}}</msisdn>
         <message>{{.Message}}</message>
         <flash>{{.Flash}}</flash>
         <PartnId>{{.PartnID}}</PartnId>
         <PartnName>{{.PartnName}}</PartnName>
      </ser:sendSMS>
   </soapenv:Body>
</soapenv:Envelope>
`

// sendsmsres REPONSE
var sendsmsres = `
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
`
