(function(e){function t(t){for(var r,i,o=t[0],u=t[1],s=t[2],f=0,b=[];f<o.length;f++)i=o[f],Object.prototype.hasOwnProperty.call(c,i)&&c[i]&&b.push(c[i][0]),c[i]=0;for(r in u)Object.prototype.hasOwnProperty.call(u,r)&&(e[r]=u[r]);l&&l(t);while(b.length)b.shift()();return a.push.apply(a,s||[]),n()}function n(){for(var e,t=0;t<a.length;t++){for(var n=a[t],r=!0,o=1;o<n.length;o++){var u=n[o];0!==c[u]&&(r=!1)}r&&(a.splice(t--,1),e=i(i.s=n[0]))}return e}var r={},c={app:0},a=[];function i(t){if(r[t])return r[t].exports;var n=r[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,i),n.l=!0,n.exports}i.m=e,i.c=r,i.d=function(e,t,n){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(i.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)i.d(n,r,function(t){return e[t]}.bind(null,r));return n},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/static/dist/";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],u=o.push.bind(o);o.push=t,o=o.slice();for(var s=0;s<o.length;s++)t(o[s]);var l=u;a.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"19ce":function(e,t,n){},"19e3":function(e,t,n){var r=n("c973").default;n("96cf");var c=n("bc3a");e.exports={base:"",getFiles:function(){var e=r(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,c.get(this.base+"/files?parent_id="+t);case 2:return n=e.sent,e.abrupt("return",n.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),getRoot:function(){var e=r(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,c.get(this.base+"/root");case 2:return e.abrupt("return",e.sent.data);case 3:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}(),createDir:function(){var e=r(regeneratorRuntime.mark((function e(t,n){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,c.post(this.base+"/mkdir",{name:t,parent_id:n});case 2:return e.abrupt("return",e.sent.data);case 3:case"end":return e.stop()}}),e,this)})));function t(t,n){return e.apply(this,arguments)}return t}(),getFile:function(){var e=r(regeneratorRuntime.mark((function e(t){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,c.get(this.base+"/file?id="+t);case 2:return e.abrupt("return",e.sent.data);case 3:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),deleteFile:function(){var e=r(regeneratorRuntime.mark((function e(t){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,c.delete(this.base+"/delete_file?id="+t);case 2:return e.abrupt("return",e.sent.data);case 3:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()}},"37f7":function(e,t,n){"use strict";n("4bee")},"4bee":function(e,t,n){},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var r=n("7a23"),c=function(e){return Object(r["K"])("data-v-33a56e92"),e=e(),Object(r["I"])(),e},a=c((function(){return Object(r["n"])("span",null,"fileCLoud",-1)})),i=Object(r["p"])("Upload File"),o=Object(r["p"])("Create Directory"),u=Object(r["p"])("Download File"),s=["onClick"],l=Object(r["p"])("Download"),f=Object(r["p"])("Delete");function b(e,t,n,c,b,d){var p=Object(r["Q"])("el-col"),h=Object(r["Q"])("el-button"),O=Object(r["Q"])("el-upload"),j=Object(r["Q"])("el-row"),g=Object(r["Q"])("el-breadcrumb-item"),m=Object(r["Q"])("el-breadcrumb"),_=Object(r["Q"])("el-header"),w=Object(r["Q"])("folder"),v=Object(r["Q"])("el-icon"),y=Object(r["Q"])("document"),k=Object(r["Q"])("el-table-column"),x=Object(r["Q"])("el-table"),F=Object(r["Q"])("el-main"),q=Object(r["Q"])("el-container");return Object(r["H"])(),Object(r["k"])(q,{style:{height:"100%",margin:"0 auto"}},{default:Object(r["fb"])((function(){return[Object(r["q"])(_,{style:{height:"15%",border:"salmon solid 1px","padding-top":"1%"}},{default:Object(r["fb"])((function(){return[Object(r["q"])(j,{gutter:20},{default:Object(r["fb"])((function(){return[Object(r["q"])(p,{span:2},{default:Object(r["fb"])((function(){return[a]})),_:1}),Object(r["q"])(p,{span:2},{default:Object(r["fb"])((function(){return[Object(r["q"])(O,{style:{margin:"0 auto"},action:b.upload_action,"show-file-list":!1,"auto-upload":!0,"on-progress":d.uploadProgress,"on-success":function(){d.flushFiles()}},{default:Object(r["fb"])((function(){return[Object(r["q"])(h,{type:"primary"},{default:Object(r["fb"])((function(){return[i]})),_:1})]})),_:1},8,["action","on-progress","on-success"])]})),_:1}),Object(r["q"])(p,{span:2},{default:Object(r["fb"])((function(){return[Object(r["q"])(h,{type:"primary",onClick:d.createDir},{default:Object(r["fb"])((function(){return[o]})),_:1},8,["onClick"])]})),_:1}),Object(r["q"])(p,{span:2},{default:Object(r["fb"])((function(){return[Object(r["q"])(h,{type:"primary"},{default:Object(r["fb"])((function(){return[u]})),_:1})]})),_:1}),Object(r["q"])(p,{span:16})]})),_:1}),Object(r["q"])(j,{style:{"margin-top":"1.5%"}},{default:Object(r["fb"])((function(){return[Object(r["q"])(p,{span:"24"},{default:Object(r["fb"])((function(){return[Object(r["q"])(m,{separator:"/"},{default:Object(r["fb"])((function(){return[(Object(r["H"])(!0),Object(r["m"])(r["b"],null,Object(r["O"])(b.path,(function(e){return Object(r["H"])(),Object(r["k"])(g,{key:e.id},{default:Object(r["fb"])((function(){return[Object(r["n"])("span",{class:"bre-item",onClick:function(t){return d.breadcrumbClick(e)},style:{"font-size":"30px"}},Object(r["U"])(e.file_name),9,s)]})),_:2},1024)})),128))]})),_:1})]})),_:1})]})),_:1})]})),_:1}),Object(r["q"])(F,{style:{height:"84.5%",border:"seagreen solid 1px"}},{default:Object(r["fb"])((function(){return[Object(r["q"])(x,{data:b.files,onRowDblclick:d.rowClick,style:{width:"100%"}},{default:Object(r["fb"])((function(){return[Object(r["q"])(k,{label:"Type"},{default:Object(r["fb"])((function(e){return[1===e.row.is_dir?(Object(r["H"])(),Object(r["k"])(v,{key:0},{default:Object(r["fb"])((function(){return[Object(r["q"])(w)]})),_:1})):(Object(r["H"])(),Object(r["k"])(v,{key:1},{default:Object(r["fb"])((function(){return[Object(r["q"])(y)]})),_:1}))]})),_:1}),Object(r["q"])(k,{prop:"file_name",label:"name"}),Object(r["q"])(k,{label:"File Size"},{default:Object(r["fb"])((function(e){return[Object(r["p"])(Object(r["U"])(d.get_size(e.row.size)),1)]})),_:1}),Object(r["q"])(k,{label:"Upload Time"},{default:Object(r["fb"])((function(e){return[Object(r["p"])(Object(r["U"])(d.format_time(e.row.upload_time)),1)]})),_:1}),Object(r["q"])(k,{label:"Action"},{default:Object(r["fb"])((function(e){return[Object(r["q"])(h,{size:"small",disabled:1===e.row.is_dir,onClick:function(t){return d.downloadFile(e.row.id)}},{default:Object(r["fb"])((function(){return[l]})),_:2},1032,["disabled","onClick"]),Object(r["q"])(h,{size:"small",type:"danger",onClick:function(t){return d.deleteFile(e.row.id)}},{default:Object(r["fb"])((function(){return[f]})),_:2},1032,["onClick"])]})),_:1})]})),_:1},8,["data","onRowDblclick"])]})),_:1})]})),_:1})}var d=n("1da1"),p=(n("99af"),n("b680"),n("96cf"),n("19e3")),h=n.n(p),O=n("0b38"),j=n("218b"),g=n("7864"),m={components:{Folder:O["a"],Document:j["a"]},data:function(){return{path:[],root:{},files:[],current_dir:{},searchValues:"",upload_action:h.a.base+"/upload?parent_id="}},watch:{current_dir:function(){this.upload_action=h.a.base+"/upload?parent_id="+this.current_dir.id,this.freshPath(),this.flushFiles()}},created:function(){var e=this;h.a.getRoot().then((function(t){t.ok?(e.root=t.data,e.path.push(t.data),e.current_dir=t.data,h.a.getFiles(t.data.id).then((function(t){t.ok&&(e.files=t.data)}))):g["c"].error({title:"error",message:t.msg})}))},methods:{deleteFile:function(){var e=Object(d["a"])(regeneratorRuntime.mark((function e(t){var n=this;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:g["b"].confirm("Delete the File?","Warning",{confirmButtonText:"Ok",cancelButtonText:"Cancel",type:"warning"}).then(Object(d["a"])(regeneratorRuntime.mark((function e(){var r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,h.a.deleteFile(t);case 2:r=e.sent,r.ok?n.flushFiles():g["c"].error({message:r.msg});case 4:case"end":return e.stop()}}),e)})))).catch((function(){}));case 1:case"end":return e.stop()}}),e)})));function t(t){return e.apply(this,arguments)}return t}(),downloadFile:function(e){var t=document.createElement("a");t.href=h.a.base+"/download?id="+e,t.click()},breadcrumbClick:function(e){this.current_dir=e},rowClick:function(e,t,n){console.log(e,t,n),1===e.is_dir&&(this.current_dir=e)},freshPath:function(){var e=Object(d["a"])(regeneratorRuntime.mark((function e(){var t,n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(this.path=[],this.current_dir.id!==this.root.id){e.next=5;break}this.path.push(this.root),e.next=18;break;case 5:console.log(this.current_dir),t=this.current_dir.parent_id,this.path.push(this.current_dir);case 8:if(""===t){e.next=18;break}if(void 0!==t){e.next=11;break}return e.abrupt("break",18);case 11:return e.next=13,h.a.getFile(t);case 13:n=e.sent,t=n.data.parent_id,this.path.unshift(n.data),e.next=8;break;case 18:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}(),uploadProgress:function(e,t,n){console.log(e),console.log(t),console.log(n)},createDir:function(){var e=this;g["b"].prompt("please input the directory name:","message",{confirmButtonText:"OK",cancelButtonText:"Cancel"}).then((function(t){h.a.createDir(t.value,e.current_dir.id).then((function(t){t.ok?(g["a"].info("create a directory successful"),e.flushFiles()):g["a"].error(t.msg)}))})).catch((function(){}))},flushFiles:function(){var e=this;h.a.getFiles(this.current_dir.id).then((function(t){t.ok&&(e.files=t.data)}))},format_time:function(e){if(e){var t=new Date(1e3*e),n=t.getFullYear(),r=t.getMonth()+1<10?"0".concat(t.getMonth()+1):t.getMonth()+1,c=t.getDate()<10?"0".concat(t.getDate()):t.getDate(),a=t.getHours()<10?"0".concat(t.getHours()):t.getHours(),i=t.getMinutes()<10?"0".concat(t.getMinutes()):t.getMinutes(),o=t.getSeconds()<10?"0".concat(t.getSeconds()):t.getSeconds();return"".concat(n,"-").concat(r,"-").concat(c," ").concat(a,":").concat(i,":").concat(o)}return""},get_size:function(e){if(!e)return"";var t=1024;return e<t?e+"B":e<Math.pow(t,2)?(e/t).toFixed(2)+"K":e<Math.pow(t,3)?(e/Math.pow(t,2)).toFixed(2)+"M":e<Math.pow(t,4)?(e/Math.pow(t,3)).toFixed(2)+"G":(e/Math.pow(t,4)).toFixed(2)+"T"}}},_=(n("77b7"),n("37f7"),n("6b0d")),w=n.n(_);const v=w()(m,[["render",b],["__scopeId","data-v-33a56e92"]]);var y=v,k=(n("c69f"),n("3ef0")),x=n.n(k),F=function(e){e.use(g["d"],{locale:x.a})},q=n("6c02"),R=[],M=Object(q["a"])({history:Object(q["b"])("/static/dist/"),routes:R}),C=M,D=Object(r["j"])(y).use(C);F(D),D.mount("#app")},"77b7":function(e,t,n){"use strict";n("19ce")},c69f:function(e,t,n){}});
//# sourceMappingURL=app.18da458f.js.map