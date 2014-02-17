package godust

import (
	"fmt"
	"io/ioutil"

	Otto "github.com/robertkrimen/otto"
)

const (
	DUSTJS_RUNTIME = `function getGlobal(){return function(){return this.dust}.call(null)}var dust={};(function(dust){function Context(e,t,n,r){this.stack=e;this.global=t;this.blocks=n;this.templateName=r}function Stack(e,t,n,r){this.tail=t;this.isObject=!dust.isArray(e)&&e&&typeof e==="object";this.head=e;this.index=n;this.of=r}function Stub(e){this.head=new Chunk(this);this.callback=e;this.out=""}function Stream(){this.head=new Chunk(this)}function Chunk(e,t,n){this.root=e;this.next=t;this.data=[];this.flushable=false;this.taps=n}function Tap(e,t){this.head=e;this.tail=t}if(!dust){return}var ERROR="ERROR",WARN="WARN",INFO="INFO",DEBUG="DEBUG",levels=[DEBUG,INFO,WARN,ERROR],logger=function(){};dust.isDebug=false;dust.debugLevel=INFO;if(typeof window!=="undefined"&&window&&window.console&&window.console.log){logger=window.console.log}else if(typeof console!=="undefined"&&console&&console.log){logger=console.log}dust.log=function(e,t){var t=t||INFO;if(dust.isDebug&&levels.indexOf(t)>=levels.indexOf(dust.debugLevel)){if(!dust.logQueue){dust.logQueue=[]}dust.logQueue.push({message:e,type:t});logger.call(console||window.console,"[DUST "+t+"]: "+e)}};dust.onError=function(e,t){dust.log(e.message||e,ERROR);if(dust.isDebug){throw e}else{return t}};dust.helpers={};dust.cache={};dust.register=function(e,t){if(!e){return}dust.cache[e]=t};dust.render=function(e,t,n){var r=(new Stub(n)).head;try{dust.load(e,r,Context.wrap(t,e)).end()}catch(i){dust.onError(i,r)}};dust.stream=function(e,t){var n=new Stream;dust.nextTick(function(){try{dust.load(e,n.head,Context.wrap(t,e)).end()}catch(r){dust.onError(r,n.head)}});return n};dust.renderSource=function(e,t,n){return dust.compileFn(e)(t,n)};dust.compileFn=function(e,t){var n=dust.loadSource(dust.compile(e,t));return function(e,r){var i=r?new Stub(r):new Stream;dust.nextTick(function(){if(typeof n==="function"){n(i.head,Context.wrap(e,t)).end()}else{dust.onError(new Error("Template ["+t+"] cannot be resolved to a Dust function"))}});return i}};dust.load=function(e,t,n){var r=dust.cache[e];if(r){return r(t,n)}else{if(dust.onLoad){return t.map(function(t){dust.onLoad(e,function(r,i){if(r){return t.setError(r)}if(!dust.cache[e]){dust.loadSource(dust.compile(i,e))}dust.cache[e](t,n).end()})})}return t.setError(new Error("Template Not Found: "+e))}};dust.loadSource=function(source,path){return eval(source)};if(Array.isArray){dust.isArray=Array.isArray}else{dust.isArray=function(e){return Object.prototype.toString.call(e)==="[object Array]"}}dust.nextTick=function(){if(typeof process!=="undefined"){return process.nextTick}else{return function(e){setTimeout(e,0)}}}();dust.isEmpty=function(e){if(dust.isArray(e)&&!e.length){return true}if(e===0){return false}return!e};dust.filter=function(e,t,n){if(n){for(var r=0,i=n.length;r<i;r++){var s=n[r];if(s==="s"){t=null;dust.log("Using unescape filter on ["+e+"]",DEBUG)}else if(typeof dust.filters[s]==="function"){e=dust.filters[s](e)}else{dust.onError(new Error("Invalid filter ["+s+"]"))}}}if(t){e=dust.filters[t](e)}return e};dust.filters={h:function(e){return dust.escapeHtml(e)},j:function(e){return dust.escapeJs(e)},u:encodeURI,uc:encodeURIComponent,js:function(e){if(!JSON){dust.log("JSON is undefined.  JSON stringify has not been used on ["+e+"]",WARN);return e}else{return JSON.stringify(e)}},jp:function(e){if(!JSON){dust.log("JSON is undefined.  JSON parse has not been used on ["+e+"]",WARN);return e}else{return JSON.parse(e)}}};dust.makeBase=function(e){return new Context(new Stack,e)};Context.wrap=function(e,t){if(e instanceof Context){return e}return new Context(new Stack(e),{},null,t)};Context.prototype.get=function(e,t){if(typeof e==="string"){if(e[0]==="."){t=true;e=e.substr(1)}e=e.split(".")}return this._get(t,e)};Context.prototype._get=function(e,t){var n=this.stack,r=1,i,s,o,u;dust.log("Searching for reference [{"+t.join(".")+"}] in template ["+this.templateName+"]",DEBUG);s=t[0];o=t.length;if(e&&o===0){u=n;n=n.head}else{if(!e){while(n){if(n.isObject){u=n.head;i=n.head[s];if(i!==undefined){break}}n=n.tail}if(i!==undefined){n=i}else{n=this.global?this.global[s]:undefined}}else{n=n.head[s]}while(n&&r<o){u=n;n=n[t[r]];r++}}if(typeof n==="function"){var a=function(){return n.apply(u,arguments)};a.isFunction=true;return a}else{if(n===undefined){dust.log("Cannot find the value for reference [{"+t.join(".")+"}] in template ["+this.templateName+"]")}return n}};Context.prototype.getPath=function(e,t){return this._get(e,t)};Context.prototype.push=function(e,t,n){return new Context(new Stack(e,this.stack,t,n),this.global,this.blocks,this.templateName)};Context.prototype.rebase=function(e){return new Context(new Stack(e),this.global,this.blocks,this.templateName)};Context.prototype.current=function(){return this.stack.head};Context.prototype.getBlock=function(e,t,n){if(typeof e==="function"){var r=new Chunk;e=e(r,this).data.join("")}var i=this.blocks;if(!i){dust.log("No blocks for context[{"+e+"}] in template ["+this.templateName+"]",DEBUG);return}var s=i.length,o;while(s--){o=i[s][e];if(o){return o}}};Context.prototype.shiftBlocks=function(e){var t=this.blocks,n;if(e){if(!t){n=[e]}else{n=t.concat([e])}return new Context(this.stack,this.global,n,this.templateName)}return this};Stub.prototype.flush=function(){var e=this.head;while(e){if(e.flushable){this.out+=e.data.join("")}else if(e.error){this.callback(e.error);dust.onError(new Error("Chunk error ["+e.error+"] thrown. Ceasing to render this template."));this.flush=function(){};return}else{return}e=e.next;this.head=e}this.callback(null,this.out)};Stream.prototype.flush=function(){var e=this.head;while(e){if(e.flushable){this.emit("data",e.data.join(""))}else if(e.error){this.emit("error",e.error);dust.onError(new Error("Chunk error ["+e.error+"] thrown. Ceasing to render this template."));this.flush=function(){};return}else{return}e=e.next;this.head=e}this.emit("end")};Stream.prototype.emit=function(e,t){if(!this.events){dust.log("No events to emit",INFO);return false}var n=this.events[e];if(!n){dust.log("Event type ["+e+"] does not exist",WARN);return false}if(typeof n==="function"){n(t)}else if(dust.isArray(n)){var r=n.slice(0);for(var i=0,s=r.length;i<s;i++){r[i](t)}}else{dust.onError(new Error("Event Handler ["+n+"] is not of a type that is handled by emit"))}};Stream.prototype.on=function(e,t){if(!this.events){this.events={}}if(!this.events[e]){dust.log("Event type ["+e+"] does not exist. Using just the specified callback.",WARN);if(t){this.events[e]=t}else{dust.log("Callback for type ["+e+"] does not exist. Listener not registered.",WARN)}}else if(typeof this.events[e]==="function"){this.events[e]=[this.events[e],t]}else{this.events[e].push(t)}return this};Stream.prototype.pipe=function(e){this.on("data",function(t){try{e.write(t,"utf8")}catch(n){dust.onError(n,e.head)}}).on("end",function(){try{return e.end()}catch(t){dust.onError(t,e.head)}}).on("error",function(t){e.error(t)});return this};Chunk.prototype.write=function(e){var t=this.taps;if(t){e=t.go(e)}this.data.push(e);return this};Chunk.prototype.end=function(e){if(e){this.write(e)}this.flushable=true;this.root.flush();return this};Chunk.prototype.map=function(e){var t=new Chunk(this.root,this.next,this.taps),n=new Chunk(this.root,t,this.taps);this.next=n;this.flushable=true;e(n);return t};Chunk.prototype.tap=function(e){var t=this.taps;if(t){this.taps=t.push(e)}else{this.taps=new Tap(e)}return this};Chunk.prototype.untap=function(){this.taps=this.taps.tail;return this};Chunk.prototype.render=function(e,t){return e(this,t)};Chunk.prototype.reference=function(e,t,n,r){if(typeof e==="function"){e.isFunction=true;e=e.apply(t.current(),[this,t,null,{auto:n,filters:r}]);if(e instanceof Chunk){return e}}if(!dust.isEmpty(e)){return this.write(dust.filter(e,n,r))}else{return this}};Chunk.prototype.section=function(e,t,n,r){if(typeof e==="function"){e=e.apply(t.current(),[this,t,n,r]);if(e instanceof Chunk){return e}}var i=n.block,s=n["else"];if(r){t=t.push(r)}if(dust.isArray(e)){if(i){var o=e.length,u=this;if(o>0){if(t.stack.head){t.stack.head["$len"]=o}for(var a=0;a<o;a++){if(t.stack.head){t.stack.head["$idx"]=a}u=i(u,t.push(e[a],a,o))}if(t.stack.head){t.stack.head["$idx"]=undefined;t.stack.head["$len"]=undefined}return u}else if(s){return s(this,t)}}}else if(e===true){if(i){return i(this,t)}}else if(e||e===0){if(i){return i(this,t.push(e))}}else if(s){return s(this,t)}dust.log("Not rendering section (#) block in template ["+t.templateName+"], because above key was not found",DEBUG);return this};Chunk.prototype.exists=function(e,t,n){var r=n.block,i=n["else"];if(!dust.isEmpty(e)){if(r){return r(this,t)}}else if(i){return i(this,t)}dust.log("Not rendering exists (?) block in template ["+t.templateName+"], because above key was not found",DEBUG);return this};Chunk.prototype.notexists=function(e,t,n){var r=n.block,i=n["else"];if(dust.isEmpty(e)){if(r){return r(this,t)}}else if(i){return i(this,t)}dust.log("Not rendering not exists (^) block check in template ["+t.templateName+"], because above key was found",DEBUG);return this};Chunk.prototype.block=function(e,t,n){var r=n.block;if(e){r=e}if(r){return r(this,t)}return this};Chunk.prototype.partial=function(e,t,n){var r;r=dust.makeBase(t.global);r.blocks=t.blocks;if(t.stack&&t.stack.tail){r.stack=t.stack.tail}if(n){r=r.push(n)}if(typeof e==="string"){r.templateName=e}r=r.push(t.stack.head);var i;if(typeof e==="function"){i=this.capture(e,r,function(e,t){dust.load(e,t,r).end()})}else{i=dust.load(e,this,r)}return i};Chunk.prototype.helper=function(e,t,n,r){var i=this;try{if(dust.helpers[e]){return dust.helpers[e](i,t,n,r)}else{return dust.onError(new Error("Invalid helper ["+e+"]"),i)}}catch(s){return dust.onError(s,i)}};Chunk.prototype.capture=function(e,t,n){return this.map(function(r){var i=new Stub(function(e,t){if(e){r.setError(e)}else{n(t,r)}});e(i.head,t).end()})};Chunk.prototype.setError=function(e){this.error=e;this.root.flush();return this};Tap.prototype.push=function(e){return new Tap(e,this)};Tap.prototype.go=function(e){var t=this;while(t){e=t.head(e);t=t.tail}return e};var HCHARS=new RegExp(/[&<>\"\']/),AMP=/&/g,LT=/</g,GT=/>/g,QUOT=/\"/g,SQUOT=/\'/g;dust.escapeHtml=function(e){if(typeof e==="string"){if(!HCHARS.test(e)){return e}return e.replace(AMP,"&").replace(LT,"&lt;").replace(GT,"&gt;").replace(QUOT,"&quot;").replace(SQUOT,"&#39;")}return e};var BS=/\\/g,FS=/\//g,CR=/\r/g,LS=/\u2028/g,PS=/\u2029/g,NL=/\n/g,LF=/\f/g,SQ=/'/g,DQ=/"/g,TB=/\t/g;dust.escapeJs=function(e){if(typeof e==="string"){return e.replace(BS,"\\\\").replace(FS,"\\/").replace(DQ,'\\"').replace(SQ,"\\'").replace(CR,"\\r").replace(LS,"\\u2028").replace(PS,"\\u2029").replace(NL,"\\n").replace(LF,"\\f").replace(TB,"\\t")}return e}})(dust);if(typeof exports!=="undefined"){if(typeof process!=="undefined"){require("./server")(dust)}module.exports=dust}`
)

var otto = Otto.New()

//for any setup needed to perform rendering()
func setup() {
	//dust-runtime needs the window object to report errors, this is a hack-solution
	otto.Set("window", "{}")
}

func RenderDustjs(templatePath string, templateName string, jsonDataString string) string {
	setup()
	//open the template file
	template := getTemplateFile(templatePath)

	//add the Runtime and template
	code := DUSTJS_RUNTIME + " " + template

	code += `var _jsRenderedResults;
			dust.render("` + templateName + `", ` + jsonDataString + `, function(err, out){
				//console.log(err);
				//console.log(out);
				_jsRenderedResults = out;
			});
			`
	//I don't directly want the value from otto.Run(...) because it's export via exportRenderedResults()
	_, err := otto.Run(code)

	if err != nil {
		fmt.Println("err: ")
		fmt.Println(err)
	}

	return exportRenderedResults()
}

func getTemplateFile(path string) string {
	var contents []byte
	contents, err := ioutil.ReadFile(path)
	fileContents := string(contents)

	if err != nil {
		fmt.Println(err)
	}
	return fileContents
}

func exportRenderedResults() string {
	value, _ := otto.Get("_jsRenderedResults")
	results, _ := value.ToString()
	return results
}
