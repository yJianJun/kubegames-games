window.__require=function e(t,r,o){function n(i,p){if(!r[i]){if(!t[i]){var u=i.split("/");if(u=u[u.length-1],!t[u]){var a="function"==typeof __require&&__require;if(!p&&a)return a(u,!0);if(c)return c(u,!0);throw new Error("Cannot find module '"+i+"'")}i=u}var s=r[i]={exports:{}};t[i][0].call(s.exports,function(e){return n(t[i][1][e]||e)},s,s.exports,e,t,r,o)}return r[i].exports}for(var c="function"==typeof __require&&__require,i=0;i<o.length;i++)n(o[i]);return n}({particleDelay:[function(e,t,r){"use strict";cc._RF.push(t,"43c5baRHj9HtIiF5ThIDKFM","particleDelay"),Object.defineProperty(r,"__esModule",{value:!0});var o=cc._decorator,n=o.ccclass,c=o.property,i=function(e){function t(){var t=null!==e&&e.apply(this,arguments)||this;return t.delayTime=0,t.repeatCount=1,t}return __extends(t,e),t.prototype.start=function(){var e=this;this.node.runAction(cc.repeat(cc.sequence(cc.delayTime(this.delayTime),cc.callFunc(function(){e.node.getComponent(cc.ParticleSystem).resetSystem()})),this.repeatCount))},__decorate([c({type:Number,tooltip:"\u95f4\u9694\u5ef6\u8fdf\u65f6\u95f4"})],t.prototype,"delayTime",void 0),__decorate([c({type:Number,tooltip:"\u91cd\u590d\u51e0\u6b21"})],t.prototype,"repeatCount",void 0),__decorate([n],t)}(cc.Component);r.default=i,cc._RF.pop()},{}],version:[function(e,t,r){"use strict";cc._RF.push(t,"9cd1aLsQHBBoa1PSon+9t52","version"),Object.defineProperty(r,"__esModule",{value:!0});var o=cc._decorator,n=o.ccclass,c=o.property,i=function(e){function t(){var t=null!==e&&e.apply(this,arguments)||this;return t.text=null,t}return __extends(t,e),t.prototype.onLoad=function(){},t.prototype.start=function(){var e=this.text.text.split("\r\n")[0];this.node.getComponent(cc.Label).string=e},__decorate([c(cc.TextAsset)],t.prototype,"text",void 0),__decorate([n],t)}(cc.Component);r.default=i,cc._RF.pop()},{}]},{},["particleDelay","version"]);