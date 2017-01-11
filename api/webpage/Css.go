package webpage
import "gopkg.in/gin-gonic/gin.v1"
func ServeCss(ctx *gin.Context) {
	content := `
		.card-meta,.chip .chip-content,.chip-sm .chip-name,.text-ellipsis,.tooltip::after{text-overflow:ellipsis}.hidden{display:none}.sidenav a,.ui.progress,article,aside,details,figcaption,figure,footer,header,main,menu,nav,section,summary{display:block}.bg-grey{background-color:#DDD;border-bottom:1px solid silver}h2{font-size:2.4rem!important}h3{font-size:2.2rem!important}.hamburger{font-size:2em;cursor:pointer}.sidenav{height:100%;position:fixed;z-index:99;top:0;right:0;background-color:#111;overflow-x:hidden;padding-top:60px;width:0;opacity:.97}.sidenav a{padding:8px 8px 8px 32px;text-decoration:none;font-size:1.8em;color:#818181}.offcanvas a:focus,.sidenav a:hover{color:#f1f1f1}.sidenav>.close{position:absolute;top:0;right:25px;font-size:36px!important;margin-left:50px}.columns.cards.flex{-webkit-flex-wrap:wrap;-ms-flex-wrap:wrap;flex-wrap:wrap;margin-bottom:30px}.card{color:rgba(51,51,51,.9);min-width:300px;cursor:pointer}.card.material{background:#fff;box-shadow:0 1px 3px rgba(0,0,0,.12),0 1px 2px rgba(0,0,0,.24);transition:all .3s cubic-bezier(.25,.8,.25,1)}.card.material:hover{box-shadow:0 14px 28px rgba(0,0,0,.25),0 10px 10px rgba(0,0,0,.22)}h4.card-title{margin-top:10px!important}.transaction.amount{padding:0;margin:15px 0 5px;font-size:1.2em}.card-meta{overflow:hidden}.card-body,pre,textarea{overflow:auto}.card-body{word-break:break-all}.toast.details>.header{font-size:1.5em}.toast.details>.meta{font-size:1.1em}.toast.details>.footer{font-size:1.3em}.margin-t-10{margin-top:10px}.margin-t-50{margin-top:50px}.margin-0{margin:0}.margin-b-10{margin-bottom:10px}.ui.tiny.progress{position:fixed;top:0;left:0;right:0;bottom:auto;width:100%;z-index:102;font-size:.85714286rem}.ui.progress{position:relative;max-width:100%;border:none;box-shadow:none;padding:0;border-radius:0}.ui.tiny.progress .bar{height:.5em}.ui.active.progress .bar{position:relative;width:100%;-webkit-animation:progress-active 0s ease infinite;animation:progress-active 0s ease infinite;background:#5764c6}.ui.progress .bar{height:1.75em;display:block;line-height:1;position:relative;width:0}@keyframes progress-active{0%{opacity:.7;width:0}100%{opacity:.3;width:100%}}.form-group.buttons{padding-top:10px;padding-bottom:20px}.btn.btn-secondary{background:#8683e2;border-color:#7875d1;color:#fff}.btn.btn-secondary:focus{background:#4f5dc3}.btn.btn-secondary:hover{background:#4d4b87;color:#fff}.btn.btn-secondary.active,.btn.btn-secondary:active{background:#8683e2;border-color:#7875d1;color:#fff}.validation.error{color:red}.btn-fixed{position:fixed;bottom:15px;right:15px;z-index:2}.btn-circle{border-radius:25px!important;font-size:2rem!important;height:50px!important;width:50px}/*! normalize.css v4.1.1 | MIT License | github.com/necolas/normalize.css */a,abbr[title]{text-decoration:underline}progress,sub,sup{vertical-align:baseline}button,hr,input{overflow:visible}html,legend{box-sizing:border-box}blockquote p:last-child,pre code{margin-bottom:0}.btn-group .btn:focus,.btn-group .btn:hover,.input-group .form-input:focus,.input-group .input-group-addon:focus,.input-group .input-group-btn:focus{z-index:99}html{font-family:sans-serif;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%}audio,canvas,progress,video{display:inline-block}audio:not([controls]){display:none;height:0}[hidden],template{display:none}a{background-color:transparent;-webkit-text-decoration-skip:objects;color:#5764c6}a:active,a:hover{outline-width:0}abbr[title]{border-bottom:none;text-decoration:underline dotted}.breadcrumb .breadcrumb-item a,.btn,.chip-sm,.menu .menu-item,.menu .menu-item a,.navbar .navbar-brand,.tab .tab-item a{text-decoration:none}b,strong{font-weight:bolder}dfn{font-style:italic}h1{margin:.67em 0}mark{background-color:#ff0;color:#000}small{font-size:80%}sub,sup{font-size:75%;line-height:0;position:relative}sub{bottom:-.25em}sup{top:-.5em}img{border-style:none}svg:not(:root){overflow:hidden}code,kbd,pre,samp{font-family:monospace,monospace;font-size:1em}figure{margin:1em 40px}hr{box-sizing:content-box;height:0}button,input,optgroup,select,textarea{font:inherit;margin:0}optgroup{font-weight:700}button,select{text-transform:none}[type=reset],[type=submit],button,html [type=button]{-webkit-appearance:button}[type=button]::-moz-focus-inner,[type=reset]::-moz-focus-inner,[type=submit]::-moz-focus-inner,button::-moz-focus-inner{border-style:none;padding:0}[type=button]:-moz-focusring,[type=reset]:-moz-focusring,[type=submit]:-moz-focusring,button:-moz-focusring{outline:ButtonText dotted 1px}fieldset{border:1px solid silver;margin:0 2px;padding:.35em .625em .75em}legend{color:inherit;display:table;max-width:100%;padding:0;white-space:normal}.btn,.chip .chip-content,.chip-sm .chip-name,.label,.text-clip,.text-ellipsis,code{white-space:nowrap}[type=checkbox],[type=radio]{box-sizing:border-box;padding:0}[type=number]::-webkit-inner-spin-button,[type=number]::-webkit-outer-spin-button{height:auto}[type=search]{-webkit-appearance:textfield;outline-offset:-2px}[type=search]::-webkit-search-cancel-button,[type=search]::-webkit-search-decoration{-webkit-appearance:none}::-webkit-input-placeholder{color:inherit;opacity:.54}::-webkit-file-upload-button{-webkit-appearance:button;font:inherit}.btn,.form-input,.form-select{-webkit-appearance:none;outline:0}*,::after,::before{box-sizing:inherit}html{font-size:10px;line-height:1.42857143;-webkit-tap-highlight-color:transparent}body{margin:0;background:#fff;color:#333;font-family:"Helvetica Neue","PingFang SC","Hiragino Sans GB","Microsoft YaHei","Hiragino Kaku Gothic Pro",Meiryo,"Malgun Gothic",sans-serif;font-size:1.4rem;overflow-x:hidden}a:focus,a:hover{color:#283176}.disabled,[disabled]{cursor:default;opacity:.75;pointer-events:none}.btn,.form-switch,.hand{cursor:pointer}.btn .icon,.menu .icon,.toast .icon{font-size:1.3333em;line-height:.8em;margin-right:.2rem;vertical-align:-20%}pre,pre code{line-height:1.8rem}h1,h2,h3,h4,h5,h6{color:inherit;font-weight:300;line-height:1.1;margin-bottom:1.5rem;margin-top:2.5rem}h1{font-size:5rem}h4{font-size:2.4rem}h5{font-size:2rem}h6{font-size:1.6rem}p{margin:0 0 1rem}blockquote{border-left:.2rem solid #ddd;margin-left:0;padding:1rem 2rem}blockquote cite{color:#b3b3b3}ol,ul{margin:2rem 0 2rem 2rem;padding:0}ol ol,ol ul,ul ol,ul ul{margin:1.5rem 0 1.5rem 2rem}ol li,ul li{margin-top:1rem}ul{list-style:disc inside}ul ul{list-style-type:circle}ol{list-style:decimal inside}ol ol{list-style-type:lower-alpha}dl dt{font-weight:700}dl dd{margin:.5rem 0 1.5rem}.lead{font-size:120%}.highlight,code,mark{border-radius:.2rem;display:inline;font-size:1em;padding:.1em .3em;vertical-align:baseline}.highlight,mark{background:#ffe5a3}pre{background:#f9f9f9;border-left:.2rem solid #5764c6;margin-bottom:1em;margin-top:1em;padding:1.5rem}code{background:#efefef}.btn,.form-select{vertical-align:middle}pre code{background:0 0;border-left:0;margin-top:0}.btn,.form-input{line-height:1.6rem}.img-responsive{display:block;height:auto;max-width:100%}.video-responsive{height:0;overflow:hidden;padding-bottom:56.25%;padding-top:3rem;position:relative}.video-responsive embed,.video-responsive iframe,.video-responsive object{height:100%;left:0;position:absolute;top:0;width:100%}.video-responsive video{height:auto;max-width:100%;width:100%}.video-responsive-4-3{padding-bottom:75%}.table{border-collapse:collapse;border-spacing:0;width:100%}.table.table-striped tbody tr:nth-of-type(odd){background:#fcfcfc}.table.table-hover tbody tr:hover{background:#f4f4f4}.table.table-hover tbody tr.selected{background:#f2f2f2}.table td,.table th{border-bottom:.1rem solid #efefef;padding:1.5rem 1rem;text-align:left}.btn,.empty{text-align:center}.table th{border-color:#c9c9c9}.btn{background:0 0;border:.1rem solid #5764c6;border-radius:.3rem;color:#5764c6;display:inline-block;font-size:1.4rem;height:3.2rem;padding:.7rem 1.5rem;-webkit-user-select:none;-moz-user-select:none;-ms-user-select:none;user-select:none}.btn:focus{background:#f3f4fb}.btn:hover{background:#5764c6;border-color:#4452c0;color:#fff}.btn.active,.btn:active{background:#4452c0;border-color:#3b49af;color:#fff}.btn.btn-primary{background:#5764c6;border-color:#4452c0;color:#fff}.btn.btn-primary:focus{background:#4f5dc3}.btn.btn-primary:hover{background:#4452c0;border-color:#3b49af;color:#fff}.btn.btn-primary.active,.btn.btn-primary:active{background:#3b49af;border-color:#35419c;color:#fff}.btn.btn-primary.loading::after{border-color:transparent transparent #fff #fff}.btn.btn-link{background:0 0;border-color:transparent;color:#5764c6}.btn.btn-link:focus,.btn.btn-link:hover{color:#35419c}.btn.btn-link.active,.btn.btn-link:active{color:#283176}.btn.btn-sm{border-radius:.2rem;font-size:1.2rem;height:2.4rem;line-height:1.4rem;padding:.4rem 1rem}.btn.btn-lg{border-radius:.4rem;font-size:1.8rem;height:4.2rem;line-height:2rem;padding:1rem 1.8rem}.btn.btn-block{display:block;width:100%}.btn.btn-clear{background:0 0;border:0;color:#666;height:2rem;margin-left:.3rem;opacity:.45;padding:0}.btn.btn-clear:hover{opacity:.85}.btn.btn-clear::before{content:"\00d7";font-size:2rem}.btn-group{display:inline-flex;display:-ms-inline-flexbox;display:-webkit-inline-flex;-webkit-flex-wrap:wrap;-ms-flex-wrap:wrap;flex-wrap:wrap}.btn-group .btn{-webkit-flex:1 0 auto;-ms-flex:1 0 auto;flex:1 0 auto}.btn-group .btn:first-child:not(:last-child){border-bottom-right-radius:0;border-top-right-radius:0}.btn-group .btn:not(:first-child):not(:last-child){border-radius:0;margin-left:-.1rem}.btn-group .btn:last-child:not(:first-child){border-bottom-left-radius:0;border-top-left-radius:0;margin-left:-.1rem}.btn-group.btn-group-block{display:flex;display:-ms-flexbox;display:-webkit-flex}.form-group:not(:last-child){margin-bottom:1rem}.form-input{background:#fff;border:.1rem solid #c5c5c5;border-radius:.3rem;color:#333;display:block;font-size:1.4rem;height:3.2rem;max-width:100%;padding:.7rem .8rem;position:relative;width:100%}.form-input:focus{border-color:#5764c6}.form-input[disabled]{background:#eeeff2}.form-input.input-sm{border-radius:.2rem;font-size:1.2rem;height:2.4rem;padding:.3rem .6rem}.form-input.input-lg{border-radius:.4rem;font-size:1.6rem;height:4.2rem;line-height:2rem;padding:1rem .8rem}.form-input.input-inline{display:inline-block;vertical-align:middle;width:auto}textarea.form-input{height:auto;line-height:2rem}.form-input.is-success,.has-success .form-input{border-color:#32b643}.form-input.is-danger,.has-danger .form-input{border-color:#e85600}.form-label{display:block;line-height:1.6rem;margin-bottom:.5rem}.form-select{-moz-appearance:none;appearance:none;border:.1rem solid #c5c5c5;border-radius:.3rem;font-size:1.4rem;line-height:1.6rem;padding:.5rem .8rem}.form-select:not([multiple]){background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAUCAMAAACzvE1FAAAADFBMVEUzMzMzMzMzMzMzMzMKAG/3AAAAA3RSTlMAf4C/aSLHAAAAPElEQVR42q3NMQ4AIAgEQTn//2cLdRKppSGzBYwzVXvznNWs8C58CiussPJj8h6NwgorrKRdTvuV9v16Afn0AYFOB7aYAAAAAElFTkSuQmCC) right .75rem center/.8rem 1rem no-repeat #fff;height:3.2rem;padding-right:2.4rem}.form-select:focus{border-color:#5764c6}.form-select::-ms-expand{display:none}.form-select.select-sm{border-radius:.2rem;font-size:1.2rem;height:2.4rem;padding:.4rem 2rem .4rem .6rem}.form-select.select-lg{font-size:1.6rem;height:4.2rem;line-height:2rem;padding:1rem 2.4rem 1rem .8rem}.form-checkbox input,.form-radio input,.form-switch input{clip:rect(0,0,0,0);height:1px;margin:-1px;overflow:hidden;position:absolute;width:1px}.form-checkbox input:focus+.form-icon,.form-radio input:focus+.form-icon,.form-switch input:focus+.form-icon{box-shadow:0 0 .3rem .1rem #efefef}.form-checkbox,.form-radio{cursor:pointer;display:inline-block;line-height:1.8rem;padding:.3rem 2rem;position:relative}.form-checkbox .form-icon,.form-radio .form-icon{border:.1rem solid #c5c5c5;display:inline-block;font-size:1.4rem;height:1.4rem;left:0;line-height:2.4rem;outline:0;padding:0;position:absolute;top:.5rem;transition:all .15s ease;vertical-align:top;width:1.4rem}.form-checkbox:hover .form-icon,.form-radio:hover .form-icon{border-color:#9f9f9f}.form-checkbox input:checked+.form-icon,.form-radio input:checked+.form-icon{background:#5764c6;border-color:#5764c6}.form-checkbox .form-icon{border-radius:.2rem}.form-checkbox input:checked+.form-icon::after{background-clip:padding-box;border:.2rem solid #fff;border-left-width:0;border-top-width:0;content:"";height:1rem;left:50%;margin-left:-.3rem;margin-top:-.6rem;position:absolute;top:50%;-webkit-transform:rotate(45deg);transform:rotate(45deg);width:.6rem}.form-radio .form-icon{border-radius:.7rem}.form-radio input:checked+.form-icon::after{background:#fff;border-radius:.2rem;content:"";height:.4rem;left:50%;margin-left:-.2rem;margin-top:-.2rem;position:absolute;top:50%;width:.4rem}.form-switch{display:inline-block;line-height:1.8rem;padding:.3rem 2rem .3rem 3.6rem;position:relative}.form-switch .form-icon{background:#c5c5c5;background-clip:padding-box;border:.1rem solid #c5c5c5;border-radius:.9rem;display:inline-block;height:1.6rem;left:0;line-height:2.4rem;outline:0;padding:0;position:absolute;top:.4rem;vertical-align:top;width:2.6rem}.form-switch .form-icon::after{background:#fff;border-radius:.8rem;content:"";display:block;height:1.4rem;left:0;position:absolute;top:0;transition:left .15s ease;width:1.4rem}.form-switch input:checked+.form-icon{background:#5764c6;border-color:#5764c6}.form-switch input:checked+.form-icon::after{left:1rem}.input-group{display:flex;display:-ms-flexbox;display:-webkit-flex}.input-group .input-group-addon{background:#f9f9f9;border:.1rem solid #c5c5c5;border-radius:.3rem;line-height:1.6rem;padding:.7rem .8rem}.input-group .input-group-addon.addon-sm{font-size:1.2rem;padding:.3rem .6rem}.input-group .input-group-addon.addon-lg{font-size:1.6rem;line-height:2rem;padding:1rem .8rem}.input-group .input-group-addon,.input-group .input-group-btn{-webkit-flex:1 0 auto;-ms-flex:1 0 auto;flex:1 0 auto}.input-group .form-input:first-child:not(:last-child),.input-group .input-group-addon:first-child:not(:last-child),.input-group .input-group-btn:first-child:not(:last-child){border-bottom-right-radius:0;border-top-right-radius:0}.input-group .form-input:not(:first-child):not(:last-child),.input-group .input-group-addon:not(:first-child):not(:last-child),.input-group .input-group-btn:not(:first-child):not(:last-child){border-radius:0;margin-left:-.1rem}.input-group .form-input:last-child:not(:first-child),.input-group .input-group-addon:last-child:not(:first-child),.input-group .input-group-btn:last-child:not(:first-child){border-bottom-left-radius:0;border-top-left-radius:0;margin-left:-.1rem}.container{margin-left:auto;margin-right:auto;padding-left:1rem;padding-right:1rem;width:100%}@media screen and (min-width:980px){.grid-960{width:98rem}}@media screen and (min-width:500px){.grid-480{width:50rem}}.columns{display:flex;display:-ms-flexbox;display:-webkit-flex;margin-left:-1rem;margin-right:-1rem}.columns.col-gapless{margin-left:0;margin-right:0}.columns.col-gapless .column{padding-left:0;padding-right:0}.column{-webkit-flex:1;-ms-flex:1;flex:1;padding:1rem}.column.col-1,.column.col-10,.column.col-11,.column.col-12,.column.col-2,.column.col-3,.column.col-4,.column.col-5,.column.col-6,.column.col-7,.column.col-8,.column.col-9{-webkit-flex:none;-ms-flex:none;flex:none}.col-12{width:100%}.col-11{width:91.66666667%}.col-10{width:83.33333333%}.col-9{width:75%}.col-8{width:66.66666667%}.col-7{width:58.33333333%}.col-6{width:50%}.col-5{width:41.66666667%}.col-4{width:33.33333333%}.col-3{width:25%}.col-2{width:16.66666667%}.col-1{width:8.33333333%}@media screen and (min-width:481px){.col-xs-1,.col-xs-10,.col-xs-11,.col-xs-12,.col-xs-2,.col-xs-3,.col-xs-4,.col-xs-5,.col-xs-6,.col-xs-7,.col-xs-8,.col-xs-9{-webkit-flex:none;-ms-flex:none;flex:none}.col-xs-12{width:100%}.col-xs-11{width:91.66666667%}.col-xs-10{width:83.33333333%}.col-xs-9{width:75%}.col-xs-8{width:66.66666667%}.col-xs-7{width:58.33333333%}.col-xs-6{width:50%}.col-xs-5{width:41.66666667%}.col-xs-4{width:33.33333333%}.col-xs-3{width:25%}.col-xs-2{width:16.66666667%}.col-xs-1{width:8.33333333%}}@media screen and (min-width:601px){.col-sm-1,.col-sm-10,.col-sm-11,.col-sm-12,.col-sm-2,.col-sm-3,.col-sm-4,.col-sm-5,.col-sm-6,.col-sm-7,.col-sm-8,.col-sm-9{-webkit-flex:none;-ms-flex:none;flex:none}.col-sm-12{width:100%}.col-sm-11{width:91.66666667%}.col-sm-10{width:83.33333333%}.col-sm-9{width:75%}.col-sm-8{width:66.66666667%}.col-sm-7{width:58.33333333%}.col-sm-6{width:50%}.col-sm-5{width:41.66666667%}.col-sm-4{width:33.33333333%}.col-sm-3{width:25%}.col-sm-2{width:16.66666667%}.col-sm-1{width:8.33333333%}}@media screen and (min-width:841px){.col-md-1,.col-md-10,.col-md-11,.col-md-12,.col-md-2,.col-md-3,.col-md-4,.col-md-5,.col-md-6,.col-md-7,.col-md-8,.col-md-9{-webkit-flex:none;-ms-flex:none;flex:none}.col-md-12{width:100%}.col-md-11{width:91.66666667%}.col-md-10{width:83.33333333%}.col-md-9{width:75%}.col-md-8{width:66.66666667%}.col-md-7{width:58.33333333%}.col-md-6{width:50%}.col-md-5{width:41.66666667%}.col-md-4{width:33.33333333%}.col-md-3{width:25%}.col-md-2{width:16.66666667%}.col-md-1{width:8.33333333%}.form-horizontal{padding:1rem}.form-horizontal .form-group{display:flex;display:-ms-flexbox;display:-webkit-flex}.form-horizontal .form-label{margin-bottom:0;padding:.8rem .4rem}.form-horizontal .form-checkbox,.form-horizontal .form-radio{margin:.5rem 0}}@media screen and (max-width:480px){.columns{display:block}.columns .column{width:100%}.hide-xs{display:none!important}}@media screen and (max-width:600px){.hide-sm{display:none!important}}@media screen and (max-width:840px){.hide-md{display:none!important}}@media screen and (max-width:960px){.hide-lg{display:none!important}}@media screen and (max-width:1280px){.hide-xl{display:none!important}}.navbar{-webkit-align-items:center;align-items:center;display:flex;display:-ms-flexbox;display:-webkit-flex;-ms-flex-align:center;-ms-flex-pack:justify;-webkit-justify-content:space-between;justify-content:space-between;padding:1rem}.chip,.chip-sm{-webkit-align-items:center}.navbar .navbar-brand{font-size:2rem;vertical-align:middle}.empty{background:#f8f8f8;border-radius:.3rem;padding:4rem}.empty .empty-title{font-size:1.8rem;margin:1.5rem 0 .5rem}.empty .empty-meta{color:#888}.empty .empty-action{margin-top:1.5rem}.avatar{border-radius:50%;display:inline-block;font-size:1.4rem;font-weight:300;height:3.2rem;line-height:1;margin:0;position:relative;vertical-align:middle;width:3.2rem}.avatar.avatar-xs{font-size:.8rem;height:1.6rem;width:1.6rem}.avatar.avatar-sm{font-size:1rem;height:2.4rem;width:2.4rem}.avatar.avatar-lg{font-size:2rem;height:4.8rem;width:4.8rem}.avatar.avatar-xl{font-size:2.6rem;height:6.4rem;width:6.4rem}.avatar img{border-radius:50%;height:100%;width:100%}.avatar .avatar-icon{background:#fff;bottom:-.2em;height:50%;padding:5%;position:absolute;right:-.2em;width:50%}.avatar[data-initial]::after{color:#fff;content:attr(data-initial);left:50%;position:absolute;top:50%;-webkit-transform:translate(-50%,-50%);transform:translate(-50%,-50%)}@-webkit-keyframes loading{0%{-webkit-transform:rotate(0);transform:rotate(0)}100%{-webkit-transform:rotate(360deg);transform:rotate(360deg)}}@keyframes loading{0%{-webkit-transform:rotate(0);transform:rotate(0)}100%{-webkit-transform:rotate(360deg);transform:rotate(360deg)}}@-webkit-keyframes slide-down{0%{margin-top:-3rem;opacity:0}100%{margin-top:0;opacity:1}}@keyframes slide-down{0%{margin-top:-3rem;opacity:0}100%{margin-top:0;opacity:1}}.form-autocomplete{position:relative}.form-autocomplete .form-autocomplete-input{background:#fff;border:.1rem solid #c5c5c5;border-radius:.3rem;color:#333;display:block;font-size:1.4rem;line-height:1.6rem;max-width:100%;min-height:3.2rem;outline:0;padding:.3rem .3rem 0;width:100%}.form-autocomplete .form-autocomplete-input .chip-sm{margin-bottom:.3rem}.form-autocomplete .form-autocomplete-input .form-input{background:#fff;border-color:transparent;display:inline-block;height:2.4rem;margin-bottom:.3rem;padding:.3rem;vertical-align:top;width:auto}.form-autocomplete .form-autocomplete-list{background:#fff;border:.1rem solid #d2d2d2;border-radius:.3rem;box-shadow:0 .1rem .2rem rgba(0,0,0,.15);display:block;height:auto;left:0;margin:.3rem 0 0;padding:.5rem;position:absolute;top:100%;width:100%;z-index:1988}.menu,.modal-container,.shadow{box-shadow:0 .1rem .4rem rgba(0,0,0,.3)}.form-autocomplete .form-autocomplete-item{border-radius:.3rem;display:block;margin-top:.1rem;padding:.2rem 1rem}.form-autocomplete .form-autocomplete-item:focus,.form-autocomplete .form-autocomplete-item:hover{background:#fff}.form-autocomplete .form-autocomplete-item.active{background:#eff1fa}.badge{position:relative}.badge[data-badge]::after{background:#5764c6;background-clip:padding-box;border:.1rem solid #fff;border-radius:1rem;color:#fff;content:attr(data-badge);display:inline-block;font-size:1.1rem;height:1.8rem;line-height:1.2rem;min-width:1.8rem;padding:.2rem .5rem;text-align:center;-webkit-transform:translate(-.2rem,-.8rem);transform:translate(-.2rem,-.8rem);white-space:nowrap}.card,.menu{display:block;z-index:999}.card,.menu,.text-left{text-align:left}.card{background:#fff;border:.1rem solid #efefef;border-radius:.3rem;margin:0;padding:0}.card .card-body,.card .card-footer,.card .card-header{padding:1.5rem 1.5rem 0}.card .card-body:last-child,.card .card-footer:last-child,.card .card-header:last-child{padding-bottom:1.5rem}.card .card-image{padding-top:1.5rem}.card .card-image:first-child{padding-top:0}.card .card-image:first-child img{border-top-left-radius:.3rem;border-top-right-radius:.3rem}.card .card-image:last-child img{border-bottom-left-radius:.3rem;border-bottom-right-radius:.3rem}.card .card-title{font-size:1.4em;line-height:1;margin-bottom:.5rem;margin-top:0}.card .card-meta{color:#b3b3b3;font-size:1em;margin-bottom:0;margin-top:0}.chip{-webkit-align-content:space-around;align-content:space-around;align-items:center;border:.1rem solid transparent;border-radius:.3rem;display:flex;display:-ms-flexbox;display:-webkit-flex;-ms-flex-align:center;-ms-flex-line-pack:distribute;list-style:none;margin:0;padding:.5rem 0}.chip .chip-icon{-webkit-flex:0 0 auto;-ms-flex:0 0 auto;flex:0 0 auto}.chip .chip-content{-webkit-flex:1 1 auto;-ms-flex:1 1 auto;flex:1 1 auto;overflow:hidden;padding:0 1rem}.chip .chip-action{-webkit-flex:0 0 auto;-ms-flex:0 0 auto;flex:0 0 auto}.chip .chip-title{font-size:1.4rem}.chip .chip-meta{color:#b3b3b3;font-size:1.2rem}.chip-sm{align-items:center;background:#eff1fa;border-radius:.3rem;color:#666;display:-ms-inline-flexbox;display:inline-flex;display:-webkit-inline-flex;-ms-flex-align:center;font-size:1.2rem;height:2.4rem;max-width:100%;padding:.3rem .6rem;vertical-align:middle}.chip-sm:focus,.chip-sm:hover{background:#e8eaf7}.chip-sm .btn-clear{margin-top:-.2rem}.chip-sm .btn-clear::before{color:#3b49af;font-size:1.6rem}.chip-sm.selected{background:#5764c6;color:#fff}.chip-sm.selected .btn-clear::before{color:#eff1fa}.chip-sm .chip-name{margin-left:.4rem;overflow:hidden}.chip-sm .avatar{font-size:.8rem;height:1.6rem;width:1.6rem}.label{background:#efefef;border-radius:.2rem;color:#666;display:inline;font-size:1em;padding:.1em .3em;vertical-align:baseline}.menu,.menu .menu-item,.menu .menu-item a,.modal-container,.tooltip::after{border-radius:.3rem}.label.label-primary{background:#5764c6;border-color:#4f5dc3;color:#fff}.menu{background:#fff;margin:0;padding:.5rem}.menu .menu-header,.menu .menu-item,.menu .menu-item a{display:block;padding:.2rem 1rem}.menu .menu-item{color:#333;line-height:2.4rem;margin-top:.1rem;-webkit-user-select:none;-moz-user-select:none;-ms-user-select:none;user-select:none}.menu .menu-item a{color:inherit;margin:-.2rem -1rem}.menu .menu-item a:focus,.menu .menu-item a:hover{color:#5764c6}.menu .menu-item a.active,.menu .menu-item a:active{background:#eff1fa;color:#4452c0}.menu .menu-header{color:#ccc;font-size:1.2rem;line-height:1.8rem;margin-top:0}.menu .menu-header .menu-header-text{background:#fff;display:inline-block;margin-left:-.6rem;padding:0 .6rem;position:relative;z-index:99}.menu .menu-header::after{border-bottom:.1rem solid #efefef;content:"";display:block;height:.1rem;-webkit-transform:translateY(-1rem);transform:translateY(-1rem);width:100%}.modal{-webkit-align-items:center;align-items:center;bottom:0;display:none;-ms-flex-align:center;-ms-flex-pack:center;-ms-grid-row-align:center;-webkit-justify-content:center;justify-content:center;left:0;opacity:0;overflow:hidden;position:fixed;right:0;top:0}.modal.active{display:flex;display:-ms-flexbox;display:-webkit-flex;opacity:1;z-index:1988}.modal.active .modal-overlay{background:rgba(0,0,0,.75);bottom:0;display:block;left:0;position:absolute;right:0;top:0}.modal.active .modal-container{-webkit-animation:slide-down .216s;animation:slide-down .216s}.modal-container{-webkit-animation:slide-up .216s;animation:slide-up .216s;background:#fff;display:block;margin:0 auto;padding:0;text-align:left;z-index:1988}.modal-container .modal-header{padding:1.5rem}.modal-container .modal-header .modal-title{font-size:1.5rem;margin:0}.modal-container .modal-body{max-height:50vh;overflow-y:auto;padding:1.5rem;position:relative}.modal-container .modal-footer{padding:1.5rem;text-align:right}@media screen and (min-width:640px){.modal-container{width:64rem}}@media screen and (min-width:320px){.modal-sm .modal-container{width:32rem}}.breadcrumb,.pagination,.tab{list-style:none;margin:.5rem 0}.breadcrumb{padding:1.2rem}.breadcrumb .breadcrumb-item{display:inline-block;margin:0}.breadcrumb .breadcrumb-item:last-child,.breadcrumb .breadcrumb-item:last-child a{color:#666;pointer-events:none}.breadcrumb .breadcrumb-item:not(:last-child)::after{color:#c5c5c5;content:"/";padding:0 .4rem}.tab{border-bottom:.1rem solid #c5c5c5;display:flex;display:-ms-flexbox;display:-webkit-flex}.tab .tab-item{margin-bottom:-.1rem;margin-top:0}.tab .tab-item a{border-bottom:.2rem solid transparent;color:#333;display:block;padding:.5rem 1.5rem}.pagination,.pagination .page-item,.pagination .page-item a{display:inline-block}.tab .tab-item a:focus,.tab .tab-item a:hover{border-bottom-color:#5764c6;color:#5764c6}.tab .tab-item.active a{border-bottom-color:#3b49af;color:#3b49af}.tab.tab-block .tab-item{-webkit-flex:1 0 auto;-ms-flex:1 0 auto;flex:1 0 auto;text-align:center}.tab.tab-block .tab-item .badge[data-badge]::after{position:absolute;right:1.5rem;top:.6rem;-webkit-transform:translate(50%,-.8rem);transform:translate(50%,-.8rem)}.pagination{padding:1.2rem}.pagination .page-item span{display:inline-block;padding:.6rem .5rem}.pagination .page-item a{border-radius:.3rem;margin:0 .1rem;padding:.6rem 1.2rem;text-decoration:none}.pagination .page-item a:focus,.pagination .page-item a:hover{background:#eff1fa}.pagination .page-item.active a{background:#5764c6;color:#fff}.toast{background:#efefef;border:.1rem solid #eaeaea;border-radius:.3rem;color:#666;display:block;padding:1.4rem;width:100%}.toast.toast-primary{background:#5764c6;border-color:#4f5dc3;color:#fff}.toast.toast-success{background:#32b643;border-color:#30ae40;color:#fff}.toast.toast-danger{background:#e85600;border-color:#de5200;color:#fff}.toast.toast-danger .btn-clear,.toast.toast-primary .btn-clear,.toast.toast-success .btn-clear{color:#fff}.tooltip{position:relative}.tooltip::after{background:rgba(51,51,51,.9);bottom:100%;color:#fff;content:attr(data-tooltip);display:block;font-size:1.2rem;left:50%;line-height:1.6rem;max-width:32rem;opacity:0;overflow:hidden;padding:.6rem 1rem;pointer-events:none;position:absolute;-webkit-transform:translate(-50%,0);transform:translate(-50%,0);transition:all .216s ease;z-index:99}.tooltip:focus::after,.tooltip:hover::after{opacity:1;-webkit-transform:translate(-50%,-.5rem);transform:translate(-50%,-.5rem)}.tooltip.disabled,.tooltip[disabled]{pointer-events:auto}.tooltip.tooltip-right::after{bottom:50%;left:100%;-webkit-transform:translate(0,50%);transform:translate(0,50%)}.tooltip.tooltip-right:focus::after,.tooltip.tooltip-right:hover::after{-webkit-transform:translate(.5rem,50%);transform:translate(.5rem,50%)}.tooltip.tooltip-bottom::after{bottom:auto;top:100%;-webkit-transform:translate(-50%,0);transform:translate(-50%,0)}.tooltip.tooltip-bottom:focus::after,.tooltip.tooltip-bottom:hover::after{-webkit-transform:translate(-50%,.5rem);transform:translate(-50%,.5rem)}.tooltip.tooltip-left::after{bottom:50%;left:auto;right:100%;-webkit-transform:translate(0,50%);transform:translate(0,50%)}.tooltip.tooltip-left:focus::after,.tooltip.tooltip-left:hover::after{-webkit-transform:translate(-.5rem,50%);transform:translate(-.5rem,50%)}.clearfix::after,.container::after{clear:both;content:"";display:table}.block,.centered{display:block}.float-left{float:left!important}.float-right{float:right!important}.rel{position:relative}.abs{position:absolute}.fixed{position:fixed}.centered{float:none;margin-left:auto;margin-right:auto}.mt-10{margin-top:1rem}.mr-10{margin-right:1rem}.mb-10{margin-bottom:1rem}.ml-10{margin-left:1rem}.mt-5{margin-top:.5rem}.mr-5{margin-right:.5rem}.mb-5{margin-bottom:.5rem}.ml-5{margin-left:.5rem}.pt-10{padding-top:1rem}.pr-10{padding-right:1rem}.pb-10{padding-bottom:1rem}.pl-10{padding-left:1rem}.pt-5{padding-top:.5rem}.pr-5{padding-right:.5rem}.pb-5{padding-bottom:.5rem}.pl-5{padding-left:.5rem}.inline{display:inline}.inline-block{display:inline-block}.flex{display:flex;display:-ms-flexbox;display:-webkit-flex}.inline-flex{display:inline-flex;display:-ms-inline-flexbox;display:-webkit-inline-flex}.hide{display:none!important}.visible{visibility:visible}.invisible{visibility:hidden}.text-hide{background:0 0;border:0;color:transparent;font:0/0 a;text-shadow:none}.text-right{text-align:right}.text-center{text-align:center}.text-justify{text-align:justify}.text-lowercase{text-transform:lowercase}.text-uppercase{text-transform:uppercase}.text-capitalize{text-transform:capitalize}.text-normal{font-weight:400}.text-bold{font-weight:700}.text-italic{font-style:italic}.text-ellipsis{overflow:hidden}.text-clip{overflow:hidden;text-overflow:clip}.text-break{-webkit-hyphens:auto;-ms-hyphens:auto;hyphens:auto;word-break:break-word;word-wrap:break-word}.light-shadow{box-shadow:0 .1rem .2rem rgba(0,0,0,.15)}.rounded{border-radius:.3rem}.circle{border-radius:50%}.divider{border-bottom:.1rem solid #efefef;display:block;margin:.5rem}.loading{color:transparent!important;min-height:1.6rem;pointer-events:none;position:relative}.loading::after{-webkit-animation:loading .5s infinite linear;animation:loading .5s infinite linear;border:.2rem solid #5764c6;border-radius:.8rem;border-right-color:transparent;border-top-color:transparent;content:"";display:block;height:1.6rem;left:50%;margin-left:-.8rem;margin-top:-.8rem;position:absolute;top:50%;width:1.6rem}
	`
	ctx.Writer.Header().Set("Content-Type", "text/css")
	ctx.String(200, content)
}
