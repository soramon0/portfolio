import{k as z,l as B}from"./scheduler.a6abe46d.js";class K{constructor(){this.listeners=new Set,this.subscribe=this.subscribe.bind(this)}subscribe(e){const n={listener:e};return this.listeners.add(n),this.onSubscribe(),()=>{this.listeners.delete(n),this.onUnsubscribe()}}hasListeners(){return this.listeners.size>0}onSubscribe(){}onUnsubscribe(){}}const M=typeof window>"u"||"Deno"in window;function ie(){}function re(t,e){return typeof t=="function"?t(e):t}function H(t){return typeof t=="number"&&t>=0&&t!==1/0}function oe(t,e){return Math.max(t+(e||0)-Date.now(),0)}function ue(t,e,n){return O(t)?typeof e=="function"?{...n,queryKey:t,queryFn:e}:{...e,queryKey:t}:t}function ae(t,e,n){return O(t)?typeof e=="function"?{...n,mutationKey:t,mutationFn:e}:{...e,mutationKey:t}:typeof t=="function"?{...e,mutationFn:t}:{...t}}function le(t,e,n){return O(t)?[{...e,queryKey:t},n]:[t||{},e]}function ce(t,e){const{type:n="all",exact:s,fetchStatus:i,predicate:u,queryKey:l,stale:f}=t;if(O(l)){if(s){if(e.queryHash!==V(l,e.options))return!1}else if(!A(e.queryKey,l))return!1}if(n!=="all"){const h=e.isActive();if(n==="active"&&!h||n==="inactive"&&h)return!1}return!(typeof f=="boolean"&&e.isStale()!==f||typeof i<"u"&&i!==e.state.fetchStatus||u&&!u(e))}function fe(t,e){const{exact:n,fetching:s,predicate:i,mutationKey:u}=t;if(O(u)){if(!e.options.mutationKey)return!1;if(n){if(g(e.options.mutationKey)!==g(u))return!1}else if(!A(e.options.mutationKey,u))return!1}return!(typeof s=="boolean"&&e.state.status==="loading"!==s||i&&!i(e))}function V(t,e){return((e==null?void 0:e.queryKeyHashFn)||g)(t)}function g(t){return JSON.stringify(t,(e,n)=>P(n)?Object.keys(n).sort().reduce((s,i)=>(s[i]=n[i],s),{}):n)}function A(t,e){return D(t,e)}function D(t,e){return t===e?!0:typeof t!=typeof e?!1:t&&e&&typeof t=="object"&&typeof e=="object"?!Object.keys(e).some(n=>!D(t[n],e[n])):!1}function k(t,e){if(t===e)return t;const n=R(t)&&R(e);if(n||P(t)&&P(e)){const s=n?t.length:Object.keys(t).length,i=n?e:Object.keys(e),u=i.length,l=n?[]:{};let f=0;for(let h=0;h<u;h++){const y=n?h:i[h];l[y]=k(t[y],e[y]),l[y]===t[y]&&f++}return s===u&&f===s?t:l}return e}function he(t,e){if(t&&!e||e&&!t)return!1;for(const n in t)if(t[n]!==e[n])return!1;return!0}function R(t){return Array.isArray(t)&&t.length===Object.keys(t).length}function P(t){if(!j(t))return!1;const e=t.constructor;if(typeof e>"u")return!0;const n=e.prototype;return!(!j(n)||!n.hasOwnProperty("isPrototypeOf"))}function j(t){return Object.prototype.toString.call(t)==="[object Object]"}function O(t){return Array.isArray(t)}function q(t){return new Promise(e=>{setTimeout(e,t)})}function L(t){q(0).then(t)}function de(){if(typeof AbortController=="function")return new AbortController}function ye(t,e,n){return n.isDataEqual!=null&&n.isDataEqual(t,e)?t:typeof n.structuralSharing=="function"?n.structuralSharing(t,e):n.structuralSharing!==!1?k(t,e):e}class J extends K{constructor(){super(),this.setup=e=>{if(!M&&window.addEventListener){const n=()=>e();return window.addEventListener("visibilitychange",n,!1),window.addEventListener("focus",n,!1),()=>{window.removeEventListener("visibilitychange",n),window.removeEventListener("focus",n)}}}}onSubscribe(){this.cleanup||this.setEventListener(this.setup)}onUnsubscribe(){if(!this.hasListeners()){var e;(e=this.cleanup)==null||e.call(this),this.cleanup=void 0}}setEventListener(e){var n;this.setup=e,(n=this.cleanup)==null||n.call(this),this.cleanup=e(s=>{typeof s=="boolean"?this.setFocused(s):this.onFocus()})}setFocused(e){this.focused!==e&&(this.focused=e,this.onFocus())}onFocus(){this.listeners.forEach(({listener:e})=>{e()})}isFocused(){return typeof this.focused=="boolean"?this.focused:typeof document>"u"?!0:[void 0,"visible","prerender"].includes(document.visibilityState)}}const W=new J,T=["online","offline"];class X extends K{constructor(){super(),this.setup=e=>{if(!M&&window.addEventListener){const n=()=>e();return T.forEach(s=>{window.addEventListener(s,n,!1)}),()=>{T.forEach(s=>{window.removeEventListener(s,n)})}}}}onSubscribe(){this.cleanup||this.setEventListener(this.setup)}onUnsubscribe(){if(!this.hasListeners()){var e;(e=this.cleanup)==null||e.call(this),this.cleanup=void 0}}setEventListener(e){var n;this.setup=e,(n=this.cleanup)==null||n.call(this),this.cleanup=e(s=>{typeof s=="boolean"?this.setOnline(s):this.onOnline()})}setOnline(e){this.online!==e&&(this.online=e,this.onOnline())}onOnline(){this.listeners.forEach(({listener:e})=>{e()})}isOnline(){return typeof this.online=="boolean"?this.online:typeof navigator>"u"||typeof navigator.onLine>"u"?!0:navigator.onLine}}const Q=new X;function Y(t){return Math.min(1e3*2**t,3e4)}function N(t){return(t??"online")==="online"?Q.isOnline():!0}class G{constructor(e){this.revert=e==null?void 0:e.revert,this.silent=e==null?void 0:e.silent}}function pe(t){return t instanceof G}function Z(t){let e=!1,n=0,s=!1,i,u,l;const f=new Promise((o,c)=>{u=o,l=c}),h=o=>{s||(b(new G(o)),t.abort==null||t.abort())},y=()=>{e=!0},r=()=>{e=!1},d=()=>!W.isFocused()||t.networkMode!=="always"&&!Q.isOnline(),E=o=>{s||(s=!0,t.onSuccess==null||t.onSuccess(o),i==null||i(),u(o))},b=o=>{s||(s=!0,t.onError==null||t.onError(o),i==null||i(),l(o))},C=()=>new Promise(o=>{i=c=>{const p=s||!d();return p&&o(c),p},t.onPause==null||t.onPause()}).then(()=>{i=void 0,s||t.onContinue==null||t.onContinue()}),m=()=>{if(s)return;let o;try{o=t.fn()}catch(c){o=Promise.reject(c)}Promise.resolve(o).then(E).catch(c=>{var p,x;if(s)return;const v=(p=t.retry)!=null?p:3,w=(x=t.retryDelay)!=null?x:Y,S=typeof w=="function"?w(n,c):w,a=v===!0||typeof v=="number"&&n<v||typeof v=="function"&&v(n,c);if(e||!a){b(c);return}n++,t.onFail==null||t.onFail(n,c),q(S).then(()=>{if(d())return C()}).then(()=>{e?b(c):m()})})};return N(t.networkMode)?m():C().then(m),{promise:f,cancel:h,continue:()=>(i==null?void 0:i())?f:Promise.resolve(),cancelRetry:y,continueRetry:r}}const _=console;function $(){let t=[],e=0,n=r=>{r()},s=r=>{r()};const i=r=>{let d;e++;try{d=r()}finally{e--,e||f()}return d},u=r=>{e?t.push(r):L(()=>{n(r)})},l=r=>(...d)=>{u(()=>{r(...d)})},f=()=>{const r=t;t=[],r.length&&L(()=>{s(()=>{r.forEach(d=>{n(d)})})})};return{batch:i,batchCalls:l,schedule:u,setNotifyFunction:r=>{n=r},setBatchNotifyFunction:r=>{s=r}}}const ee=$();class te{destroy(){this.clearGcTimeout()}scheduleGc(){this.clearGcTimeout(),H(this.cacheTime)&&(this.gcTimeout=setTimeout(()=>{this.optionalRemove()},this.cacheTime))}updateCacheTime(e){this.cacheTime=Math.max(this.cacheTime||0,e??(M?1/0:5*60*1e3))}clearGcTimeout(){this.gcTimeout&&(clearTimeout(this.gcTimeout),this.gcTimeout=void 0)}}class ve extends te{constructor(e){super(),this.defaultOptions=e.defaultOptions,this.mutationId=e.mutationId,this.mutationCache=e.mutationCache,this.logger=e.logger||_,this.observers=[],this.state=e.state||ne(),this.setOptions(e.options),this.scheduleGc()}setOptions(e){this.options={...this.defaultOptions,...e},this.updateCacheTime(this.options.cacheTime)}get meta(){return this.options.meta}setState(e){this.dispatch({type:"setState",state:e})}addObserver(e){this.observers.includes(e)||(this.observers.push(e),this.clearGcTimeout(),this.mutationCache.notify({type:"observerAdded",mutation:this,observer:e}))}removeObserver(e){this.observers=this.observers.filter(n=>n!==e),this.scheduleGc(),this.mutationCache.notify({type:"observerRemoved",mutation:this,observer:e})}optionalRemove(){this.observers.length||(this.state.status==="loading"?this.scheduleGc():this.mutationCache.remove(this))}continue(){var e,n;return(e=(n=this.retryer)==null?void 0:n.continue())!=null?e:this.execute()}async execute(){const e=()=>{var a;return this.retryer=Z({fn:()=>this.options.mutationFn?this.options.mutationFn(this.state.variables):Promise.reject("No mutationFn found"),onFail:(F,I)=>{this.dispatch({type:"failed",failureCount:F,error:I})},onPause:()=>{this.dispatch({type:"pause"})},onContinue:()=>{this.dispatch({type:"continue"})},retry:(a=this.options.retry)!=null?a:0,retryDelay:this.options.retryDelay,networkMode:this.options.networkMode}),this.retryer.promise},n=this.state.status==="loading";try{var s,i,u,l,f,h,y,r;if(!n){var d,E,b,C;this.dispatch({type:"loading",variables:this.options.variables}),await((d=(E=this.mutationCache.config).onMutate)==null?void 0:d.call(E,this.state.variables,this));const F=await((b=(C=this.options).onMutate)==null?void 0:b.call(C,this.state.variables));F!==this.state.context&&this.dispatch({type:"loading",context:F,variables:this.state.variables})}const a=await e();return await((s=(i=this.mutationCache.config).onSuccess)==null?void 0:s.call(i,a,this.state.variables,this.state.context,this)),await((u=(l=this.options).onSuccess)==null?void 0:u.call(l,a,this.state.variables,this.state.context)),await((f=(h=this.mutationCache.config).onSettled)==null?void 0:f.call(h,a,null,this.state.variables,this.state.context,this)),await((y=(r=this.options).onSettled)==null?void 0:y.call(r,a,null,this.state.variables,this.state.context)),this.dispatch({type:"success",data:a}),a}catch(a){try{var m,o,c,p,x,v,w,S;throw await((m=(o=this.mutationCache.config).onError)==null?void 0:m.call(o,a,this.state.variables,this.state.context,this)),await((c=(p=this.options).onError)==null?void 0:c.call(p,a,this.state.variables,this.state.context)),await((x=(v=this.mutationCache.config).onSettled)==null?void 0:x.call(v,void 0,a,this.state.variables,this.state.context,this)),await((w=(S=this.options).onSettled)==null?void 0:w.call(S,void 0,a,this.state.variables,this.state.context)),a}finally{this.dispatch({type:"error",error:a})}}}dispatch(e){const n=s=>{switch(e.type){case"failed":return{...s,failureCount:e.failureCount,failureReason:e.error};case"pause":return{...s,isPaused:!0};case"continue":return{...s,isPaused:!1};case"loading":return{...s,context:e.context,data:void 0,failureCount:0,failureReason:null,error:null,isPaused:!N(this.options.networkMode),status:"loading",variables:e.variables};case"success":return{...s,data:e.data,failureCount:0,failureReason:null,error:null,status:"success",isPaused:!1};case"error":return{...s,data:void 0,error:e.error,failureCount:s.failureCount+1,failureReason:e.error,isPaused:!1,status:"error"};case"setState":return{...s,...e.state}}};this.state=n(this.state),ee.batch(()=>{this.observers.forEach(s=>{s.onMutationUpdate(e)}),this.mutationCache.notify({mutation:this,type:"updated",action:e})})}}function ne(){return{context:void 0,data:void 0,error:null,failureCount:0,failureReason:null,isPaused:!1,status:"idle",variables:void 0}}const U="$$_queryClient",be=()=>{const t=z(U);if(!t)throw new Error("No QueryClient was found in Svelte context. Did you forget to wrap your component with QueryClientProvider?");return t},me=t=>{B(U,t)};export{ve as M,te as R,K as S,ee as a,N as b,Z as c,_ as d,fe as e,W as f,de as g,V as h,pe as i,ue as j,re as k,g as l,ce as m,ie as n,Q as o,le as p,A as q,ye as r,me as s,oe as t,he as u,ne as v,be as w,ae as x};