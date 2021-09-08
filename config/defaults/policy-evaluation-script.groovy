function getIdmClientDetails() {
  return { 
    "id": "{{IDM_CLIENT_ID}}",
    "secret": "{{IDM_CLIENT_SECRET}}",
    "endpoint": "http://am/am/oauth2/realms/root/realms/alpha/access_token",
    "scope": "fr:idm:*",
    "idmAdminUsername": "{{IG_IDM_USER}}",
    "idmAdminPassword": "{{IG_IDM_PASSWORD}}"
  }
}

STATUS_AUTHORISED = "Authorised"

logger.message("OB_Policy starting")

function parseResourceUri() {
  var elements = resourceURI.split("/");
  return {
    "api": elements[6],
    "account": (elements.length > 7) ? elements[7] : null,
    "data" : (elements.length > 8) ? elements[8] : null
  }
}

var p="=";
var tab="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

function base64encode(ba) {
		var s=[], l=ba.length;
		var rm=l%3;
		var x=l-rm;
		for (var i=0; i<x;){
			var t=ba[i++]<<16|ba[i++]<<8|ba[i++];
			s.push(tab.charAt((t>>>18)&0x3f));
			s.push(tab.charAt((t>>>12)&0x3f));
			s.push(tab.charAt((t>>>6)&0x3f));
			s.push(tab.charAt(t&0x3f));
		}
		//	deal with trailers, based on patch from Peter Wood.
		switch(rm){
			case 2:{
				var t=ba[i++]<<16|ba[i++]<<8;
				s.push(tab.charAt((t>>>18)&0x3f));
				s.push(tab.charAt((t>>>12)&0x3f));
				s.push(tab.charAt((t>>>6)&0x3f));
				s.push(p);
				break;
			}
			case 1:{
				var t=ba[i++]<<16;
				s.push(tab.charAt((t>>>18)&0x3f));
				s.push(tab.charAt((t>>>12)&0x3f));
				s.push(p);
				s.push(p);
				break;
			}
		}
		return s.join("");	//	string
}

function base64decode(str) {
		var s=str.split(""), out=[];
		var l=s.length;
		while(s[--l]==p){ }	//	strip off trailing padding
		for (var i=0; i<l;){
			var t=tab.indexOf(s[i++])<<18;
			if(i<=l){ t|=tab.indexOf(s[i++])<<12 };
			if(i<=l){ t|=tab.indexOf(s[i++])<<6 };
			if(i<=l){ t|=tab.indexOf(s[i++]) };
			out.push((t>>>16)&0xff);
			out.push((t>>>8)&0xff);
			out.push(t&0xff);
		}
		//	strip off any null bytes
		while(out[out.length-1]==0){ out.pop(); }
		return out;	//	byte[]
}

function stringFromArray(data) {
    var count = data.length;
    var str = "";
    
    for(var index = 0; index < count; index += 1)
      str += String.fromCharCode(data[index]);
    
    return str;
}

function logResponse(response) {
    logger.warning("OB_Policy User REST Call. Status: " + response.getStatus() + ", Body: " + response.getEntity().getString());
}

function getIdmAccessToken() {

    var clientInfo = getIdmClientDetails();
    var request = new org.forgerock.http.protocol.Request();
    request.setUri(clientInfo.endpoint);
  	request.setMethod("POST");
    request.getHeaders().add("Content-Type","application/x-www-form-urlencoded");
    var formvars = "grant_type=password" +
        "&client_id=" + clientInfo.id +
        "&client_secret=" + clientInfo.secret +
        "&scope=" + clientInfo.scope +
        "&username=" + clientInfo.idmAdminUsername +
        "&password=" + clientInfo.idmAdminPassword;
    request.setEntity(formvars);

    var response = httpClient.send(request).get();

  
    logResponse(response);

    var oauth2response = JSON.parse(response.getEntity().getString());

    var accessToken = oauth2response.access_token
    logger.warning("OB_Policy Got acess token " + accessToken);
    return accessToken
}

function getIntent(intentId,api) {
  var obj = null
  
  switch (api) {
    case "accounts": 
      obj = "accountAccessIntent"
      break;
    case "domestic-payments":
      obj = "domesticPaymentIntent"
      break;
  }
   
  var accessToken = getIdmAccessToken();
  var request = new org.forgerock.http.protocol.Request();
  var uri = "http://idm/openidm/managed/" + obj + "/" + intentId
  logger.message("OB_Policy IDM fetch " + uri)
  
  request.setMethod('GET');
  request.setUri(uri)
  request.getHeaders().add("Authorization","Bearer " + accessToken);

  
  var response = httpClient.send(request).get();
  logResponse(response);

  var intent = JSON.parse(response.getEntity().getString());
  return intent
}

function dataAuthorised(permissions,dataRequest) {
  switch (dataRequest) {
    case "balances":
      authorised = (permissions.indexOf("ReadBalances") > -1);
      break;
    case "transactions":
      authorised = (permissions.indexOf("ReadTransactionsDetail") > -1);
      break
    default:
      authorised = false
      
  }
  
  return authorised
}

function initiationMatch(initiationRequest,initiation) {
  
  // TODO: do comparison at object level, like JSONAssert()
  
  var initiationRequestObj = JSON.parse(stringFromArray(base64decode(initiationRequest))) 
  
  var match = 
      (initiationRequestObj.InstructionIdentification == initiation.InstructionIdentification) &&
      (initiationRequestObj.EndToEndIdentification == initiation.EndToEndIdentification) &&
      (initiationRequestObj.InstructedAmount.Amount == initiation.InstructedAmount.Amount) &&
      (initiationRequestObj.InstructedAmount.Currency == initiation.InstructedAmount.Currency) &&
      (initiationRequestObj.CreditorAccount.SchemeName == initiation.CreditorAccount.SchemeName) &&
      (initiationRequestObj.CreditorAccount.Name == initiation.CreditorAccount.Name) &&
      (initiationRequestObj.CreditorAccount.SecondaryIdentification == initiation.CreditorAccount.SecondaryIdentification) &&
      (initiationRequestObj.RemittanceInformation.Reference == initiation.RemittanceInformation.Reference) &&
      (initiationRequestObj.RemittanceInformation.Unstructured == initiation.RemittanceInformation.Unstructured)
      
      
  if (!match) {
    logger.warning("Mismatch between request [" + JSON.stringify(initiationRequestObj) + "] and consent [" + JSON.stringify(initiation) + "]");
  }
                   
  return match
}

var intentId = environment.get("intent_id").iterator().next();

var apiRequest = parseResourceUri()

logger.warning("OB_Policy req " + apiRequest.api + ":" + apiRequest.account + ":" + apiRequest.data);
               
var intent = getIntent(intentId,apiRequest.api);

if (apiRequest.api == "accounts") {

  var status = intent.Data.Status
  var permissions = intent.Data.Permissions
  var accounts = intent.accounts

  if (status != STATUS_AUTHORISED) {
    logger.warning("Rejecting request - status [" + status + "]")
    authorized = false
    
  }
  else if (apiRequest.account == null) {
    logger.message("OB_POLICY accounts " + accounts);
    responseAttributes.put("grantedAccounts",accounts);
    responseAttributes.put("grantedPermissions",permissions);
    authorized = true
  }
  else if (apiRequest.data == null) {
    logger.message("OB_POLICY account info for " + apiRequest.account);
    // RS server expects granted accounts and permissions even though we're checking as well
    responseAttributes.put("grantedAccounts",accounts);
    responseAttributes.put("grantedPermissions",permissions);
    authorized = (accounts.indexOf(apiRequest.account) > -1) &&
                 (permissions.indexOf("ReadAccountsDetail") > -1)
  }
  else {
    logger.message("OB_POLICY account request for " + apiRequest.account + ":" + apiRequest.data);
    // RS server expects granted accounts and permissions even though we're checking as well
    responseAttributes.put("grantedAccounts",accounts);
    responseAttributes.put("grantedPermissions",permissions);
    authorized = (accounts.indexOf(apiRequest.account) > -1) && 
                 dataAuthorised(permissions,apiRequest.data)
  }
  
}
else if (apiRequest.api == "domestic-payments") {
  
  var status = intent.Data.Status
  var permissions = intent.Data.Permissions
  var account = intent.account

  if (status != STATUS_AUTHORISED) {
    logger.warning("Rejecting request - status [" + status + "]")
    authorized = false
  }
  else {
    authorized = initiationMatch(environment.get("initiation").iterator().next(),intent.Data.Initiation)
  }
  
}
else {
  authorized = false
}