
var LastUserName =  "admin";
var MACRO_LAST_USER_LOGIN = "user";
var refresh_html_flag = 0;
var lang_id = "en";
var multi_language_string = "";
var hardwarePlatform = navigator.platform.toLowerCase();
var isMac = hardwarePlatform.indexOf("mac") != -1;
var vendor_name = "";
var footer_height;

function HideFrm(sFormName, iFormType)
{
    if (iFormType == null)
    {
        iFormType = 0;
    }

    this.oForm = sFormName;
    this.oForm.method = "POST";
    
    this.hasCallBack = false;
    
    if (iFormType != 1)
    {
        //this.oForm.innerHTML = "";
        for(i=0; i < this.oForm.elements.length; i++) 
        {
            this.oForm.elements[i].disabled = true;
        }
    }

    this.addParameter = function(sName, sValue)
    {
        var i=0;
        
        for(i=0; i < this.oForm.elements.length; i++) 
        {
            if((this.oForm.elements[i].name == sName) && this.oForm.elements[i].disabled) 
            {
                this.oForm.elements[i].value = sValue;
                this.oForm.elements[i].disabled = false;
            }
        }
        
        if(i >= this.oForm.elements.length-1) 
        {
            var oParam = document.createElement("INPUT");
            oParam.type = "hidden";
            oParam.name = sName;
            oParam.value = sValue;
    
            this.oForm.appendChild(oParam);
        }
    };

    this.setMethod = function(sMethod) 
    {
        if(sMethod.toUpperCase() == "GET")
            this.oForm.method = "GET";
        else
            this.oForm.method = "POST";
    };

    this.setAction = function(sUrl)
    {
        this.oForm.action = sUrl;
    }

    this.setTarget = function(sTarget) 
    {
        this.oForm.target = sTarget;
    };

    this.setCallBack = function(szCBFunc, aHirerPar) 
    {
        var vTmp = window.location.pathname;
        vTmp = vTmp.split("/");
        top.g_szHireUrl = vTmp[vTmp.length-1];
        if (szCBFunc != null)
        {
            top.g_szHireCBFun = szCBFunc;
        }
        if (aHirerPar != null)
        {
            top.g_aHirePar = arguments;
        }
        this.hasCallBack = true;
    };

    this.submit = function(sURL, sMethod)
    {
        if( sURL != null && sURL != "" ) this.setAction(sURL);
        if( sMethod != null && sMethod!= "" ) this.setMethod(sMethod);

        if (iFormType != 1)
        {
            if (iFormType == 2)
            {
                this.setTarget(top.g_oCmdFrame.name);
                if (!this.hasCallBack)
                {
                    this.setCallBack();
                }
            }
        }
        else
        {
            var sTmp = this.oForm.action;
            if ( sTmp.indexOf("?") > 0 )
            {
                sTmp += "&";
            }
            else
            {
                sTmp += "?";
            }
            this.setAction(sTmp + "uid=" + top.g_iUID);
        }
        
        
        if (iFormType == 0)
        {
        }
        this.oForm.submit();
    };
}

function LoginForm()
{
    var userLevel = "";
    var userType = "";      

    userLevel = document.form1.user_admin.checked;
    if (!userLevel)
    {
        userLevel = "?user=1";
        userType = "user";
    }
    else
    {
        userLevel = "";
        userType = "admin";	
    }        
   
   var encode_passwd = encodeURIComponent(document.form1.password.value);

   document.cookie = "Basic=" + userType + ':' + encode_passwd + ':' + "0" + ";path=/"; 
   if(navigator.userAgent.indexOf('Opera') >= 0)
   {
       if (vendor_name == "Syriatel")
       {
         location.replace( "en" + "/content_syriatel_opera.asp" + userLevel );
       }
       else
       {
         location.replace( "en" + "/content_opera.asp" + userLevel );
       }
   }
   else
   {
       if (vendor_name=="Syriatel")
       {
         location.replace( "en" + "/content_syriatel.asp" + userLevel );
       }
       else
       {
         location.replace( "en" + "/content.asp" + userLevel );
       }
   } 
}

function isEnterPressed( evt)
{
	try
	{
	    var evt = window.event ? window.event : evt;
	    
	    /* Check if enter is pressed down */
           if(isMac)
           {
               switch( evt.keyCode )
       	        {
       	            case 13:  /* key code of "enter" */
       	            case 3:
       	                LoginForm();
       	                return false;            
       	            default:	            
       	                return true;
       		 }
           }
           else
       	    {
       	        switch( evt.keyCode )
       	        {
       	            case 13:  /* key code of "enter" */
       	                LoginForm();
       	                return false;            
       	            default:	            
       	                return true;
       		 }
            }
	}
	catch(e)
	{
	    return true;
	}

}

function CloseLogin()
{
	setTimeout("CloseLogin()", 200);
	window.close();
}
function init()
{  
   document.cookie = "Basic=index;path=/";
   try
   {
       if ( window.opener != null )
       {
           CloseLogin();
       }
   }
   catch(e)
   {
   }
   try 
   {
       document.form1.user_type.onkeypress = isEnterPressed;
       document.form1.password.onkeypress = isEnterPressed;
       document.form1.lang.onkeypress = isEnterPressed;      
       
       if (LastUserName == MACRO_LAST_USER_LOGIN)
       {
           document.form1.user_admin.checked = false;
           document.form1.user_sms.checked = true;
       }
       else
       {
           document.form1.user_admin.checked = true;
           document.form1.user_sms.checked = false;
       }
       
   }
   catch(e)
   {
   }
   document.form1.password.focus();
   change_title();
   if (1 == refresh_html_flag )
   {
       refresh_all();
   }

   with(document.forms[1])
   {
       if( multi_language_string == "on" )
       {
           language_type.value = lang_id;
       }
   }
}

function CleanPassword()
{ 
    document.forms[1].password.value = "";
    var userLevel = "";
    userLevel = document.form1.user_admin.checked;
}

function change_title()
{
    document.title = TITLE;
}

function get_footer_height()
{
    if( screen.height <= 600 )
    {
        footer_height = 40;
    }
    else
    {
        footer_height = 79;
    }
}

function show_head()
{
    get_footer_height();
    var screen_hgt = document.body.scrollHeight - 79 - footer_height;
    if( vendor_name == "Etisalat" )
    {
        document.write('<div id="main"><table width = "100%" height ="78px" border = "0" cellpadding = "0" cellspacing = "0"><tr><td ><img src = "images/logo_custom.gif" width = "100%" height = "78px" style="top:0; bottom:0; margin:0; padding:0;"></td></tr></table>');
    }
    else if( vendor_name == "pl_ptc" )
    {
        document.write('<div id="header"><table width = "100%" height ="79px" border = "0" cellpadding = "0" cellspacing = "0" background = "images/title.gif" ><tr><td ><width = "280px" height = "79px" style="top:0;bottom:0; margin:0; padding:0;"></td></tr></table></div><div id="main" style="height:'+screen_hgt+'px;">');
    }
    else if( vendor_name == "ni_movistar")
    {
        document.write('<div id="header"><table width = "100%" height ="79px" border = "0" cellpadding = "0" cellspacing = "0" background = "images/title.gif" ><tr><td ><img src = "images/logo.gif" width = "165px" height = "79px" style="top:0;bottom:0; margin:0; padding:0;"></td><td style="text-align:right"><img src = "images/logo_movistar.gif" style="padding:0 0 0 0;"></td></tr></table></div><div id="main">');
    }
    else if ( vendor_name == "Syriatel" )
    {
        document.write('<div id="header"><table width = "100%" height ="79px" border = "0" cellpadding = "0" cellspacing = "0" background = "images/title_custom.gif" ><tr><td ><img src = "images/logo_custom.gif" width = "165px" height = "79px" style="top:0;bottom:0; margin:0; padding:0;"></td><td style="text-align:right"><img src = "images/logo.gif" style="padding:0 0 0 0;"></td></tr></table></div><div id="main">');
    }
    else
    {
        document.write('<div id="header"><table width = "100%" height ="79px" border = "0" cellpadding = "0" cellspacing = "0" background = "images/title.gif" ><tr><td ><img src = "images/logo.gif" width = "165px" height = "79px" style="top:0;bottom:0; margin:0; padding:0;"></td></tr></table></div><div id="main">');
    }
}

function show_user_type()
{
    if( vendor_name == "Etisalat" )
    {
        document.write('<td class="label" style="border-left: #000000 1px solid;">User Type<!--User Type--></td>');
    }
    else
    {
        document.write('<td class="label" >User Type<!--User Type--></td>');
    }
}

function show_radio()
{
    if( vendor_name == "Etisalat" )
    {
        document.write('<td style="border-right: #000000 1px solid;">');
    }
    else
    {
        document.write('<td>');
    }
}

function show_password()
{
    if( vendor_name == "Etisalat" )
    {
        document.write('<td class="label" width="40%" style="border-left: #000000 1px solid;">Password<!--Password--></td>');
    }
    else
    {
        document.write('<td class="label" width="40%">Password<!--Password--></td>');
    }
}

function show_input_password()
{
    if( vendor_name == "Etisalat" )
    {
        document.write('<td style="border-right: #000000 1px solid;"><input id="password" name="password" class="input"  size = "17" maxlength="15" type="password"></td>');
    }
    else
    {
        document.write('<td ><input id="password" name="password" class="input"  size = "17" maxlength="15" type="password"></td>');
    }
}
function show_footer()
{
    if( vendor_name == "ni_movistar")
    {
        document.write('<div id="footer"><table width = "100%" style="height:'+footer_height+'px;" background = "images/footer.gif"><tr><td ></td></tr></table></div>');
    }
}
function SearchSelectIndex(labContentSelectID)
{
    var index;
    for(i=1;i<ContentSelectInfo.length;i++)//Firef…