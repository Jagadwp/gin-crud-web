(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{31:function(e,t,a){e.exports=a(65)},37:function(e,t,a){},65:function(e,t,a){"use strict";a.r(t);var n=a(1),s=a.n(n),r=a(9),o=a.n(r),l=(a(37),a(19)),i=a(10),m=a(11),c=a(13),d=a(12),u=a(14),h=a(75),p=a(76),g=a(77),E=a(71),f=a(78),b=a(72),v=a(73),y=a(17),S=a(66),j=a(67),O=a(68),C=a(69),w=a(70),k=a(7),D=a.n(k),I=a(16),A=a.n(I),Y=a(22),x=a.n(Y),F=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,s=new Array(n),r=0;r<n;r++)s[r]=arguments[r];return(a=Object(c.a)(this,(e=Object(d.a)(t)).call.apply(e,[this].concat(s)))).state={id:0,name:"",email:"",phone:"",dob:"",gender:"",address:"",errors:{name:{status:!1,message:""},email:{status:!1,message:""},phone:{status:!1,message:""},dob:{status:!1,message:""},gender:{status:!1,message:""},address:{status:!1,message:""}}},a.validDate=function(e){return e.isBefore(x.a.moment())},a.onChange=function(e){a.setState(Object(y.a)({},e.target.name,e.target.value))},a.submitFormAdd=function(e){e.preventDefault(),fetch("/api/v1/users",{method:"post",headers:{"Content-Type":"application/json"},body:JSON.stringify({name:a.state.name,email:a.state.email,phone:a.state.phone,dob:a.state.dob,gender:a.state.gender,address:a.state.address})}).then(function(e){422===e.status?e.json().then(function(e){a.handleApiErrors(e.error)}):e.json().then(function(e){e.success?(e.data&&a.props.addItemToState(e.data),a.props.toggle(),D.a.success(e.message,"Success")):D.a.error(e.message,"Error")})}).catch(function(e){D.a.error("Can't connect to server","Error")})},a.submitFormEdit=function(e){e.preventDefault(),fetch("/api/v1/users/"+a.state.id,{method:"put",headers:{"Content-Type":"application/json"},body:JSON.stringify({id:a.state.id,name:a.state.name,email:a.state.email,phone:a.state.phone,dob:a.state.dob,gender:a.state.gender,address:a.state.address})}).then(function(e){422===e.status?e.json().then(function(e){a.handleApiErrors(e.error)}):e.json().then(function(e){e.success?e.data&&(a.props.updateState(e.data),a.props.toggle(),D.a.success(e.message,"Success")):D.a.error(e.message,"Error")})}).catch(function(e){D.a.error("Can't connect to server","Error")})},a}return Object(u.a)(t,e),Object(m.a)(t,[{key:"handleApiErrors",value:function(e){var t={name:{status:!1,message:""},email:{status:!1,message:""},phone:{status:!1,message:""},dob:{status:!1,message:""},gender:{status:!1,message:""},address:{status:!1,message:""}};e.forEach(function(e){switch(e.id){case"name":t.name.status=!0,t.name.message=e.message;break;case"email":t.email.status=!0,t.email.message=e.message;break;case"phone":t.phone.status=!0,t.phone.message=e.message;break;case"dob":t.dob.status=!0,t.dob.message=e.message;break;case"gender":t.gender.status=!0,t.gender.message=e.message;break;case"address":t.address.status=!0,t.address.message=e.message}}),this.setState({errors:t})}},{key:"componentDidMount",value:function(){if(this.props.item){var e=this.props.item,t=e.id,a=e.name,n=e.email,s=e.phone,r=e.dob,o=e.gender,l=e.address;this.setState({id:t,name:a,email:n,phone:s,dob:r,gender:o,address:l})}}},{key:"render",value:function(){var e=this;return s.a.createElement(S.a,{onSubmit:this.props.item?this.submitFormEdit:this.submitFormAdd},s.a.createElement(j.a,null,s.a.createElement(O.a,{for:"name"},"Name"),s.a.createElement(C.a,{type:"text",name:"name",id:"name",onChange:this.onChange,value:null===this.state.name?"":this.state.name,invalid:this.state.errors.name.status}),s.a.createElement(w.a,null,this.state.errors.name.message)),s.a.createElement(j.a,null,s.a.createElement(O.a,{for:"email"},"Email"),s.a.createElement(C.a,{type:"text",name:"email",id:"email",onChange:this.onChange,value:null===this.state.email?"":this.state.email,invalid:this.state.errors.email.status}),s.a.createElement(w.a,null,this.state.errors.email.message)),s.a.createElement(j.a,null,s.a.createElement(O.a,{for:"phone"},"Phone"),s.a.createElement(C.a,{type:"phone",name:"phone",id:"phone",onChange:this.onChange,value:null===this.state.phone?"":this.state.phone,placeholder:"ex. 081122225555",invalid:this.state.errors.phone.status}),s.a.createElement(w.a,null,this.state.errors.phone.message)),s.a.createElement(j.a,null,s.a.createElement(O.a,{for:"dob"},"Date Of Birth"),s.a.createElement(x.a,{onChange:function(t){e.setState({dob:t.format("YYYY-MM-DD")})},value:this.state.dob?A()(this.state.dob).format("YYYY-MM-DD"):"",renderInput:function(t,a,n){return s.a.createElement("div",null,s.a.createElement(C.a,Object.assign({invalid:e.state.errors.dob.status},t)),s.a.createElement(w.a,null,e.state.errors.dob.message))},timeFormat:!1,dateFormat:"YYYY-MM-DD",max:new Date,isValidDate:this.validDate})),s.a.createElement(j.a,null,s.a.createElement(O.a,{for:"gender"},"Gender"),s.a.createElement(C.a,{type:"select",name:"gender",id:"gender",onChange:this.onChange,value:null===this.state.gender?"":this.state.gender,invalid:this.state.errors.gender.status},s.a.createElement("option",{value:""},"Please choose ..."),s.a.createElement("option",{value:"m"},"Male"),s.a.createElement("option",{value:"f"},"Female")),s.a.createElement(w.a,null,this.state.errors.gender.message)),s.a.createElement(j.a,null,s.a.createElement(O.a,{for:"address"},"Address"),s.a.createElement(C.a,{type:"textarea",name:"address",id:"address",onChange:this.onChange,value:this.state.address,invalid:this.state.errors.address.status}),s.a.createElement(w.a,null,this.state.errors.address.message)),s.a.createElement(E.a,null,"Submit"))}}]),t}(s.a.Component),M=function(e){function t(e){var a;return Object(i.a)(this,t),(a=Object(c.a)(this,Object(d.a)(t).call(this,e))).toggle=function(){a.setState(function(e){return{modal:!e.modal}})},a.state={modal:!1},a}return Object(u.a)(t,e),Object(m.a)(t,[{key:"render",value:function(){var e=s.a.createElement("button",{className:"close",onClick:this.toggle},"\xd7"),t=this.props.buttonLabel,a="",n="";return"Edit"===t?(a=s.a.createElement(E.a,{color:"warning",onClick:this.toggle,style:{float:"left",marginRight:"10px"}},t),n="Edit Item"):(a=s.a.createElement(E.a,{color:"success",onClick:this.toggle,style:{float:"left",marginRight:"10px"}},t),n="Add New Item"),s.a.createElement("div",null,a,s.a.createElement(f.a,{isOpen:this.state.modal,toggle:this.toggle,className:this.props.className},s.a.createElement(b.a,{toggle:this.toggle,close:e},n),s.a.createElement(v.a,null,s.a.createElement(F,{addItemToState:this.props.addItemToState,updateState:this.props.updateState,toggle:this.toggle,item:this.props.item}))))}}]),t}(n.Component),N=a(74),T=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,s=new Array(n),r=0;r<n;r++)s[r]=arguments[r];return(a=Object(c.a)(this,(e=Object(d.a)(t)).call.apply(e,[this].concat(s)))).deleteItem=function(e){window.confirm("Delete item forever?")&&fetch("/api/v1/users/"+e,{method:"delete",headers:{"Content-Type":"application/json"},body:JSON.stringify({id:e})}).then(function(e){return e.json()}).then(function(t){t.success?(a.props.deleteItemFromState(e),D.a.success(t.message,"Success")):D.a.error(t.message,"Error")}).catch(function(e){D.a.error("Can't connect to server","Error")})},a}return Object(u.a)(t,e),Object(m.a)(t,[{key:"render",value:function(){var e=this,t=this.props.items.map(function(t){return s.a.createElement("tr",{key:t.id},s.a.createElement("th",{scope:"row"},t.id),s.a.createElement("td",null,t.name),s.a.createElement("td",null,t.email),s.a.createElement("td",null,t.phone),s.a.createElement("td",null,A()(t.dob).format("YYYY-MM-DD")),s.a.createElement("td",null,"m"===t.gender?"Male":"Female"),s.a.createElement("td",null,t.address),s.a.createElement("td",null,s.a.createElement("div",{style:{width:"110px"}},s.a.createElement(M,{buttonLabel:"Edit",item:t,updateState:e.props.updateState})," ",s.a.createElement(E.a,{color:"danger",onClick:function(){return e.deleteItem(t.id)}},"Del"))))});return s.a.createElement(N.a,{responsive:!0,hover:!0},s.a.createElement("thead",null,s.a.createElement("tr",null,s.a.createElement("th",null,"ID"),s.a.createElement("th",null,"Name"),s.a.createElement("th",null,"Email"),s.a.createElement("th",null,"Phone"),s.a.createElement("th",null,"Date Of Birth"),s.a.createElement("th",null,"Gender"),s.a.createElement("th",null,"Address"),s.a.createElement("th",null,"Actions"))),s.a.createElement("tbody",null,t))}}]),t}(n.Component),J=a(30),R=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,s=new Array(n),r=0;r<n;r++)s[r]=arguments[r];return(a=Object(c.a)(this,(e=Object(d.a)(t)).call.apply(e,[this].concat(s)))).state={items:[]},a.addItemToState=function(e){a.setState(function(t){return{items:[].concat(Object(l.a)(t.items),[e])}})},a.updateState=function(e){var t=a.state.items.findIndex(function(t){return t.id===e.id}),n=[].concat(Object(l.a)(a.state.items.slice(0,t)),[e],Object(l.a)(a.state.items.slice(t+1)));a.setState({items:n})},a.deleteItemFromState=function(e){var t=a.state.items.filter(function(t){return t.id!==e});a.setState({items:t})},a}return Object(u.a)(t,e),Object(m.a)(t,[{key:"getItems",value:function(){var e=this;fetch("/api/v1/users").then(function(e){return e.json()}).then(function(t){t.success?t.data&&e.setState({items:t.data}):D.a.error(t.message,"Error")}).catch(function(e){D.a.error("Can't connect to server","Error")})}},{key:"componentDidMount",value:function(){this.getItems()}},{key:"render",value:function(){return s.a.createElement(h.a,{className:"App"},s.a.createElement(p.a,null,s.a.createElement(g.a,null,s.a.createElement("h2",{style:{margin:"20px 0"}},"CRUD REACTJS - GO REST API"))),s.a.createElement(p.a,null,s.a.createElement(g.a,null,s.a.createElement(T,{items:this.state.items,updateState:this.updateState,deleteItemFromState:this.deleteItemFromState}))),s.a.createElement(p.a,null,s.a.createElement(g.a,null,s.a.createElement(J.CSVLink,{filename:"db.csv",color:"primary",style:{float:"left",marginRight:"10px"},className:"btn btn-primary",data:this.state.items},"Download CSV"),s.a.createElement(M,{buttonLabel:"Add Item",addItemToState:this.addItemToState}))))}}]),t}(n.Component);Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));a(62),a(63),a(64);o.a.render(s.a.createElement(R,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(function(e){e.unregister()})}},[[31,1,2]]]);
//# sourceMappingURL=main.0b7d4f51.chunk.js.map