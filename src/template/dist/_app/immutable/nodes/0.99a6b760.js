import{s as we,c as ye,u as Oe,g as Ce,d as Ee,a as Me,e as Ie}from"../chunks/scheduler.a6620314.js";import{S as Ae,i as Pe,d as Re,t as Be}from"../chunks/index.1ddc4d5b.js";import{w as ke}from"../chunks/index.0f8a9725.js";var Q=typeof globalThis<"u"?globalThis:typeof window<"u"?window:typeof global<"u"?global:typeof self<"u"?self:{};function $e(n){return n&&n.__esModule&&Object.prototype.hasOwnProperty.call(n,"default")?n.default:n}var U={exports:{}};/*!
 * Platform.js v1.3.6
 * Copyright 2014-2020 Benjamin Tan
 * Copyright 2011-2013 John-David Dalton
 * Available under MIT license
 */U.exports;(function(n,a){(function(){var b={function:!0,object:!0},f=b[typeof window]&&window||this,l=a,S=n&&!n.nodeType&&n,O=l&&S&&typeof Q=="object"&&Q;O&&(O.global===O||O.window===O||O.self===O)&&(f=O);var G=Math.pow(2,53)-1,E=/\bOpera/,y=Object.prototype,C=y.hasOwnProperty,M=y.toString;function A(t){return t=String(t),t.charAt(0).toUpperCase()+t.slice(1)}function $(t,c,h){var v={"10.0":"10","6.4":"10 Technical Preview","6.3":"8.1","6.2":"8","6.1":"Server 2008 R2 / 7","6.0":"Server 2008 / Vista","5.2":"Server 2003 / XP 64-bit","5.1":"XP","5.01":"2000 SP1","5.0":"2000","4.0":"NT","4.90":"ME"};return c&&h&&/^Win/i.test(t)&&!/^Windows Phone /i.test(t)&&(v=v[/[\d.]+$/.exec(t)])&&(t="Windows "+v),t=String(t),c&&h&&(t=t.replace(RegExp(c,"i"),h)),t=R(t.replace(/ ce$/i," CE").replace(/\bhpw/i,"web").replace(/\bMacintosh\b/,"Mac OS").replace(/_PowerPC\b/i," OS").replace(/\b(OS X) [^ \d]+/i,"$1").replace(/\bMac (OS X)\b/,"$1").replace(/\/(\d)/," $1").replace(/_/g,".").replace(/(?: BePC|[ .]*fc[ \d.]+)$/i,"").replace(/\bx86\.64\b/gi,"x86_64").replace(/\b(Windows Phone) OS\b/,"$1").replace(/\b(Chrome OS \w+) [\d.]+\b/,"$1").split(" on ")[0]),t}function q(t,c){var h=-1,v=t?t.length:0;if(typeof v=="number"&&v>-1&&v<=G)for(;++h<v;)c(t[h],h,t);else W(t,c)}function R(t){return t=z(t),/^(?:webOS|i(?:OS|P))/.test(t)?t:A(t)}function W(t,c){for(var h in t)C.call(t,h)&&c(t[h],h,t)}function _(t){return t==null?A(t):M.call(t).slice(8,-1)}function se(t,c){var h=t!=null?typeof t[c]:"number";return!/^(?:boolean|number|string|undefined)$/.test(h)&&(h=="object"?!!t[c]:!0)}function P(t){return String(t).replace(/([ -])(?!$)/g,"$1?")}function T(t,c){var h=null;return q(t,function(v,X){h=c(h,v,X,t)}),h}function z(t){return String(t).replace(/^ +| +$/g,"")}function K(t){var c=f,h=t&&typeof t=="object"&&_(t)!="String";h&&(c=t,t=null);var v=c.navigator||{},X=v.userAgent||"";t||(t=X);var le=h?!!v.likeChrome:/\bChrome\b/.test(t)&&!/internal|\n/i.test(M.toString()),J="Object",ce=h?J:"ScriptBridgingProxyObject",de=h?J:"Environment",fe=h&&c.java?"JavaPackage":_(c.java),pe=h?J:"RuntimeObject",L=/\bJava/.test(fe)&&c.java,ue=L&&_(c.environment)==de,be=L?"a":"α",he=L?"b":"β",Y=c.document||{},B=c.operamini||c.opera,V=E.test(V=h&&B?B["[[Class]]"]:_(B))?V:B=null,e,j=t,d=[],D=null,k=t==X,o=k&&B&&typeof B.version=="function"&&B.version(),Z,p=me([{label:"EdgeHTML",pattern:"Edge"},"Trident",{label:"WebKit",pattern:"AppleWebKit"},"iCab","Presto","NetFront","Tasman","KHTML","Gecko"]),r=ge(["Adobe AIR","Arora","Avant Browser","Breach","Camino","Electron","Epiphany","Fennec","Flock","Galeon","GreenBrowser","iCab","Iceweasel","K-Meleon","Konqueror","Lunascape","Maxthon",{label:"Microsoft Edge",pattern:"(?:Edge|Edg|EdgA|EdgiOS)"},"Midori","Nook Browser","PaleMoon","PhantomJS","Raven","Rekonq","RockMelt",{label:"Samsung Internet",pattern:"SamsungBrowser"},"SeaMonkey",{label:"Silk",pattern:"(?:Cloud9|Silk-Accelerated)"},"Sleipnir","SlimBrowser",{label:"SRWare Iron",pattern:"Iron"},"Sunrise","Swiftfox","Vivaldi","Waterfox","WebPositive",{label:"Yandex Browser",pattern:"YaBrowser"},{label:"UC Browser",pattern:"UCBrowser"},"Opera Mini",{label:"Opera Mini",pattern:"OPiOS"},"Opera",{label:"Opera",pattern:"OPR"},"Chromium","Chrome",{label:"Chrome",pattern:"(?:HeadlessChrome)"},{label:"Chrome Mobile",pattern:"(?:CriOS|CrMo)"},{label:"Firefox",pattern:"(?:Firefox|Minefield)"},{label:"Firefox for iOS",pattern:"FxiOS"},{label:"IE",pattern:"IEMobile"},{label:"IE",pattern:"MSIE"},"Safari"]),s=te([{label:"BlackBerry",pattern:"BB10"},"BlackBerry",{label:"Galaxy S",pattern:"GT-I9000"},{label:"Galaxy S2",pattern:"GT-I9100"},{label:"Galaxy S3",pattern:"GT-I9300"},{label:"Galaxy S4",pattern:"GT-I9500"},{label:"Galaxy S5",pattern:"SM-G900"},{label:"Galaxy S6",pattern:"SM-G920"},{label:"Galaxy S6 Edge",pattern:"SM-G925"},{label:"Galaxy S7",pattern:"SM-G930"},{label:"Galaxy S7 Edge",pattern:"SM-G935"},"Google TV","Lumia","iPad","iPod","iPhone","Kindle",{label:"Kindle Fire",pattern:"(?:Cloud9|Silk-Accelerated)"},"Nexus","Nook","PlayBook","PlayStation Vita","PlayStation","TouchPad","Transformer",{label:"Wii U",pattern:"WiiU"},"Wii","Xbox One",{label:"Xbox 360",pattern:"Xbox"},"Xoom"]),g=Se({Apple:{iPad:1,iPhone:1,iPod:1},Alcatel:{},Archos:{},Amazon:{Kindle:1,"Kindle Fire":1},Asus:{Transformer:1},"Barnes & Noble":{Nook:1},BlackBerry:{PlayBook:1},Google:{"Google TV":1,Nexus:1},HP:{TouchPad:1},HTC:{},Huawei:{},Lenovo:{},LG:{},Microsoft:{Xbox:1,"Xbox One":1},Motorola:{Xoom:1},Nintendo:{"Wii U":1,Wii:1},Nokia:{Lumia:1},Oppo:{},Samsung:{"Galaxy S":1,"Galaxy S2":1,"Galaxy S3":1,"Galaxy S4":1},Sony:{PlayStation:1,"PlayStation Vita":1},Xiaomi:{Mi:1,Redmi:1}}),i=ve(["Windows Phone","KaiOS","Android","CentOS",{label:"Chrome OS",pattern:"CrOS"},"Debian",{label:"DragonFly BSD",pattern:"DragonFly"},"Fedora","FreeBSD","Gentoo","Haiku","Kubuntu","Linux Mint","OpenBSD","Red Hat","SuSE","Ubuntu","Xubuntu","Cygwin","Symbian OS","hpwOS","webOS ","webOS","Tablet OS","Tizen","Linux","Mac OS X","Macintosh","Mac","Windows 98;","Windows "]);function me(x){return T(x,function(m,u){return m||RegExp("\\b"+(u.pattern||P(u))+"\\b","i").exec(t)&&(u.label||u)})}function Se(x){return T(x,function(m,u,I){return m||(u[s]||u[/^[a-z]+(?: +[a-z]+\b)*/i.exec(s)]||RegExp("\\b"+P(I)+"(?:\\b|\\w*\\d)","i").exec(t))&&I})}function ge(x){return T(x,function(m,u){return m||RegExp("\\b"+(u.pattern||P(u))+"\\b","i").exec(t)&&(u.label||u)})}function ve(x){return T(x,function(m,u){var I=u.pattern||P(u);return!m&&(m=RegExp("\\b"+I+"(?:/[\\d.]+|[ \\w.]*)","i").exec(t))&&(m=$(m,I,u.label||u)),m})}function te(x){return T(x,function(m,u){var I=u.pattern||P(u);return!m&&(m=RegExp("\\b"+I+" *\\d+[.\\w_]*","i").exec(t)||RegExp("\\b"+I+" *\\w+-[\\w]*","i").exec(t)||RegExp("\\b"+I+"(?:; *(?:[a-z]+[_-])?[a-z]+\\d+|[^ ();-]*)","i").exec(t))&&((m=String(u.label&&!RegExp(I,"i").test(u.label)?u.label:m).split("/"))[1]&&!/[\d.]+/.test(m[0])&&(m[0]+=" "+m[1]),u=u.label||u,m=R(m[0].replace(RegExp(I,"i"),u).replace(RegExp("; *(?:"+u+"[_-])?","i")," ").replace(RegExp("("+u+")[-_.]?(\\w)","i"),"$1 $2"))),m})}function re(x){return T(x,function(m,u){return m||(RegExp(u+"(?:-[\\d.]+/|(?: for [\\w-]+)?[ /-])([\\d.]+[^ ();/_-]*)","i").exec(t)||0)[1]||null})}function xe(){return this.description||""}if(p&&(p=[p]),/\bAndroid\b/.test(i)&&!s&&(e=/\bAndroid[^;]*;(.*?)(?:Build|\) AppleWebKit)\b/i.exec(t))&&(s=z(e[1]).replace(/^[a-z]{2}-[a-z]{2};\s*/i,"")||null),g&&!s?s=te([g]):g&&s&&(s=s.replace(RegExp("^("+P(g)+")[-_.\\s]","i"),g+" ").replace(RegExp("^("+P(g)+")[-_.]?(\\w)","i"),g+" $2")),(e=/\bGoogle TV\b/.exec(s))&&(s=e[0]),/\bSimulator\b/i.test(t)&&(s=(s?s+" ":"")+"Simulator"),r=="Opera Mini"&&/\bOPiOS\b/.test(t)&&d.push("running in Turbo/Uncompressed mode"),r=="IE"&&/\blike iPhone OS\b/.test(t)?(e=K(t.replace(/like iPhone OS/,"")),g=e.manufacturer,s=e.product):/^iP/.test(s)?(r||(r="Safari"),i="iOS"+((e=/ OS ([\d_]+)/i.exec(t))?" "+e[1].replace(/_/g,"."):"")):r=="Konqueror"&&/^Linux\b/i.test(i)?i="Kubuntu":g&&g!="Google"&&(/Chrome/.test(r)&&!/\bMobile Safari\b/i.test(t)||/\bVita\b/.test(s))||/\bAndroid\b/.test(i)&&/^Chrome/.test(r)&&/\bVersion\//i.test(t)?(r="Android Browser",i=/\bAndroid\b/.test(i)?i:"Android"):r=="Silk"?(/\bMobi/i.test(t)||(i="Android",d.unshift("desktop mode")),/Accelerated *= *true/i.test(t)&&d.unshift("accelerated")):r=="UC Browser"&&/\bUCWEB\b/.test(t)?d.push("speed mode"):r=="PaleMoon"&&(e=/\bFirefox\/([\d.]+)\b/.exec(t))?d.push("identifying as Firefox "+e[1]):r=="Firefox"&&(e=/\b(Mobile|Tablet|TV)\b/i.exec(t))?(i||(i="Firefox OS"),s||(s=e[1])):!r||(e=!/\bMinefield\b/i.test(t)&&/\b(?:Firefox|Safari)\b/.exec(r))?(r&&!s&&/[\/,]|^[^(]+?\)/.test(t.slice(t.indexOf(e+"/")+8))&&(r=null),(e=s||g||i)&&(s||g||/\b(?:Android|Symbian OS|Tablet OS|webOS)\b/.test(i))&&(r=/[a-z]+(?: Hat)?/i.exec(/\bAndroid\b/.test(i)?i:e)+" Browser")):r=="Electron"&&(e=(/\bChrome\/([\d.]+)\b/.exec(t)||0)[1])&&d.push("Chromium "+e),o||(o=re(["(?:Cloud9|CriOS|CrMo|Edge|Edg|EdgA|EdgiOS|FxiOS|HeadlessChrome|IEMobile|Iron|Opera ?Mini|OPiOS|OPR|Raven|SamsungBrowser|Silk(?!/[\\d.]+$)|UCBrowser|YaBrowser)","Version",P(r),"(?:Firefox|Minefield|NetFront)"])),(e=p=="iCab"&&parseFloat(o)>3&&"WebKit"||/\bOpera\b/.test(r)&&(/\bOPR\b/.test(t)?"Blink":"Presto")||/\b(?:Midori|Nook|Safari)\b/i.test(t)&&!/^(?:Trident|EdgeHTML)$/.test(p)&&"WebKit"||!p&&/\bMSIE\b/i.test(t)&&(i=="Mac OS"?"Tasman":"Trident")||p=="WebKit"&&/\bPlayStation\b(?! Vita\b)/i.test(r)&&"NetFront")&&(p=[e]),r=="IE"&&(e=(/; *(?:XBLWP|ZuneWP)(\d+)/i.exec(t)||0)[1])?(r+=" Mobile",i="Windows Phone "+(/\+$/.test(e)?e:e+".x"),d.unshift("desktop mode")):/\bWPDesktop\b/i.test(t)?(r="IE Mobile",i="Windows Phone 8.x",d.unshift("desktop mode"),o||(o=(/\brv:([\d.]+)/.exec(t)||0)[1])):r!="IE"&&p=="Trident"&&(e=/\brv:([\d.]+)/.exec(t))&&(r&&d.push("identifying as "+r+(o?" "+o:"")),r="IE",o=e[1]),k){if(se(c,"global"))if(L&&(e=L.lang.System,j=e.getProperty("os.arch"),i=i||e.getProperty("os.name")+" "+e.getProperty("os.version")),ue){try{o=c.require("ringo/engine").version.join("."),r="RingoJS"}catch{(e=c.system)&&e.global.system==c.system&&(r="Narwhal",i||(i=e[0].os||null))}r||(r="Rhino")}else typeof c.process=="object"&&!c.process.browser&&(e=c.process)&&(typeof e.versions=="object"&&(typeof e.versions.electron=="string"?(d.push("Node "+e.versions.node),r="Electron",o=e.versions.electron):typeof e.versions.nw=="string"&&(d.push("Chromium "+o,"Node "+e.versions.node),r="NW.js",o=e.versions.nw)),r||(r="Node.js",j=e.arch,i=e.platform,o=/[\d.]+/.exec(e.version),o=o?o[0]:null));else _(e=c.runtime)==ce?(r="Adobe AIR",i=e.flash.system.Capabilities.os):_(e=c.phantom)==pe?(r="PhantomJS",o=(e=e.version||null)&&e.major+"."+e.minor+"."+e.patch):typeof Y.documentMode=="number"&&(e=/\bTrident\/(\d+)/i.exec(t))?(o=[o,Y.documentMode],(e=+e[1]+4)!=o[1]&&(d.push("IE "+o[1]+" mode"),p&&(p[1]=""),o[1]=e),o=r=="IE"?String(o[1].toFixed(1)):o[0]):typeof Y.documentMode=="number"&&/^(?:Chrome|Firefox)\b/.test(r)&&(d.push("masking as "+r+" "+o),r="IE",o="11.0",p=["Trident"],i="Windows");i=i&&R(i)}if(o&&(e=/(?:[ab]|dp|pre|[ab]\d+pre)(?:\d+\+?)?$/i.exec(o)||/(?:alpha|beta)(?: ?\d)?/i.exec(t+";"+(k&&v.appMinorVersion))||/\bMinefield\b/i.test(t)&&"a")&&(D=/b/i.test(e)?"beta":"alpha",o=o.replace(RegExp(e+"\\+?$"),"")+(D=="beta"?he:be)+(/\d+\+?/.exec(e)||"")),r=="Fennec"||r=="Firefox"&&/\b(?:Android|Firefox OS|KaiOS)\b/.test(i))r="Firefox Mobile";else if(r=="Maxthon"&&o)o=o.replace(/\.[\d.]+/,".x");else if(/\bXbox\b/i.test(s))s=="Xbox 360"&&(i=null),s=="Xbox 360"&&/\bIEMobile\b/.test(t)&&d.unshift("mobile mode");else if((/^(?:Chrome|IE|Opera)$/.test(r)||r&&!s&&!/Browser|Mobi/.test(r))&&(i=="Windows CE"||/Mobi/i.test(t)))r+=" Mobile";else if(r=="IE"&&k)try{c.external===null&&d.unshift("platform preview")}catch{d.unshift("embedded")}else(/\bBlackBerry\b/.test(s)||/\bBB10\b/.test(t))&&(e=(RegExp(s.replace(/ +/g," *")+"/([.\\d]+)","i").exec(t)||0)[1]||o)?(e=[e,/BB10/.test(t)],i=(e[1]?(s=null,g="BlackBerry"):"Device Software")+" "+e[0],o=null):this!=W&&s!="Wii"&&(k&&B||/Opera/.test(r)&&/\b(?:MSIE|Firefox)\b/i.test(t)||r=="Firefox"&&/\bOS X (?:\d+\.){2,}/.test(i)||r=="IE"&&(i&&!/^Win/.test(i)&&o>5.5||/\bWindows XP\b/.test(i)&&o>8||o==8&&!/\bTrident\b/.test(t)))&&!E.test(e=K.call(W,t.replace(E,"")+";"))&&e.name&&(e="ing as "+e.name+((e=e.version)?" "+e:""),E.test(r)?(/\bIE\b/.test(e)&&i=="Mac OS"&&(i=null),e="identify"+e):(e="mask"+e,V?r=R(V.replace(/([a-z])([A-Z])/g,"$1 $2")):r="Opera",/\bIE\b/.test(e)&&(i=null),k||(o=null)),p=["Presto"],d.push(e));(e=(/\bAppleWebKit\/([\d.]+\+?)/i.exec(t)||0)[1])&&(e=[parseFloat(e.replace(/\.(\d)$/,".0$1")),e],r=="Safari"&&e[1].slice(-1)=="+"?(r="WebKit Nightly",D="alpha",o=e[1].slice(0,-1)):(o==e[1]||o==(e[2]=(/\bSafari\/([\d.]+\+?)/i.exec(t)||0)[1]))&&(o=null),e[1]=(/\b(?:Headless)?Chrome\/([\d.]+)/i.exec(t)||0)[1],e[0]==537.36&&e[2]==537.36&&parseFloat(e[1])>=28&&p=="WebKit"&&(p=["Blink"]),!k||!le&&!e[1]?(p&&(p[1]="like Safari"),e=(e=e[0],e<400?1:e<500?2:e<526?3:e<533?4:e<534?"4+":e<535?5:e<537?6:e<538?7:e<601?8:e<602?9:e<604?10:e<606?11:e<608?12:"12")):(p&&(p[1]="like Chrome"),e=e[1]||(e=e[0],e<530?1:e<532?2:e<532.05?3:e<533?4:e<534.03?5:e<534.07?6:e<534.1?7:e<534.13?8:e<534.16?9:e<534.24?10:e<534.3?11:e<535.01?12:e<535.02?"13+":e<535.07?15:e<535.11?16:e<535.19?17:e<536.05?18:e<536.1?19:e<537.01?20:e<537.11?"21+":e<537.13?23:e<537.18?24:e<537.24?25:e<537.36?26:p!="Blink"?"27":"28")),p&&(p[1]+=" "+(e+=typeof e=="number"?".x":/[.+]/.test(e)?"":"+")),r=="Safari"&&(!o||parseInt(o)>45)?o=e:r=="Chrome"&&/\bHeadlessChrome/i.test(t)&&d.unshift("headless")),r=="Opera"&&(e=/\bzbov|zvav$/.exec(i))?(r+=" ",d.unshift("desktop mode"),e=="zvav"?(r+="Mini",o=null):r+="Mobile",i=i.replace(RegExp(" *"+e+"$"),"")):r=="Safari"&&/\bChrome\b/.exec(p&&p[1])?(d.unshift("desktop mode"),r="Chrome Mobile",o=null,/\bOS X\b/.test(i)?(g="Apple",i="iOS 4.3+"):i=null):/\bSRWare Iron\b/.test(r)&&!o&&(o=re("Chrome")),o&&o.indexOf(e=/[\d.]+$/.exec(i))==0&&t.indexOf("/"+e+"-")>-1&&(i=z(i.replace(e,""))),i&&i.indexOf(r)!=-1&&!RegExp(r+" OS").test(i)&&(i=i.replace(RegExp(" *"+P(r)+" *"),"")),p&&!/\b(?:Avant|Nook)\b/.test(r)&&(/Browser|Lunascape|Maxthon/.test(r)||r!="Safari"&&/^iOS/.test(i)&&/\bSafari\b/.test(p[1])||/^(?:Adobe|Arora|Breach|Midori|Opera|Phantom|Rekonq|Rock|Samsung Internet|Sleipnir|SRWare Iron|Vivaldi|Web)/.test(r)&&p[1])&&(e=p[p.length-1])&&d.push(e),d.length&&(d=["("+d.join("; ")+")"]),g&&s&&s.indexOf(g)<0&&d.push("on "+g),s&&d.push((/^on /.test(d[d.length-1])?"":"on ")+s),i&&(e=/ ([\d.+]+)$/.exec(i),Z=e&&i.charAt(i.length-e[0].length-1)=="/",i={architecture:32,family:e&&!Z?i.replace(e[0],""):i,version:e?e[1]:null,toString:function(){var x=this.version;return this.family+(x&&!Z?" "+x:"")+(this.architecture==64?" 64-bit":"")}}),(e=/\b(?:AMD|IA|Win|WOW|x86_|x)64\b/i.exec(j))&&!/\bi686\b/i.test(j)?(i&&(i.architecture=64,i.family=i.family.replace(RegExp(" *"+e),"")),r&&(/\bWOW64\b/i.test(t)||k&&/\w(?:86|32)$/.test(v.cpuClass||v.platform)&&!/\bWin64; x64\b/i.test(t))&&d.unshift("32-bit")):i&&/^OS X/.test(i.family)&&r=="Chrome"&&parseFloat(o)>=39&&(i.architecture=64),t||(t=null);var w={};return w.description=t,w.layout=p&&p[0],w.manufacturer=g,w.name=r,w.prerelease=D,w.product=s,w.ua=t,w.version=r&&o,w.os=i||{architecture:null,family:null,version:null,toString:function(){return"null"}},w.parse=K,w.toString=xe,w.version&&d.unshift(o),w.name&&d.unshift(r),i&&r&&!(i==String(i).split(" ")[0]&&(i==r.split(" ")[0]||s))&&d.push(s?"("+i+")":"on "+i),d.length&&(w.description=d.join(" ")),w}var ee=K();l&&S?W(ee,function(t,c){l[c]=t}):f.platform=ee}).call(Q)})(U,U.exports);var We=U.exports;const F=$e(We),_e=typeof window<"u",Te=function(n={}){const a={};return a.detailed=n.detailed===!0,a.ignoreLocalhost=n.ignoreLocalhost!==!1,a.ignoreOwnVisits=n.ignoreOwnVisits!==!1,a},Fe=function(n){return n===""||n==="localhost"||n==="127.0.0.1"||n==="::1"},Ge=function(n){return/bot|crawler|spider|crawling/i.test(n)},H=function(n){return n==="88888888-8888-8888-8888-888888888888"},ie=function(){return document.visibilityState==="hidden"},Le=function(){const n=(location.search.split("source=")[1]||"").split("&")[0];return n===""?void 0:n},oe=function(n=!1){const a={siteLocation:window.location.href,siteReferrer:document.referrer,source:Le()},b={siteLanguage:(navigator.language||navigator.userLanguage).substr(0,2),screenWidth:screen.width,screenHeight:screen.height,screenColorDepth:screen.colorDepth,deviceName:F.product,deviceManufacturer:F.manufacturer,osName:F.os.family,osVersion:F.os.version,browserName:F.name,browserVersion:F.version,browserWidth:window.outerWidth,browserHeight:window.outerHeight};return{...a,...n===!0?b:{}}},Ne=function(n,a){return{query:`
			mutation createRecord($domainId: ID!, $input: CreateRecordInput!) {
				createRecord(domainId: $domainId, input: $input) {
					payload {
						id
					}
				}
			}
		`,variables:{domainId:n,input:a}}},ne=function(n){return{query:`
			mutation updateRecord($recordId: ID!) {
				updateRecord(id: $recordId) {
					success
				}
			}
		`,variables:{recordId:n}}},Ke=function(n,a){return{query:`
			mutation createAction($eventId: ID!, $input: CreateActionInput!) {
				createAction(eventId: $eventId, input: $input) {
					payload {
						id
					}
				}
			}
		`,variables:{eventId:n,input:a}}},Xe=function(n,a){return{query:`
			mutation updateAction($actionId: ID!, $input: UpdateActionInput!) {
				updateAction(id: $actionId, input: $input) {
					success
				}
			}
		`,variables:{actionId:n,input:a}}},Ve=function(n){const a=n.substr(-1)==="/";return n+(a===!0?"":"/")+"api"},N=function(n,a,b,f){const l=new XMLHttpRequest;l.open("POST",n),l.onload=()=>{if(l.status!==200)throw new Error("Server returned with an unhandled status");let S=null;try{S=JSON.parse(l.responseText)}catch{throw new Error("Failed to parse response from server")}if(S.errors!=null)throw new Error(S.errors[0].message);if(typeof f=="function")return f(S)},l.setRequestHeader("Content-Type","application/json;charset=UTF-8"),l.withCredentials=b.ignoreOwnVisits,l.send(JSON.stringify(a))},je=function(){const n=document.querySelector("[data-ackee-domain-id]");if(n==null)return;const a=n.getAttribute("data-ackee-server")||"",b=n.getAttribute("data-ackee-domain-id"),f=n.getAttribute("data-ackee-opts")||"{}";ae(a,JSON.parse(f)).record(b)},ae=function(n,a){a=Te(a);const b=Ve(n),f=()=>{},l={record:()=>({stop:f}),updateRecord:()=>({stop:f}),action:f,updateAction:f};return a.ignoreLocalhost===!0&&Fe(location.hostname)===!0?(console.warn("Ackee ignores you because you are on localhost"),l):Ge(navigator.userAgent)===!0?(console.warn("Ackee ignores you because you are a bot"),l):{record:(y,C=oe(a.detailed),M)=>{let A=!1;const $=()=>{A=!0};return N(b,Ne(y,C),a,q=>{const R=q.data.createRecord.payload.id;if(H(R)===!0)return console.warn("Ackee ignores you because this is your own site");const W=setInterval(()=>{if(A===!0){clearInterval(W);return}ie()!==!0&&N(b,ne(R),a)},15e3);if(typeof M=="function")return M(R)}),{stop:$}},updateRecord:y=>{let C=!1;const M=()=>{C=!0};if(H(y)===!0)return console.warn("Ackee ignores you because this is your own site"),{stop:M};const A=setInterval(()=>{if(C===!0){clearInterval(A);return}ie()!==!0&&N(b,ne(y),a)},15e3);return{stop:M}},action:(y,C,M)=>{N(b,Ke(y,C),a,A=>{const $=A.data.createAction.payload.id;if(H($)===!0)return console.warn("Ackee ignores you because this is your own site");if(typeof M=="function")return M($)})},updateAction:(y,C)=>{if(H(y)===!0)return console.warn("Ackee ignores you because this is your own site");N(b,Xe(y,C),a)}}};_e===!0&&je();const De=function(n){const a=ke({current:void 0,previous:void 0});return a.subscribe(b=>{n(!b.previous||!b.current||b.previous.pathname!==b.current.pathname)}),a},He=function(n,a,{server:b,domainId:f},l={}){let S=!1,O=De(E=>S=E),G=ae(b,l);return n(()=>{typeof window<"u"&&O.update(E=>({previous:E.current,current:{...window.location}}))}),a(()=>{if(S){let E=window.location.pathname;const y=oe(l.detailed),C=new URL(E,location);G.record(f,{...y,siteLocation:C.href}).stop}}),G},Ue="https://ackee-developemnt.up.railway.app",qe="7684f35d-6e48-4386-ab83-fa251b73a790";function ze(n){let a;const b=n[1].default,f=ye(b,n,n[0],null);return{c(){f&&f.c()},l(l){f&&f.l(l)},m(l,S){f&&f.m(l,S),a=!0},p(l,[S]){f&&f.p&&(!a||S&1)&&Oe(f,b,l,l[0],a?Ee(b,l[0],S,null):Ce(l[0]),null)},i(l){a||(Re(f,l),a=!0)},o(l){Be(f,l),a=!1},d(l){f&&f.d(l)}}}function Je(n,a,b){let{$$slots:f={},$$scope:l}=a;return He(Ie,Me,{server:Ue,domainId:qe},{ignoreLocalhost:!1}),n.$$set=S=>{"$$scope"in S&&b(0,l=S.$$scope)},[l,f]}class et extends Ae{constructor(a){super(),Pe(this,a,Je,ze,we,{})}}export{et as component};
